package conf

var ErrorMsg = make(map[int]string)

const (
	Success                = 0
	ParamsInvalid          = 4001
	WebsiteListGetError    = 5001
	WebsiteListAddError    = 5002
	WebsiteListUpdateError = 5003
	WebsiteListDeleteError = 5004
)

func init() {
	ErrorMsg[Success] = "请求成功"
	ErrorMsg[ParamsInvalid] = "无效请求参数"
	ErrorMsg[WebsiteListGetError] = "获取网站失败"
	ErrorMsg[WebsiteListAddError] = "添加网站失败"
	ErrorMsg[WebsiteListUpdateError] = "更新网站失败"
	ErrorMsg[WebsiteListDeleteError] = "删除网站失败"
}
