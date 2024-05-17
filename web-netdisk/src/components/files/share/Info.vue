<template>
  <div class="form-div">
    <el-row>

      <el-col :span="24" v-if="!validated">
        <div style="position: absolute; left: 10%;color: #adabab;
    font: 800 23px Arial, sans-serif; line-height: 100%;">提取文件
        </div>
        <div v-if="ownerInfo.data.shareStatus === shareExpired"
             class="small-zi">
          当前分享已过期！😣
        </div>
        <div v-if="ownerInfo.data.shareStatus === shareNotExistOrDeleted"
             class="small-zi">
          当前分享已被删除或者不存在！😣
        </div>
        <div v-else-if="ownerInfo.data.shareStatus === shareIllegal"
             class="small-zi">
          当前分享已被违法封禁！😡
        </div>
        <div v-else-if="ownerInfo.data.userStatus != userStatus.ok"
             class="small-zi">
          当前用户已被违法封禁！😡
        </div>

        <div  v-if="ownerInfo.data.shareStatus === shareNotExpired && !validated">
          <div class="pwd-box">
            <el-image class="big-pic"
                      :src="ownerInfo.data.avatar"
                      fit="cover"
            />
            <div style="position: relative; top: -100px; right: -115px">
              <div style="font-size: 2rem; font-weight: 700">
                {{ ownerInfo.data.name }}
              </div>
              <div style="position: absolute; margin-top: 20px">
                <span v-if="ownerInfo.data.signature == ''">暂无签名</span>
                <span v-else>{{ ownerInfo.data.signature }}</span>
              </div>
            </div>

            <el-form label-position="top">
              <el-form-item label="请先输入提取码：" size="large">
                <el-input v-model="pwdInput"/>
              </el-form-item>
              <el-form-item>
                <el-button size="large" style="width: 100%;"
                           type="primary"
                           @click="listItems(pwdInput)">提取文件
                </el-button>
              </el-form-item>
            </el-form>
          </div>
        </div>
      </el-col>

      <el-col  v-if="list.items && list.items.length!=0
            && ownerInfo.data.shareStatus === shareNotExpired
            && validated" :span="24" style="margin-bottom: 100px">
        <!--                <el-empty v-if="!list.items || list.items.length==0"-->
        <!--                          description="当前分享文件夹没有文件 😺"/>-->

        <div style="margin: 20px 0">
          <div style="font-size: 1.8rem; font-weight: 700; margin-bottom: 10px">
            <div>
              {{ list.name }}
            </div>
            <span style="float: right;">
                            <el-button v-if="ownerInfo.data.userId != list.owner"
                                       size="large" type="primary"
                                       @click="downloadFiles">
                                <el-icon>
                                    <Download/>
                                </el-icon>&nbsp;
                                下载
                            </el-button>
                            <el-button-group v-if="ownerInfo.data.userId == list.owner">
                                <el-button plain size="large" type="primary"
                                           @click="downloadFiles">
                                <el-icon>
                                    <Download/>
                                </el-icon>&nbsp;
                                下载
                            </el-button>
                                <el-button v-if="ownerInfo.data.userId == userInfo.id"
                                           plain size="large" type="danger"
                                           @click="dialogVisible = true">
                                <el-icon><CircleClose/></el-icon>&nbsp;
                                取消分享
                            </el-button>
                                <el-button v-if="ownerInfo.data.userId !== userInfo.id"
                                           plain size="large" type="danger"
                                           @click="tipoff=true; reason=''"
                                >
                                    举报
                                </el-button>
                            </el-button-group>
                        </span>
          </div>
          <div style="font-size: 1.4rem">
            <el-icon>
              <Clock/>
            </el-icon>&nbsp;
            <span>{{ list.created }}</span>
            <span style="margin-left: 50px;">{{ state }}</span>
          </div>
        </div>

        <el-table
            ref="fileTableRef"
            :data="list.items" style="width: 100%"
        >
          <el-table-column type="selection" width="55"/>
          <el-table-column label="文件名" min-width="250">
            <template #default="scope">
              <div style="display: flex; align-items: center">
                <el-image :src="`/src/assets/alt_type${scope.row.type}.jpg`"
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
      </el-col>
      <el-footer >Copyright © 2024 咪咪网盘</el-footer>
    </el-row>

    <el-dialog v-model="dialogVisible" title="取消分享">
      <h3>
        <el-icon>
          <Warning/>
        </el-icon>
        确定取消这个分享吗😶
      </h3>
      <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false;">取消</el-button>
        <el-button type="primary" @click="cancelShare()">
          确定
        </el-button>
      </span>
      </template>
    </el-dialog>

    <el-dialog v-model="tipoff" title="确认举报">
      <h3>
        <el-icon>
          <Warning/>
        </el-icon>
        确定举报这个分享吗😶
      </h3>
      <el-form label-position="left">
        <el-form-item label="举报理由：">
          <el-input v-model="reason"></el-input>
        </el-form-item>
      </el-form>
      <template #footer>
      <span class="dialog-footer">
        <el-button @click="tipoff = false;">取消</el-button>
        <el-button type="primary" @click="tipoffCommit()">
          确定
        </el-button>
      </span>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts" setup>
import {ElTable} from "element-plus";
import {onMounted, reactive, ref} from "vue";
import {
  downloadCount,
  getOwnerInfoByShareId,
  listFilesByShareId,
  shareCancel,
  ShareItem,
  shareReport
} from "@/components/files/share/Info.ts";
import {formatSize, formatState} from "@/utils/util.ts";
import {
  Clock, Download, CircleClose, Warning
} from "@element-plus/icons-vue";
import {useRoute} from "vue-router";
import {codeOk, promptError, promptSuccess} from "@/utils/apis/base.ts";
import {
  fileStatus,
  fileStatusMap, shareExpired,
  shareIllegal,
  shareNotExistOrDeleted,
  shareNotExpired,
  userStatus
} from "@/utils/constant.ts";
import {useBaseStore} from "@/store";

const query = useRoute().query,
    userStore = useBaseStore()

let pwd = ref(''),
    pwdInput = ref(''),
    validated = ref(false),
    userInfo = {id: -1, username: '', name: '', avatar: '', email: '', signature: '', used: 0, capacity: 0, status: 0}

const props = defineProps(['shareId']),
    fileTableRef = ref<InstanceType<typeof ElTable>>(),
    list = reactive<{
      name: string
      created: string
      expired: number
      owner: number
      items: ShareItem[]
    }>({
      name: '',
      created: '',
      expired: 0,
      owner: 0,
      items: []
    }),
    ownerInfo = reactive({
      data:
          {
            shareStatus: 0,
            userId: -1,
            name: '',
            avatar: '',
            signature: '',
            userStatus: 0,
          }
    })

let state = '',
    selected = [],
    dialogVisible = ref(false),
    tipoff = ref(false),
    reason = ref('')

const listItems = async (pwdStr: string) => {
  pwd.value = pwdStr
  let resp = await listFilesByShareId(props.shareId, pwdStr)
  if (resp.code === codeOk && resp.data) {
    list.name = resp.data.name
    list.created = resp.data.created
    list.items = resp.data.items
    list.owner = resp.data.owner
    state = formatState(resp.data.expired)
    list.items.forEach(item => {
      item.sizeStr = formatSize(item.size)
    })
    validated.value = true
  }
}

async function downloadFiles() {
  if (state == '已过期') {
    promptError('当前分享已经过期了🥲')
    return
  }
  selected = fileTableRef.value!.getSelectionRows()
  for (const file of selected) {
    if (file.status !== fileStatus.ok) {
      promptError(`文件${fileStatusMap[file.status]}`)
      continue
    }
    await window.open(file.url)
  }
  await  downloadCount(props.shareId)
}

async function getOwnerInfo() {
  const resp = await getOwnerInfoByShareId(props.shareId)
  if (resp.code === codeOk) {
    ownerInfo.data = resp.data
    console.log(ownerInfo.data.userStatus, ownerInfo.data.shareStatus)
  }
}

async function cancelShare() {
  const resp = await shareCancel([props.shareId])
  if (resp.code === codeOk) {
    dialogVisible.value = false
    promptSuccess('操作成功，窗口即将关闭')
    setTimeout(() => {
      window.close()
    }, 1000)
    return
  }
  promptError(`取消失败，${resp.msg}`)
}

async function tipoffCommit() {
  const resp = await shareReport(reason.value, props.shareId)
  if (resp.code === codeOk) {
    tipoff.value = false
    promptSuccess('操作成功，窗口即将关闭')
    return
  }
  promptError(`提交失败，${resp.msg}`)
}

onMounted(async () => {
  await getOwnerInfo()

  console.log(query.pwd, 123)
  if (ownerInfo.data.shareStatus === shareNotExpired) {
    console.log(ownerInfo.data.shareStatus)
    if (query.pwd != undefined) {
      pwd.value = query.pwd
      await listItems(pwd.value)
    }
  }
  try {
    userInfo = await userStore.getUserInfo()
  } catch (e) {
    console.log(userInfo, "aaaadsaddas")
    userInfo.id = -1
  }
})
</script>

<style scoped>
.small-pic {
  width: 40px;
  height: 40px;
  border-radius: 5px;
}

.big-pic {
  width: 100px;
  height: 100px;
  border-radius: 50%;
}

.form-div {
  background: rgba(133, 233, 255, 0.5);
  padding: 10px;
  border: 1px solid lightgray;
  border-radius: 10px;
  display: flex;
  justify-content: center;
}

.pwd-box {
  margin: 200px 0;
}

.small-zi {
  font-weight: 600;
  font-size: 3rem;
  margin: 250px 0;
}
</style>