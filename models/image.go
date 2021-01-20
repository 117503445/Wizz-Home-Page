package models

//图片
type Image struct {
	//image's id
	ID int `json:"id" example:"1"`
	//成员类型,0 - Vread,1 - Vtalk，2 - 轮播图
	ImageType int `enums:"0,1,2" example:"0"`
	//图片的url
	ImageUrl string `example:"http://wuygewfuyd/weiug.jpg"`
}
