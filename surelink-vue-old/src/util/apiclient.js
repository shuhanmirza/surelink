import axios from "axios";

export default class ApiClient {
    basePath;
    constructor(basePath) {
        this.basePath = basePath
        this.axios = axios.create({
            baseURL:basePath,
            headers: {
                'Content-Type': 'application/json'
            }
        })
    }

    async post(url, payload){
        return this.axios.post(url, payload)
    }

    async get(url, payload){
        return this.axios.get(url, payload)
    }

}