<template>
  <el-row class="el-row">
    <el-col :span="23">
      <div class="grid-content ep-bg-purple"></div>
    </el-col>
    <el-col :span="1">
      <div class="grid-content ep-bg-purple-light">
        <div class="demo-basic--circle">
          <el-popover
              placement="top-start"
              width="10vw"
              trigger="hover"
          >
            <template #reference>
              <div class="block head-pic">
                <el-avatar :size="50" :src="user.data.avatar"/>
              </div>
            </template>
            <div class="name">{{ user.data.name }}</div>

            <el-progress type="line"
                         :percentage="percentage"
                         :stroke-width="10">
              <template #default>
                <div style="font-size: 1.2rem">{{ used }} / {{ capacity }}</div>
              </template>
            </el-progress>

            <el-button-group style="margin-top: 10px;display:flex; justify-content: center">
              <el-button type="warning" plain size="small" @click="toFileFolder">
                <el-icon>
                  <Place/>
                </el-icon>&nbsp;个人信息
              </el-button>
              <el-button type="warning" plain size="small" @click="baseStore.clearToken">
                退出登录&nbsp;<el-icon>
                <Right/>
              </el-icon>
              </el-button>
            </el-button-group>
          </el-popover>

        </div>
      </div>
    </el-col>
  </el-row>
</template>

<script lang="ts" setup>
import {onMounted, reactive, ref} from "vue";
import {useBaseStore} from "@/store";
import {
  Place, Right
} from "@element-plus/icons-vue";
import {formatSize} from "@/utils/util.ts";

const baseStore = useBaseStore()

let user = reactive({
      data:
          {id: 0, username: '', name: '', avatar: '', email: '', signature: '', used: 0, capacity: 0, status: 0}
    }),
    percentage = ref(0),
    status = ref('success'),
    used = ref(''),
    capacity = ref('')

async function showUserInfo() {
  if (user.data.id === 0) {
    user.data = await useBaseStore().getUserInfo()
    percentage.value = Math.floor((user.data.used / user.data.capacity) * 100)
    used.value = formatSize(user.data.used)
    capacity.value = formatSize(user.data.capacity)
    if (percentage.value > 90) {
      status.value = 'exception'
    }
    console.log(user.data, percentage.value)
  }
}

function toFileFolder() {
  window.location.href = '/info/user'
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

.name {
  margin-bottom: 5px;
  color: #213547;
  font-family: Arial, sans-serif;
  font-size: 1.8rem;
  font-weight: 700;
}
</style>
