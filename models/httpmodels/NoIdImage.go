package httpModels

//没有ID的图片
type NoIdImage struct {
	//成员类型,0 - Vread,1 - Vtalk，2-轮播图
	ImageType int `enums:"0,1,2" example:"0"`
	//图片的url
	ImageUrl string `example:"http://wuygewfuyd/weiug.jpg"`
}
