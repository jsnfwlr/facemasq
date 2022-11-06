import { createRouter, createWebHistory } from "vue-router"

import Home from "@/pages/dashboard.vue"
import DevicePage from "@/pages/devices.vue"
import ParamPage from "@/pages/params.vue"
import InfoPage from "@/pages/info.vue"

import { setupI18n, setI18nLanguage, loadLocaleMessages } from "@/i18n"
import en from "@/i18n/facemasq.en.json"
import fr from "@/i18n/facemasq.fr.json"

const i18n = setupI18n({
  locale: "en",
  messages: {
    en: en,
    fr: fr,
  },
})

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
router.beforeEach(async (to, from, next) => {
  const defaultLocale = "en"
  let useLocale = ""
  if (typeof to.params?.locale === "string" && to.params.locale !== "") {
    useLocale = to.params.locale
  } else if (typeof to.params?.locale === "object" && Array.isArray(to.params.locale)) {
    useLocale = to.params?.locale[0]
  } else {
    useLocale = defaultLocale
  }
  console.log(to.params, useLocale)

  // use locale if paramsLocale is not in SUPPORT_LOCALES
  // if (paramsLocale === "" || !SUPPORT_LOCALES.includes(paramsLocale)) {
  //   return next(`/${locale}`)
  // }

  // load locale messages
  if (!i18n.global.availableLocales.includes(useLocale)) {
    await loadLocaleMessages(i18n, useLocale)
  }

  // set i18n language
  setI18nLanguage(i18n, useLocale)

  return next()
})

router.afterEach((to) => {
  /* Set document title from route meta */
  if (to.meta && to.meta.title) {
    document.title = `${to.meta.title} - ${defaultDocumentTitle}`
  } else {
    document.title = defaultDocumentTitle
  }
})

export default router
