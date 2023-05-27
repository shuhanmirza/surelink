<template>
    <section class="section m-5">
        <div class="container">
            <div class="columns is-centered">
                <div class="column custom-card is-12-mobile">
                    <div class="card shadow-lg is-cursor-pointer has-text-centered">
                        <div v-if="isLoading && !loadingFail">
                            <br/>
                            <Loader/>
                            <br/>
                        </div>
                        <div v-else-if="loadingFail">
                            <figure class="">
                                <img src="../assets/images/sad.svg" class="is-64x64" alt="reload">
                            </figure>
                            <div class="content has-text-weight-bold">
                                <h3></h3>
                                {{ message }}
                            </div>
                            <br/>
                        </div>
                        <div v-else>
                            <div>
                                <figure class="card-image">
                                    <img :src="image" class="is-64x64" alt="Placeholder image">
                                </figure>
                            </div>
                            <div class="card-content">
                                <div class="content has-text-weight-bold is-marginless">
                                    <h3></h3>
                                    {{ title }}
                                </div>
                                <div class="content has-text-grey-light">
                                    <h3></h3>
                                    {{ description }}
                                </div>
                            </div>
                        </div>
                        <footer class="card-footer has-background-white-bis">
                            <span @click="redirectToLink()"
                                  :class="{
                                           'has-background-success' : !this.isLoading && !this.loadingFail ,
                                           'has-background-danger' : !this.isLoading && this.loadingFail
                                  }"
                                  class="card-footer-item p-5 has-text-weight-bold is-uppercase is-text-wide-1">
                                {{ proceedText }}
                            </span>
                        </footer>
                    </div>
                </div>
            </div>
        </div>
    </section>
</template>

<script>

import Vue from "vue";
import axios from "axios";
import Loader from "@/components/Loader.vue";
import Toast from "vue-toastification";
import "vue-toastification/dist/index.css";

Vue.use(Toast);
export default {
    name: "PreviewRedirection",
    components: {Loader},
    data() {
        return {
            redirectLink: '',
            title: '',
            description: '',
            image: '',
            url: '',
            isLoading: true,
            message: '',
            loadingFail: false,
            proceedText: 'Generating Link!'
        }
    },
    mounted() {
        this.fetchOriginalLink();
        setTimeout(() => {
            if (this.title !== '') {
                this.proceedText = 'Proceed to Link';
                this.loadingFail = false;
                this.isLoading = false;

            } else {
                this.message = 'Could not load preview';
                this.loadingFail = true;
                this.isLoading = false;
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
                    this.loadingFail = false;
                    this.isLoading = false;
                    this.proceedText = 'Proceed to Link';
                    this.redirectLink = response.data['url'];
                    this.generatePreview(link);
                })
                .catch(error => {
                    this.toastFailure("Incorrect Short Link!");
                    this.proceedText = 'Incorrect Link';
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
                })
                .catch(error => {
                    this.title = url
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

body {
    color: #333;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
}

.cover-image img {
    transition: transform 1s, filter 2s ease-in-out;
    -webkit-transition: transform 1s, filter 2s ease-in-out;
    transform: scale(1.1);
    -webkit-transform: scale(1.1);
}

.card:hover .cover-image img {
    cursor: pointer;
    transform: scale(1.2);
    -webkit-transform: scale(1.2);
}

.is-cursor-pointer {
    cursor: pointer !important;
}

.shadow-lg {
    box-shadow: 0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05) !important;
}

.custom-card {
    flex: none;
    width: 40%;
}

.is-64x64 {
    margin-top: 1.5rem;
}

.card-content {
    padding: -1rem;
}

</style>