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

// WebsitesStoreItem 存储的站点信息
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
	// WebsitesGroupDefault 默认分组名称
	WebsitesGroupDefault = "default"
)

// List 获取站点列表
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

// Add 添加站点
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
		if item.Group == "" {
			item.Group = WebsitesGroupDefault
		}
	}
	list = append(list, data)
	if err := w.save(list, fileSync); err != nil {
		return 0, err
	}
	return 1, nil
}

// Update 更新站点
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
	if data.Group == "" {
		data.Group = WebsitesGroupDefault
	}
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

// RequestWebsitesDelete 删除请求参数
type RequestWebsitesDelete struct {
	IDS string `json:"ids"`
}

// Delete 删除站点
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

// Groups 获取站点分组列表
func (w *WebsitesModel) Groups(fileSync *util.FileSync) ([]string, error) {
	// 读取
	list := make([]WebsitesStoreItem, 0)
	list, err := w.List(fileSync)
	if err != nil {
		return nil, err
	}
	groups := make([]string,0)
	mapGroups := make(map[string]string)
	for _, item := range list {
		if item.Group == "" {
			continue
		}
		if _,ok:= mapGroups[item.Group];ok {
			continue
		}
		mapGroups[item.Group] = item.Group
		groups = append(groups,item.Group)
	}
	return groups,nil
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