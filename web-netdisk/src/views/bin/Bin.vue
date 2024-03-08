<template>
  <el-row>
    <el-col :span="24">
      <div class="file-table">
        <el-upload v-if="fileList && fileList.data.length!=0">
          <el-button v-if="fileButtonsState === 0"
                     type="primary" :icon="Download" round
                     @click="dialogButtons(0)">清空回收站
          </el-button>

          <el-button-group v-if="fileButtonsState === 1">
            <el-button type="warning" round plain
                       :icon="RefreshRight" @click="dialogButtons(1)">恢复文件
            </el-button>
            <el-button type="danger" round plain
                       :icon="Delete" @click="dialogButtons(2)">删除文件
            </el-button>
          </el-button-group>
          <template #trigger v-show="false"></template>
        </el-upload>

        <el-empty v-if="!fileList.data || fileList.data.length==0"
                  description="回收站暂时为空😚"/>

        <el-table v-if="fileList && fileList.data.length!=0"
                  ref="fileTableRef"
                  :data="fileList.data" style="width: 100%"
                  @selection-change="fileSelectionChange"
        >
          <el-table-column type="selection" width="55"/>
          <el-table-column label="文件名" min-width="400">
            <template #default="scope">
              <div class="file-folder-row" style="display: flex; align-items: center">
                <span>{{ scope.row.name }}</span>
              </div>
            </template>
          </el-table-column>

          <el-table-column label="大小" min-width="100">
            <template #default="scope">
              <div>{{ scope.row.sizeStr }}</div>
            </template>
          </el-table-column>
          <el-table-column label="删除时间" min-width="100">
            <template #default="scope">
              <div>{{ scope.row.delTimeStr }}</div>
            </template>
          </el-table-column>
          <el-table-column label="剩余时间" min-width="100">
            <template #default="scope">
              <div>{{ scope.row.left }}</div>
            </template>
          </el-table-column>
          <el-table-column label="来源文件夹" min-width="100">
            <template #default="scope">
              <div>{{ scope.row.folderName }}</div>
            </template>
          </el-table-column>
          <el-table-column min-width="100">
            <template #default="scope">
                            <span @click="dialogButtons(1); singleSelectedFile=scope.row.id">
                                <el-icon>
                                    <RefreshRight/>
                                </el-icon>
                            </span>
              &nbsp;&nbsp;
              <span @click="dialogButtons(2); singleSelectedFile=scope.row.id">
                                <el-icon>
                                    <Delete/>
                                </el-icon>
                            </span>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-col>
  </el-row>

  <el-dialog v-model="dialogVisible.option[0]" title="清空回收站">
    <h3>
      <el-icon>
        <Warning/>
      </el-icon>
      确定清空你的回收站吗，删除的文件不可再被找回！😶
    </h3>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible.option[0] = false">取消</el-button>
        <el-button type="primary" @click="clearBin">
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>

  <el-dialog v-model="dialogVisible.option[1]" title="恢复文件">
    <h3>
      <el-icon>
        <Warning/>
      </el-icon>
      确定恢复这个文件吗？😺
    </h3>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible.option[1] = false; singleSelectedFile=0">取消</el-button>
        <el-button type="primary" @click="recoverOrDelete(1)">
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>

  <el-dialog v-model="dialogVisible.option[2]" title="删除文件">
    <h3>
      <el-icon>
        <Warning/>
      </el-icon>
      确定删除该文件吗，删除后不可再被找回！😶
    </h3>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible.option[2] = false; singleSelectedFile=0">取消</el-button>
        <el-button type="primary" @click="recoverOrDelete(2)">
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import {ElTable} from "element-plus";
import {onMounted, reactive, ref} from "vue";
import {deleteAllFilesTruly, DeleteFile, deleteFilesTruly, getDeletedFiles, recoverFiles} from "./bin.ts";
import {
  Download, Warning,
  RefreshRight, Delete
} from "@element-plus/icons-vue";
import {codeOk, promptError} from "@/utils/apis/base.ts";
import {formatLeft, formatSize, formatTime} from "@/utils/util.ts";
import {File} from "@/components/files/file.ts";

const fileList = reactive<{ data: DeleteFile[] }>({
      data: []
    }),
    dialogVisible = reactive({option: [false, false, false]}),
    fileTableRef = ref<InstanceType<typeof ElTable>>()

let singleSelectedFile = ref(0),
    selectedFiles: File[],
    fileButtonsState = ref(0)

function dialogButtons(option: number) {
  selectedFiles = fileTableRef.value!.getSelectionRows()
  console.log(selectedFiles, typeof selectedFiles)
  dialogVisible.option = [false]
  dialogVisible.option[option] = true
}

function fileSelectionChange(files: File[]) {
  fileButtonsState.value = 1
  if (!files || files.length == 0) {
    fileButtonsState.value = 0
  }
}

async function listDeletedFiles() {
  const resp = await getDeletedFiles()
  if (resp.code === codeOk) {
    fileList.data = resp.data
    fileList.data.forEach(file => {
      file.sizeStr = formatSize(file.size)
      file.delTimeStr = formatTime(file.delTime)
      file.left = formatLeft(file.delTime + 30 * 24 * 60 * 60)
      console.log(file)
    })
  } else {
    promptError(resp.msg)
  }
}

async function clearBin() {
  const resp = await deleteAllFilesTruly()
  if (resp.code === codeOk) {
    dialogVisible.option[0] = false
    fileList.data = []
  } else {
    promptError(resp.msg)
  }
}

async function recoverOrDelete(option: number) {
  let ids, resp
  if (singleSelectedFile.value !== 0) {
    ids = [singleSelectedFile.value]
  } else {
    ids = selectedFiles.map(file => file.id)
  }
  if (option === 1) {
    resp = await recoverFiles(ids)
  } else {
    resp = await deleteFilesTruly(ids)
  }
  if (resp.code === codeOk) {
    dialogVisible.option[option] = false
    await listDeletedFiles()
  } else {
    promptError(resp.msg)
  }
}

onMounted(() => {
  listDeletedFiles()
})

</script>

<style scoped>

</style>
