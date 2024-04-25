<template>
    <el-form :model="loginForm" :rules="rules" label-width="120px">
        <el-form-item label="账号" size="large" prop="username">
            <el-input v-model="loginForm.username"/>
        </el-form-item>
        <el-form-item label="密码" size="large" prop="password">
            <el-input type="password" v-model="loginForm.password"/>
        </el-form-item>
        <el-form-item>
            <el-button type="primary" size="large" @click="login">登录</el-button>
        </el-form-item>
    </el-form>
</template>

<script lang="ts" setup>
import {reactive} from "vue";
import {loginPost, LoginReq} from "./loginForm.ts";
import {useBaseStore} from "@/store";
import router from "../router";

let loginForm = reactive<LoginReq>({
    password: "",
    username: "",
})

const login = async () => {
  const resp = await loginPost(loginForm)
  if (resp.code === 0) {
    const baseStore = useBaseStore()
    baseStore.updateToken(resp.data.token)
    await baseStore.updateUserInfo(resp.data.userInfo)
    router.push('/file/folder/0')
  }
}
const rules = {
    username: [
        {required: true, message: '请输入账号', trigger: 'blur'},
        {min: 4, max: 20, message: '账号的长度在4-20之间', trigger: 'blur'},
    ],
    password: [
        {required: true, message: '请输入密码', trigger: 'blur'},
        {min: 4, max: 20, message: '密码的长度在4-20之间', trigger: 'blur'},
    ],
}


// const resetForm = (formEl: FormInstance | undefined) => {
//   if (!formEl) return
//   formEl.resetFields()
// }

</script>

<style scoped>

</style>