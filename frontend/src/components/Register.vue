<template>
  <el-form :model="loginForm" ref="loginForm" :rules="loginRules" class="register-container">
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

    <el-form-item prop="passwordConfirm">
      <el-input
        v-model="loginForm.passwordConfirm"
        name="password"
        placeholder="确认密码"
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
          style="width: auto; height: 41px"
          :src="captchaurl"
          fit="contain"
          @click="getCaptcha()"
        ></el-image>
      </el-col>
    </el-row>

    <el-form-item>
      <el-button type="primary" @click="onRegister()">注册</el-button>
      <el-button @click="$emit('switch','login')">登录</el-button>
    </el-form-item>
  </el-form>
</template>


<script>
import { requestRegister, requestCaptcha } from "../utils";
export default {
  name: "register",
  data() {
    const validatePassConfirm = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("请再次输入密码"));
      } else if (value !== this.loginForm.password) {
        callback(new Error("两次输入密码不一致!"));
      } else {
        callback();
      }
    };
    return {
      loginForm: {
        username: "",
        password: "",
        passwordConfirm: ""
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
          },
          {
            pattern: "^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{8,}$",
            message: "密码需要为大于8位的密码数字组合",
            trigger: "blur"
          }
        ],
        passwordConfirm: [
          {
            validator: validatePassConfirm,
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
    onRegister() {
      this.$refs.loginForm.validate(valid => {
        if (valid) {
          requestRegister(
            this.loginForm.username,
            this.loginForm.password,
            this.loginForm.code,
            this.captchaid
          )
            .then(res => {
              if (res.data.code == 0) {
                this.$message({
                  message: "注册成功",
                  type: "success"
                });
              } else if (res.data.code == -2) {
                this.$message({
                  message: "验证码错误",
                  type: "error"
                });
              } else {
                this.$message({
                  message: "注册失败，用户已存在",
                  type: "error"
                });
              }
              this.getCaptcha();
            })
            .catch(err => {
              this.$message({
                message: "注册失败，数据库错误",
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
      });
    }
  }
};
</script>
<style lang="scss" scoped>
.register-container {
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