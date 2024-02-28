<template>
    <el-form :model="registerForm" label-width="120px">
        <el-form-item label="账号" size="large">
            <el-input v-model="registerForm.name"/>
        </el-form-item>
        <el-form-item label="密码" size="large">
            <el-input v-model="registerForm.password"/>
        </el-form-item>
      <el-form-item label="邮箱" size="large">
        <el-input v-model="registerForm.email"/>
      </el-form-item>
      <el-form-item label="验证码" prop="code">
        <el-row :gutter="20">
          <el-col :span="12">
            <el-input v-model="registerForm.code"/>
          </el-col>
          <el-col :span="12" style="text-align:right;">
            <el-button disabled v-if="remainTime>0&&remainTime<60">{{ remainTime }}秒</el-button>
            <el-button @click="sendCode" v-else type="primary">发送验证码</el-button>
          </el-col>
        </el-row>
      </el-form-item>
        <el-form-item>
            <el-button type="primary" size="large" @click="register">注册</el-button>
        </el-form-item>
    </el-form>
</template>

<script lang="ts" setup>
import {reactive,ref} from "vue";
import {registerPost, RegisterReq, sendEmailPost} from "./registerForm.ts";
import router from "../router";
import {promptSuccess} from "../utils/apis/base.ts";

let registerForm = reactive<RegisterReq>({
    name: '',
    password: '',
    email: '',
    code: '',
})
const remainTime = ref(60)

const register = async () => {
    const resp = await registerPost(registerForm)
    if (resp.code === 0) {
      promptSuccess()
    }

}
const sendCode = async () => {
  const resp = await sendEmailPost(registerForm.email)
  if (resp.code === 0) {

  }
}

</script>

<style scoped>

</style>