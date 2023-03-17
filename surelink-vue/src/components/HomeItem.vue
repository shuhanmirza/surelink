<template>
  <div class="is-family-primary">
    <h2 class="title is-size-2">Shorten Your URLS</h2>
    <h3 class="is-size-5">Transforming long, ugly links into Shorten URLs</h3>
    <div class="columns is-centered">
      <div class="card column is-8">
        <div class="card-content">
          <div class="content">
            <h4 class="subtitle is-size-4">Shorten URL Is Just Simple</h4>
            <div class="field">
              <div class="control is-medium is-centered columns is-variable is-4">
                <span class="column is-6">
                  <input class="input is-medium" type="text" placeholder="Your Link Here">
                </span>
                <span class="column is-2">
                  <button class="button is-rounded is-medium is-primary">
                    Shorten
                  </button>
                </span>
              </div>
            </div>
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
h2 {
  margin: 40px 0 0 0;
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
