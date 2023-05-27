<template>
    <section class="class-404">
      <div class="container">
          <img src="../../src/assets/images/logo.svg" width="20" height="20"> SureLink
          <h3 class="title">404 Not Found</h3>
          <div class="class-404-image"></div>
          <h3 class="subtitle">The link you're looking for doesn't exist!</h3>
          <a href="/"><button class="button is-primary">Go Back Home</button></a>
      </div>
    </section>
</template>

<script>

export default {
    name: "ErrorItem.vue",
}
</script>

<style scoped>
.class-404 {
    padding:40px 0; background:#fff;
}
.class-404-image{
    background-image: url('../../src/assets/images/404.gif');
    height: 400px;
    background-position: center;
}
.title {
    padding-top: 1rem;
}
</style>