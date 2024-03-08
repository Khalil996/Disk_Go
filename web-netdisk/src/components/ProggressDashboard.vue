<template>
  <el-progress :type="type" :percentage="80" :status="status" :stroke-width="15">
    <!--    当status 是 exception 时，是红色，success是绿色-->
    <template #default="{ percentage }">
      <p class="percentage-value" v-bind:style="{color:fontColor}">{{ percentage }}%</p>
      <p class="percentage-label">302402304</p>
    </template>
  </el-progress>
</template>

<script setup>
let type = 'dashboard'
let percentage = 30
let status = 'success'
let fontColor = 'lightseagreen'
</script>

<style lang="css" scoped>
.percentage-label{
  color: #213547;
  font-family: Arial,sans-serif;
  font-size: 12px;
  font-weight: 700;
}
.percentage-value{
  font: 1000 22px Arial,sans-serif;
}
</style>