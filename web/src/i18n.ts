import { computed, nextTick } from "vue"
import { I18nOptions, I18n, createI18n } from "vue-i18n"
import { mande } from "mande"
// export const SUPPORT_LOCALES = ["en", "fr"]
const url = import.meta.env.DEV ? "http://192.168.0.41:6135/i18n" : "/api/i18n"
const fetchI18n = mande(url)

export function setupI18n(options: I18nOptions = { legacy: false, locale: "" }) {
  const i18n = createI18n(options)
  const locale = typeof options.locale === "string" && options.locale !== "" ? options.locale : "xy"
  setI18nLanguage(i18n, locale)
  return i18n
}

export function setI18nLanguage(i18n: I18n, locale: string) {
  if (i18n.mode === "legacy") {
    i18n.global.locale = locale
  } else {
    i18n.global.locale = computed(() => locale)
  }
  /**
   * NOTE:
   * If you need to specify the language setting for headers, such as the `fetch` API, set it here.
   * The following is an example for axios.
   *
   * axios.defaults.headers.common['Accept-Language'] = locale
   */
  document.querySelector("html")?.setAttribute("lang", locale)
}

export async function loadLocaleMessages(i18n: I18n, locale: string) {
  if (locale === "fr") {
    const messages = {
      sidebar: {
        dashboard: "Dashboard",
        devices: "Devices",
        taxonomy: "Taxonomy",
        categories: "Categories",
        statuses: "Statuses",
        maintainers: "Maintainers",
        locations: "Locations",
        deviceTypes: "Device Types",
        osTypes: "OS Types",
        cpuTypes: "CPU Types",
        vLANs: "VLANs",
        access: "Access",
        users: "Users",
        about: "About",
        info: "Info",
      },
    }
    i18n.global.setLocaleMessage("fr", messages)
    console.log(i18n.global.availableLocales)
  } else {
    console.log("!!", locale)
    fetchI18n.get("/facemasq." + locale + ".json").then((messages) => {
      console.log(messages)
      // set locale and locale message
      i18n.global.setLocaleMessage(locale, messages)
    })
  }
  return nextTick()
}
