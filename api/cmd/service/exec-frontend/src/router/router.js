import AppLayout from "../components/AppLayout.vue";
import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/creator",
      component: AppLayout,
      children: [
        {
          path: "",
          name: "dashboard",
          component: () => import("../views/About.vue"),
        },
        {
          path: "/qc",
          name: "qc",
          component: () => import("../views/QcQuestion.vue"),
        },
      ],
    },
    {
      path: "/examViewer",
      name: "examviewer",
      component: () => import("../components/EditorLayout.vue"),
    },
    {
      path: "/",
      name: "landing",
      component: () => import("../views/Landing.vue"),
    },
    {
      path: "/leaderboard",
      name: "leaderboard",
      component: () => import("../views/Leaderboard.vue"),
    },
    {
      path: "/login",
      name: "login",
      component: () => import("../views/login.vue"),
    },
    {
      path: "/access-denied",
      name: "access denied",
      component: () => import("../views/auth.vue"),
    },
    // {
    //     path: '/pages/notfound',
    //     name: 'notfound',
    //     component: () => import('@/views/pages/NotFound.vue')
    // },

    // {
    //     path: '/auth/login',
    //     name: 'login',
    //     component: () => import('@/views/pages/auth/Login.vue')
    // },
    // {
    //     path: '/auth/access',
    //     name: 'accessDenied',
    //     component: () => import('@/views/pages/auth/Access.vue')
    // },
    // {
    //     path: '/auth/error',
    //     name: 'error',
    //     component: () => import('@/views/pages/auth/Error.vue')
    // }
  ],
});

export default router;
