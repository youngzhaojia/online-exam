import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import Layout from "@/layout/index.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/user/login",
    component: import("@/views/user/login.vue"),
  },
  {
    path: "/user/register",
    component: import("@/views/user/register.vue"),
  },
  {
    path: "/",
    component: Layout,
    redirect: "index",
    children: [
      {
        path: "index",
        component: () => import("@/views/index.vue"),
      },
      {
        path: "question/list",
        component: () => import("@/views/question/list.vue"),
      },
      {
        path: "question/edit",
        component: () => import("@/views/question/edit.vue"),
      },
      {
        path: "exam/list",
        component: () => import("@/views/exam/list.vue"),
      },
      {
        path: "exam/edit",
        component: () => import("@/views/exam/edit.vue"),
      },
      {
        path: "answer/list",
        component: () => import("@/views/answer/list.vue"),
      },
      {
        path: "answer/detail",
        component: () => import("@/views/answer/detail.vue"),
      },
    ],
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
