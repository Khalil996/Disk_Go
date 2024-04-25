<template>
    <div class="card-list">
        <el-card>
            <el-input style="width:440px" @clear="searchUser" clearable v-model="searchForm.name"
                      placeholder="请输入用户姓名" class="input-with-select">
                <template #append>
                    <el-button icon="Search" @click="searchUser"/>
                </template>
            </el-input>
            <el-table :data="shares" border style="width: 100%; margin-top:20px">
                <el-table-column prop="id" label="ID" min-width="100"/>
                <el-table-column prop="name" label="名称" min-width="200"/>
                <el-table-column prop="created" label="创建时间" min-width="150"/>
                <el-table-column label="类型" min-width="60">
                    <template #default="scope">
                        {{ typeMap[scope.row.type] }}
                    </template>
                </el-table-column>
                <el-table-column prop="downloadNum" label="下载次数" min-width="60"/>
                <el-table-column prop="clickNum" label="点击次数" min-width="60"/>
              <el-table-column prop="reason" label="举报理由"/>
                <el-table-column prop="state" label="状态" min-width="180"/>
                <el-table-column label="操作" width="200">
                    <template #default="scope">
                        <el-button v-if="scope.row.type === typeMulti"
                                   type="primary" plain size="small"
                                   @click="openInfo(`/share/${scope.row.id}`)">进入查看
                        </el-button>
                        <el-button v-if="scope.row.type !== typeMulti"
                                   type="primary" plain size="small"
                                   @click="getUrl(scope.row.id)">下载查看
                        </el-button>
                        <el-button v-if="scope.row.status !== shareIllegal"
                                   type="danger" size="small"
                                   @click="buttonClick(1, scope.row.id, shareIllegal, scope.row.type)">封禁
                        </el-button>
                        <el-button v-if="scope.row.status === shareIllegal"
                                   type="primary" size="small"
                                   @click="buttonClick(0, scope.row.id, shareNotExpired, scope.row.type)">恢复
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>
            <!-- 分页 -->
            <el-pagination style="margin-top:20px" :current-page="searchForm.current" :page-size="searchForm.size"
                           :page-sizes="[10, 20, 30, 40]" layout="->,total, sizes, prev, pager, next, jumper"
                           :total="total"
                           @size-change="handleSizeChange" @current-change="handleCurrentChange"/>
        </el-card>
    </div>

    <el-dialog v-model="dialogVisible.option[0]" title="恢复分享">
        <h3>
            <el-icon>
                <Warning/>
            </el-icon>
            需要恢复这个分享吗？😶
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

    <el-dialog v-model="dialogVisible.option[1]" title="封禁分享">
        <h3>
            <el-icon>
                <Warning/>
            </el-icon>
            确定封禁这个分享吗😶，如果包含多个文件需要进入详情操作，单个文件可被直接封禁
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

<script setup>
import shareApi from "@/api/share.js";
import {onMounted, reactive, ref} from "vue";
import {ElMessage, ElMessageBox} from 'element-plus';
import {useRouter} from 'vue-router'
import {formatState} from "@/utils/util.js";
import {shareExpired, shareIllegal, shareNotExpired, typeMap, typeMulti} from "@/utils/constant.js";
import {codeOk, promptError, promptSuccess} from "@/utils/http/base.js";

const router = useRouter();

onMounted(() => {
    listShares()
})

let shares = ref([])
let total = ref(0)
let urlMap = new Map()

const searchForm = reactive({
        current: 1,
        size: 10,
        name: ''
    }),
    dialogVisible = reactive({option: [false, false]}),
    setStatusObj = {
        id: 0, status: 0, type: 0
    }

const listShares = async () => {
    const res = await shareApi.getShareList({'page': 0, 'size': 100});
    console.log(res.data);
    shares.value = res.data.data
    shares.value.forEach(share => {
        switch (share.status) {
            case shareNotExpired:
                share.state = formatState(share.expired)
                break
            case shareIllegal:
                share.state = '内涵非法内容，已封禁'
                break
            case shareExpired:
                share.state = '已过期'
                break
        }
    })
    total.value = res.data.data.total;
}

function openInfo(link) {
    window.open(link)
}

async function getUrl(id) {
    if (urlMap.has(id)) {
        window.open(urlMap.get(id))
        return
    }
    const resp = await shareApi.getUrl(id, 0)
    if (resp.data.code === codeOk) {
        urlMap.set(id, resp.data.data.url)
        window.open(resp.data.data)
        return
    }
    promptError(`获取链接失败，${resp.data.msg}`)
}

function buttonClick(option, id, status, type) {
    setStatusObj.id = id
    setStatusObj.status = status
    setStatusObj.type = type
    dialogVisible.option[option] = true
}

async function setStatus(option) {
    console.log(setStatusObj)
    const resp = await shareApi.setStatus(setStatusObj)
    if (resp.data.code === codeOk) {
        await listShares()
        promptSuccess('操作成功')
        dialogVisible.option[option] = false
        return
    }
    promptError(`操作失败，${resp.data.msg}`)
}

const handleSizeChange = (size) => {
    searchForm.size = size;
    listShares();
}
const handleCurrentChange = (current) => {
    searchForm.current = current;
    listShares();
}
const searchUser = () => {
    searchForm.current = 1;
    listShares();
}
// 删除用户
const deleteUser = (id) => {
    ElMessageBox.confirm(
        '确定要删除该用户信息吗?',
        {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning',
        }
    ).then(async () => {
        const res = await userApi.delUser({id: id});
        if (res.data.success) {
            ElMessage.success("删除成功")
            getUserList();
        } else {
            ElMessage.error("删除失败")
        }
    }).catch(() => {
        ElMessage({
            type: 'info',
            message: '取消删除',
        })
    })
}
</script>

<style lang="scss" scoped>
</style>