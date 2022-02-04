package models

import (
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/go-co-op/gocron"
)

func StartTweetScheduler() error {

	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(20).Seconds().Do(func() {
		tconfigs, err := AllTwitterConfigs()
		if err != nil {
			logs.Warn("failed to get twitter configs")
			logs.Error(err)
			return
		}

		if len(tconfigs) <= 0 {
			logs.Warn("no twitter config present, not starting the tweet scheduler")
			return
		}

		var tconf = tconfigs[0]

		config := oauth1.NewConfig(tconf.ApiKey, tconf.ApiKeySecret)
		token := oauth1.NewToken(tconf.AccessToken, tconf.AccessTokenSecret)
		httpClient := config.Client(oauth1.NoContext, token)

		// Twitter client
		client := twitter.NewClient(httpClient)

		// fetch all tweets
		tweets, err := AllTweets()
		if err != nil {
			logs.Warn("failed to get all tweets for scheduler")
			logs.Error(err)
			return
		}

		for _, tweet := range tweets {
			if !tweet.Done {
				// tweet is not done
				if tweet.SendAt.After(time.Now()) || tweet.SendAt.Equal(time.Now()) {
					// Send a Tweet
					if stweet, resp, err := client.Statuses.Update(tweet.Content, nil); err == nil {
						logs.Info("tweet sent: ", stweet)
						logs.Info("resp: ", resp)
						// just delete the tweet once sent
						_ = tweet.Delete()
					} else {
						logs.Info("failed to send tweet")
						logs.Warn("resp", resp)
						logs.Warn("err", err)
					}
				}
			}
		}
	})

	if err != nil {
		logs.Error(err)
		return err
	}

	s.StartAsync()
	logs.Info("started tweet scheduler ...")
	return nil
}

func Init() {
	go func() {
		err := StartTweetScheduler()
		if err != nil {
			logs.Error(err)
		}
	}()
}
