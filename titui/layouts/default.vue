<template>
<div class="min-h-screen">
  <TheNavBar />
  <Nuxt />
</div>
</template>

<script>
export default {
  name: "default"
}
</script>

<style>

html {
  @apply font-sans bg-zinc-900 !important;
}
/* Page-transition styles must be global */
/* page is used by default */
.page-leave-active,
.page-enter-active {
  transition: opacity 0.3s;
}
.page-enter,
.page-leave-to {
  opacity: 0;
}

::-moz-selection { @apply bg-blue-600 text-zinc-200; }
::selection { @apply bg-blue-600 text-zinc-200; }
</style>
