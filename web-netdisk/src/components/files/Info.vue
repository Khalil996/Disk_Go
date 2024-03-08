<template>
    <el-row justify="center">
        <el-col :span="22">
            <div class="search-box">
                <el-input placeholder="搜索我的文件"
                          v-model="searchStr" clearable>
                    <template  >
                        <el-button @click="searchConfirm">
                            <el-icon>
                                <search/>
                            </el-icon>
                        </el-button>
                    </template>
                </el-input>
            </div>

            <div v-if="showCase[0]">
                <div style="margin: 3% 0 3% 0">文件详情</div>
                <div>已选择 {{ selectedNum }} 个项目</div>
            </div>

            <div v-if="showCase[1]">
                <div style="margin: 3% 0 3% 0">文件详情</div>
                <el-image style="border-radius: 5px; width: 100%; max-height: 600px;"
                          :src="fileDetail.data.url"
                          :fit="'cover'"/>
                <div style="margin-top: 4%; padding-left: 4%;">
                    <div class="file-name">{{ fileDetail.data.name }}</div>
                    <div class="file-info">创建时间：{{ fileDetail.data.created }}</div>
                    <div class="file-info">修改时间：{{ fileDetail.data.updated }}</div>
                    <div class="file-info">文件格式：{{ fileDetail.data.ext }}</div>
                    <div class="file-info">文件大小：{{ fileDetail.data.sizeStr }}</div>
                </div>
            </div>

            <div v-if="showCase[2]">
                <div style="margin: 3% 0 3% 0">文件夹内容</div>

            </div>
        </el-col>
    </el-row>
</template>

<script lang="ts" setup>
import {Search} from '@element-plus/icons-vue';
import {onUnmounted, reactive, ref} from "vue";
import {useFileFolderStore} from "../../store/fileFolder.ts";
import {getFileDetailById, listFolderDetailById, search} from "./Info.ts";
import {codeOk} from "../../utils/apis/base.ts";
import {formatSize} from "../../utils/util.ts";

const fileFolderStore = useFileFolderStore()

let showCase = ref([true, false, false])
const url = 'https://fuss10.elemecdn.com/e/5d/4a731a90594a4af544c0c25941171jpeg.jpeg'
const fileDetail: { data: {} } = reactive({
    data: {}
})

let searchStr = ref('')
let selectedNum = ref(0)

async function searchConfirm() {
    const resp = await search(searchStr.value)
    if (resp && resp.code === codeOk) {
// TODO
    }
}

async function getFileDetail(fileId: number) {
    const resp = await getFileDetailById(fileId)
    if (resp.code === codeOk) {
        fileDetail.data = resp.data
        fileDetail.data.sizeStr = formatSize(resp.data.size)
        console.log(fileDetail.data, formatSize(resp.data.size))
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
</style>