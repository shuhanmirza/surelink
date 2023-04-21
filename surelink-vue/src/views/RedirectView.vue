<template>
    <div id="counter">
        <CountDown/>
    </div>
</template>

<script>

import CountDown from "../components/CountDown.vue";

export default {
    name: "RedirectView",
    components: {
        CountDown
    }
};
</script>

<style>
body {
    background-color: #fff;
}

#counter {
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    margin-top: 60px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    width: 100%;
}
</style>
