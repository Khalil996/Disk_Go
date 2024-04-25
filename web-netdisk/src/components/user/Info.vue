<template>

  <div class="form-div">
    <el-form
        label-position="top"
        label-width="100px"
        class="form"
        :show-file-list="false"
    >
      <el-form-item label="当前头像">
        <el-avatar :size="120"
                   :src="user.data.avatar"
                   style="margin-right:20px"
        />
        <el-upload
            class="upload-demo"
            :limit="1"
            :http-request="changeAvatar"
            :auto-upload="true"
        >
          <el-button type="primary" style="margin-top: 70px">上传头像</el-button>
          <template #tip>
            <div class="el-upload__tip">
              接受 jpg/png 格式
            </div>
          </template>
        </el-upload>

        <div class="progress">
          <el-progress type="dashboard" :percentage="percentage"
                       :status="status" :stroke-width="20"
                       style="position:relative; top: 20px">
            <template #default="{ percentage }">
              <div>
                <div>{{ percentage }}%</div>
              </div>
            </template>
          </el-progress>
          <div style="font-weight: 700;color: #707070"> {{ used }} / {{ capacity }}</div>
        </div>
      </el-form-item>

      <el-form-item label="名称" size="large">
        <el-input v-model="user.data.name"/>
      </el-form-item>
      <el-form-item label="账号" size="large">
        <el-input disabled v-model="user.data.username"/>
      </el-form-item>
      <el-form-item label="邮箱" size="large">
        <el-input v-model="user.data.email">
          <template #append>
            <el-button @click="sendCode2Email(user.data.email)">获取验证码</el-button>
          </template>
        </el-input>
      </el-form-item>
      <el-form-item label="验证码" size="large">
        <el-input v-model="code"/>
      </el-form-item>
      <el-form-item label="个性简介" size="large">
        <el-input
            v-model="user.data.signature"
            :autosize="{ minRows: 4, maxRows: 10 }"
            type="textarea"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="warning" size="large" style="margin-top: 20px"
                   @click="updateInfo">确定
        </el-button>
      </el-form-item>
    </el-form>
  </div>

</template>

<script lang="ts" setup>
import {onMounted, reactive, ref} from "vue";
import {extM, updateAvatar} from "@/components/user/info.ts";
import {codeOk, promptError, promptSuccess} from "@/utils/apis/base.ts";
import {useBaseStore, UserInfo} from "@/store";

import {UploadRequestOptions} from "element-plus";
import {formatSize, sendCode2Email} from "@/utils/util.ts";

const baseStore = useBaseStore(),
    user: { data: UserInfo } = reactive({
      data:
          {id: 0, username: '', name: '', avatar: '', email: '', signature: '', used: 0, capacity: 0, status: 0}
    })

let code = ref(''),
    percentage = ref(0),
    status = ref('success'),
    used = ref(''),
    capacity = ref('')

async function showUserInfo() {
  if (user.data.id === 0) {
    user.data = await baseStore.getUserInfo()
    percentage.value = Math.floor((user.data.used / user.data.capacity) * 100)
    used.value = formatSize(user.data.used)
    capacity.value = formatSize(user.data.capacity)
    if (percentage.value > 90) {
      status.value = 'exception'
    }
    console.log(user.data, user.data.name)
  }
}

async function changeAvatar(param: UploadRequestOptions) {
  const file = param.file
  let ext = file.name.substring(file.name.lastIndexOf('.'))
  if (!extM[ext]) {
    promptError(`不支持${ext}，仅支持 jpg/png`)
    return
  }
  if (file.size > 1024 * 1025) {
    promptError('文件大小不能超过1MB')
    return
  }
  const formData = new FormData()
  formData.append('file', file)
  const resp = await updateAvatar(formData)
  if (resp.code === codeOk) {
    user.data.avatar = resp.data.url
    await baseStore.updateUserInfo(user.data, false)
  }
}

async function updateInfo() {
  user.data.code = code.value
  const resp = await baseStore.updateUserInfo(user.data, true)
  if (resp.code === codeOk) {
    promptSuccess('修改成功')
  }
}

onMounted(() => {
  showUserInfo()
})

</script>

<style lang="scss" scoped>

.progress {
  position: absolute;
  float: right;
  right: 1px;
}

.form-div {
  background: rgba(133, 227, 255, 0.5);
  padding: 10px;
  border: 1px solid lightgray;
  border-radius: 10px;
  display: flex;
  justify-content: center;
}

.form {
  width: 60%;
  padding: 50px 0 0 0;
}
</style>