<template>
    <div>
        <h2 class="title">Shorten Your URLS</h2>
        <h3>Transforming long, ugly links into Shorten URLs</h3>
        <div class="columns is-centered">
        <div class="card column is-8">
          <div class="card-content">
            <div class="content">
                <h4 class="subtitle"> Shorten URL Is Just Simple</h4>
                <div class="field">
              <div class="control is-medium is-centered columns is-variable is-4">
                <span class="column is-6">
                    <input class="input is-medium" type="text" placeholder="Input your ugly link, we'll do the rest">
                </span>
                <span class="column is-2">
                    <button class="button is-rounded is-medium is-primary">Shorten</button>
                </span>
              </div>
            </div>
            </div>
          </div>
        </div>
        </div>
    </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped lang="scss">
h2{
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
