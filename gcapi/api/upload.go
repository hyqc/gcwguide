package api

import (
	"errors"
	"fmt"
	conf "gcapi/config"
	"gcapi/extend/util"
	"github.com/gin-gonic/gin"
	"log"
	"path/filepath"
	"strings"
)

// 上传单个图片
func Image(c *gin.Context) {
	app := conf.App
	output := conf.JsonOutput{}
	_, headers, err := c.Request.FormFile("file")
	if err != nil {
		log.Println(err)
		output.Debug = err.Error()
		output.Code = conf.UploadFileError
		output.Data = nil
		output.Message = conf.ErrorMsg[conf.UploadFileError]
		response(c, output)
		return
	}
	//headers.Size 获取文件大小
	if headers.Size > app.Static.Upload.Maxsize {
		errMsg := fmt.Sprintf(conf.ErrorMsg[conf.UploadFileSizeOutRange], int(app.Static.Upload.Maxsize/1024/1024))
		err := errors.New(errMsg)
		log.Println(err)
		output.Debug = err.Error()
		output.Code = conf.UploadFileSizeOutRange
		output.Data = nil
		output.Message = errMsg
		response(c, output)
		return
	}

	fileType := headers.Header.Get("Content-Type")
	if strings.HasPrefix(fileType, "image/") == false {
		errMsg := fmt.Sprintf(conf.ErrorMsg[conf.UploadFileTypeNotAllow], fileType, "image/*")
		err := errors.New(errMsg)
		log.Println(err)
		output.Debug = err.Error()
		output.Code = conf.UploadFileTypeNotAllow
		output.Data = nil
		output.Message = errMsg
		response(c, output)
		return
	}
	fileUtil := util.FileUtil{}
	if err := fileUtil.CreateFolderIfNotExist(getStaticUploadImagePath(), 0755); err != nil {
		log.Println(err)
		output.Debug = err.Error()
		output.Code = conf.UploadFileError
		output.Data = nil
		output.Message = err.Error()
		response(c, output)
		return
	}
	fileNewName := makeNewFileName(headers.Filename)
	if err := c.SaveUploadedFile(headers, fileNewName); err != nil {
		log.Println(err)
		output.Debug = err.Error()
		output.Code = conf.UploadFileError
		output.Data = nil
		output.Message = err.Error()
		response(c, output)
		return
	}
	data := make(map[string]interface{})
	data["url"] = makeNewFileUrl(fileNewName)
	output.Debug = ""
	output.Code = conf.Success
	output.Data = data
	output.Message = conf.ErrorMsg[conf.Success]
	response(c, output)
	return
}

func makeNewFileName(filename string) string {
	return getStaticUploadImagePath() + util.CreateMD5(filename, true) + filepath.Ext(filename)
}

func makeNewFileUrl(filepath string) string {
	return conf.App.Static.Upload.BaseUrl + "/"+ filepath
}

func getStaticPath() string {
	return conf.App.Static.Static + "/"
}

func getStaticUploadImagePath() string {
	return getStaticPath() + conf.App.Static.Upload.Path + "/"
}
