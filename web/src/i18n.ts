import { nextTick } from "vue"
import { I18nOptions, I18n, createI18n } from "vue-i18n"

export const SUPPORT_LOCALES = ["en", "fr"]

export function setupI18n(options: I18nOptions = { locale: "en" }) {
  const i18n = createI18n(options)
  const locale = typeof options.locale !== "undefined" ? options.locale : "en"
  setI18nLanguage(i18n, locale)
  return i18n
}

export function setI18nLanguage(i18n: I18n, locale: string) {
  if (i18n.mode === "legacy") {
    i18n.global.locale = locale
  } else {
    if (typeof i18n.global.locale === "string") {
      i18n.global.locale = locale
    } else {
      i18n.global.locale.value = locale
    }
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
  const messages = await import(/* webpackChunkName: "locale-[request]" */ `../../i18n/web/facemasq.${locale}.json`)

  // set locale and locale message
  i18n.global.setLocaleMessage(locale, messages.default)

  return nextTick()
}
