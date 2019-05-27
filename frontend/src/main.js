import Vue from "vue";
import './plugins/vuetify'
import App from "./App.vue";
import router from "./router/";
import store from "./store";
import "./registerServiceWorker";
import 'roboto-fontface/css/roboto/roboto-fontface.css'
import 'font-awesome/css/font-awesome.css'
import axios from 'axios';

Vue.config.productionTip = false;

Vue.mixin({
    methods: {
        getApiUrl(slug = "") {
            return ApiBaseUrl ? ApiBaseUrl + slug : window.location.protocol + '//' + window.location.hostname + slug
        },
        apiCall(method, url, payload = {}) {
            const instance = axios.create({
                //baseURL: 'https://some-domain.com/api/',
                timeout: 2000,
                headers: {'X-Custom-Header': 'foobar'}
            });

            switch (method) {
                case 'get' || 'delete':
                    return instance.get(url)
                        /*
                        .then(response => {
                            return response
                        })
                        .catch(e => {
                            alert(e)
                        })*/
                case 'post' || 'put':
                    break;
                default:
                    alert('Invalid Method for API client')
            }
        }
    }
})

new Vue({
    router,
    store,
    render: h => h(App)
}).$mount("#app");
