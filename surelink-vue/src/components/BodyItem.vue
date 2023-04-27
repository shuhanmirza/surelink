<template>
    <div class="container">
        <br/>
        <div class="columns is-centered">
            <div class="column is-7">
                <h1 class="title">What We DO?</h1>
                <div class="grid-container">
                  <div class="grid-item">
                    <div class="image-text is-centered">
                      <img src="../../src/assets/images/minimize.png" width="64" height="64"/> &nbsp &nbsp
                      <h2 class="subtitle is-size-3">SHORTEN</h2>
                    </div>
                    <h3 class="is-size-6 has-text-centered">SureLink is a platform where you can quickly and easily shorten your long and complex URLs into short and memorable links that are easy to share and manage.</h3>
                  </div>
                  <div class="grid-item">
                    <div class="image-text is-centered">
                      <img src="../../src/assets/images/stars.png" width="64" height="64"/> &nbsp &nbsp
                      <h2 class="subtitle is-size-3">BEAUTIFY</h2>
                    </div>
                    <h3 class="is-size-6 has-text-centered">Our platform offers a user-friendly interface that makes it simple for anyone to create custom links. Use our generated links for free in any website.</h3>
                  </div>
                  <div class="grid-item">
                    <div class="image-text is-centered">
                      <img src="../../src/assets/images/light-bulb.png" width="64" height="64"/> &nbsp &nbsp
                      <h2 class="subtitle is-size-3">MEMORIALIZE</h2>
                    </div>
                    <h3 class="is-size-6 has-text-centered">Our service is fast, reliable, and secure, ensuring that your links are always accessible. We constantly monitor our system to ensure that your links remain secure and private.</h3>
                  </div>
                  <div class="grid-item">
                    <div class="image-text is-centered">
                      <img src="../../src/assets/images/graph.png" width="64" height="64"/> &nbsp &nbsp
                      <h2 class="subtitle is-size-3">PROMOTE</h2>
                    </div>
                    <h3 class="is-size-6 has-text-centered">Whether you're a business owner looking to promote your products or services online or on social media who wants to share content with followers, our service can help.</h3>
                  </div>
                </div>
            </div>  
        </div>
    </div>
</template>
<script>
export default {
    name: 'BodyItem'
}
</script>

<style scoped lang="scss">
.image-text {
    display: flex;
    align-items: center;
    justify-content: center;
}

.grid-container {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  grid-gap: 20px;
  margin-top: 20px;
}

@media screen and (min-width: 768px) {
  .grid-container {
    grid-template-columns: repeat(2, 1fr);
    grid-auto-rows: minmax(100px, auto);
    grid-gap: 30px;
  }
}

.grid-item {
  padding: 20px;
  text-align: center;
}

</style>