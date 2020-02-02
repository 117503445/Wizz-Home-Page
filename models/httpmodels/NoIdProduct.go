package httpModels
//没有 ID 的 Product
type NoIdProduct struct {
	// 产品名称
	Name string `example:"一起来开黑"`
	//一句话介绍
	LittleDescribe string `example:"一键约游戏"`
	//项目介绍
	Describe string `example:"大学生喜闻乐见的跨校交友平台，一键匹配聊天约游戏，一起来开黑帮你聊天&游戏两不误~"`
	//合作方
	Partner string `example:"腾讯"`
	//项目类型,0 - 校企合作,1 - 校园合作,2 - 校内自研
	ProjectType int `example:"" enums:"0,1"`
	//项目图标的 Url
	UrlAvatar string `example:"http://qiniu.com/UrlAvatar.png"`
	//合作方Logo的 Url
	UrlPartnerLogo string `example:"http://qiniu.com/UrlPartnerLogo.png"`
	//项目背景图片的 Url
	UrlBackground string `example:"http://qiniu.com/UrlBackground.png"`
	//项目截图的 Url
	UrlScreenshot string `example:"http://qiniu.com/UrlScreenshot.png"`
	//二维码的 Url
	UrlProCode string `example:"http://qiniu.com/UrlProCode.png"`
}
