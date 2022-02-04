<template>
<div>
  <div class="shadow-md border border-zinc-800 rounded-md p-4">
    <div class="text-sm text-zinc-500">
      <div>api key: <span class="text-zinc-300">{{ config.apiKey }}</span></div>
      <div>api key secret: <span class="text-zinc-300">{{ config.apiKeySecret }}</span></div>
      <div>access token: <span class="text-zinc-300">{{ config.accessToken }}</span></div>
      <div>access token secret: <span class="text-zinc-300">{{ config.accessTokenSecret }}</span></div>
    </div>
  </div>
</div>
</template>

<script>
export default {
  name: "Config",
  props: {
    config: {
      type: Object,
      default: () => {
        return {
          id: '',
          bearerToken: '',
          default: true,
          created: '',
          updated: '',
          apiKey: '',
          apiKeySecret: '',
          accessToken: '',
          accessTokenSecret: ''
        }
      }
    }
  },
  methods: {

  },
  filters: {
    formatDate(d) {
      const date = new Date(d)
      return date.toString()
    }
  }
}
</script>

<style scoped>

</style>
