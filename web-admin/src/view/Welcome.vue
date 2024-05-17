<template>
  <el-card>
    <div style="display: flex; justify-content: space-around; align-items: flex-start">
      <el-card>
        <template #header>
          <div class="header">用户数量：</div>
        </template>
        <div class="user-num">
          {{ data.cnt.user }} 名
        </div>
      </el-card>


      <el-card>
        <template #header>
          <div class="header">系统容量：</div>
        </template>
        <div style="position: relative">
          <span>{{ text.used }} / {{ text.capacity }}</span>
          <span class="to-right">已用 {{ text.capPercentage }}</span>
        </div>
        <div class="cap-pie pie"></div>
      </el-card>

      <el-card>
        <template #header>
          <div class="header">用户文件：</div>
        </template>
        <div style="position: relative">
          <span>文件数：{{ data.cnt.file.file }}</span>
          <span style="margin-left: 50px">实际文件数：{{ data.cnt.file.fs }}</span>
          <span style="margin-left: 50px">复用率：{{ text.filePercentage }}</span>
        </div>
        <div style="margin: 10px 0">
          <span>文件夹数：{{ data.cnt.file.folder }}</span>
        </div>
        <div>类型分布：</div>
        <div class="file-pie pie"></div>
      </el-card>

      <el-card>
        <template #header>
          <div class="header">分享统计：</div>
        </template>
        <div style="position: relative">
          <span>下载总计：{{ data.cnt.download }} 次</span>
          <span style="margin-left: 50px">浏览总计：{{ data.cnt.click }} 次</span>
        </div>
      </el-card>
    </div>
  </el-card>
</template>

<script lang="js" setup>
import * as echarts from "echarts"
import {onBeforeMount, onMounted, reactive} from "vue";
import http from "../utils/http/http.js";
import {codeOk} from "../utils/http/base.js";
import {formatSize} from "../utils/util.js";
import {typeMap} from "../utils/constant.js";

const data = reactive({
      'cnt': {
        'user': 0,
        'download': 0,
        'click': 0,
        'file': {
          'file': 0,
          'type': [],
          'fs': 0,
          'folder': 0,
        }
      },
      'cap': {
        'used': 0,
        'capacity': 0
      }
    }),
    text = reactive({
      'used': '',
      'capacity': '',
      'capPercentage': '',
      'filePercentage': ''
    })

async function getStatisticData() {
  const resp = await http.get("/admin/statistic")
  if (resp.data.code === codeOk) {
    Object.assign(data, resp.data.data)
    format()
  }
}

function format() {
  data.cap.capacity = Number.parseInt(data.cap.capacity, 0)
  text.used = formatSize(data.cap.used)
  text.capacity = formatSize(data.cap.capacity)
  text.capPercentage = '< 0.001%'
  text.filePercentage = '< 0.0001%'

  let percentage = data.cap.used / data.cap.capacity
  if (percentage > 0.001) {
    text.capPercentage = percentage*100 + '%'
  }

  percentage = (data.cnt.file.file - data.cnt.file.fs) / data.cnt.file.file
  if (percentage > 0.0001) {
    text.filePercentage = percentage*100 + '%'
  }
}

onBeforeMount(async () => {
  await getStatisticData()

  initCapPie()
  initFilePie()
})

function initCapPie() {
  const capPie = echarts.init(document.querySelector('.cap-pie'))
  let option = {
    series: [
      {
        type: 'pie',
        data: [
          {
            value: data.cap.used,
            name: '已用容量',
          },
          {
            value: data.cap.capacity,
            name: '总容量',
          }
        ],
        radius: [100, 60],
      }
    ]
  }
  capPie.setOption(option)
}

function initFilePie() {
  const filePie = echarts.init(document.querySelector('.file-pie'))
  const pieData = []
  data.cnt.file.type.forEach(cnt => {
    pieData.push({
      value: cnt.cnt,
      name: typeMap[cnt.type]
    })
  })
  console.log(pieData)
  let option = {
    series: [
      {
        type: 'pie',
        data: pieData,
      }
    ]
  }
  filePie.setOption(option)
}
</script>

<style scoped>
.header {
  font-size: 1.3rem;
  font-weight: 700;
}

.pie {
  width: 400px;
  height: 400px;
}

.to-right {
  position: absolute;
  right: 0;
}
</style>