<template>
    <el-upload ref="uploadFiles"
               class="upload-demo"
               action="actionUrl"
               multiple
               :limit="20"
               :auto-upload="true"
               :http-request="handleUpload"
    >
        <!--               :on-change="change"-->

        <el-button type="primary" :icon="Upload" round>选择上传</el-button>
        <!--        <div v-if="!progress" class="el-upload__text">-->
        <!--            Drop file here or <em>click to upload</em>-->
        <!--        </div>-->

        <!--        <el-progress-->

        <!--                v-else-->
        <!--                :text-inside="true"-->
        <!--                :stroke-width="24"-->
        <!--                :percentage="progress"-->
        <!--                status="success"-->
        <!--        />-->

    </el-upload>
</template>

<script lang="ts" setup>
import {Upload} from "@element-plus/icons-vue";
import {UploadRequestOptions} from "element-plus";
import {UploadRawFile} from "element-plus/es/components/upload/src/upload";
import {useFileFolderStore} from "../../store/fileFolder.ts";
import SparkMD5 from 'spark-md5'
import {checkChunk, checkFile, upload, uploadChunk, uploadConst} from "./uploading.ts";
import {codeOk, promptError, promptSuccess} from "../../utils/apis/base.ts";
import api from "../../utils/apis/request.ts";

const fileFolderStore = useFileFolderStore()

async function handleUpload(param: UploadRequestOptions) {
    const res = await checkBeforeUpload(param.file)
    if (res.success) {
        if (res.status === uploadConst.codeNeedUpload) {
            if (
                param.file.size > uploadConst.shardingFloor &&
                res.confirmShard == uploadConst.shardConfirmed
            ) {
                await uploadSlice(param.file, res.fileId, res.hash)
            } else {
                await uploadSingle(param.file, res.fileId)
            }
        } else {
            promptSuccess(param.file.name + ' 上传成功！😺')
        }
    } else {
        promptError('请检查文件是否合法！')
    }
}

async function checkBeforeUpload(file: UploadRawFile) {
    const md5 = await genMd5(file);
    const resp = await checkFile({
        folderId: fileFolderStore.folderId,
        name: file.name,
        size: file.size,
        ext: file.name.substring(file.name.lastIndexOf('.')),
        hash: md5
    })

    const res = {
        success: true,
        fileId: resp.data.fileId,
        status: resp.data.status,
        confirmShard: resp.data.confirmShard,
        hash: md5
    }
    if (resp && resp.code === codeOk) {
        return res
    }
    res.success = false
    return res
}

async function uploadSingle(file: UploadRawFile, fileId: number) {
    const formData = new FormData();
    formData.append('file', file)
    formData.append('fileId', fileId.toString())
    const resp = await upload(formData)
    if (resp && resp.code === codeOk) {
        promptSuccess(file.name + ' 上传成功！😺')
    }
}

async function uploadSlice(file: UploadRawFile, fileId: number, hash: string) {
    const chunkNum = Math.ceil(file.size / uploadConst.shardingSize)
    let start = 0
    let end = 0
    const chunks = []
    while (start < file.size) {
        end = Math.min(start + uploadConst.shardingSize, file.size)
        chunks.push({chunk: file.slice(start, end), fileId, hash})
        start = end
    }
    if (chunks.length != chunkNum) {
        promptError('上传' + file.name + '过程出错！😿')
    }
    // await Promise.all(chunks.map(checkChunkAndUpload))

    for (let i = 0; i < chunkNum; i++) {
        await checkChunkAndUpload(chunks[i], i)
    }
}

async function checkChunkAndUpload({chunk, fileId, hash}: any, chunkSeq: number) {
    let resp = await checkChunk({
        fileId: fileId,
        hash: hash,
        chunkSeq: chunkSeq
    })
    if (resp.code === codeOk && resp.data.status === 1) {
        return
    }
    const formData = new FormData();
    formData.append('file', chunk)
    formData.append('fileId', fileId.toString())
    formData.append('chunkSeq', chunkSeq.toString())
    await uploadChunk(formData)
    // TODO
}

async function genMd5(file: UploadRawFile) {
    const spark = new SparkMD5.ArrayBuffer()
    try {
        spark.append(await file.arrayBuffer())
        return spark.end()
    } catch (e) {
        promptError(`上传文件 ${file.name} 失败！😨，${e}`)
    }
}

function formatSize(file) {
    //console.log("size",file.size);
    let size = file.size;
    let unit;
    const units = [" B", " K", " M", " G"];
    let pointLength = 2;
    while ((unit = units.shift()) && size > 1024) {
        size = size / 1024;
    }
    return (
        (unit === "B"
            ? size
            : size.toFixed(pointLength === undefined ? 2 : pointLength)) + unit
    );
}

// function asd(e: Event) {
//     const target = e.target
//     if (target instanceof HTMLInputElement) {
//         const file = target.files
//         if (file) {
//             const form = new FormData()
//             for (let i = 1; i < file.length; i++) {
//                 form.append("file", file[i])
//             }
//             axios.post("/", form, {
//                 onUploadProgress: (progressEvent: AxiosProgressEvent) => {
//                     Math.round((progressEvent.loaded / (progressEvent.total as number) * 100))
//                 }
//             })
//         }
//     }
// }

// const uploadProcedure = (options: UploadRequestOptions) => {
//     console.log(options.files)
//     options.files
//     return XMLHttpRequest
// }
</script>

<style scoped>

</style>