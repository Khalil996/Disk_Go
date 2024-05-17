<template>
  <div>
    <el-container class="home-container">
      <!-- header -->
      <el-header>
        <el-row>
          <el-col :span="4">
            <img src="@/assets/img/ji_tui_cai.png" alt=""/>
            <p class="system-name">咪咪网盘管理后台</p>
          </el-col>
          <el-col :offset="12" :span="8" style="min-width: 150px">
            <el-dropdown style="float: right; margin: 20px 10px">
                            <span class="el-dropdown-link" style="color: #fff; cursor: pointer; font-weight: 600">
                                {{ admin.data.Name }} &nbsp;&nbsp; <el-icon class="el-icon--right">
                                    <arrow-down/>
                                </el-icon>
                            </span>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item @click.native="logout">退出系统</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
            <el-avatar shape="square" :src="avatar" style="margin: 10px; float: right"></el-avatar>
          </el-col>
        </el-row>
      </el-header>

      <el-container style="overflow: auto">
        <!-- 菜单 -->
        <el-aside>
          <div class="toggle-button" @click="isCollapse = !isCollapse">
            <el-icon :size="20">
              <Expand v-if="isCollapse"/>
              <Fold v-if="!isCollapse"/>
            </el-icon>
          </div>
          <el-menu router :default-active="activePath"
                   class="el-menu-vertical-demo"
                   :collapse="isCollapse"
                   :default-openeds="['1']"
          >
            <el-menu-item index="/log" @click="saveActiveNav('/log')">
              <el-icon>
                <Memo/>
              </el-icon>
              <span>系统日志</span>
            </el-menu-item>
            <el-menu-item index="/index" @click="saveActiveNav('/index')">
              <el-icon>
                <house/>
              </el-icon>
              <span>首页</span>
            </el-menu-item>
            <el-menu-item index="/user/list" @click="saveActiveNav('/user/list')">
              <el-icon>
                <user/>
              </el-icon>
              <span>用户列表</span>
            </el-menu-item>
            <el-menu-item index="/share" @click="saveActiveNav('/user/list')">
              <el-icon>
                <share/>
              </el-icon>
              <span>分享管理</span>
            </el-menu-item>

            <el-menu-item v-if="admin.data.Status === 1"
                          index="/admin" @click="saveActiveNav('/admin')">
              <el-icon>
                <Sunny/>
              </el-icon>
              <span>管理员列表</span>
            </el-menu-item>
          </el-menu>
        </el-aside>

        <el-container>
          <el-main>
            <router-view></router-view>
          </el-main>
          <el-footer>Copyright © 2024 咪咪网盘</el-footer>
        </el-container>
      </el-container>
    </el-container>
  </div>
</template>
<script setup>
import {
  Memo, User, Share, House, ArrowDown, Sunny
} from '@element-plus/icons-vue'
import {onBeforeMount, reactive, ref} from 'vue';
import avatar from "../assets/img/img.png"
import {useRouter} from 'vue-router'
import {useBaseStore} from "../../store/index.js";
import adminApi from "@/api/admin.js";
import {codeOk, promptError} from "@/utils/http/base.js";

function qwe() {
  window.open("http://127.0.0.1:5601/app/discover#/")
}

const router = useRouter(),
    baseStore = useBaseStore(),
    admin = reactive({
      data: {}
    })

async function getAdminInfo() {
  const resp = await adminApi.getAdminInfo()
  if (resp.data.code === codeOk) {
    admin.data = resp.data.data
    console.log(admin.data.Name)
  } else {
    promptError(resp.data.msg)
  }
}

// 挂载 DOM 之前
onBeforeMount(async () => {
  await getAdminInfo()

  activePath.value = sessionStorage.getItem("activePath")
      ? sessionStorage.getItem("activePath")
      : "/index"
})

let isCollapse = ref(false);
let activePath = ref("");
// 保存链接的激活状态
const saveActiveNav = (path) => {
  sessionStorage.setItem("activePath", path);
  activePath.value = path;
}
const logout = () => {
  // 清除缓存
  baseStore.clearToken()
  router.push("/login");
}
</script>

<style scoped>
.home-container {
  position: absolute;
  height: 100%;
  top: 0px;
  left: 0px;
  width: 100%;
  background: url("@/assets/img/background_grid.png");
}

.el-header {
  background: linear-gradient(120deg, #0699f5, #83c6ea, #2dd4f1);
  padding: 0 10px;
  overflow: hidden;
}

.system-name {
  color: #b0e6f8;
  font: 800 23px Arial, sans-serif;
  line-height: 100%;
  position: absolute;
  left: 5%;
  top: 1px
}

.el-aside {
  background: #23a9f1;
  width: auto !important;
}

.el-menu-vertical-demo {
  background-color: #1a9fe7;
}

.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}

.el-footer {
  color: #cccccc;
  text-align: center;
  line-height: 60px;
}

.el-footer:hover {
  color: lightyellow;
}

.toggle-button {
  background-color: #2fecdc;
  font-size: 18px;
  line-height: 24px;
  color: #fff;
  text-align: center;
  letter-spacing: 0.2em;
  cursor: pointer;
  color: black;
  border-bottom: 1px solid darkred;
}

.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}

.el-menu-item.is-active {
  color: black !important;
  font-size: 15px;
  font-weight: bold;
  background-color: #63b7f3 !important;
  border-radius: 2px;
  height: 50px;
  line-height: 50px;
  box-sizing: border-box;
  margin: 2px 5px 0 2px;
}
</style>