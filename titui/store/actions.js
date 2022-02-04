import { Notification } from "~/constants"

export default {

  createTweet({state, commit}, data) {
    return new Promise((resolve, reject) => {
      this.$axios.post('/tweets', data)
        .then((resp) => {
          const tweet = resp.data
          commit('setItem', {resource: 'tweets', item: tweet, id: tweet.id})
          resolve(tweet)
        }).catch((err) => {
          reject(err)
      })
    })
  },

  createTwitterConfig({state, commit}, data) {
    return new Promise((resolve, reject) => {
      this.$axios.post('/tconfigs', data)
        .then((resp) => {
          const tconfig = resp.data
          commit('setItem', {resource: 'tconfigs', item: tconfig, id: tconfig.id})
          resolve(tconfig)
        }).catch((err) => {
          reject(err)
      })
    })
  },

  deleteTweet({state, commit}, { id }) {
    return new Promise((resolve, reject) => {
      this.$axios.delete('/tweets/' + id)
        .then(() => {
          commit('deleteItem', {resource: 'tweets', id})
          resolve()
        }).catch((err) => {
          reject(err)
      })
    })
  },

  fetchAllTweets({state, commit}) {
    return new Promise((resolve, reject) => {
      this.$axios.get('/tweets')
        .then((resp) => {
          const tweets = resp.data.tweets
          tweets.forEach((tweet) => {
            commit('setItem', {resource: 'tweets', item: tweet, id: tweet.id})
          })
          resolve(tweets)
        }).catch((err) => {
          reject(err)
      })
    })
  },

  fetchAllTwitterConfigs({state, commit}) {
    return new Promise((resolve, reject) => {
      this.$axios.get('/tconfigs')
        .then((resp) => {
          const tconfigs = resp.data.twitterConfigs
          tconfigs.forEach((tconfig) => {
            commit('setItem', {resource: 'tconfigs', item: tconfig, id: tconfig.id})
          })
          resolve(tconfigs)
        }).catch((err) => {
          reject(err)
      })
    })
  },

  info({ commit }, { title, message }) {
    commit("setResource", {
      resource: "notification", value: {
        type: Notification.Info, title, message, show: true
      }
    })
  },

  success({ commit }, { title, message }) {
    commit("setResource", {
      resource: "notification",
      value: { type: Notification.Success, title, message, show: true }
    })
  },

  danger({ commit }, { title, message }) {
    commit("setResource", {
      resource: "notification",
      value: { type: Notification.Danger, title, message, show: true }
    })
  },
}
