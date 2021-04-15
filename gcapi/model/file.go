package model

import (
	"encoding/json"
	"gcapi/extend/util"
	"strings"
	"time"
)

type WebsitesModel struct {
	Model
}

type WebsitesStoreItem struct {
	ID     string `json:"id"`     // 网站ID
	Group  string `json:"group"`  // 网站分组
	Name   string `json:"name"`   // 网站名称
	Pic    string `json:"pic"`    // 网站图标
	Host   string `json:"host"`   // 网站地址
	Desc   string `json:"desc"`   // 网站描述
	Create string `json:"create"` // 添加时间
	Update string `json:"update"` // 更新时间
}

const (
	WebsitesGroupDefault = "default"
)

func (w *WebsitesModel) List(fileSync *util.FileSync) ([]WebsitesStoreItem, error) {
	content, err := fileSync.ReadJSON()
	if err != nil {
		return nil, err
	}
	if content == nil {
		return nil, nil
	}
	list := make([]WebsitesStoreItem, 0)
	if err = json.Unmarshal(content, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (w *WebsitesModel) Add(fileSync *util.FileSync, data WebsitesStoreItem, backupsDir string) (int, error) {
	// 读取
	list := make([]WebsitesStoreItem, 0)
	list, err := w.List(fileSync)
	if err != nil {
		return 0, err
	}
	data.Create = time.Now().Format(time.RFC3339)
	data.Update = data.Create
	data.Group = strings.TrimSpace(data.Group)
	if data.Group == "" {
		data.Group = WebsitesGroupDefault
	}
	data.ID = util.CreateMD5(data.Host, true)
	for _, item := range list {
		if item.ID == data.ID {
			return 0, nil
		}
	}
	list = append(list, data)
	if err := w.save(list, fileSync); err != nil {
		return 0, err
	}
	return 1, nil
}

func (w *WebsitesModel) Update(fileSync *util.FileSync, data WebsitesStoreItem, backupsDir string) (int, error) {
	list := make([]WebsitesStoreItem, 0)
	list, err := w.List(fileSync)
	if err != nil {
		return 0, err
	}
	// 遍历更新
	oldID := data.ID
	data.Update = time.Now().Format(time.RFC3339)
	data.ID = util.CreateMD5(data.Host, true)
	for index, item := range list {
		if item.ID == oldID {
			list[index] = data
		}
	}
	if err := w.save(list, fileSync); err != nil {
		return 0, err
	}
	return 1, nil
}

type RequestWebsitesDelete struct {
	IDS string `json:"ids"`
}

func (w *WebsitesModel) Delete(fileSync *util.FileSync, ids []string, backupsDir string) (int, error) {
	// 读取
	list := make([]WebsitesStoreItem, 0)
	list, err := w.List(fileSync)
	if err != nil {
		return 0, err
	}
	// 遍历删除
	newList := make([]WebsitesStoreItem, 0)
	success := 0
	for _, item := range list {
		if util.StringInArray(item.ID, ids) == false {
			newList = append(newList, item)
		} else {
			success++
		}
	}
	if err := w.save(newList, fileSync); err != nil {
		return 0, err
	}
	return success, nil
}

func (w *WebsitesModel) save(list []WebsitesStoreItem, fileSync *util.FileSync) error {
	if len(list) == 0 {
		return nil
	}
	content, err := json.Marshal(list)
	if err != nil {
		return err
	}
	if err := fileSync.CoverJSON(content); err != nil {
		return err
	}
	return nil
}