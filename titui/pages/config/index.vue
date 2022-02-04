<template>
  <div class="mx-auto max-w-3xl p-4">
    <form @submit.prevent="createTwitterConfig" class="space-y-4">
      <div>
        <label for="api-key" class="block text-sm font-medium text-zinc-500">api key</label>
        <div class="mt-1">
          <input v-model="form.apiKey" name="api-key" id="api-key" type="text" class="text-field-style"/>
        </div>
      </div>
      <div>
        <label for="api-key-secret" class="block text-sm font-medium text-zinc-500">api key secret</label>
        <div class="mt-1">
          <input v-model="form.apiKeySecret" name="api-key-secret" id="api-key-secret" type="text" class="text-field-style"/>
        </div>
      </div>
      <div>
        <label for="access-token" class="block text-sm font-medium text-zinc-500">access token</label>
        <div class="mt-1">
          <input v-model="form.accessToken" name="access-token" id="access-token" type="text" class="text-field-style"/>
        </div>
      </div>
      <div>
        <label for="access-token-secret" class="block text-sm font-medium text-zinc-500">access token secret</label>
        <div class="mt-1">
          <input v-model="form.accessTokenSecret" name="access-token-secret" id="access-token-secret" type="text" class="text-field-style"/>
        </div>
      </div>
      <div class="pt-2">
        <button type="submit" class="w-full text-center flex justify-center items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-transparent focus:ring-offset-transparent focus:ring-blue-500">update config</button>
      </div>
    </form>

    <div class="space-y-2 py-4">
      <Config v-for="config in tconfigs" :key="config.id" :config="config" />
    </div>
  </div>
</template>

<script>
export default {
  fetch() {
    this.$store.dispatch('fetchAllTwitterConfigs')
  },
  computed: {
    tconfigs() {
      return this.$store.getters.tconfigs
    }
  },
  data() {
    return {
      form: {
        apiKey: '',
        apiKeySecret: '',
        accessToken: '',
        accessTokenSecret: ''
      }
    }
  },
  methods: {
    createTwitterConfig() {
      this.$store.dispatch('createTwitterConfig', { ...this.form})
        .then(() => {
          this.form = {
            apiKey: '',
            apiKeySecret: '',
            accessToken: '',
            accessTokenSecret: ''
          }
        })
    }
  }
}
</script>
