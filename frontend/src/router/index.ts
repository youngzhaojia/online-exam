import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";
import Layout from "@/layout/index.vue"

const routes: Array<RouteRecordRaw> = [
  {
    path: '/user/login',
    component: import("@/views/user/login.vue"),
  },
  {
    path: '/user/register',
    component: import("@/views/user/register.vue"),
  },
  {
    path: '/',
    component: Layout,
    redirect: 'index',
    children: [
      {
        path: 'index',
        component: () => import('@/views/index.vue'),
      },{
      path: 'question/list',
      component: () => import('@/views/question/list.vue'),
    }]
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
});

export default router;
