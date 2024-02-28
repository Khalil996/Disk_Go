<template>
  <div class="common-layout">
    <el-container>
      <el-header style="background-color: lightsalmon">
        <avatar/>

      </el-header>
      <el-container>
        <el-aside style="width: 10%;">
          <navigator/>
        </el-aside>
        <el-main style="background-color: #fff2e3; padding: 15px 0 0 0">
          <router-view/>
        </el-main>
      </el-container>
    </el-container>
  </div>

</template>

<script lang="ts" setup>
import Avatar from "../components/header/Avatar.vue";
import Navigator from "../components/aside/Navigator.vue";

</script>

<style scoped>
.common-layout{
    margin: 0;
}
.el-header{
    height: 6vh;
}
</style>
