<template>
  <div class="login-container">
    <el-form ref="loginForm"
             :model="loginForm"
             class="login-form"
             auto-complete="on"
             label-position="right"
             label-width="80px">

      <div class="title-container">
        <h3 class="title">Login Form</h3>
      </div>

      <el-form-item prop="mobile"
                    label="手机号">
        <el-input ref="mobile"
                  v-model="loginForm.mobile"
                  placeholder="mobile"
                  name="mobile"
                  type="text"
                  tabindex="1"
                  auto-complete="on" />
      </el-form-item>

      <el-form-item prop="password"
                    label="密码">
        <el-input ref="password"
                  v-model="loginForm.password"
                  placeholder="密码"
                  name="password"
                  tabindex="2"
                  auto-complete="on"
                  @keyup.enter="handleLogin" />
      </el-form-item>

      <el-form-item>
        <el-button :loading="loading"
                 type="primary"
                 style="width:100%;"
                 @click.prevent="handleLogin">Login</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  name: "Login",
  data() {
    return {
      loginForm: {
        mobile: "13100000002",
        password: "123456",
      },
      loading: false,
      passwordType: "password",
      redirect: undefined,
    };
  },
  watch: {
    $route: {
      handler: function (route) {
        this.redirect = route.query && route.query.redirect;
      },
      immediate: true,
    },
  },
  methods: {
    handleLogin() {
      this.$refs.loginForm.validate((valid) => {
        if (valid) {
          this.loading = true;
          this.$store
            .dispatch("user/login", this.loginForm)
            .then(() => {
              this.$router.push({ path: this.redirect || "/" });
              this.loading = false;
            })
            .catch(() => {
              this.loading = false;
            });
        } else {
          console.log("error submit!!");
          return false;
        }
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
