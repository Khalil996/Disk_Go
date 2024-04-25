package common

import (
	"cloud_go/Disk/common/minio"
	"cloud_go/Disk/common/xorm"
	"cloud_go/Disk/define"
	"cloud_go/Disk/models"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"io"

	"os"
	"strconv"
)

type MergeStruct struct {
	FsId       int64 `xorm:"fsId"`
	SId        int64 `xorm:"sId"`
	ObjectName string
	Hash       string
	ChunkNum   int64
	Size       int64
}

func MergeLogic() {
	logx.Infof("合并任务开始")

	var mss []*MergeStruct
	cols := "a.id as sId, b.id as fsId, b.object_name, b.hash, b.chunk_num, b.size"
	if err := xorm.Xorm.Select(cols).
		Table(&models.FileSchedule{}).Alias("a").
		Join("LEFT", []string{"file_fs", "b"}, "a.fs_id = b.id").
		Where("a.stage = ?", define.StageNeedMerge).
		Limit(1000).Find(&mss); err != nil {
		logx.Errorf("MergeLogic，查询fileSchedule出错，ERR：[%v]", err)
		return
	}

	q := make(chan struct{}, 5)
	for _, ms1 := range mss {
		ms2 := ms1
		go func() {
			q <- struct{}{}
			Merge(ms2, func(int64) {})
			<-q
		}()
	}
}

func Merge(ms *MergeStruct, errCallBack func(int64)) {
	var err error
	defer func() {
		if err != nil {
			errCallBack(ms.SId)
		}
	}()

	minioSvc := minio.Minio.NewService()
	exist, size, _ := minioSvc.IfExist(ms.ObjectName)
	if exist && size == ms.Size {
		logx.Infof("MergeLogic, 文件：[%v] 已上传1", ms.ObjectName)
		return
	}

	//bigFile, err := os.CreateTemp("", "netdisk")
	bigFileName := "D:/netdisk_tmp/bigFile" + ms.Hash
	stat, err := os.Stat(bigFileName)
	if err == nil && stat.Size() == ms.Size {
		uploadAndUpdateDb(minioSvc, ms, bigFileName)
		logx.Infof("MergeLogic, 文件：[%v] 已上传2", ms.ObjectName)
		_ = os.Remove(bigFileName)
		return
	}

	bigFile, err := os.Create(bigFileName)
	if err != nil {
		logx.Errorf("MergeLogic，创建临时文件出错，ERR：[%v]", err)
		return
	}
	defer DeleteTemp(bigFile)

	var (
		chunks          []*os.File
		chunkObjectName []string
	)
	defer deleteChunks(minioSvc, chunks, chunkObjectName)
	for i := 0; int64(i) < ms.ChunkNum; i++ {
		objectName := ms.ObjectName + "@" + strconv.Itoa(i)
		//fileName := os.TempDir() + "/" + ms.Hash + "@" + strconv.Itoa(i)
		name := ms.Hash + "@" + strconv.Itoa(i)
		fileName := "D:/netdisk_tmp/" + name
		logx.Infof("MergeLogic，分片：[%v] 开始合并", fileName)
		if err2 := minioSvc.DownloadChunk2Local(context.TODO(), objectName, fileName); err2 != nil {
			logx.Errorf("MergeLogic，文件%v 下载分片[%v] 失败，ERR：[%v]", ms.Hash, i, err2)
			err = err2
			return
		}

		chunk, err2 := os.Open(fileName)
		fileInfo, err2 := chunk.Stat()
		if err2 != nil {
			logx.Errorf("MergeLogic，文件%v 下载分片[%v]有误，ERR：[%v]", ms.Hash, i, err2)
			err = err2
			return
		}

		chunkObjectName = append(chunkObjectName, objectName)
		chunks = append(chunks, chunk)
		buffer := make([]byte, fileInfo.Size())
		_, err2 = io.CopyBuffer(bigFile, chunk, buffer)
		if err2 != nil {
			logx.Errorf("MergeLogic，文件%v 合并分片[%v]出错，ERR：[%v]", ms.Hash, i, err2)
			err = err2
			return
		}
	}

	CloseTemp(bigFile)
	uploadAndUpdateDb(minioSvc, ms, bigFileName)
}

func uploadAndUpdateDb(minioSvc *minio.Service, ms *MergeStruct, bigFileName string) {
	_, _ = xorm.Xorm.DoTransaction(func(session *xorm.Session) (interface{}, error) {
		fs1 := &models.FileFs{Status: define.StatusFsUploaded}
		if _, err := session.ID(ms.FsId).Update(fs1); err != nil {
			logx.Errorf("MergeLogic，文件%v 更新fs状态出错，ERR：[%v]", ms.Hash, err)
			return nil, err
		}

		file := &models.File{Status: define.StatusFileUploaded}
		if _, err := session.Where("fs_id = ?", ms.FsId).
			Update(file); err != nil {
			logx.Errorf("MergeLogic，文件%v 更新file状态出错，ERR：[%v]", ms.Hash, err)
			return nil, err
		}

		fs2 := &models.FileSchedule{Stage: define.StageMergeDone}
		if _, err := session.ID(ms.SId).Update(fs2); err != nil {
			logx.Errorf("MergeLogic，文件%v 更新fileSchedule状态出错，ERR：[%v]", ms.Hash, err)
			return nil, err
		}

		if err := minioSvc.FUpload(context.TODO(), ms.ObjectName, bigFileName); err != nil {
			logx.Errorf("MergeLogic，上传大文件：[%v]，出错，ERR：[%v]", ms.ObjectName, err)
			return nil, err
		}
		return nil, nil
	})
}

func DeleteTemp(temp *os.File) {
	CloseTemp(temp)
	err := os.Remove(temp.Name())
	if err != nil {
		logx.Errorf("DeleteTemp，删除临时文件 %v 出错，ERR：[%v]", temp.Name(), err)
	}
}

func CloseTemp(temp *os.File) {
	err := temp.Close()
	if err != nil {
		logx.Errorf("DeleteTemp，关闭临时文件 %v 出错，ERR：[%v]", temp.Name(), err)
	}
}

func deleteChunks(minioSvc *minio.Service, chunks []*os.File, chunkObjectName []string) {
	logx.Infof("MergeLogic->deleteChunks，准备删除分片，删除文件中")
	for _, chunk := range chunks {
		name := chunk.Name()
		if err := os.Remove(name); err != nil {
			logx.Errorf("MergeLogic->deleteChunks，删除分片临时文件：[%v]，出错，ERR：[%v]", name, err)
			return
		}
		logx.Infof("MergeLogic->deleteChunks，删除分片：[%v]中", name)
	}

	logx.Infof("MergeLogic->deleteChunks，删除minio分片中")
	for _, objectName := range chunkObjectName {
		_ = minioSvc.DeleteFile(objectName)
		logx.Infof("MergeLogic->deleteChunks，删除minio分片：[%v]中", objectName)
	}
}
