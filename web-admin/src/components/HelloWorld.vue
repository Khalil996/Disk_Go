<template>
  <div>
    {{ name }}
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue';
import userApi from '../api/user'
const name = ref("xiaopang");
  onMounted(() => {
    userApi.getUserList().then(res=>{
      console.log('------->');
      console.log(res);
    })
  })
</script>

