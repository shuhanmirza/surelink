<template>
  <v-layout>
    <h1>
      You will be redirected to
    </h1>
    <h2>
      {{ url}}
    </h2>
  </v-layout>
</template>

<script>
export default {
  name: "RedirectionConfirmation",
  data() {
    return {
      url: "",
    }
  }, computed: {
    uid() {
      return this.$route.params.uid
    }
  },
  methods: {
    async getUrlRedirectionMap() {
      this.$apiClient.post("redirection/get-map", {
        uid: this.uid
      }).then(({data}) => {
        console.log(data)
        this.url = data.url
      }).catch(({response}) => {
        console.log(response)
      });
    }
  },
  created() {
    this.getUrlRedirectionMap()
  }
}
</script>

<style scoped>

</style>