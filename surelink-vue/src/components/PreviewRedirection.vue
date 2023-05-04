<template>
    <div class="container">
        <div class="columns">
            <div class="column is-offset-half">
                <div class="columns is-flex-mobile">
                    <div class="column is-left is-1 is-2-mobile">
                        <a href="/"><img class="back" src="../assets/icons/left-arrow.svg" alt="reload"/></a>
                    </div>
                    <div class="column is-11">
                        <h3 class="subtitle">You are being redirected to: </h3>
                    </div>
                </div>
                <div class="card">
                    <div v-if="!isLoading">
                        <div class="card-image">
                            <img v-if="image" :src="image" class="is-64x64" alt="">
                        </div>
                    </div>
                    <div class="card-content">
                        <div v-if="!isLoading">
                            <div class="columns is-flex-mobile">
                                <div class="column is-left is-2 is-2-mobile">
                                    <img v-if="icon" :src="icon" class="icon margin" alt="">
                                </div>
                                <div class="column is-11 is-flex is-align-items-center is-justify-content-start">
                                    <h3 class="title is-5">{{ title }}</h3>
                                </div>
                            </div>

                            <div class="content has-text-justified">
                                {{ description }}
                                <br>
                            </div>
                        </div>
                        <div v-if="isLoading && !loadingFail">
                            <Loader/>
                            <h3 class="margin">Loading Preview</h3>
                        </div>
                        <div v-if="loadingFail">
                            <img class="sad" src="../assets/images/sad.svg" alt="reload"/>
                            <h3 class="margin">{{message}}</h3>
                        </div>
                    </div>
                </div>
                <button class="button is-6 is-primary" @click="redirectToLink()">Proceed to link</button>
            </div>
        </div>
    </div>
</template>

<script>

import axios from "axios";
import Loader from "@/components/Loader.vue";
export default {
    name: "PreveiwRedirection",
    components: {Loader},
    data(){
        return {
            redirectLink: '',
            title: '',
            description: '',
            image: '',
            url: '',
            icon: '',
            isLoading: true,
            message: '',
            loadingFail: false
        }
    } ,
    mounted() {
        this.isLoading = true;
        this.fetchOriginalLink();
        setTimeout(()=>{
            if(this.title !== ''){
                this.isLoading = false;
            }
            else {
                this.message = 'Could not load preview';
                this.loadingFail = true;
            }
        },2000)
    },
    methods: {
        fetchOriginalLink(){
            const link = this.$route.params.link;
            axios.get('https://api.surel.ink/redirection/get-map', {
                params: {
                    "uid": link
                },
                headers: {
                    'Content-Type': 'application/json'
                },
            })
                .then(response => {
                    this.redirectLink = response.data['url'];
                    this.generatePreview();
                    // this.demoPreview();
                })
                .catch(error => {
                    console.log(error.response);
                });
        },
        generatePreview(){
            axios.post(
                'https://api.peekalink.io/',
                { link: this.redirectLink },
                { headers: { 'X-API-Key': process.env.VUE_APP_PREVIEW_API } }
            ).then(response => {
                this.title = response.data['title'];
                this.description = response.data['description'];
                this.image = response.data['image'].url;
                this.icon = response.data['icon'].url;
            }).catch(error => {
                console.error(error);
            });
        },
        redirectToLink(){
            window.location.href = this.redirectLink;
        },
        //TODO: used for testing only
        demoPreview(){
            this.title = 'Github';
            this.description = 'Lorem Ipsum Lorent'
            this.image = 'https://cdn.peekalink.io/public/images/5cfe8a0d-1844-49f2-8e79-8c77bc62e4ec/b00639ed-1948-4c2d-9eb1-0b784876ea05.jpg';
            this.icon = 'https://cdn.peekalink.io/public/images/fa594927-5453-4c4d-9ac6-1fc59f3c6704/7effbc39-74ca-4ad9-bb27-989cba6856d0.jpg';
        }
    }
}
</script>

<style scoped>
.container {
    margin: 1rem;
}
.back {
    width: 20px;
    height: 20px;
}
.card {
    padding: 2rem;
    margin: 1rem;
}
.margin {
    margin: 1rem;
}
.sad {
    width: 250px;
    height: 250px;
}
.icon {
    width: 2rem;
    height: 2rem;
}
</style>