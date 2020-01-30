package models

//产品
type Product struct {
	//product's id
	ID uint `json:"id" example:"1"`
	// 产品名称
	Name string `example:"一起来开黑"`
	//一句话介绍
	LittleDescribe string `example:"一键约游戏"`
	//项目介绍
	Describe string `example:"大学生喜闻乐见的跨校交友平台，一键匹配聊天约游戏，一起来开黑帮你聊天&游戏两不误~"`
	//合作方
	Partner string `example:"腾讯"`
	//项目类型
	ProjectType int `example:"" enums:"0,1"`

	//项目图标的 Base64 字符串
	ImgAvatar string `example:""`
	//项目图标的 Url
	UrlAvatar string `example:""`

	//合作方Logo的 Base64 字符串
	ImgPartnerLogo string `example:""`
	//合作方Logo的 Url
	UrlPartnerLogo string `example:""`

	//项目背景图片的 Id
	IdBackGroundImg int `example:""`
	//项目背景图片的 Url
	UrlBackground string `example:""`

	//项目截图的 Base64 字符串
	ImgScreenshot []string `gorm:"type:string[]" example:""`
	//项目截图的 Url
	UrlScreenshot []string `gorm:"type:string[]" example:""`

	//二维码的 Base64 字符串
	ImgProCode string `example:""`
	//二维码的 Url
	UrlProCode string `example:""`
}
