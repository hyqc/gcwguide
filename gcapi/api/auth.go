package api

import (
	"encoding/json"
	conf "gcapi/config"
	"gcapi/extend/util"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type RequestLoginParams struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

func Login(c *gin.Context) {
	app := conf.App
	output := conf.JsonOutput{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		output.Debug = err.Error()
		output.Code = conf.ParamsInvalid
		output.Data = nil
		output.Message = conf.ErrorMsg[conf.ParamsInvalid]
		response(c, output)
		return
	}
	data := RequestLoginParams{}
	if err := json.Unmarshal(body, &data); err != nil {
		output.Debug = err.Error()
		output.Code = conf.AuthLoginAccountInvalid
		output.Data = nil
		output.Message = conf.ErrorMsg[conf.AuthLoginAccountInvalid]
		response(c, output)
		return
	}
	if data.Name != app.Account.Name || data.Pwd != app.Account.Password {
		output.Debug = "account or password invalid"
		output.Code = conf.AuthLoginAccountInvalid
		output.Data = nil
		output.Message = conf.ErrorMsg[conf.AuthLoginAccountInvalid]
		response(c, output)
		return
	}
	j := util.JWT{}
	token, err := j.Make(app.Account.Name,app.Account.Password,app.Account.CookieExpireSeconds)
	if err != nil {
		output.Debug = err.Error()
		output.Code = conf.Error
		output.Data = nil
		output.Message = conf.ErrorMsg[conf.Error]
		response(c, output)
		return
	}
	res := make(map[string]interface{})
	res["token"] = token
	res["expire"] = app.Account.CookieExpireSeconds
	output.Debug = ""
	output.Code = conf.Success
	output.Data = res
	output.Message = conf.ErrorMsg[conf.Success]
	response(c, output)
	return
}
