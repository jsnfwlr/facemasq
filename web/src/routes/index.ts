import { createRouter, createWebHistory } from "vue-router"

import Home from "@/pages/dashboard.vue"
import DevicePage from "@/pages/devices.vue"
import ParamPage from "@/pages/params.vue"
import InfoPage from "@/pages/info.vue"

// import { setupI18n, setI18nLanguage, loadLocaleMessages } from "@/i18n"
// const defaultLocale = /(\w+)[\w-]*/.test(navigator.language) ? RegExp.$1 : "en"
// const i18n = setupI18n({
//   legacy: false,
//   locale:
//     fallbackLocale: defaultLocale,
// })

const routes = [
  {
    meta: {
      title: "Dashboard",
      layout: "Standard",
    },
    path: "/:locale*",
    alias: "/",
    name: "home",
    component: Home,
  },
  {
    meta: {
      title: "Devices",
      layout: "Standard",
    },
    path: "/:locale*/devices",
    name: "Devices",
    component: DevicePage,
  },
  {
    meta: {
      title: "Manage params",
      layout: "Standard",
    },
    path: "/:locale*/manage/:param",
    name: "Manage",
    component: ParamPage,
  },
  {
    meta: {
      title: "Manage params",
      layout: "Standard",
    },
    path: "/:locale*/admin/:param",
    name: "Admin",
    component: ParamPage,
  },
  {
    meta: {
      title: "Information",
      layout: "Standard",
    },
    path: "/:locale*/about/info",
    name: "Info",
    component: InfoPage,
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    return savedPosition || { top: 0 }
  },
  linkActiveClass: "active",
  linkExactActiveClass: "exact",
})

const defaultDocumentTitle = "faceMasq"

router.afterEach((to) => {
  /* Set document title from route meta */
  if (to.meta && to.meta.title) {
    document.title = `${to.meta.title} - ${defaultDocumentTitle}`
  } else {
    document.title = defaultDocumentTitle
  }
})

export default router
