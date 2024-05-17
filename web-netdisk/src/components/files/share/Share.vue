<template>
  <el-row>
    <el-col :span="24">
      <div class="file-table">
        <div style="margin-bottom: 15px;">
          <div v-if="fileButtonsState === 0" style="height: 32px; line-height: 30px;
                    font-size: 1.4rem; font-weight: 700;">我的全部分享
          </div>
          <el-button-group v-else-if="fileButtonsState==1">
            <el-button type="primary" round plain :icon="Link"
                       @click="copyLink('', true)">复制链接
            </el-button>
            <el-button type="danger" round plain :icon="DeleteFilled"
                       @click="dialogVisible.option[0]=true">
              删除分享
            </el-button>
          </el-button-group>
          <el-button v-else-if="fileButtonsState==2"
                     type="danger" :icon="DeleteFilled" round
                     @click="dialogVisible.option[0]=true">删除分享
          </el-button>
        </div>

        <el-empty v-if="!shareList.data || shareList.data.length==0"
                  description="分享列表为空，创建你的第一个分享吧！😺"/>

        <el-table v-if="shareList && shareList.data.length!=0"
                  ref="fileTableRef"
                  :data="shareList.data" style="width: 100%"
                  @selection-change="fileSelectionChange"
        >
          <el-table-column type="selection" width="55"/>
          <el-table-column label="分享文件" min-width="500">
            <template #default="scope">
              <div class="share-row share-row2"
                   v-if="scope.row.type === typeMulti"
                   @click="toShareInfo(`/info/share/${scope.row.id}?pwd=${scope.row.pwd}`)">
                <el-image
                    class="small-pic"
                    src="/src/assets/alt_folder.jpg"
                    :fit="'cover'"/>
                &nbsp;<span style="margin-left: 5px">{{ scope.row.name }}</span>
              </div>

              <div v-else class="share-row">
                <el-image
                    class="small-pic"
                    :src="`/src/assets/alt_type${scope.row.type}.jpg`"
                    :fit="'cover'"/>
                &nbsp;<span style="margin-left: 5px">{{ scope.row.name }}</span>
              </div>
            </template>
          </el-table-column>
          <el-table-column label="创建时间" min-width="100">
            <template #default="scope">
              <div>{{ scope.row.created }}</div>
            </template>
          </el-table-column>
          <el-table-column label="状态" min-width="100">
            <template #default="scope">
              <div>{{ scope.row.state }}</div>
            </template>
          </el-table-column>
          <el-table-column prop="pwd" label="提取码" min-width="100"/>
          <el-table-column label="类型" min-width="100">
            <template #default="scope">
              <div>{{ typeMap[scope.row.type] }}</div>
            </template>
          </el-table-column>
          <el-table-column label="点击次数" min-width="100">
            <template #default="scope">
              <div>{{ scope.row.clickNum }}</div>
            </template>
          </el-table-column>
          <el-table-column label="下载次数" min-width="100">
            <template #default="scope">
              <div>{{ scope.row.downloadNum }}</div>
            </template>
          </el-table-column>
          <el-table-column min-width="100">
            <template #default="scope">
                            <span @click="copyLink(scope.row.url, false)">
                                <el-icon color="#48a3ff"><Link/></el-icon>
                            </span>
              &nbsp;&nbsp;&nbsp;
              <span @click="dialogVisible.option[0]=true; cancelId=scope.row.id">
                                <el-icon color="red"><CircleClose/></el-icon>
                            </span>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-col>
  </el-row>

  <el-dialog v-model="dialogVisible.option[0]" title="删除分享">
    <h3>
      <el-icon>
        <Warning/>
      </el-icon>
      确定删除这个分享吗😶
    </h3>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible.option[0]=false; cancelId=''">取消</el-button>
        <el-button type="primary" @click="cancel()">
          确定
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script lang="ts" setup>
import {
  CircleClose, DeleteFilled, Link, Warning
} from "@element-plus/icons-vue";
import {shareExpired, shareIllegal, shareNotExpired, typeMap, typeMulti} from "@/utils/constant.ts";
import {ElTable} from "element-plus";
import {reactive, onMounted, ref} from "vue";
import {listShareFiles, Share, shareCancel} from "@/components/files/share/share.ts";
import {codeOk, promptError, promptSuccess} from "@/utils/apis/base.ts";
import {formatState} from "@/utils/util.ts";

let shareList = reactive<{ data: Share[] }>({
      data: []
    }),
    cancelId = '',
    fileButtonsState = ref(0)

const dialogVisible = reactive({option: [false]}),
    fileTableRef = ref<InstanceType<typeof ElTable>>()


function fileSelectionChange(shares: Share[]) {
  if (!shares || shares.length == 0) {
    fileButtonsState.value = 0
  } else if (shares) {
    if (shares.length === 1) {
      fileButtonsState.value = 1
    } else {
      fileButtonsState.value = 2
    }
  }
}

const listFiles = async () => {
  const resp = await listShareFiles()
  if (resp.code === codeOk) {
    shareList.data = resp.data
    shareList.data.forEach(share => {
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
  }
}

async function cancel() {
  let ids = [cancelId]
  if (cancelId == '') {
    const selected = fileTableRef.value!.getSelectionRows()
    ids = selected.map(share => share.id)
  }
  console.log(ids)
  const resp = await shareCancel(ids)
  if (resp.code === codeOk) {
    promptSuccess(`删除分享成功，${resp.msg}`)
    dialogVisible.option[0] = false // 确保对话框关闭
    cancelId = '' // 重置 cancelId
   await listFiles() // 重新加载分享列表
  } else {
    promptError(`删除分享失败，${resp.msg}`)
  }
}

async function copyLink(link: string, button: boolean) {
  if (button) {
    link = (fileTableRef.value!.getSelectionRows())[0].url
  }
  try {
    link = '分享链接：http://' + link + '\n复制后点击链接即可'
    if (link.includes('pwd')) {
      link += '，无需输入提取码'
    }
    await navigator.clipboard.writeText(link)
    promptSuccess('已将链接复制到剪贴板')
  } catch (e) {
    promptError(`复制链接失败，${e}`)
  }
}

async function toShareInfo(url: string) {
  window.open(url)
}

onMounted(() => {
  listFiles()
})

</script>

<style scoped>
.small-pic {
  width: 40px;
  height: 40px;
  border-radius: 5px;
}

.share-row {
  display: flex;
  align-items: center;
}

.share-row2:hover {
  cursor: pointer;
  background-color: rgb(230, 230, 245);
  border-radius: 5px;
}
</style>