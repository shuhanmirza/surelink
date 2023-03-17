<template>
  <nav class="navbar is-transparent" role="navigation" aria-label="main navigation">
    <div class="navbar-brand columns">
      <a class="column is-3" href="/">
        <img src="https://cdn-icons-png.flaticon.com/512/1011/1011407.png" width="28" height="28"> SureLink
      </a>
      <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
      </a>
      <div class="navbar-menu column is-6 content" id="navbarBasicExample">
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
      <div class="navbar-end navbar-item column is-3">
        <div class="buttons">
          <button class="button is-rounded is-primary">
            Sign up
          </button>
          <button class="button is-rounded is-light">
            Log in
          </button>
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
.navbar .navbar-brand {
  text-align: center;
  display: block;
  width: 100%;
}

nav.navbar {
  min-height: unset;
  height: 0;
  font-family: 'Roboto';
  padding: 0;

}

a {
  text-decoration: none !important;
}

.navbar .navbar-brand>.navbar-item,
.navbar .navbar-brand {
  display: inline-block;
}
</style>