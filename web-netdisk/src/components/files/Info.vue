<template>
  <el-row justify="center">
      <el-col :span="22">
        <div class="search-box">
          <el-row :gutter="10">
            <el-col :span="21">
              <el-input placeholder="搜索我的文件" v-model="searchStr" clearable />
            </el-col>
            <el-col :span="3">
              <el-button @click="searchConfirm"  :icon="Search" circle></el-button>
            </el-col>
          </el-row>
        </div>
          <div v-if="showCase[0]">
              <div style="margin: 3% 0 3% 0">文件详情</div>
              <div>已选择 {{ selectedNum }} 个项目</div>
          </div>

          <div v-if="showCase[1]">
              <div style="margin: 3% 0 3% 0">文件详情</div>

              <el-image v-if="fileDetail.data.type === typeImage"
                        :src="fileDetail.data.url"
                        :fit="'cover'" class="small-img"/>
              <el-image v-else
                        :src="`/src/assets/alt_type${fileDetail.data.type}.jpg`"
                        :fit="'cover'" class="small-img"/>
              <div style="margin-top: 4%; padding-left: 4%;">
                  <div class="file-name">{{ fileDetail.data.name }}</div>
                  <div class="file-info">创建时间：{{ fileDetail.data.created }}</div>
                  <div class="file-info">修改时间：{{ fileDetail.data.updated }}</div>
                  <div class="file-info">文件格式：{{ fileDetail.data.ext }}</div>
                  <div class="file-info">文件大小：{{ fileDetail.data.sizeStr }}</div>
                  <div class="file-info">文件类型：{{ typeMap[fileDetail.data.type] }}</div>
              </div>
          </div>

          <div v-if="showCase[2]">
              <div style="margin: 3% 0 3% 0">文件夹内容</div>

          </div>

          <div v-if="showCase[3]">
              <div style="margin: 3% 0 3% 0">搜索结果</div>

              <el-table v-if="files.data && files.data.length!=0"
                        ref="fileTableRef"
                        :data="files.data" style="width: 100%"
              >
                  <el-table-column width="50">
                      <template #default="scope">
                          <el-image :src="`/src/assets/alt_type${scope.row.type}.jpg`"
                                    class="tiny-img"
                                    :fit="'cover'"/>
                      </template>
                  </el-table-column>

                  <el-table-column prop="name"/>
                  <el-table-column>
                      <template #default="scope">
                          <el-button type="primary" plain size="small"
                                     @click="toFolder(scope.row.folderId)">位置
                          </el-button>
                          <el-button type="primary" plain
                                     size="small">下载
                          </el-button>
                      </template>
                  </el-table-column>
              </el-table>
          </div>
      </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import {Search} from '@element-plus/icons-vue';
import {onUnmounted, reactive, ref} from "vue";
import {useFileFolderStore} from "@/store/fileFolder.ts";
import {getFileDetailById, listFolderDetailById, search} from "./Info.ts";
import {codeOk} from "@/utils/apis/base.ts";
import {formatSize} from "@/utils/util.ts";
import {typeImage, typeMap} from "@/utils/constant.ts";
import {ElTable} from "element-plus";

const fileFolderStore = useFileFolderStore()

let showCase = ref([true, false, false, false])
const fileDetail: { data: {} } = reactive({data: {}}),
  files = reactive({data: []})

let searchStr = ref('')
let selectedNum = ref(0)

async function searchConfirm() {
  const resp = await search(searchStr.value)
  if (resp && resp.code === codeOk) {
      files.data = resp.data
      showCase.value = [false, false, false, true]

  }
}

async function getFileDetail(fileId: number) {
  const resp = await getFileDetailById(fileId)
  if (resp.code === codeOk) {
      fileDetail.data = resp.data
      fileDetail.data.sizeStr = formatSize(resp.data.size)
  }
}

async function getFolderDetail(fileId: number) {
  const resp = await listFolderDetailById(fileId)
  if (resp.code === codeOk) {
      fileDetail.data = resp.data
  }
}

const unsubscribe = fileFolderStore.$subscribe((_, state) => {
  const filesLen = state.selectedItems.files.length
  const foldersLen = state.selectedItems.folders.length;
  if (filesLen === 1 && foldersLen < 1) {
      showCase.value = [false, true, false]
      getFileDetail(state.selectedItems.files[0])
  } else if (foldersLen == 1 && filesLen < 1) {
      showCase.value = [false, false, true]
      getFolderDetail(state.selectedItems.folders[0])
  } else {
      selectedNum.value = filesLen + foldersLen
      showCase.value = [true, false, false]
  }
})

function toFolder(folderId: number) {
  window.location.href = `/file/folder/${folderId}`
}

onUnmounted(() => {
  unsubscribe()
})

</script>

<style scoped>
.search-box {
}

.file-name {
  color: #454d5a;
  font-size: 1.5rem;
  font-weight: 600;
}

.file-info {
  margin-top: 3%;
  font-size: 1.2rem;
  color: #878c9c;
}

.small-img {
  border-radius: 5px;
  width: 100%;
  max-height: 600px;
}

.tiny-img {
  border-radius: 5px;
  width: 100%;
  max-height: 50px;
}
</style>