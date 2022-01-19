import { createApp } from 'vue'
import App from './App.vue'
import './styles/index.scss'
import router from './routes'
import { createPinia } from 'pinia'

createApp(App).use(router).use(createPinia()).mount('#app')
