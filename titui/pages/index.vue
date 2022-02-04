<template>
  <div class="mx-auto max-w-3xl p-4">
    <form @submit.prevent="createTweet" class="space-y-4">
      <div>
        <label for="content" class="block text-sm font-medium text-zinc-500">new tweet</label>
        <div class="mt-1">
          <textarea v-model="form.content" placeholder="max length is 280 chars" maxlength="280" rows="10" name="content" id="content" class="text-field-style"></textarea>
        </div>
      </div>
      <div>
        <label for="send-at" class="block text-sm font-medium text-zinc-500">send at</label>
        <div class="mt-1">
          <input v-model="form.sendAt" name="send-at" id="send-at" type="text" class="text-field-style" placeholder="5 min from now"/>
        </div>
      </div>
      <div class="pt-2">
        <button type="submit" class="w-full text-center flex justify-center items-center px-3 py-2 border border-transparent text-sm leading-4 font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-transparent focus:ring-offset-transparent focus:ring-blue-500">schedule</button>
      </div>
    </form>

    <div class="space-y-2 py-4">
      <Tweet v-for="tweet in tweets" :key="tweet.id" :tweet="tweet" />
    </div>
  </div>
</template>

<script>
export default {
  fetch() {
    this.$store.dispatch('fetchAllTweets')
  },
  computed: {
    tweets() {
      return this.$store.getters.sortedTweets
    }
  },
  data() {
    return {
      form: {
        content: '',
        sendAt: ''
      }
    }
  },
  methods: {
    createTweet() {
      this.$store.dispatch('createTweet', { ...this.form})
        .then(() => {
          this.form = {
            content: '',
            sendAt: ''
          }
        })
    }
  }
}
</script>
