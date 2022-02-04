<template>
<div>
  <div class="shadow-md border border-zinc-800 rounded-md p-4">
    <div class="text-sm text-zinc-300">
      {{ tweet.content }}
    </div>
    <div class="flex pt-4 text-xs text-zinc-500 flex-row justify-between items-center">
      <div>
        send at: {{ tweet.sendAt | formatDate }}
      </div>
      <div class="">
        <button @click="deleteTweet" class="hover:text-blue-500 text-zinc-500">delete</button>
      </div>
    </div>
  </div>
</div>
</template>

<script>
export default {
  name: "Tweet",
  props: {
    tweet: {
      type: Object,
      default: () => {
        return {
          id: 'some id',
          content: 'some content',
          done: false,
          sendAt: new Date().toISOString(),
          created: new Date().toISOString(),
          updated: new Date().toISOString()
        }
      }
    }
  },
  methods: {
    deleteTweet() {
      this.$store.dispatch('deleteTweet', { id: this.tweet.id })
    }
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
