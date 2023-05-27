<template>
  <div class="is-family-primary columns is-centered">
    <div class="column is-centered is-9">
      <h2 class="title is-size-1 ">SHORTEN YOUR URLS</h2>
      <h3 class="is-size-5">Streamline your links with ease!</h3>
      <div class="columns is-centered">
        <div class="card column is-9">
          <div class="card-content">
            <div class="content">
              <h4 class="subtitle is-size-4">Shorten URL Is Just Simple</h4>
              <div class="field">
                <div class="control is-medium is-centered columns is-variable is-4">
                  <span class="column is-7">
                    <input class="input is-medium is-primary" type="text" placeholder="Your Link Here"
                           v-model="targetLink"/>
                  </span>
                  <span class="column is-2">
                    <button class="button is-medium is-primary" @click="generateShortLink(targetLink)">
                      Shorten
                    </button>
                  </span>
                </div>
              </div>
              <div class="columns is-centered" v-if="isVerified && success">
                <span class="column is-5">
                    <input class="input is-medium" ref="shortenField" v-model="shortenUrl" type="text"/>
                </span>
                <span class="column is-2">
                  <button class="button is-medium is-small" @click="copyLink">
                    Copy
                  </button>
                </span>
              </div>
              <br/>
              <div v-if="statusExists" class="columns is-centered">
                <div class="column is-3">
                  <h2 class="is-size-2">
                    <VueJsCounter start="0" :end=totalRedirectedLifetime duration="1500" thousand=","
                                  decimal="."></VueJsCounter>
                  </h2>
                  <h4>Links Redirected in Total</h4>
                </div>
                <div class="column is-3">
                  <h2 class="is-size-2">
                    <VueJsCounter start="0" :end=totalLinksLifetime duration="1500" thousand=","
                                  decimal="."></VueJsCounter>
                  </h2>
                  <h4>Links Generated in Total</h4>
                </div>
                <div class="column is-3">
                  <h2 class="is-size-2">
                    <VueJsCounter start="0" :end=Math.ceil(totalRedirectedLifetime/365) duration="1500" thousand=","
                                  decimal="."></VueJsCounter>
                  </h2>
                  <h4>Links Generated per Day</h4>
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
  components: {VueJsCounter},
  data() {
    return {
      isVerified: false,
      uuid: null,
      img: null,
      captchaValue: '',
      targetLink: '',
      shortenUrl: '',
      success: false,
      showDiv: true,
      totalLinksLifetime: 0,
      totalRedirectedLifetime: 0,
      statusExists: false,
    };
  },
  beforeMount() {
    this.getStatus();
  },
  mounted() {
    console.log("Hello Visitor!");
  },
  methods: {
    generateShortLink(link) {
      //TODO: the url is validated on the backend. So we can ease this regex check.
      if (/^(http|https):\/\/[a-zA-Z0-9\-\.]+\.[a-zA-Z]{2,}(\/\S*)?$/.test(link) === false) {
        this.toastFailure('Insert a Valid Link!');
        return;
      }

      grecaptcha.ready(function () {
            //TODO: fetch site key from environment
            grecaptcha.execute('6LcPjUMmAAAAAEOc7hMHqPMCRTisgzCy68jzUUcF', {action: 'submit'})
                .then(this.recaptchaTokenCallback);
          }.bind(this)
      );
    },
    recaptchaTokenCallback(token) {
      const requestBody = {
        "recaptcha_token": token,
        "url": this.targetLink
      }

      axios.post('https://api.surel.ink/redirection/set-map/v2', requestBody)
          .then(response => {
            this.shortenUrl = 'https://surel.ink/' + response.data['short_url'];
            this.success = true;
            this.isVerified = true;
            this.targetLink = '';
            this.toastSuccess('Link Generated!');
          })
          .catch(error => {
            console.log(error.response);
            let errorMsg = error.response.data.hasOwnProperty("error") ? error.response.data.error : "An error occurred! Please try again."
            this.toastFailure(errorMsg);
          });
    },
    copyLink() {
      const inputField = this.$refs.shortenField;
      navigator.clipboard.writeText(inputField.value);
      inputField.setSelectionRange(0, 0);
      this.toastSuccess('Link copied to clipboard!');
    },
    toastSuccess(message) {
      this.$toast.success(message, {
        timeout: 2000,
        position: 'bottom-center'
      });
    },
    toastFailure(message) {
      this.$toast.error(message, {
        timeout: 2000,
        position: 'bottom-center'
      });
    },
    async getStatus() {
      await axios.get('https://api.surel.ink/stat/home')
          .then(response => {
            this.totalLinksLifetime = response.data['num_url_map_created_lifetime'] + 4000;
            this.totalRedirectedLifetime = response.data['num_url_map_redirected_lifetime'] + 8000;
            this.statusExists = true;
          })
          .catch(error => {
            console.log(error)
          });
    }
  },
  watch: {
    success: function (newValue, oldValue) {
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

.reload {
  width: 20px;
  height: 20px;
  margin: 1rem;
}
</style>
