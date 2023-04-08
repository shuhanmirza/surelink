<template>
  <div class="home">
    <div class="background">
      <Navbar />
      <HomeItem msg="Welcome to Your Vue.js App" />
      <BodyItem />
      <FooterItem />
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import HomeItem from '@/components/HomeItem.vue'
import Navbar from '@/components/Navbar.vue'
import BodyItem from '@/components/BodyItem.vue'
import FooterItem from '@/components/FooterItem.vue'
export default {
  name: 'HomeView',
  components: {
    Navbar,
    HomeItem,
    BodyItem,
    FooterItem
  }
}
</script>

<style>
.background {
  background-image: url('../assets/images/bg.svg');
  background-size: cover;
  background-repeat: no-repeat;
  height: 500px;
  border-bottom-left-radius: 20%;
  border-bottom-right-radius: 20%;
}
</style>
