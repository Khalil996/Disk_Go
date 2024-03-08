<template>
    <el-row class="tac">
        <el-menu
                active-text-color="blue"
                background-color="#71dae3"
                class="el-menu-vertical-demo"
                default-active="2"
                text-color="black"
                style="width: 100%"
        >
            <el-sub-menu index="1">
                <template #title>
                    <el-icon>
                        <Files/>
                    </el-icon>
                    文件
                </template>
                <el-menu-item index="1-1">
                    <router-link active-class="a" to="/file/folder/0">
                        <el-icon>
                            <Folder/>
                        </el-icon>
                        全部文件
                    </router-link>
                </el-menu-item>
                <el-menu-item index="1-1">
                    <router-link active-class="a" to="/file/0">
                        <el-icon>
                            <Document/>
                        </el-icon>
                        文档
                    </router-link>
                </el-menu-item>
                <el-menu-item index="1-2">
                    <router-link active-class="a" to="/file/1">
                        <el-icon>
                            <PictureFilled/>
                        </el-icon>
                        图片
                    </router-link>
                </el-menu-item>
                <el-menu-item index="1-3">
                    <router-link active-class="a" to="/file/2">
                        <el-icon>
                            <VideoCameraFilled/>
                        </el-icon>
                        视频
                    </router-link>
                </el-menu-item>
                <el-menu-item index="1-4">
                    <router-link class="a" to="/file/3">
                        <el-icon>
                            <Headset/>
                        </el-icon>
                        音频
                    </router-link>
                </el-menu-item>
                <el-menu-item index="1-5">
                    <router-link class="a" to="/file/4">
                        <el-icon>
                            <More/>
                        </el-icon>
                        其他
                    </router-link>
                </el-menu-item>
            </el-sub-menu>
          <el-sub-menu  index="2">
            <template #title>
              <el-icon>
                <Delete/>
              </el-icon>
              回收站
            </template>
            <el-menu-item index="2-1">
              <router-link class="a" to="/bin">
                <el-icon>
                  <Delete/>
                </el-icon>
                <span>回收站</span>
              </router-link>
            </el-menu-item>
          </el-sub-menu>
          <el-sub-menu  index="3">
            <template #title>
              <el-icon>
                <Delete/>
              </el-icon>
              分享
            </template>
            <el-menu-item index="3-1">
              <router-link class="a" to="/share">
                <el-icon>
                  <Delete/>
                </el-icon>
                <span>分享</span>
              </router-link>
            </el-menu-item>
          </el-sub-menu>
            <el-menu-item index="4" disabled>
                <el-icon>
                    <document/>
                </el-icon>
                <span>TODO</span>
            </el-menu-item>
        </el-menu>
        <div class="cat-banner">
            <img src="@/assets/cat_scare.png" alt="" style="width: 30%"/>
        </div>
    </el-row>
</template>

<script lang="ts" setup>
import {
  Files,
  Folder, Document, PictureFilled, Headset, More,
  Delete, VideoCameraFilled
} from '@element-plus/icons-vue'

</script>
<style>
.tac {
  overflow: hidden;
  height: 94vh;
}

.cat-banner {
  background-color: #71dae3;
  position: absolute;
  bottom: 1px;
}

.cat-banner img {
  animation-duration: 1s;
  animation-name: slideInRight-enter;
  animation-iteration-count: infinite;
}

@keyframes slideInRight-enter {
  from {
    margin-left: 0;
  }
  25% {
    width: 50%;
  }
  to {
    margin-left: 100%;
  }
}


router-link {
  text-decoration: underline;
}

router-link:hover {
  text-decoration: none;
}
</style>