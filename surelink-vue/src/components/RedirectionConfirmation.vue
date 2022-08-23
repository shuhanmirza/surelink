<template>
<v-layout>
  <h1>
    You will be redirected to
  </h1>
  <h2>
    {{this.$route.params.link_uuid}}
  </h2>
</v-layout>
</template>

<script>
export default {
  name: "RedirectionConfirmation"
}
</script>

<style scoped>

</style>