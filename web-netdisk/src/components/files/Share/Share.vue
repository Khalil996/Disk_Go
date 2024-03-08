<template>
  <el-row>
    <el-col :span="24">
      <div class="file-table">

        <el-empty v-if="!shareList.data || shareList.data.length==0"
                  description="文件列表为空，上传你的第一个文件吧！😺"/>

        <el-table v-if="shareList && shareList.data.length!=0"
                  ref="fileTableRef"
                  :data="shareList.data" style="width: 100%"
                  @selection-change="fileSelectionChange"
        >
          <el-table-column type="selection" width="55"/>
          <el-table-column label="分享文件" min-width="500">
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
                            <span @click="">
                                <el-icon><Link/></el-icon> 复制链接
                            </span>
              &nbsp;&nbsp;&nbsp;
              <span @click="">
                                <el-icon><CircleClose/></el-icon> 取消分享
                            </span>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </el-col>
  </el-row>

</template>

<script lang="ts" setup>
import {CircleClose, Link} from "@element-plus/icons-vue";
import {shareExpired, shareIllegal, shareNotExpired, typeImage} from "@/utils/constant.ts";
import {ElTable} from "element-plus";
import {reactive, onMounted} from "vue";
import {listShareFiles, Share} from "@/components/files/share/share.ts";
import {codeOk} from "@/utils/apis/base.ts";
import {formatLeft, formatTime} from "@/utils/util.ts";

let shareList = reactive<{ data: Share[] }>({
  data: []
})

const listFiles = async () => {
  const resp = await listShareFiles()
  if (resp.code === codeOk) {
    shareList.data = resp.data
    shareList.data.forEach(share => {
      switch (share.status) {
        case shareNotExpired:
          share.state = formatState(share)
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

function formatState(share: Share) {
  const now = new Date().getTime() / 1000
  console.log(now, share.expired)
  if (now >= share.expired) {
    share.status = shareExpired
    return '已过期'
  }
  return formatLeft(share.expired) + '后过期'
}

onMounted(() => {
  listFiles()
})

</script>

<style scoped>

</style>