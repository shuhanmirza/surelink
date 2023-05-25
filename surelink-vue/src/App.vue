<template>
  <div id="app">
    <meta name="title" content="SureLink - Free Link Shortener" />
    <meta name="description" content="It is a short link generator tool that helps you generate tinyurls. It provides security from harmful links and makes the urls human friendly" />
    <meta name="keywords" content="url shortener free, bitly url shortener, google short url, tinyurl generator, url shorten" />
    <meta property="og:title" content="SureLink - Free Link Shortener" />
    <meta property="og:description" content="It is a short link generator tool that helps you generate tinyurls. It provides security from harmful links and makes the urls human friendly" />
    <meta property="og:image" content="https://surel.ink/img/logo.b4c96294.svg" />
    <router-view/>
  </div>
</template>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>
