package controllers

import (
	"encoding/json"
	"net/http"
	"time"
	"titd/models"

	"github.com/tj/go-naturaldate"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// TweetController - Operations about Tweets
type TweetController struct {
	beego.Controller
}

// Post
// @Title CreateTweet
// @Description create tweets
// @Param	body		body 	models.Tweet	true		"body for tweet content"
// @Success 200 {object} models.Tweet
// @router / [post]
func (tc *TweetController) Post() {
	var etweet models.ETweet
	var err error

	if err = json.Unmarshal(tc.Ctx.Input.RequestBody, &etweet); err != nil {
		logs.Error(err)
		tc.Ctx.Output.SetStatus(http.StatusUnprocessableEntity)
		return
	}

	tweet := &models.Tweet{
		Content: etweet.Content,
	}

	if tweet.SendAt, err = naturaldate.Parse(etweet.SendAt, time.Now()); err != nil {
		logs.Info(err)
		tc.Ctx.Output.SetStatus(http.StatusBadRequest)
		return
	}

	if err = tweet.Insert(); err != nil {
		logs.Error(err)
		tc.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}

	tc.Ctx.Output.SetStatus(http.StatusCreated)
	_ = tc.Ctx.Output.JSON(tweet, false, false)
}

// GetAll
// @Title GetAll
// @Description get all Tweets
// @Success 200 {object} models.TweetsResponse
// @router / [get]
func (tc *TweetController) GetAll() {

	var tweets []*models.Tweet
	var err error

	if tweets, err = models.AllTweets(); err != nil {
		logs.Error(err)
		tc.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}

	tc.Ctx.Output.SetStatus(http.StatusOK)
	_ = tc.Ctx.Output.JSON(models.TweetsResponse{Tweets: tweets}, false, false)
}

// Delete
// @Title Delete
// @Description deletes the tweet
// @Param	tid		path 	string	true		"The tid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 tid is empty
// @router /:tid [delete]
func (tc *TweetController) Delete() {
	var err error
	var tweet = &models.Tweet{
		Id: tc.GetString(":tid"),
	}

	if err = tweet.Delete(); err != nil {
		logs.Error(err)
		tc.Ctx.Output.SetStatus(http.StatusBadRequest)
		return
	}

	tc.Ctx.Output.SetStatus(http.StatusNoContent)
}
