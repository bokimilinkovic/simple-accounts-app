import { createApp } from 'vue'
import App from './App.vue'
// Import Bootstrap and BootstrapVue CSS files (order is important)
import router from './router'
import store from "./store";


createApp(App).use(router).use(store).mount('#app')
