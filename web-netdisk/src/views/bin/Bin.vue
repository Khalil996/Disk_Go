<template>
  <el-row>
    <el-col :span="24">
      <div class="file-table">
        <el-button v-if="fileList && fileList.data.length!=0"
                   type="primary" round :icon="Download"
                   @click="dialogButtons(0)">清空回收站
        </el-button>

        <el-empty v-if="!fileList.data || fileList.data.length==0"
                  description="回收站暂时为空😚"/>

        <el-table v-if="fileList && fileList.data.length!=0"
                  ref="fileTableRef"
                  :data="fileList.data" style="width: 100%"
                  @selection-change="fileSelectionChange"
        >
          <el-table-column type="selection" width="55"/>
          <el-table-column label="文件名" width="180">
            <template #default="scope">
              <div class="file-folder-row" style="display: flex; align-items: center">
                <span>{{ scope.row.name }}</span>
              </div>
            </template>
          </el-table-column>
          <el-table-column label="修改时间" width="180">
            <template #default="scope">
              <div>{{ scope.row.updated }}</div>
            </template>
          </el-table-column>
          <el-table-column label="大小" width="180">
            <template #default="scope">
              <template>
                <div>{{ scope.row.size }}</div>
              </template>
            </template>
          </el-table-column>
          <el-table-column label="Operations">
            <template #default="scope">
              <el-button size="small" @click="handleEdit(scope.$index, scope.row)"
              >Edit
              </el-button>
              <el-button
                  size="small"
                  type="danger"
                  @click="handleDelete(scope.$index, scope.row)"
              >Delete
              </el-button>
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
</template>

<script lang="ts" setup>
import {ElTable} from "element-plus";
import {onMounted, reactive} from "vue";
import {deleteAllFilesTruly, DeleteFile, deleteFilesTruly, getDeletedFiles} from "./bin.ts";
import {Download, Warning} from "@element-plus/icons-vue";
import {codeOk, promptError} from "../../utils/apis/base.ts";

const fileList: { data: DeleteFile[] } = reactive({
  data: []
})

const dialogVisible = reactive({option: [false]})

function dialogButtons(option: number) {
  dialogVisible.option = [false]
  dialogVisible.option[option] = true
}

async function listDeletedFiles() {
  const resp = await getDeletedFiles()
  if (resp.code === codeOk) {
    fileList.data = resp.data
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

onMounted(() => {
  listDeletedFiles()
})

</script>

<style scoped>

</style>