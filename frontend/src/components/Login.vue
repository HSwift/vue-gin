<template>
  <el-form :model="loginForm" ref="loginForm" :rules="loginRules" class="login-container">
    <h2>网络安全原理</h2>
    <el-form-item prop="username">
      <el-input v-model="loginForm.username" name="username" placeholder="用户名" auto-complete="on"></el-input>
    </el-form-item>

    <el-form-item prop="password">
      <el-input
        v-model="loginForm.password"
        name="password"
        placeholder="密码"
        auto-complete="off"
        show-password
      ></el-input>
    </el-form-item>

    <el-row>
      <el-col :span="12">
        <el-form-item prop="code">
          <el-input v-model="loginForm.code" name="code" placeholder="验证码" auto-complete="off"></el-input>
        </el-form-item>
      </el-col>
      <el-col :span="12">
        <el-image
          v-model="loginForm.captcha"
          style="width: auto; height: 41px"
          :src="captchaurl"
          fit="contain"
          @click="getCaptcha()"
        ></el-image>
      </el-col>
    </el-row>

    <el-form-item>
      <el-button type="primary" @click="onLogin()">登录</el-button>
      <el-button @click="$emit('switch','register')">注册</el-button>
    </el-form-item>
  </el-form>
</template>


<script>
import { requestLogin, requestCaptcha } from "../utils";
export default {
  name: "login",
  data() {
    return {
      loginForm: {
        username: "",
        password: "",
        code: ""
      },
      captchaid: "",
      captchaurl: "",
      loginRules: {
        username: [
          {
            required: true,
            message: "请输入用户名",
            trigger: "blur"
          }
        ],
        password: [
          {
            required: true,
            message: "请输入密码",
            trigger: "blur"
          }
        ],
        code: [
          {
            required: true,
            message: "请验证码",
            trigger: "blur"
          }
        ]
      }
    };
  },
  created: function() {
    this.getCaptcha();
  },
  methods: {
    onLogin() {
      this.$refs.loginForm.validate(valid => {
        if (valid) {
          requestLogin(
            this.loginForm.username,
            this.loginForm.password,
            this.loginForm.code,
            this.captchaid
          )
            .then(res => {
              if (res.data.code == 0) {
                this.$message({
                  message: "登录成功",
                  type: "success"
                });
              } else if (res.data.code == -2) {
                this.$message({
                  message: "验证码错误",
                  type: "error"
                });
              } else {
                this.$message({
                  message: "登录失败，用户名或密码错误",
                  type: "error"
                });
              }
              this.getCaptcha();
            })
            .catch(err => {
              this.$message({
                message: "登录失败，数据库错误",
                type: "error"
              });
              window.console.log(err);
            });
        }
      });
    },
    getCaptcha() {
      requestCaptcha().then(res => {
        this.captchaid = res.data.captchaid;
        this.captchaurl = "/captcha/" + this.captchaid;
        this.loginForm.code = "";
      });
    }
  }
};
</script>
<style lang="scss" scoped>
.login-container {
  -webkit-border-radius: 5px;
  border-radius: 5px;
  -moz-border-radius: 5px;
  background-clip: padding-box;
  margin: auto auto;
  width: 350px;
  padding: 35px 35px 15px 35px;
  border: 1px solid #eaeaea;
  box-shadow: 4px 4px 12px #cac6c6;
}
</style>