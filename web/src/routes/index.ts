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

// navigation guards
// router.beforeEach(async (to, from, next) => {
//   // set the fallback locale to the browser locale, if that fails, use english as the final fallback

//   let useLocale = ""

//   if (typeof to.params?.locale === "string" && to.params.locale !== "") {
//     // Use the locale mentioned in the URL first
//     useLocale = to.params.locale
//   } else {
//     // If there is no locale in the URL, check with the user's settings
//   }
//   if (useLocale === "") {
//     // if everything else fails, use the fallback locale we set at the beginning
//     useLocale = defaultLocale
//   }

//   // use locale if paramsLocale is not in SUPPORT_LOCALES
//   // if (paramsLocale === "" || !SUPPORT_LOCALES.includes(paramsLocale)) {
//   //   return next(`/${locale}`)
//   // }

//   console.log("active locale: " + useLocale)
//   console.log("available locales: ", i18n.global.availableLocales)
//   console.log("fallback: ", i18n.global.fallbackLocale)

//   // load locale messages
//   if (!i18n.global.availableLocales.includes(useLocale)) {
//     await loadLocaleMessages(i18n, useLocale)
//   }

//   // set i18n language
//   setI18nLanguage(i18n, useLocale)

//   return next()
// })

router.afterEach((to) => {
  /* Set document title from route meta */
  if (to.meta && to.meta.title) {
    document.title = `${to.meta.title} - ${defaultDocumentTitle}`
  } else {
    document.title = defaultDocumentTitle
  }
})

export default router
