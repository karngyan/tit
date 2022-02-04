<template>
<div aria-live="assertive" class="lowercase dark fixed bottom-0 right-0 left-0 z-50 flex items-end px-4 py-6 pointer-events-none sm:p-6 sm:items-start">
  <div class="w-full flex flex-col items-center space-y-4 sm:items-end">
    <transition enter-active-class="transform ease-out duration-300 transition" enter-from-class="translate-y-2 opacity-0 sm:translate-y-0 sm:translate-x-2" enter-to-class="translate-y-0 opacity-100 sm:translate-x-0" leave-active-class="transition ease-in duration-100" leave-from-class="opacity-100" leave-to-class="opacity-0">
      <div v-if="notification.show" class="max-w-sm w-full bg-white dark:bg-zinc-900 shadow-lg shadow-black rounded-lg pointer-events-auto ring-1 ring-zinc-100 dark:ring-zinc-800 ring-opacity-5 overflow-hidden">
        <div class="p-4">
          <div class="flex items-start">
            <div class="flex-shrink-0">
              <!-- Heroicon name: outline/check-circle -->
              <svg v-if="notification.type === 'success'" class="h-6 w-6 text-green-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              <svg v-else-if="notification.type === 'danger'" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-red-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
              <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-zinc-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
            <div class="ml-3 w-0 flex-1 pt-0.5">
              <p class="text-sm font-medium text-zinc-900 dark:text-zinc-200">
                {{ notification.title }}
              </p>
              <p class="mt-1 text-sm text-zinc-500 dark:text-zinc-400">
                {{ notification.message }}
              </p>
            </div>
            <div class="ml-4 flex-shrink-0 flex">
              <button class="bg-white dark:bg-zinc-900 rounded-md inline-flex text-zinc-400 hover:text-zinc-500 focus:outline-none focus:ring-transparent focus:ring-offset-transparent focus:ring-blue-500" @click="close">
                <span class="sr-only">Close</span>
                <!-- Heroicon name: solid/x -->
                <svg class="h-5 w-5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                  <path fill-rule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clip-rule="evenodd" />
                </svg>
              </button>
            </div>
          </div>
        </div>
        <div class="h-1 w-full rounded-md bg-white mt-0.5 dark:bg-zinc-900">
          <div class="h-1 transition-all ease-in-out duration-200 rounded-md bg-blue-500 dark:bg-blue-600 progress"></div>
        </div>
      </div>
    </transition>
  </div>
</div>

</template>

<script>
export default {
  data() {
    return {
      cent: 0
    }
  },
  computed: {
    notification() {
      return this.$store.getters.notification
    }
  },
  watch: {
    notification: {
      deep: true,
      handler(newVal, oldVal) {
        if (!oldVal.show && newVal.show) {
          this.triggerCloseTimeout()
        }
      }
    }
  },
  methods: {
    close() {
      this.$store.commit('setResource', {resource: 'notification', value: {
        show: false
        }})
    },
    triggerCloseTimeout() {
      setTimeout(() => {
        this.close()
      }, 5000)
    }
  }
}
</script>

<style scoped>
.progress {
  width: 0;
  animation: runProgress 5s linear forwards 0.1s;
}

@keyframes runProgress {
	0%	{ width: 0; }
	100% { width: 100%; }
}
</style>
