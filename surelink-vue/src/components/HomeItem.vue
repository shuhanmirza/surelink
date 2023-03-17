<template>
  <div class="is-family-primary columns is-centered">
    <div class="column is-centered is-9">
      <h2 class="title is-size-1 ">SHORTEN YOUR URLS</h2>
      <h3 class="is-size-5">Streamline your links with ease!</h3>
      <!-- Transforming long, ugly links into Shorten URLs -->
      <div class="columns is-centered">
        <div class="column is-3">
          <p class="is-size-3"><b>SHORTEN</b></p>
        </div>
        <div class="column is-3">
          <p class="is-size-3"><b>BEAUTIFY</b></p>
        </div>
        <div class="column is-3">
          <p class="is-size-3"><b>MEMORIALIZE</b></p>
        </div>
      </div>
      <div class="columns is-centered">
        <div class="card column is-9">
          <div class="card-content">
            <div class="content">
              <h4 class="subtitle is-size-4">Shorten URL Is Just Simple</h4>
              <div class="field">
                <div class="control is-medium is-centered columns is-variable is-4">
                  <span class="column is-7">
                    <input class="input is-medium is-primary" type="text" placeholder="Your Link Here">
                  </span>
                  <span class="column is-2">
                    <button class="button is-rounded is-medium is-primary">
                      Shorten
                    </button>
                  </span>
                </div>
              </div>
              <br />
              <div class="columns">
                <div class="column is-4">
                  <h2 class="is-size-2">
                    <VueJsCounter start="0" end="5231" duration="1500" thousand="," decimal="."></VueJsCounter>
                  </h2>
                  <h4>Total Users</h4>
                </div>
                <div class="column is-4">
                  <h2 class="is-size-2">
                    <VueJsCounter start="0" end="6340" duration="1500" thousand="," decimal="."></VueJsCounter>
                  </h2>
                  <h4>Links Generated in Total</h4>
                </div>
                <div class="column is-4">
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

import VueJsCounter from 'vue-js-counter';

export default {
  name: "HomeItem",
  components: { VueJsCounter }
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
  height: 400px;
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
</style>
