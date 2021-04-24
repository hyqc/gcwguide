package conf

var ErrorMsg = make(map[int]string)

const (
	Success = 0
	Error   = 1000

	ParamsInvalid          = 4001
	WebsiteListGetError    = 5001
	WebsiteListAddError    = 5002
	WebsiteListUpdateError = 5003
	WebsiteListDeleteError = 5004

	UploadFileError        = 6000
	UploadFileSizeOutRange = 60001
	UploadFileTypeNotAllow = 60002
)

func init() {
	ErrorMsg[Success] = "请求成功"
	ErrorMsg[Error] = "请求失败"

	ErrorMsg[ParamsInvalid] = "无效请求参数"
	ErrorMsg[WebsiteListGetError] = "获取网站失败"
	ErrorMsg[WebsiteListAddError] = "添加网站失败"
	ErrorMsg[WebsiteListUpdateError] = "更新网站失败"
	ErrorMsg[WebsiteListDeleteError] = "删除网站失败"

	ErrorMsg[UploadFileError] = "上传文件失败"
	ErrorMsg[UploadFileSizeOutRange] = "上传文件的大小不能超出%dM"
	ErrorMsg[UploadFileTypeNotAllow] = "上传的文件类型错误%s，应为%s类型"
}
