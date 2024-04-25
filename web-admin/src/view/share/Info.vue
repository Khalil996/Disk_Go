<template>
    <el-card>
        <el-table
                ref="fileTableRef" border
                :data="list.items" style="width: 100%"
        >
            <el-table-column type="selection" width="55"/>
            <el-table-column label="ID" min-width="150">
                <template #default="scope">
                    <div style="display: flex; align-items: center">
                        <el-image :src="`/src/assets/img/alt_type${scope.row.type}.jpg`"
                                  class="small-pic"
                                  :fit="'cover'"/>
                        <span style="margin-left: 5px">{{ scope.row.id }}</span>
                    </div>
                </template>
            </el-table-column>
            <el-table-column prop="name" label="文件名" min-width="400"/>
            <el-table-column label="修改时间" min-width="200">
                <template #default="scope">
                    <div>{{ scope.row.updated }}</div>
                </template>
            </el-table-column>
            <el-table-column label="大小" min-width="200">
                <template #default="scope">
                    <div>{{ scope.row.sizeStr }}</div>
                </template>
            </el-table-column>
            <el-table-column label="类型" min-width="100">
                <template #default="scope">
                    <div>{{ typeMap[scope.row.type] }}</div>
                </template>
            </el-table-column>
            <el-table-column prop="state" label="状态" min-width="100"/>
            <el-table-column label="操作" min-width="250">
                <template #default="scope">
                    <el-button type="primary" plain size="small" @click="getUrl(scope.row.id)">下载查看</el-button>
                    <el-button type="primary" plain size="small"
                               @click="buttonClick(0, scope.row.id, fileStatus.ok)">恢复
                    </el-button>
                    <template v-if="scope.row.status !== fileStatus.banned">
                        <el-button type="danger" size="small"
                                   @click="buttonClick(0, scope.row.id, fileStatus.banned)">封禁
                        </el-button>
                    </template>
                </template>
            </el-table-column>
        </el-table>
    </el-card>

    <el-dialog v-model="dialogVisible.option[0]" title="恢复文件">
        <h3>
            <el-icon>
                <Warning/>
            </el-icon>
            需要恢复这个文件吗？文件将会在所有分享中生效😶
        </h3>
        <template #footer>
              <span class="dialog-footer">
                <el-button @click="dialogVisible.option[0]=false">取消</el-button>
                <el-button type="primary" @click="setStatus(0)">
                  确定
                </el-button>
              </span>
        </template>
    </el-dialog>

    <el-dialog v-model="dialogVisible.option[1]" title="封禁文件">
        <h3>
            <el-icon>
                <Warning/>
            </el-icon>
            确定封禁这个文件吗？文件将会在所有分享中被封禁😶
        </h3>
        <template #footer>
              <span class="dialog-footer">
                <el-button @click="dialogVisible.option[1]=false">取消</el-button>
                <el-button type="primary" @click="setStatus(1)">
                  确定
                </el-button>
              </span>
        </template>
    </el-dialog>
</template>

<script lang="js" setup>
import {onBeforeMount, reactive} from "vue";
import {formatSize} from "@/utils/util.js";
import fileApi from "@/api/file.js";
import shareApi from "@/api/share.js";
import {fileStatus, fileStatusMap, typeMap} from "@/utils/constant.js";
import {codeOk, promptError, promptSuccess} from "@/utils/http/base.js";

const props = defineProps(["shareId"]),
    list = reactive({
        name: '',
        created: '',
        expired: 0,
        owner: 0,
        items: []
    }),
    dialogVisible = reactive({
        option: [false, false]
    }),
    setStatusObj = {
        ids: 0,
        status: 0
    }
let urlMap = new Map()

const listItems = async () => {
    let resp = await shareApi.getShareFilesByShareId(props.shareId)
    if (resp.data.code === 0 && resp.data.data) {
        list.items = resp.data.data.items

        list.items.forEach(item => {
            item.sizeStr = formatSize(item.size)
            item.state = fileStatusMap[item.status]
            if (item.delFlag === fileStatus.deleted) {
                item.state = fileStatusMap[item.delFlag]
            }
        })
    }
}

function buttonClick(option, id, status) {
    setStatusObj.ids = [id]
    setStatusObj.status = status
    dialogVisible.option[option] = true
}

async function setStatus(option) {
    console.log(setStatusObj)
    const resp = await fileApi.setStatus(setStatusObj)
    if (resp.data.code === codeOk) {
        await listItems()
        promptSuccess('操作成功')
        dialogVisible.option[option] = false
        return
    }
    promptError(`操作失败，${resp.data.msg}`)
}

async function getUrl(id) {
    if (urlMap.has(id)) {
        window.open(urlMap.get(id))
        return
    }
    console.log(id)
    const resp = await shareApi.getUrl(id, 1)
    if (resp.data.code === codeOk) {
        urlMap.set(id, resp.data.data.url)
        window.open(resp.data.data)
        return
    }
    promptError(`获取链接失败，${resp.data.msg}`)
}

onBeforeMount(async () => {
    await listItems()
})

</script>

<style scoped>
.small-pic {
    width: 40px;
    height: 40px;
    border-radius: 5px;
}</style>