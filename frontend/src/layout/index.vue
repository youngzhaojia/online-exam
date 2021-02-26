<template>
  <div class="app-container">
    <div class="app-menu">
      <div class="user-info">
        132****0000
      </div>
      <el-menu :uniqueOpened="true"
               :default-active="defaultActive"
               class="el-menu-vertical-demo"
               background-color="#545c64"
               text-color="#fff"
               active-text-color="#ffd04b"
               @select="handleSelect">
        <el-menu-item v-for="(menu,index) in menuList"
                      :index="index+''"
                      :key="index">
          <i :class="menu.icon"></i>
          <template #title>{{ menu.name }}</template>
        </el-menu-item>
      </el-menu>
    </div>

    <!-- 页面内容 -->
    <div class="app-body">
      <router-view :key="$router.path" />
    </div>
  </div>
</template>

<script>
export default {
  name: "Layout",
  components: {},
  props: {},
  computed: {
    defaultActive() {
      return this.menuList
        .findIndex((item) => item.url === this.$route.path)
        .toString();
    },
  },
  data() {
    return {
      menuList: [
        {
          url: "/index",
          icon: "el-icon-s-home",
          name: "首页",
        },
        {
          url: "/question/list",
          icon: "el-icon-s-operation",
          name: "题目管理",
        },
        {
          url: "/exam/list",
          icon: "el-icon-s-order",
          name: "试卷管理",
        },
        {
          url: "/answer/list",
          icon: "el-icon-s-ticket",
          name: "答题管理",
        },
        {
          url: "/logout",
          icon: "el-icon-setting",
          name: "退出",
        },
      ],
    };
  },
  watch: {},
  created() {},
  mounted() {},
  methods: {
    handleSelect(index) {
      this.$router.push(this.menuList[index].url);
    },
  },
};
</script>

<style lang="scss" scoped>
.app-container {
  min-height: 100vh;
  width: 100vh;
  .app-menu {
    float: left;
    width: 150px;

    .user-info {
      width: 149px;
      height: 50px;
      background: #545c64;
      color: #fff;
      display: flex;
      align-items: center;
      justify-content: center;
    }

    .el-menu {
      min-height: 100vh;
    }
  }

  .app-body {
    float: left;
    width: calc(100% - 150px);
    box-sizing: border-box;
    padding: 10px;
  }
}
</style>
