<template>
    <div>
      <div :style="'color:'+color">
          <vue-loaders-ball-beat color="currentColor" scale="1"></vue-loaders-ball-beat>
      </div>
    </div>
</template>

<script>
import Vue from "vue";
import 'vue-loaders/dist/vue-loaders.css';
import VueLoaders from 'vue-loaders';

Vue.use(VueLoaders);
export default {
    name: "Loader",
    components : {VueLoaders},
    data(){
        return {
            color: 'black'
        }
    }
}
</script>

<style scoped>

</style>