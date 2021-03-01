<template>
  <div class="register-container">
    <el-form ref="registerForm"
             :model="registerForm"
             class="register-form"
             auto-complete="on"
             label-position="right"
             label-width="80px">

      <div class="title-container">
        <h3 class="title">注册</h3>
      </div>

      <el-form-item prop="name"
                    label="姓名:">
        <el-input ref="name"
                  v-model="registerForm.name"
                  placeholder="请输入姓名"
                  name="name"
                  type="text"
                  tabindex="1"
                  auto-complete="on" />
      </el-form-item>

      <el-form-item prop="mobile"
                    label="手机号:">
        <el-input ref="mobile"
                  v-model="registerForm.mobile"
                  placeholder="请输入手机号"
                  name="mobile"
                  type="text"
                  tabindex="1"
                  auto-complete="on" />
      </el-form-item>

      <el-form-item prop="password"
                    label="密码:">
        <el-input ref="password"
                  v-model="registerForm.password"
                  placeholder="请输入密码"
                  name="password"
                  type="password"
                  tabindex="2"
                  auto-complete="on"
                  @keyup.enter="handleRegister" />
      </el-form-item>

      <el-form-item prop="password2"
                    label="密码确认:">
        <el-input ref="password"
                  v-model="registerForm.password2"
                  placeholder="请再次输入密码"
                  name="password"
                  type="password"
                  tabindex="2"
                  auto-complete="on"
                  @keyup.enter="handleRegister" />
      </el-form-item>

      <el-form-item>
        <el-button :loading="loading"
                   type="primary"
                   style="width:45%;"
                   @click.prevent="handleRegister">确认注册</el-button>

        <el-button type="default"
                   style="width:45%;float: right;"
                   @click.prevent="$router.push('/user/login')">已有账号，去登录</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { register } from "@/api/user";
import { ElMessage } from "element-plus";

export default {
  name: "Register",
  data() {
    return {
      loading: false,
      registerForm: {
        name: "",
        mobile: "",
        password: "",
        password2: "",
      },
    };
  },
  methods: {
    handleRegister() {
      if (
        !this.registerForm.name ||
        !this.registerForm.mobile ||
        !this.registerForm.password
      ) {
        ElMessage.error("信息不能为空");
        return;
      }
      if (this.registerForm.password !== this.registerForm.password2) {
        ElMessage.error("两次输入的密码不一致");
        return;
      }
      this.loading = true;
      register(this.registerForm)
        .then((resp) => {
          const { ret, msg, data } = resp;
          if (ret) {
            ElMessage.error(msg);
            return;
          }
          ElMessage.success("注册成功，去登录");
          this.$router.push({ path: "/user/login" });
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

.register-container {
  min-height: 100vh;
  min-width: 100vh;
  background-color: $bg;
  overflow: hidden;

  .register-form {
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
