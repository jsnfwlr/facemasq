import { createApp } from "vue"
import App from "./App.vue"
import "./styles/index.scss"
import router from "./routes"

import { createPinia } from "pinia"

const app = createApp(App)
app.use(router)
// app.use(i18n)
app.use(createPinia())
app.mount("#app")
