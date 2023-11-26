import './index.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import { createI18n } from 'vue-i18n'

import en from './locales/en.json'

const i18n = createI18n({
    legacy: false,
    fallbackLocale: 'en',
    messages: {
        "en": en,
        "en-GB": en,
        "en-US": en,
    }
})

const app = createApp(App)

app.use(i18n)
app.use(createPinia())
app.use(router)

app.mount('#app')
