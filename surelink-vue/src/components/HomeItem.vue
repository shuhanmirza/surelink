<template>
  <div class="is-family-primary columns is-centered">
    <div class="column is-centered is-9">
      <h2 class="title is-size-1 ">SHORTEN YOUR URLS</h2>
      <h3 class="is-size-5">Streamline your links with ease!</h3>
      <!-- Transforming long, ugly links into Shorten URLs -->
<!--      <div class="columns is-centered">-->
<!--        <div class="column is-3">-->
<!--          <p class="is-size-3"><b>SHORTEN</b></p>-->
<!--        </div>-->
<!--        <div class="column is-3">-->
<!--          <p class="is-size-3"><b>BEAUTIFY</b></p>-->
<!--        </div>-->
<!--        <div class="column is-3">-->
<!--          <p class="is-size-3"><b>MEMORIALIZE</b></p>-->
<!--        </div>-->
<!--      </div>-->
      <div class="columns is-centered">
        <div class="card column is-9">
          <div class="card-content">
            <div class="content">
              <h4 class="subtitle is-size-4">Shorten URL Is Just Simple</h4>
              <div class="field">
                <div class="control is-medium is-centered columns is-variable is-4">
                  <span class="column is-7">
                    <input class="input is-medium is-primary" type="text" placeholder="Your Link Here" v-model="targetLink"/>
                  </span>
                  <span class="column is-2">
                    <button class="button is-rounded is-medium is-primary" @click="verifyCaptcha(captchaValue, targetLink)">
                      Shorten
                    </button>
                  </span>
                </div>
              </div>
              <div v-if="!isVerified && img && showDiv">
                  <img :src="'data:image/png;base64,' + img" alt="Base64 Image"/> <br/>
<!--                  <button class="button"> .-->
<!--                      <img src="../assets/icons/rotate-right-solid.svg" alt="reload"/>-->
<!--                  </button>-->
                  <input class="input is-rounded is-primary is-small captcha-input" type="text" v-model="captchaValue" placeholder="Enter Captcha Value"/> <br/>
              </div>
              <div class="columns is-centered" v-if="isVerified && success">
                <span class="column is-5">
                    <input class="input is-medium" ref="shortenField" v-model="shortenUrl" type="text"/>
                </span>
                <span class="column is-2">
                  <button class="button is-rounded is-medium is-small" @click="copyLink">
                    Copy
                  </button>
                </span>
              </div>
              <br />
              <div class="columns is-centered">
                <div class="column is-3">
                  <h2 class="is-size-2">
                    <VueJsCounter start="0" end="5231" duration="1500" thousand="," decimal="."></VueJsCounter>
                  </h2>
                  <h4>Total Users</h4>
                </div>
                <div class="column is-3">
                  <h2 class="is-size-2">
                    <VueJsCounter start="0" end="6340" duration="1500" thousand="," decimal="."></VueJsCounter>
                  </h2>
                  <h4>Links Generated in Total</h4>
                </div>
                <div class="column is-3">
                  <h2 class="is-size-2">
                    <VueJsCounter start="0" end="232" duration="1500" thousand="," decimal="."></VueJsCounter>
                  </h2>
                  <h4>Links Generated Per Day</h4>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>

import Vue from 'vue';
import VueJsCounter from 'vue-js-counter';
import axios from 'axios';
import Toast from "vue-toastification";
import "vue-toastification/dist/index.css";

Vue.use(Toast);
export default {
  name: "HomeItem",
  components: { VueJsCounter },
  data() {
      return {
          isVerified: false,
          uuid: null,
          img: null,
          captchaValue: '',
          targetLink: '',
          shortenUrl: '',
          success: false,
          showDiv: true
      };
  },
  mounted() {
      this.generateCaptcha();
  },
  methods: {
      generateCaptcha(){
          axios.get('https://api.surel.ink/captcha/new')
              .then(response => {
                  this.uuid = response.data['uuid'];
                  this.img = response.data['img'];
              })
              .catch(error => {

              });
      },
      verifyCaptcha(captcha, link) {
          const requestBody = {
              "captcha_uuid": this.uuid,
              "captcha_value": captcha,
              "url": link
          }
          axios.post('https://api.surel.ink/redirection/set-map', requestBody)
              .then(response => {
                  this.shortenUrl = response.data['short_url'];
                  this.success = true;
                  this.isVerified = true;
                  this.toastSuccess('Link Generated!');
              })
              .catch(error => {
                  this.toastFailure('Wrong Captcha, Please Retry');
                  this.generateCaptcha();
              });
      },
      copyLink(){
          const inputField = this.$refs.shortenField;
          navigator.clipboard.writeText(inputField.value);
          inputField.setSelectionRange(0, 0);
          this.toastSuccess('Link copied to clipboard!');
      },
      toastSuccess(message){
          this.$toast.success(message, {
              timeout: 2000, // duration of the toast message in milliseconds
              position: 'bottom-center' // position of the toast message on the screen
          });
      },
      toastFailure(message){
          this.$toast.error(message, {
              timeout: 2000, // duration of the toast message in milliseconds
              position: 'bottom-center' // position of the toast message on the screen
          });
      }
  },
  watch: {
      success: function(newValue, oldValue) {
          this.showDiv = false;
      }
  }
}



</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
.title {
  padding: 8rem 0 1rem 0;
}

h3 {
  margin: 0 0 40px 0;
}

.card {
  min-height: 300px;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
.captcha-input {
  max-width: 13rem;
  margin: 1rem 0 1rem 0;
}
</style>
