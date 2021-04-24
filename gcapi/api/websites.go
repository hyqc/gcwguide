package api

import (
	"encoding/json"
	conf "gcapi/config"
	"gcapi/model"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
)

func WebSiteList(c *gin.Context) {
	app := conf.App
	output := conf.JsonOutput{}
	data := make(map[string]interface{})
	fileModel := model.WebsitesModel{}
	list, err := fileModel.List(app.Store.FileSync)
	if err != nil {
		output.Debug = err.Error()
		output.Code = conf.WebsiteListGetError
		data["rows"] = nil
		data["total"] = 0
		output.Data = data
		output.Message = conf.ErrorMsg[conf.WebsiteListGetError]
		response(c, output)
		return
	}
	if list == nil {
		data["rows"] = nil
		data["total"] = 0
		output.Data = data
	}else{
		mapList := make(map[string][]model.WebsitesStoreItem)
		for _, item := range list {
			if item.Group == "" {
				item.Group = model.WebsitesGroupDefault
			}
			if _, ok := mapList[item.Group];ok==false {
				mapList[item.Group] = make([]model.WebsitesStoreItem,0)
			}
			mapList[item.Group] = append(mapList[item.Group],item)
		}
		data["rows"] = mapList
		data["total"] = len(list)
		output.Data = data
	}
	output.Debug = ""
	output.Code = conf.Success
	output.Message = conf.ErrorMsg[conf.Success]
	response(c, output)
	return
}

func WebSiteAdd(c *gin.Context) {
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
	data := model.WebsitesStoreItem{}
	if err := json.Unmarshal(body, &data); err != nil {
		output.Debug = err.Error()
		output.Code = conf.WebsiteListAddError
		output.Data = nil
		output.Message = conf.ErrorMsg[conf.WebsiteListAddError]
		response(c, output)
		return
	}
	fileModel := model.WebsitesModel{}
	success := 0
	if success,err = fileModel.Add(app.Store.FileSync, data, app.Store.BackupsDir); err != nil {
		output.Debug = err.Error()
		output.Code = conf.WebsiteListAddError
		output.Data = nil
		output.Message = conf.ErrorMsg[conf.WebsiteListAddError]
		response(c, output)
		return
	}
	info := make(map[string]interface{})
	info["success"] = success
	output.Debug = ""
	output.Code = conf.Success
	output.Data = info
	output.Message = conf.ErrorMsg[conf.Success]
	response(c, output)
	return
}

func WebSiteUpdate(c *gin.Context) {
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
	data := model.WebsitesStoreItem{}
	if err := json.Unmarshal(body, &data); err != nil {
		output.Debug = err.Error()
		output.Code = conf.WebsiteListUpdateError
		output.Data = nil
		output.Message = conf.ErrorMsg[conf.WebsiteListUpdateError]
		response(c, output)
		return
	}
	fileModel := model.WebsitesModel{}
	success := 0
	if success,err = fileModel.Update(app.Store.FileSync, data, app.Store.BackupsDir); err != nil {
		output.Debug = err.Error()
		output.Code = conf.WebsiteListUpdateError
		output.Data = nil
		output.Message = conf.ErrorMsg[conf.WebsiteListUpdateError]
		response(c, output)
		return
	}
	info := make(map[string]interface{})
	info["success"] = success
	output.Debug = ""
	output.Code = conf.Success
	output.Data = info
	output.Message = conf.ErrorMsg[conf.Success]
	response(c, output)
	return
}

func WebSiteDelete(c *gin.Context) {
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
	data := model.RequestWebsitesDelete{}
	if err := json.Unmarshal(body, &data); err != nil {
		output.Debug = err.Error()
		output.Code = conf.WebsiteListDeleteError
		output.Data = nil
		output.Message = conf.ErrorMsg[conf.WebsiteListDeleteError]
		response(c, output)
		return
	}
	data.IDS = strings.TrimSpace(data.IDS)
	idsArr := strings.Split(data.IDS, ",")
	if len(idsArr) == 0 {
		output.Debug = "params ids is invalid"
		output.Code = conf.ParamsInvalid
		output.Data = nil
		output.Message = conf.ErrorMsg[conf.ParamsInvalid]
		response(c, output)
		return
	}

	fileModel := model.WebsitesModel{}
	success := 0
	if success,err = fileModel.Delete(app.Store.FileSync, idsArr, app.Store.BackupsDir); err != nil {
		output.Debug = err.Error()
		output.Code = conf.WebsiteListDeleteError
		output.Data = nil
		output.Message = conf.ErrorMsg[conf.WebsiteListDeleteError]
		response(c, output)
		return
	}
	info := make(map[string]interface{})
	info["success"] = success
	output.Debug = ""
	output.Code = conf.Success
	output.Data = info
	output.Message = conf.ErrorMsg[conf.Success]
	response(c, output)
	return
}

func WebsiteGroups(c *gin.Context)  {
	app := conf.App
	output := conf.JsonOutput{}
	fileModel := model.WebsitesModel{}
	list, err := fileModel.Groups(app.Store.FileSync)
	if  err != nil {
		output.Debug = err.Error()
		output.Code = conf.WebsiteListDeleteError
		output.Data = nil
		output.Message = conf.ErrorMsg[conf.WebsiteListDeleteError]
		response(c, output)
		return
	}
	output.Debug = ""
	output.Code = conf.Success
	output.Data = list
	output.Message = conf.ErrorMsg[conf.Success]
	response(c, output)
	return
}
