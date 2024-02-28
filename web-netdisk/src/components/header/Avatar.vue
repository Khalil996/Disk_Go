<template>
  <!--  <el-row :gutter="20" class="demo-avatar demo-basic">-->
  <!--    <el-col :span="6" :offset="6">-->
  <!--      <div class="demo-basic&#45;&#45;circle">-->
  <!--        <div class="block head-pic">-->
  <!--          <el-avatar :size="50" :src="url" />-->
  <!--        </div>-->
  <!--      </div>-->
  <!--    </el-col>-->
  <!--  </el-row>-->
    <el-row class="el-row">
        <el-col :span="23">
            <div class="grid-content ep-bg-purple"/>
        </el-col>
        <el-col :span="1">
            <div class="grid-content ep-bg-purple-light">
                <div class="demo-basic--circle">
                    <div class="block head-pic">
                        <el-avatar :size="50" :src="url"/>
                    </div>
                </div>
            </div>
        </el-col>
    </el-row>
</template>

<script lang="ts" setup>
import {onMounted} from "vue";
import {useBaseStore, UserInfo} from "../../store";

const url = "https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"

let user: { data: UserInfo }

async function showUserInfo() {
    if (!user) {
        user = await useBaseStore().getUserInfo()
    }
}

onMounted(() => {
    showUserInfo()
})

</script>

<style scoped>
.head-pic {
    margin-top: 5px;
}

.el-row {
    margin-bottom: 20px;
}

.el-row:last-child {
    margin-bottom: 0;
}

.grid-content {
    min-height: 36px;
}
</style>
