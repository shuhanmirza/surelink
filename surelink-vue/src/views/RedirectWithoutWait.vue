<template>
    <div>
        <Navbar/>
        <PreviewRedirection/>
    </div>
</template>

<script>
import PreviewRedirection from "@/components/PreviewRedirection.vue";
import Navbar from "@/components/Navbar.vue";
import FooterItem from "@/components/FooterItem.vue";

export default {
    name: "RedirectWithoutWait.vue",
    components: {FooterItem, Navbar, PreviewRedirection}
}
</script>

<style scoped>

</style>