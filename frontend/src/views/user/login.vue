<template>
  <div class="login-container">
    <el-form ref="loginForm"
             :model="loginForm"
             class="login-form"
             auto-complete="on"
             label-position="right"
             label-width="80px">

      <div class="title-container">
        <h3 class="title">登录</h3>
      </div>

      <el-form-item prop="mobile"
                    label="手机号:">
        <el-input ref="mobile"
                  v-model="loginForm.mobile"
                  placeholder="请输入手机号"
                  name="mobile"
                  type="text"
                  tabindex="1"
                  auto-complete="on" />
      </el-form-item>

      <el-form-item prop="password"
                    label="密码:">
        <el-input ref="password"
                  v-model="loginForm.password"
                  placeholder="请输入密码"
                  name="password"
                  type="password"
                  tabindex="2"
                  auto-complete="on"
                  @keyup.enter="handleLogin" />
      </el-form-item>

      <el-form-item>
        <el-button :loading="loading"
                   type="primary"
                   style="width:45%;"
                   @click.prevent="handleLogin">登录</el-button>

        <el-button type="default"
                   style="width:45%;float:right;"
                   @click.prevent="$router.push('/user/register')">没有账号，去注册</el-button>

      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { login } from "@/api/user";
import { setToken } from "@/utils/auth";
import { ElMessage } from "element-plus";

export default {
  name: "Login",
  data() {
    return {
      loading: false,
      loginForm: {
        mobile: "",
        password: "",
      },
    };
  },
  methods: {
    handleLogin() {
      this.loading = true;
      login(this.loginForm)
        .then((resp) => {
          const { ret, msg, data } = resp;
          console.log(ret);
          if (ret) {
            ElMessage.error(msg);
            return;
          }
          ElMessage.success("登录成功");
          setToken(data.token);
          this.$router.push({ path: "/" });
        })
        .finally(() => {
          this.loading = false;
        });
    },
  },
};
</script>

<style lang="scss" scoped>
$bg: #2d3a4b;
$dark_gray: #889aa4;
$light_gray: #eee;

.login-container {
  min-height: 100vh;
  min-width: 100vh;
  background-color: $bg;
  overflow: hidden;

  .login-form {
    position: relative;
    width: 520px;
    max-width: 100%;
    padding: 160px 35px 0;
    margin: 0 auto;
    overflow: hidden;

    ::v-deep .el-form-item__label {
      color: #ffffff;
    }
  }

  .svg-container {
    padding: 6px 5px 6px 15px;
    color: $dark_gray;
    vertical-align: middle;
    width: 30px;
    display: inline-block;
  }

  .title-container {
    position: relative;

    .title {
      font-size: 26px;
      color: $light_gray;
      margin: 0px auto 40px auto;
      text-align: center;
      font-weight: bold;
    }
  }
}
</style>
