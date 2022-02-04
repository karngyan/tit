package models

import (
	"encoding/json"
	"time"
	"titd/modules/global"
)

type TwitterConfigsResponse struct {
	TwitterConfigs []*TwitterConfig `json:"twitterConfigs"`
}

// TwitterConfig - can be multiple
type TwitterConfig struct {
	Id                string    `json:"id"`
	ApiKey            string    `json:"apiKey"`
	ApiKeySecret      string    `json:"apiKeySecret"`
	BearerToken       string    `json:"bearerToken"`
	AccessToken       string    `json:"accessToken"`
	AccessTokenSecret string    `json:"accessTokenSecret"`
	Default           bool      `json:"default"`
	Created           time.Time `json:"created"` // when was it created
	Updated           time.Time `json:"updated"` // when was it last updated
}

// Insert - last inserted config automatically becomes default
// TODO: Add support for use of multiple configs in API
func (t *TwitterConfig) Insert() error {

	t.Created = time.Now()
	t.Updated = time.Now()
	//t.Id = uuid.NewString()
	t.Id = "fixed_id_no_support_for_multiple_config"
	t.Default = true

	val, err := json.Marshal(*t)
	if err != nil {
		return err
	}

	if err := global.DataStore.Put(twitterConfigBucketName, t.Id, val); err != nil {
		return err
	}

	return nil
}

func (t *TwitterConfig) Delete() error {

	var err error

	if err = global.DataStore.Delete(twitterConfigBucketName, t.Id); err != nil {
		return err
	}

	return nil
}

func AllTwitterConfigs() (tconfigs []*TwitterConfig, err error) {

	tconfigs = make([]*TwitterConfig, 0)
	tb := make([][]byte, 0)
	if err = global.DataStore.GetAll(twitterConfigBucketName, &tb); err != nil {
		return tconfigs, err
	}

	for _, tbyte := range tb {
		var tweet TwitterConfig
		if err = json.Unmarshal(tbyte, &tweet); err != nil {
			return tconfigs, err
		}

		tconfigs = append(tconfigs, &tweet)
	}

	return tconfigs, err
}
