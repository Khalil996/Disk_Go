<template>
    <el-row>
        <el-col :span="24">
            <div class="file-table">
                <template v-if="fileButtonsState === 0">
                    <uploading></uploading>
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
                            <el-button type="danger" round plain :icon="DeleteFilled" @click="fileButton(4)">删除
                            </el-button>
                        </el-button-group>
                    </template>
                </div>

                <el-empty v-if="!fileList.data || fileList.data.length==0"
                          description="文件列表为空，上传你的第一个文件吧！😺"/>

                <el-table v-if="fileList && fileList.data.length!=0"
                          ref="fileTableRef"
                          :data="fileList.data" style="width: 100%;"
                          @selection-change="fileSelectionChange"
                >
                    <el-table-column type="selection" width="55"/>
                    <el-table-column label="文件名" min-width="200">
                        <template #default="scope">
                            <div style="display: flex; align-items: center">
                                <el-image v-if="scope.row.type === typeImage"
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
                    <el-table-column label="修改时间" min-width="100">
                        <template #default="scope">
                            <div>{{ scope.row.updated }}</div>
                        </template>
                    </el-table-column>
                    <el-table-column label="大小" min-width="100">
                        <template #default="scope">
                            <div>{{ scope.row.sizeStr }}</div>
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

    <el-dialog v-model="fileCopyAndMoveDialog" title="选择文件夹">
        <el-table :data="fileMovableFolderList.data" highlight-current-row>
            <el-table-column label="" width="180">
                <template #default="scope">
                    <div style="display: flex; align-items: center">
                        <div @click="toFolder( scope.row.id)">
                            <el-icon>
                                <FolderOpened/>
                            </el-icon>
                            <span style="margin-left: 10px">{{ scope.row.name }}</span>
                        </div>
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

    <el-dialog v-model="fileDialogVisible[4]" title="删除文件">
        <h3>
            <el-icon>
                <Warning/>
            </el-icon>
            确定要<span style="color: red"> 删除 {{ selectedFiles.map(file => file.name).join('，') }} </span>吗？
            你可以在回收站中找到他们。
        </h3>
        <template #footer>
      <span class="dialog-footer">
        <el-button @click="fileDialogVisible[4] = false">取消</el-button>
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
    EditPen, FolderOpened, Rank, Warning
} from "@element-plus/icons-vue";
import {onMounted, reactive, ref} from "vue";
import {
    copyFiles, deleteFiles,
    listFilesByFolderId,
    listFileMovableFolders, moveFiles,
    updateFileName, listFilesByFileType
} from "./file.ts";
import type {File} from './file.ts'
import type {Folder} from "./folder.ts";
import {codeOk, promptSuccess, Resp} from "../../utils/apis/base.ts";
import Uploading from "./Uploading.vue";
import {useFileFolderStore} from "../../store/fileFolder.ts";
import {formatSize} from "../../utils/util.ts";
import {typeImage} from "../../utils/constant.ts";

let fileFolderStore = useFileFolderStore()
let forFolder = false
let folderId: number
let fileType: number
const props = defineProps(['fileType', 'folderId'])

let fileButtonsState = ref(0)
const fileTableRef = ref<InstanceType<typeof ElTable>>()

let folderList = reactive<{ data: Folder[] }>({
    data: []
})
let fileList = reactive<{ data: File[] }>({
    data: []
})

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
    })
}

const fileDialogVisible = reactive([false, false, false, false, false])
let renamingFile = reactive<any>({})
let listFoldersCurrentFolderId = 0
let fileCopyAndMoveDialog = ref(false)
let fileCopyAndMoveFlag: number
let selectedFiles: File[]
let fileMovableFolderList = reactive<{ data: Folder[] }>({data: folderList.data})

// 对话框
function fileButton(option: number) {
    selectedFiles = fileTableRef.value!.getSelectionRows()
    if (!selectedFiles) {
        return
    }
  if (option === 0) {

  } else if (option === 1) {
        Object.assign(renamingFile, selectedFiles[0])
    } else if (option === 2 || option === 3) {
        toFolder(0)
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
    }
}

async function deleteFilesConfirm() {
    await deleteFiles(selectedFiles.map(file => file.id), folderId)
    await listFiles()
}

function fileSelectionChange(files: File[]) {
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
    width: 35px;
    height: 35px;
    border-radius: 5px;
}

</style>