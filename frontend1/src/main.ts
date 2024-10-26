import { createApp } from 'vue'
import './assets/styles.css'
import App from './App.vue'
import router from './index'

createApp(App)
    .use(router)
    .mount('#app')
