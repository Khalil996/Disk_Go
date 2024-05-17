<template>
  <el-row>
    <el-col :span="24">
      <div class="file-table">
        <template v-if="fileButtonsState === 0">
          <uploading @list="listFiles"/>
        </template>

        <div class="button-group">
          <template v-if="fileButtonsState !== 0">
            <el-button-group>
              <el-button type="primary" round plain :icon="Download" @click="fileButton(0)">下载
              </el-button>
              <template v-if="fileButtonsState === 1">
                <el-button type="primary" round plain :icon="EditPen" @click="fileButton(1)">重命名
                </el-button>
              </template>
              <el-button type="primary" round plain :icon="Rank" @click="fileButton(2)">移动</el-button>
              <el-button type="primary" round plain :icon="CopyDocument" @click="fileButton(3)">复制
              </el-button>
              <el-button type="primary" round plain :icon="Share" @click="fileButton(4)">分享
              </el-button>
              <el-button type="danger" round plain :icon="DeleteFilled" @click="fileButton(5)">删除
              </el-button>
            </el-button-group>
            <el-button type="primary" round plain :icon="RefreshRight" @click="listFiles()"
                       style="margin-left: 20px">刷新
            </el-button>
          </template>
        </div>

        <el-empty v-if="!fileList.data || fileList.data.length==0"
                  description="文件列表为空，上传你的第一个文件吧！😺"/>

        <el-table v-if="fileList && fileList.data.length!=0"
                  ref="fileTableRef"
                  :data="fileList.data" style="width: 100%;max-height: 570px; overflow-y: auto;"
                  @selection-change="fileSelectionChange"
        >
          <el-table-column type="selection" width="55"/>
          <el-table-column label="文件名" min-width="250">
            <template #default="scope">
              <div style="display: flex; align-items: center">
                <el-image v-if="scope.row.type === typeImage && scope.row.status != fileStatus.banned"
                          class="small-pic"
                          :src="scope.row.url"
                          alt="../../assets/alt_type1.jpg"
                          :fit="'cover'"/>
                <el-image v-else
                          :src="`/src/assets/alt_type${scope.row.type}.jpg`"
                          alt=""
                          class="small-pic"
                          :fit="'cover'"/>
                <span style="margin-left: 5px">{{ scope.row.name }}</span>
              </div>
            </template>
          </el-table-column>
          <el-table-column label="修改时间" min-width="150">
            <template #default="scope">
              <div>{{ scope.row.updated }}</div>
            </template>
          </el-table-column>
          <el-table-column label="大小" min-width="100">
            <template #default="scope">
              <div>{{ scope.row.sizeStr }}</div>
            </template>
          </el-table-column>
          <el-table-column label="状态" min-width="100">
            <template #default="scope">
              <div>{{ scope.row.state }}</div>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-col>
  </el-row>

  <el-dialog v-model="fileDialogVisible[1]" title="输入要更改的文件名">
    <el-form :model="renamingFile">
      <el-form-item label="文件名">
        <el-input v-model="renamingFile.name"/>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="fileDialogVisible[1] = false">取消</el-button>
        <el-button type="primary" @click="renameFile(1)">
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>

  <el-dialog v-model="fileCopyAndMoveDialog" title="选择文件夹" width="250">
    <el-table :data="fileMovableFolderList.data" highlight-current-row width="200">
      <el-table-column label="文件夹名" width="200">
        <template #default="scope">
          <div @click="toFolder( scope.row.id)"
               style="display: flex; align-items: center">
            <el-icon>
              <FolderOpened/>
            </el-icon>
            <span style="margin-left: 10px">{{ scope.row.name }}</span>
          </div>
        </template>
      </el-table-column>
    </el-table>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="fileCopyAndMoveDialog = false">取消</el-button>
        <el-button type="primary" @click="fileCopyAndMoveConfirm()">
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>

  <el-dialog v-model="fileDialogVisible[4]" title="分享文件" width="500">

    <el-form label-position="left">
      <el-form-item>
        <div>
          分享 <span style="color: #3c9bff"> {{ selectedFiles[0].name }}</span>
          <template v-if="selectedFiles.length > 1"> 等文件吗？</template>
          <template v-else> 吗？</template>
        </div>
      </el-form-item>

      <el-form-item label="有效期：">
        <div>
          <input type="radio" checked :value="0" v-model="shareInput.radio1"/>1天 &nbsp;&nbsp;
          <input type="radio" :value="1" v-model="shareInput.radio1"/>7天 &nbsp;&nbsp;
          <input type="radio" :value="2" v-model="shareInput.radio1"/>30天 &nbsp;&nbsp;
          <input type="radio" :value="3" v-model="shareInput.radio1"/>长期有效 &nbsp;&nbsp;
        </div>
      </el-form-item>
      <el-form-item label="提取码：">
        <div>
          <input type="radio" :value="0" checked v-model="shareInput.radio2"/>系统生成 &nbsp;&nbsp;
          <input type="radio" :value="1" v-model="shareInput.radio2"/>自己填写 &nbsp;
          <input type="text" placeholder="四位数字或字母" v-model="shareInput.pwd"
                 style="height: 15px; width: 60px; position: relative; top: -2px"/>
        </div>

      </el-form-item>
      <el-form-item>
        <el-checkbox v-model="shareInput.check" label="分享链接自动填充提取码" size="large"/>
      </el-form-item>
    </el-form>

    <template #footer>
      <span class="dialog-footer">
        <el-button @click="fileDialogVisible[4] = false">取消</el-button>
        <el-button type="primary" @click="shareConfirm()">
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>

  <el-dialog v-model="fileDialogVisible[5]" title="删除文件">
    <h3>
      <el-icon>
        <Warning/>
      </el-icon>
      确定要<span style="color: red"> 删除 {{ selectedFiles.map(file => file.name).join('，') }} </span>吗？
      你可以在回收站中找到他们。
    </h3>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="fileDialogVisible[5] = false">取消</el-button>
        <el-button type="primary" @click="deleteFilesConfirm()">
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import {ElTable} from "element-plus";
import {
  CopyDocument, DeleteFilled, Download,
  EditPen, FolderOpened, Rank, Warning,
  Share, RefreshRight
} from "@element-plus/icons-vue";
import {onMounted, reactive, ref} from "vue";
import {
  copyFiles, deleteFiles,
  listFilesByFolderId,
  listFileMovableFolders, moveFiles,
  updateFileName, listFilesByFileType, share
} from "./file.ts";
import type {File} from '/file.ts'
import type {Folder} from "./folder.ts";
import {codeError, codeOk, promptError, promptSuccess, Resp} from "@/utils/apis/base.ts";
import Uploading from "./Uploading.vue";
import {useFileFolderStore} from "@/store/fileFolder.ts";
import {formatSize} from "@/utils/util.ts";
import {expireType, fileStatus, fileStatusMap, typeImage} from "@/utils/constant.ts";

let fileFolderStore = useFileFolderStore(),
    forFolder = false,
    folderId: number,
    fileType: number

const props = defineProps(['fileType', 'folderId']),
    fileTableRef = ref<InstanceType<typeof ElTable>>(),
    fileDialogVisible = reactive([false, false, false, false, false, false]),
    shareInput = reactive({
      radio1: 0,
      radio2: 0,
      pwd: '',
      check: false
    })



let fileButtonsState = ref(0),
    folderList = reactive<{ data: Folder[] }>({
      data: []
    }),
    fileList = reactive<{ data: File[] }>({
      data: []
    }),
    renamingFile = reactive<any>({}),
    listFoldersCurrentFolderId = 0,
    fileCopyAndMoveDialog = ref(false),
    fileCopyAndMoveFlag: number,
    selectedFiles: File[],
    fileMovableFolderList = reactive<{ data: Folder[] }>({data: folderList.data})

const listFiles = async () => {
  let resp
  if (forFolder) {
    resp = await listFilesByFolderId(folderId)
  } else {
    resp = await listFilesByFileType(fileType)
  }
  if (resp.code === 0 && resp.data) {
    fileList.data = resp.data
  }
  fileList.data.forEach(file => {
    file.sizeStr = formatSize(file.size)
    file.state = fileStatusMap[file.status]
  })
}
// 对话框
async function fileButton(option: number) {
  selectedFiles = fileTableRef.value!.getSelectionRows()
  if (!selectedFiles) {
    return
  }
  if (option === 0) {
    await download(selectedFiles)
    return
  } else if (option === 1) {
    Object.assign(renamingFile, selectedFiles[0])
  } else if (option === 2 || option === 3) {
    await toFolder(0)
    fileCopyAndMoveDialog.value = true
    fileCopyAndMoveFlag = option
    return
  }
  fileDialogVisible[option] = true
}

async function renameFile(option: number) {
  const resp = await updateFileName(renamingFile)
  if (resp && resp.code === codeOk) {
    for (const idx in folderList.data) {
      if (fileList.data[idx].id == renamingFile.id) {
        fileList.data[idx].name = renamingFile.name
        break
      }
    }
    promptSuccess()
    fileDialogVisible[option] = false
  }
}

async function toFolder(folderId: number) {
  const resp = await listFileMovableFolders(folderId)
  if (resp && resp.code === codeOk) {
    fileMovableFolderList.data = resp.data
    listFoldersCurrentFolderId = folderId
  }
}

async function fileCopyAndMoveConfirm() {
  const fileIds = selectedFiles.map(file => file.id);
  let resp: Resp<any>
  if (fileCopyAndMoveFlag === 2) {
    resp = await moveFiles(listFoldersCurrentFolderId, fileIds)
  } else {
    resp = await copyFiles(listFoldersCurrentFolderId, fileIds)
  }
  if (resp && resp.code == codeOk) {
    promptSuccess()
    fileCopyAndMoveDialog.value = false
    listFoldersCurrentFolderId = 0
    await listFiles()
  }
}

async function deleteFilesConfirm() {
  const resp = await deleteFiles(selectedFiles.map(file => file.id), folderId)
  if (resp.code === codeOk) {
    await listFiles()
    promptSuccess('删除成功')
    fileDialogVisible[5] = false
    return
  }
  promptError(`删除失败 ${resp.msg}`)
}

function fileSelectionChange(files: File[]) {
  selectedFiles = fileTableRef.value!.getSelectionRows()
  if (!files || files.length == 0) {
    fileButtonsState.value = 0
  } else if (files) {
    if (files.length === 1) {
      fileButtonsState.value = 1
    } else {
      fileButtonsState.value = 2
    }
  }

  fileFolderStore.selectChange(files.map(file => file.id), true)
}

async function download(files: File[]) {
  for (const file of files) {
    if (file.status !== fileStatus.ok) {
      promptError(`文件${fileStatusMap[file.status]}`)
      continue
    }
    await window.open(file.url)
   //  // 创建a标签
   //  setTimeout(() => {
   //    const a = document.createElement('a');
   //    a.href = file.url;
   //    a.download = file.name || 'download';
   //    document.body.appendChild(a);
   //    a.click();
   //    document.body.removeChild(a);
   //  }, 10000);
  }
}

async function shareConfirm() {
  let pwd = (Math.floor(Math.random() * 10000)).toString().padStart(4, '0')
  const regex = /^[a-z0-9]{4}$/i
  if (shareInput.radio2 === 1) {
    if (!regex.test(shareInput.pwd)) {
      promptError('密码只能是数字字母混合四位')
      return
    }
    pwd = shareInput.pwd
  }
  let auto = 0
  if (shareInput.check) {
    auto = 1
  }
  const prefix = `localhost:5173/info/share/`
  const resp = await share(selectedFiles.map(file => file.id), prefix, pwd, shareInput.radio1, auto)
  if (resp.code == codeOk) {
    fileDialogVisible[4] = false
    promptSuccess('分享成功')
    return
  }
  promptError(`分享失败，${resp.msg}`)
}

onMounted(() => {
  if (props.folderId != undefined) {
    folderId = props.folderId
    forFolder = true
  } else {
    fileType = props.fileType
  }

  listFiles()
})

</script>

<style scoped>
.button-group {
  margin-bottom: 15px;
}

.small-pic {
  width: 40px;
  height: 40px;
  border-radius: 5px;
}


</style>