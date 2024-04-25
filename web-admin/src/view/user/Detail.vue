<template>
    <div>
        <el-card class="box-card">
            <template #header>
                <div class="card-header">
                    <span>用户详情</span>
                    <el-button size="mini" style="float:right" @click="router.go(-1)">返回</el-button>
                </div>
            </template>
            <el-form>
                <el-form-item label="姓名:">
                    {{ userDetail.name }}
                </el-form-item>
                <el-form-item label="年龄:">
                    {{ userDetail.age }}
                </el-form-item>
            </el-form>
        </el-card>
    </div>
</template>

<script setup>
import { onBeforeMount, reactive, ref } from 'vue';
import { useRouter, useRoute } from 'vue-router'
import userApi from "../../api/user";
const route = useRoute();
const router = useRouter();
const userDetail = reactive({
    id: ''
})
onBeforeMount(async () => {
    if (route.query.id) {
        const res = await userApi.getUserDetail({ id: route.query.id })
        Object.assign(userDetail, res.data.data);
    }
})
</script>

<style lang="scss" scoped>

</style>