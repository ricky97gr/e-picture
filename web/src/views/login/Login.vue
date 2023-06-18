<template>
  <div class="login-all">
    <div class="login-container">
      <div class="login-container-left"></div>
      <div class="login-container-right">
        <div class="login-title">
          <h3>用户登录</h3>
        </div>
        <div class="login-form-container">
          <el-form :model="loginInfo" class="login-form" :rules=validator.Login>
            <el-form-item prop="phone">
              <el-input v-model="loginInfo.phone" placeholder="请输入身份信息" prefix-icon="User">
              </el-input>
            </el-form-item>
            <el-form-item prop="passwd">
              <el-input prefix-icon="lock" v-model="loginInfo.passwd" placeholder="请输入密码" show-password>
              </el-input>
            </el-form-item>

            <el-form-item>
              <el-button class="login-button" @click="login" type="primary">登录</el-button>
            </el-form-item>
            <!-- <RouterLink to="/register">注册</RouterLink>
            <RouterLink to="/login" class="forget-passwd"
              >忘记密码？</RouterLink> -->
          </el-form>
        </div>
      </div>
    </div>
  </div>
</template>


<script lang="ts" setup>
import { reactive } from "vue";
import { RouterLink } from "vue-router";
import apiUrl from "../../apis/url";
import router from "../../router";
import errorHandle from "../../util/errHandle";
import httprequest from "../../util/http";
import successHandle from "../../util/successHandle";
import validator from "../../util/validator/validator"

const loginInfo = reactive({
  phone: "",
  passwd: "",
});


function login() {
  httprequest(apiUrl.PostLogin, {
    method: "POST",
    body: {
      phone: loginInfo.phone,
      passwd: loginInfo.passwd
    }
  }).then(data => {
    if (errorHandle(data, "")) {
      return
    }
    sessionStorage.setItem("token", data.token)
    router.push("/home")
    successHandle("登陆成功")
  }
  )
}

function register() {
  router.push("/register");
}
</script>

<style scoped>
.login-all {
  width: 100%;
  height: 100%;
  /* background-image: url('../../assets/login.jpg'); */
  /* background-repeat: repeat-x; */
  /* padding-top: 100px; */
  background-size: cover;
  /* padding-top: 140px; */
  background-position: center;
  /* background-image: url("../../assets/login.jpg"); */
  /* padding-left: 200px; */
}

.login-container {
  width: 1000px;
  height: 600px;
  margin: 120px 0 0 400px;
  /* 上右下左对外边框元素的距离 */
  /* 内容和边框之间的距离，例如表格和外面class之间的距离 */
  /* padding: 35px 10px 35px 35px; */
  /* 背景颜色 */
  background-color: rgba(255, 255, 255, 0.5);
  /* 向外发散15px 当作阴影 x,y的偏移量为0 */
  box-shadow: 0 0 10px rgb(255, 255, 255);
  /* 半径30px画圆 */
  border-radius: 5px;
  /* display: flex; */
  /* margin: auto; */
}

.login-container-left {
  width: 500px;
  height: 100%;
  background-color: white;
  float: left;
  background-image: url("../../assets/login-left.png");
  background-position: cover;
}

.login-container-right {
  width: 500px;
  height: 100%;
  /* background-color: green; */
  float: right;
}

.login-title {
  margin-top: 50px;
  height: 80px;
  text-align: center;
}

.login-form-container {
  width: 300px;
  margin-top: 80px;
  margin-left: 100px;
}

.login-input {
  margin-top: 40px;
}

.login-button {
  width: 100%;
  margin-top: 40px;
  background-color: rgb(48, 146, 143);
}

.forget-passwd {
  float: right;
}
</style>