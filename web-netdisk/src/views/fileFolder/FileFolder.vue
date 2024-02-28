<template>
    <el-row>
        <el-col :span="19">
            <router-view :key="$route.fullPath"/>
        </el-col>

        <el-col :span="5">
          <file-info/>
        </el-col>
    </el-row>
</template>

<script lang="ts" setup>
import FileInfo from "../../components/files/Info.vue";

</script>

<style scoped>

</style>