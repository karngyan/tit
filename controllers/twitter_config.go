package controllers

import (
	"encoding/json"
	"net/http"
	"titd/models"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// TwitterConfigController - Operations about TwitterConfigs
type TwitterConfigController struct {
	beego.Controller
}

// Post
// @Title CreateTwitterConfig
// @Description create tconfigs
// @Param	body		body 	models.TwitterConfig	true		"body for tconfig content"
// @Success 200 {object} models.TwitterConfig
// @router / [post]
func (tc *TwitterConfigController) Post() {
	var tconfig models.TwitterConfig
	var err error

	if err = json.Unmarshal(tc.Ctx.Input.RequestBody, &tconfig); err != nil {
		logs.Error(err)
		tc.Ctx.Output.SetStatus(http.StatusUnprocessableEntity)
		return
	}

	if err = tconfig.Insert(); err != nil {
		logs.Error(err)
		tc.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}

	tc.Ctx.Output.SetStatus(http.StatusCreated)
	_ = tc.Ctx.Output.JSON(tconfig, false, false)
}

// GetAll
// @Title GetAll
// @Description get all TwitterConfigs
// @Success 200 {object} models.TwitterConfigsResponse
// @router / [get]
func (tc *TwitterConfigController) GetAll() {

	var tconfigs []*models.TwitterConfig
	var err error

	if tconfigs, err = models.AllTwitterConfigs(); err != nil {
		logs.Error(err)
		tc.Ctx.Output.SetStatus(http.StatusInternalServerError)
		return
	}

	tc.Ctx.Output.SetStatus(http.StatusOK)
	_ = tc.Ctx.Output.JSON(models.TwitterConfigsResponse{TwitterConfigs: tconfigs}, false, false)
}

// Delete
// @Title Delete
// @Description deletes the tconfig
// @Param	tid		path 	string	true		"The tid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 tid is empty
// @router /:tid [delete]
func (tc *TwitterConfigController) Delete() {
	var err error
	var tconfig = &models.TwitterConfig{
		Id: tc.GetString(":tid"),
	}

	if err = tconfig.Delete(); err != nil {
		logs.Error(err)
		tc.Ctx.Output.SetStatus(http.StatusBadRequest)
		return
	}

	tc.Ctx.Output.SetStatus(http.StatusNoContent)
}
