import { createApp } from "vue";
import ElementPlus from "element-plus";
import locale from "element-plus/lib/locale/lang/zh-cn";
import "element-plus/lib/theme-chalk/index.css";

import Vant from "vant";
import "vant/lib/index.css";

import App from "./App.vue";
import "./registerServiceWorker";
import router from "./router";
import store from "./store";

createApp(App)
  .use(store)
  .use(router)
  .use(ElementPlus, { locale })
  .use(Vant)
  .mount("#app");
