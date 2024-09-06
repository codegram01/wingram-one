import './assets/main.css'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { init_store } from './stores'
import { init_auth } from './auth'

const init = async (): Promise<void> => {
    await init_store();

    await init_auth();

    const app = createApp(App)

    app.use(router)

    app.mount('#app')
}
init();