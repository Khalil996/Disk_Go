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

      <template #trigger>
        <el-button type="primary" :icon="Upload" round>选择上传</el-button>
      </template>
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
      <el-button type="primary" round plain :icon="RefreshRight" @click="refresh()"
                 style="margin-left: 20px;position: absolute">刷新
      </el-button>
    </el-upload>
</template>

<script lang="ts" setup>
import {RefreshRight, Upload} from "@element-plus/icons-vue";
import {UploadRequestOptions} from "element-plus";
import {UploadRawFile} from "element-plus/es/components/upload/src/upload";
import {useFileFolderStore} from "../../store/fileFolder.ts";
import SparkMD5 from 'spark-md5'
import {checkChunk, checkFile, upload, uploadChunk, uploadConst} from "./uploading.ts";
import {codeOk, promptError, promptSuccess} from "../../utils/apis/base.ts";
import api from "../../utils/apis/request.ts";

const fileFolderStore = useFileFolderStore(),
    emit = defineEmits(['list'])

async function handleUpload(param: UploadRequestOptions) {
  const res = await checkBeforeUpload(param.file)
  if (res.success) {
    if (res.status === uploadConst.codeNeedUpload) {
      if (
          param.file.size > uploadConst.shardingFloor &&
          res.confirmShard == uploadConst.shardConfirmed
      ) {
        console.log(1111111)
        await uploadSlice(param.file, res.fileId, res.hash)
      } else {
        console.log(22222222)

        await uploadSingle(param.file, res.fileId)
      }
    } else {
      promptSuccess(param.file.name + ' 上传成功！😺')
      emit('list')
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

  for (let i = 0; i < chunkNum; i++) {
    await checkChunkAndUpload(chunks[i], i)
  }
  setTimeout(() => {
    emit('list')
  }, 1000)
}

async function checkChunkAndUpload({chunk, fileId, hash}: any, chunkSeq: number) {
  let resp = await checkChunk({
    fileId: fileId,
    hash: hash,
    chunkSeq: chunkSeq
  })

  if (resp.code === codeOk &&
      resp.data.status === 1) {
    return
  }
  const formData = new FormData();
  formData.append('file', chunk)
  formData.append('fileId', fileId.toString())
  formData.append('chunkSeq', chunkSeq.toString())
  await uploadChunk(formData)
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

function refresh() {
  emit('list')
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