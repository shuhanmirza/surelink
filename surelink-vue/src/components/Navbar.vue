<template>
    <nav class="navbar is-transparent mb-5" role="navigation" aria-label="main navigation">
        <div class="navbar-brand">
            <a href="/">
                <img src="../../src/assets/images/logo.svg" width="32" height="32"> SureLink
            </a>
        </div>
    </nav>
=======
<!--    <nav class="navbar is-transparent mb-5" role="navigation" aria-label="main navigation">-->
<!--        <div class="navbar-brand">-->
<!--&lt;!&ndash;            <a href="/">&ndash;&gt;-->
<!--&lt;!&ndash;                <img src="../../src/assets/images/logo.svg" width="32" height="32"> SureLink&ndash;&gt;-->
<!--&lt;!&ndash;            </a>&ndash;&gt;-->
<!--        </div>-->
<!--    </nav>-->
>>>>>>> Stashed changes
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
    font-family: 'Roboto',sans-serif;
    padding: 0;
}

a {
    text-decoration: none !important;
}

.navbar .navbar-brand > .navbar-item,
.navbar .navbar-brand {
    display: inline-block;
}
</style>
