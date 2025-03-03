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
          name: "create",
          component: () => import("../views/About.vue"),
        },
        {
          path: "/qc",
          name: "qc",
          component: () => import("../views/QcQuestion.vue"),
        },
        {
          path: "/setexam",
          name: "setexam",
          component: () => import("../views/CreateExam.vue"),
        },
      ],
    },
    {
      path: "/challengeViewer",
      name: "challengeViewer",
      component: () => import("../components/EditorLayout.vue"),
    },
    {
      path: "/examViewer",
      name: "examviewer",
      component: () => import("../views/examViewer.vue"),
    },
    {
      path: "/",
      name: "landing",
      component: () => import("../views/Landing.vue"),
    },
    {
      path: "/menu",
      name: "menu",
      component: () => import("../views/CardMenu.vue"),
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
    {
      path: "/editor",
      name: "editor",
      component: () => import("../views/Editor.vue"),
    },
    // {
    //     path: '/pages/notfound',
    //     name: 'notfound',
    //     component: () => import('../views/')
    // },

    {
        path: '/auth/login',
        name: 'login',
        component: () => import('../views/login.vue')
    },
    {
        path: '/auth/signup',
        name: 'login',
        component: () => import('../views/signup.vue')
    },
    {
        path: '/test',
        name: 'test',
        component: () => import('../views/QuestionTypes.vue')
    },
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
