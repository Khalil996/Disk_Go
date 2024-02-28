<template>
  <el-row>
    <el-col :span="14">
      <file :folderId="folderId"/>
    </el-col>

    <!--  folder  -->
    <el-col :span="10">
      <div class="folder-table">
        <el-upload v-if="folderButtonsState === 0">
          <el-button type="primary" :icon="FolderAdd" round @click="folderButton(0)">新建文件夹</el-button>
          <template #trigger v-show="false"></template>
        </el-upload>
        <div class="button-group">
          <template v-if="folderButtonsState !== 0">
            <el-button-group>
              <el-button type="primary" round plain :icon="Download" @click="folderButton(1)">下载
              </el-button>
              <template v-if="folderButtonsState === 1">
                <el-button type="primary" round plain :icon="EditPen" @click="folderButton(2)">重命名
                </el-button>
              </template>
              <el-button type="primary" round plain :icon="Rank" @click="folderButton(3)">移动</el-button>
              <el-button type="primary" round plain :icon="CopyDocument" @click="folderButton(4)">复制
              </el-button>
              <el-button type="danger" round plain :icon="DeleteFilled" @click="folderButton(5)">删除
              </el-button>
            </el-button-group>
          </template>
        </div>

        <el-empty v-if="!folderList.arr || folderList.arr.length==0"
                  description="文件夹列表为空，创建你的第一个文件夹吧！😺"/>

        <el-table v-if="folderList.arr && folderList.arr.length!=0"
                  ref="folderTableRef"
                  :data="folderList.arr" style="width: 100%"
                  @selection-change="folderSelectionChange"
        >
          <el-table-column type="selection" width="55"/>
          <el-table-column label="文件夹名" width="180">
            <template #default="scope">
              <div class="file-folder-row" @click="router.push(`/file/folder/${scope.row.id}`)">
                <el-icon>
                  <FolderOpened/>
                </el-icon>
                <span style="margin-left: 10px">{{ scope.row.name }}</span>
              </div>
            </template>
          </el-table-column>
          <el-table-column label="修改时间" width="180">
            <template #default="scope">
              <div>{{ scope.row.updated }}</div>
            </template>
          </el-table-column>
          <el-table-column label="Operations">
            <template #default="scope">
              <el-button size="small" @click=""
              >Edit
              </el-button>
              <el-button
                  size="small"
                  type="danger"
                  @click=""
              >Delete
              </el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-col>
  </el-row>

  <!--  folder  -->
  <el-dialog v-model="folderDialogVisible[0]" title="创建文件夹">
    <el-form :model="createFolderName" label-width="120px">
      <el-form-item label="文件夹名">
        <el-input v-model="createFolderName"/>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="folderDialogVisible[0] = false">取消</el-button>
        <el-button type="primary" @click="createFolderConfirm(0)">
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>

  <el-dialog v-model="folderDialogVisible[2]" title="输入要更改的文件夹名">
    <el-form :model="selectedFolders[0]">
      <el-form-item label="文件名">
        <el-input v-model="renamingFolder.name"/>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="folderDialogVisible[2] = false">取消</el-button>
        <el-button type="primary" @click="renameFolder(2)">
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>

  <el-dialog v-model="folderCopyAndMoveDialog" title="选择文件夹">
    <el-table :data="folderMovableFolderList.arr" highlight-current-row>
      <el-table-column label="" width="180">
        <template #default="scope">
          <div style="display: flex; align-items: center">
            <div @click="toFolder( scope.row.id, folderCopyAndMoveFlag)">
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
        <el-button @click="folderCopyAndMoveDialog = false">取消</el-button>
        <el-button type="primary" @click="folderCopyAndMoveConfirm()">
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>

  <el-dialog v-model="folderDialogVisible[5]" title="删除文件夹">
    <h3>
      <el-icon>
        <Warning/>
      </el-icon>
      确定要<span style="color: red"> 删除 {{ selectedFolders.map(folder => folder.name).join('，') }} </span>吗？
      你可以在回收站中找到文件。
    </h3>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="folderDialogVisible[5] = false">取消</el-button>
        <el-button type="primary" @click="deleteFoldersConfirm(5)">
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import {onMounted, reactive, ref} from 'vue'
import {ElTable} from 'element-plus'
import type {Folder} from './folder.ts'
import {
  FolderOpened, FolderAdd, Download, CopyDocument,
  EditPen, DeleteFilled, Rank, Warning
} from '@element-plus/icons-vue'
import {
  updateFolderName,
  createFolder,
  listFolderMovableFolders,
  moveFolders,
  copyFolders,
  deleteFolders, listFoldersByParentFolderId
} from "./folder.ts";
import File from './File.vue'
import router from "../../router";
import {codeOk, promptSuccess, Resp} from "../../utils/apis/base.ts";
import {useFileFolderStore} from "../../store/fileFolder.ts";

let fileFolderStore = useFileFolderStore()
let props = defineProps(['folderId']);
let folderId = Number.parseInt(props.folderId, 10)
let folderButtonsState = ref(0)
const folderTableRef = ref<InstanceType<typeof ElTable>>()

let folderList = reactive<{ arr: Folder[] }>({
  arr: []
})

const listFolders = async () => {
  const resp = await listFoldersByParentFolderId(folderId)
  if (resp.code === 0 && resp.data) {
    folderList.arr = resp.data
  }
}

let listFoldersCurrentFolderId = 0

const folderDialogVisible = reactive([false, false, false, false, false, false])
let createFolderName = ref<string>('')
let renamingFolder = reactive<any>({})
let folderCopyAndMoveDialog = ref(false)
let folderCopyAndMoveFlag: number
let selectedFolders: Folder[]
let folderMovableFolderList = reactive<{ arr: Folder[] }>({arr: folderList.arr})

async function toFolder(folderId: number, option: number) {
  let resp: Resp<any>
  if (option === 3) {
    selectedFolders = folderTableRef.value!.getSelectionRows()
    resp = await listFolderMovableFolders(folderId, selectedFolders.map(folder => folder.id))
  } else if (option === 4) {
    resp = await listFoldersByParentFolderId(folderId)
  }
  if (resp && resp.code === 0) {
    folderMovableFolderList.arr = resp.data
  }
}

// 对话框
function folderButton(option: number) {
  if (option === 0) {
    folderDialogVisible[option] = true
    return
  }
  selectedFolders = folderTableRef.value!.getSelectionRows()
  if (!selectedFolders) {
    return
  }
  if (option === 2) {
    Object.assign(renamingFolder, selectedFolders[0])
  } else if (option === 3 || option === 4) {
    toFolder(0, option)
    folderCopyAndMoveDialog.value = true
    folderCopyAndMoveFlag = option
    return
  }
  folderDialogVisible[option] = true
}

// 创建请求
async function createFolderConfirm(option: number) {
  let name
  if (createFolderName.value) {
    name = createFolderName.value
  }
  await createFolder(folderId, name)
  await listFolders()
  promptSuccess()
  folderDialogVisible[option] = false
}

// 重命名请求
async function renameFolder(option: number) {
  const resp = await updateFolderName(renamingFolder)
  if (resp && resp.code === codeOk) {
    for (const idx in folderList.arr) {
      if (folderList.arr[idx].id == renamingFolder.id) {
        folderList.arr[idx].name = renamingFolder.name
        break
      }
    }
    promptSuccess()
    folderDialogVisible[option] = false
  }
}

// 复制/移动请求
async function folderCopyAndMoveConfirm() {
  const folderIds = selectedFolders.map(folder => folder.id);
  if (folderCopyAndMoveFlag === 3) {
    await moveFolders(listFoldersCurrentFolderId, folderIds)
  } else if (folderCopyAndMoveFlag === 4) {
    await copyFolders(listFoldersCurrentFolderId, folderIds)
  }
  listFoldersCurrentFolderId = 0
}

// 删除请求
async function deleteFoldersConfirm(option: number) {
  const ids = folderTableRef.value!.getSelectionRows().map(folder => folder.id)
  await deleteFolders(ids)
  const idMap = new Map()
  ids.forEach(id => idMap.set(id, true))
  folderList.arr = folderList.arr.filter(folder => {
    if (idMap.get(folder.id) == undefined) {
      return folder
    }
  })
  folderDialogVisible[option] = false
}

function folderSelectionChange(folders: Folder[]) {
  if (!folders || folders.length == 0) {
    folderButtonsState.value = 0
  } else if (folders) {
    if (folders.length === 1) {
      folderButtonsState.value = 1
    } else {
      folderButtonsState.value = 2
    }
  }
  fileFolderStore.selectChange(folders.map(folder => folder.id), false)
}

onMounted(() => {
  fileFolderStore.setFolderId(Number.parseInt(props.folderId))
  listFolders()
})

</script>


<style scoped>
.button-group {
  margin-bottom: 15px;
}

.file-folder-row {
  display: flex;
  align-items: center;
}

.file-folder-row:hover {
  cursor: pointer;
  background-color: rgb(230, 230, 245);
  border-radius: 5px;
}
</style>