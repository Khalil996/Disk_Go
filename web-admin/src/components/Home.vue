<template>
  <div>
    <p>{{ msg }}</p>
    <p>姓名：{{ user.name }}</p>
    <p>年龄：{{ user.age }}</p>
    <button @click="changeAge">点击</button>
    <p>单独姓名：{{ name }}</p>
    <button @click="changeName">修改姓名</button>
    <p>单独年龄：{{ age }}</p>
    <button @click="changeRefAge">修改年龄</button>
    <p>诗人：{{ poet }}</p>
    <button @click="changePoet">修改诗人</button>
    <p>computed:{{ newPoet }}</p>
    <!-- 父子传值 -->
    <Son :user="user" />
  </div>
</template>
<script >
import { reactive, ref, toRef, toRefs } from "@vue/reactivity";
import { computed, watch, provide } from "vue";
import Son from "./Son.vue";
export default {
  components: {
    Son,
  },
  setup() {
    let msg = "hello";
    // reactive是一个函数，它可以将复杂数据类型变成响应式数据。
    let user = reactive({
      name: "张无忌",
      age: 12,
    });
    const changeAge = () => {
      user.age++;
    };
    // toRef是函数，将对象中某个属性为单独的响应式数据
    const name = toRef(user, "name");
    const changeName = () => {
      name.value = "赵敏";
    };
    // toRefs是函数，将对象中所有属性为单独响应式数据
    const refUser = toRefs(user);
    const changeRefAge = () => {
      refUser.age.value = 18;
    };
    //   ref函数，常用于简单数据类型定义为响应式数据，再修改值，获取值的时候，需要.value
    let poet = ref("李白");
    const changePoet = () => {
      poet.value = "杜甫";
    };
    // computed 计算
    const newPoet = computed(() => {
      return poet.value + ",白居易";
    });
    // watch 监听
    watch(user, () => {
      console.log("user改变了");
    });
    watch(
      () => user.age,
      () => {
        console.log("age改变了");
      }
    );
    // provide
    provide("poet", poet);
    provide("changePoet", changePoet);
    return {
      msg,
      user,
      name,
      changeAge,
      changeName,
      changeRefAge,
      ...refUser,
      poet,
      changePoet,
      newPoet,
    };
  },
};
</script>
