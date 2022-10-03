import { createApp } from "vue"
import App from "./App.vue"
import "./styles/index.scss"
import router from "./routes"
import { createI18n } from "vue-i18n"
import { createPinia } from "pinia"

import en from "../../i18n/web/facemasq-en.json"
import fr from "../../i18n/web/facemasq-fr.json"

type MessageSchema = typeof en | typeof fr

const i18n = createI18n<[MessageSchema], "en" | "fr">({
  locale: "en-US",
  messages: {
    en: en,
    fr: fr,
  },
})

const app = createApp(App)
app.use(router)
app.use(i18n)
app.use(createPinia())
app.mount("#app")
