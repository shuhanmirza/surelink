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
                            <div v-if="icon" class="columns is-flex-mobile">
                                <div class="column is-left is-2 is-2-mobile">
                                    <img :src="icon" class="icon margin" alt="">
                                </div>
                                <div class="column is-11 is-flex is-align-items-center is-justify-content-start">
                                    <h3 class="title is-5">{{ title }}</h3>
                                </div>
                            </div>
                            <div v-else>
                                <h3 class="title is-5">{{ title }}</h3>
                            </div>
                            <div v-if="description" class="content has-text-justified">
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
                            <h3 class="margin">{{ message }}</h3>
                        </div>
                    </div>
                </div>
                <button class="button is-6 is-primary" @click="redirectToLink()">Proceed to link</button>
            </div>
        </div>
    </div>
</template>

<script>

import Vue from "vue";
import axios from "axios";
import Loader from "@/components/Loader.vue";
import Toast from "vue-toastification";
import "vue-toastification/dist/index.css";

Vue.use(Toast);
export default {
    name: "PreveiwRedirection",
    components: {Loader},
    data() {
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
    },
    mounted() {
        this.fetchOriginalLink();
        setTimeout(() => {
            if (this.title !== '') {
                this.isLoading = false;
            } else {
                this.message = 'Could not load preview';
                this.loadingFail = true;
            }
        }, 2000)
    },
    methods: {
        fetchOriginalLink() {
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
                    this.generatePreview(link);
                })
                .catch(error => {
                    this.toastFailure("Incorrect Short Link!")
                });
        },
        generatePreview(url) {
            axios.get('https://api.surel.ink/link-preview/', {
                params: {
                    "uid": url
                },
                headers: {
                    'Content-Type': 'application/json'
                }
            })
                .then(response => {
                    this.title = response.data['title'];
                    this.description = response.data['description'];
                    this.image = response.data['image'];
                    this.icon = this.image;
                })
                .catch(error => {
                    console.log("Could not load preview");
                })

        },
        redirectToLink() {
            window.location.href = this.redirectLink;
        },
        toastFailure(message) {
            this.$toast.error(message, {
                timeout: 2000,
                position: 'bottom-center'
            });
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