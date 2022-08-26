import {createApp} from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import {loadFonts} from './plugins/webfontloader'
import router from './router'
import ApiClient from "@/util/apiclient";

loadFonts()

const app = createApp(App);
app.use(router)
    .use(vuetify)
    .mount('#app')

console.log(process.env.VUE_APP_SERVER_BASE_PATH)


app.config.globalProperties.$apiClient = new ApiClient(
    process.env.VUE_APP_SERVER_BASE_PATH)



