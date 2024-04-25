<template>

  <div class="header">
    <div class="header-content" style="width: 1000px;">
      <router-link to="/file/folder/0" class="logo">
        <el-image class="small-pic"
                  src="/src/assets/img.png"
                  alt="/sec/assets/alt_type1.jpg"
                  :fit="'cover'"/>
      </router-link>
      <!--            <router-link to="/user/info" class="logo">-->
      <!--                <span v-for="c in color7" :style="{ color: c.color }">{{ c.letter }}</span>-->
      <!--            </router-link>-->
      <div class="nav">
        <router-link class="menu-item home" to="/file/folder/0">去首页</router-link>
      </div>

      <div class="user-info-panel">
        <avatar></avatar>
      </div>
    </div>
  </div>

  <div class="body-content">
    <div class="main-outer">
      <div class="container-body">
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import Avatar from "@/components/header/Avatar.vue";
import {ref} from "vue";
import {sendCode2Email} from "@/utils/util.ts";

const color7 = ref([
  {
    letter: "我",
    color: "#3285FF"
  },
  {
    letter: "的",
    color: "#FB3624"
  },
  {
    letter: "信",
    color: "#FFBA02"
  },
  {
    letter: "息",
    color: "#24B24E"
  },
])

</script>

<style lang="scss" scoped>
.header {
  top: 0;
  width: 100%;
  position: fixed;
  box-shadow: 0 2px 6px 0 #ddd;
  background-color: #b3d0f5;
  z-index: 1000;
}

.header-content {
  margin: 0 auto;
  height: 60px;
  display: flex;
  align-items: center;
}

.logo {
  display: block;
  text-decoration: none;
  margin-right: 5px;
}

span {
  font-size: 30px;
}

.nav {
  flex: 1;
}

.menu-item {
  cursor: pointer;
  margin-left: 20px;
}

.user-info-panel {
  width: 300px;
}

.body-content {
  padding-top: 60px;
  position: relative;
}

.small-pic {
  width: 50px;
  height: 50px;
  border-radius: 5px;
}

.main-outer {
  background: url("@/assets/background_grid.png");
  width: 100%;
  height: 140vh;
  margin: 0;
  padding: 0;
  position: absolute;
  top: 6vh;
  overflow: hidden;
}

.container-body {
  margin: 5vh auto;
  width: 45%;
}

</style>