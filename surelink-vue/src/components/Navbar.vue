<template>
  <nav class="navbar" role="navigation" aria-label="main navigation">
  <div class="navbar-brand columns">
    <a class="column is-3" href="/">
      <img src="https://cdn-icons-png.flaticon.com/512/1011/1011407.png" width="28" height="28"> SureLink
    </a>
    <div class="column is-6">
        <div class="columns">
            <a class="column">
                Home
            </a>
            <a class="column">
                About
            </a>
            <a class="column">
                Contact
            </a>
        </div>
    </div>
      <div class="navbar-item column is-3">
        <div class="buttons">
          <a class="button is-rounded is-primary">
            <strong>Sign up</strong>
          </a>
          <a class="button is-rounded is-light">
            Log in
          </a>
        </div>
      </div>
  </div>
</nav>
</template>
<script>
export default {
  name: 'Navbar'
}
</script>
<style>
    .navbar .navbar-brand{
        text-align: center;
        display: block;
        width: 100%;
    }
    .navbar .navbar-brand > .navbar-item,
    .navbar .navbar-brand {
        display: inline-block;
    }
</style>