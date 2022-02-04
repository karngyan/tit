export default {
  sortedTweets(state) {
    return Object.values(state.tweets).sort((a, b) => new Date(b.created) - new Date(a.created))
  },
  tconfigs(state) {
    return Object.values(state.tconfigs)
  },
  notification(state) {
    return state.notification;
  }
}
