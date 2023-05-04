<!--<template>-->
<!--    <div id="counter">-->
<!--        <CountDown/>-->
<!--    </div>-->
<!--</template>-->

<!--<script>-->

<!--import Vue from 'vue';-->
<!--import CountDown from "../components/CountDown.vue";-->
<!--import axios from "axios";-->
<!--import Toast from "vue-toastification";-->
<!--import "vue-toastification/dist/index.css";-->

<!--Vue.use(Toast);-->
<!--export default {-->
<!--    name: "RedirectView",-->
<!--    components: {-->
<!--        CountDown-->
<!--    },-->
<!--    mounted() {-->
<!--        const link = this.$route.params.link;-->
<!--        let redirectLink = '';-->

<!--        axios.get('https://api.surel.ink/redirection/get-map', {-->
<!--            params: {-->
<!--                "uid": link-->
<!--            },-->
<!--            headers: {-->
<!--                'Content-Type': 'application/json'-->
<!--            },-->
<!--        })-->
<!--            .then(response => {-->
<!--                redirectLink = response.data['url'];-->
<!--            })-->
<!--            .catch(error => {-->
<!--                console.log(error.response);-->
<!--            });-->

<!--        setTimeout(()=> {-->
<!--            if (redirectLink !== '') {-->
<!--                window.location.href = redirectLink;-->
<!--            }-->
<!--            else {-->
<!--                this.$toast.error('Incorrect Link, Unable to redirect', {-->
<!--                    timeout: 2000,-->
<!--                    position: 'bottom-center'-->
<!--                });-->
<!--            }-->
<!--        }, 5500);-->
<!--    }-->
<!--};-->
<!--</script>-->

<!--<style>-->
<!--body {-->
<!--    background-color: #fff;-->
<!--}-->

<!--#counter {-->
<!--    -webkit-font-smoothing: antialiased;-->
<!--    -moz-osx-font-smoothing: grayscale;-->
<!--    text-align: center;-->
<!--    margin-top: 60px;-->
<!--    display: flex;-->
<!--    flex-direction: column;-->
<!--    justify-content: center;-->
<!--    align-items: center;-->
<!--    width: 100%;-->
<!--}-->
<!--</style>-->
