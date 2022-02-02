package models

import (
	"encoding/json"
	"time"
	"titd/modules/global"

	"github.com/google/uuid"
)

type TweetsResponse struct {
	Tweets []*Tweet `json:"tweets"`
}

type ETweet struct {
	Content      string `json:"content"`
	SendAt       string `json:"sendAt"`                 // based on go-natural date
	ConfigPrefix string `json:"configPrefix,omitempty"` // TODO: Add support for this
}

type Tweet struct {
	Id      string    `json:"id"`      // uuid
	Content string    `json:"content"` // tweet content
	Done    bool      `json:"done"`    // tweeted successfully
	SendAt  time.Time `json:"sendAt"`  // send time, local to the machine
	Created time.Time `json:"created"` // when was it created
	Updated time.Time `json:"updated"` // when was it last updated
}

func (t *Tweet) Insert() error {

	t.Created = time.Now()
	t.Updated = time.Now()
	t.Id = uuid.NewString()

	val, err := json.Marshal(*t)
	if err != nil {
		return err
	}

	if err := global.DataStore.Put(tweetsBucketName, t.Id, val); err != nil {
		return err
	}

	return nil
}

func (t *Tweet) MarkAsDone() error {

	t.Updated = time.Now()
	t.Done = true

	val, err := json.Marshal(*t)
	if err != nil {
		return err
	}

	if err := global.DataStore.Put(tweetsBucketName, t.Id, val); err != nil {
		return err
	}

	return nil
}

func (t *Tweet) Delete() error {

	var err error

	if err = global.DataStore.Delete(tweetsBucketName, t.Id); err != nil {
		return err
	}

	return nil
}

func AllTweets() (tweets []*Tweet, err error) {

	tweets = make([]*Tweet, 0)
	tb := make([][]byte, 0)
	if err = global.DataStore.GetAll(tweetsBucketName, &tb); err != nil {
		return tweets, err
	}

	for _, tbyte := range tb {
		var tweet Tweet
		if err = json.Unmarshal(tbyte, &tweet); err != nil {
			return tweets, err
		}

		tweets = append(tweets, &tweet)
	}

	return tweets, err
}
