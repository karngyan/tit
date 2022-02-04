<template>
<div class="mx-auto max-w-3xl p-4 flex items-center justify-center">
  <div class="space-x-4">
    <nuxt-link exact to="/" class="active-nav hover:text-blue-400 text-gray-400 text-sm">home</nuxt-link>
    <nuxt-link to="/config" class="active-nav hover:text-blue-400 text-gray-400 text-sm">config</nuxt-link>
  </div>
</div>
</template>

<script>
export default {
  name: "TheNavBar"
}
</script>

<style scoped>
.nuxt-link-active.active-nav {
  @apply underline underline-offset-2 decoration-blue-500 decoration-2;
}
</style>
