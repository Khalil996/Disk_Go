<template>
    <div class="bg">
        <div class="logo-icon">
            <img src="@/assets/img.png" alt=""/>
            <span>咪咪网盘</span>
        </div>
        <img src="@/assets/background_boy.png" alt=""/>
        <div class="form-box">
            <div class="banner-header">
                <button class="active" @click="showBannerHeader(true)">登录</button>
                <button @click="showBannerHeader(false)">注册</button>
            </div>
            <div v-show="flag">
                <login-form></login-form>
            </div>
            <div v-show="!flag">
                <register-form></register-form>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>
import LoginForm from '../components/LoginForm.vue'
import RegisterForm from "../components/RegisterForm.vue"
import {ref} from "vue";

let flag = ref<boolean>(true)

const change = (num: number) => {
    flag.value = num === 0
}

function showBannerHeader(bo) {
    let sibling = document.getElementsByClassName('banner-header')[0].children
    sibling[0].classList.remove('active')
    sibling[1].classList.remove('active')
    if (bo) {
        sibling[0].classList.add('active')
    } else {
        sibling[1].classList.add('active')
    }
    flag.value = bo
}

</script>

<style>
.logo-icon {
    position: absolute;
    left: 10px;
    top: 10px;
    display: flex;
}

.logo-icon img {
    width: 40px;
    margin-right: 5px;
}

.logo-icon span {
    font-size: 28px;
    font-weight: 700;
    height: 45px;
    line-height: 50px;
    color: orange;
}

.bg {
    display: flex;
    justify-content: center;
    align-items: center;
    background: url('@/assets/background_grid.png');
    width: 100%;
    height: 98vh;
    overflow: hidden;
}

.form-box {
    background-color: #e3f1ff;
    width: 50vh;
    border-radius: 20px;
    box-shadow: 5px 3px 3px #e5e5e5;
    padding: 2rem;
}

.text {
    margin: 2rem 3rem;
}

.big-text {
    font-size: 5rem;
    margin-right: 3rem;
    color: #787eef;
}

.banner-header {
    width: 100%;
    border-bottom: 3px solid lightseagreen;
    margin-bottom: 4%;
}

.banner-header button {
    margin-bottom: 2px;
    color: slategrey;
    font-size: 22px;
    font-weight: 700;
    margin-left: 3%;
    border: none;
    background-color: rgb(227, 241, 255);
}

.banner-header button:hover {
    background-color: lightblue;
}

.banner-header button.active {
    color: lightcoral;
    font-size: 39px;
}


</style>
