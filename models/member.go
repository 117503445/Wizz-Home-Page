package models

//成员
type Member struct {
	//member's id
	ID int `json:"id" example:"1"`
	//成员姓名
	Name string `example:"tengye"`
	//成员头像图片的url
	UrlAvatar string `example:"http://wuygewfuyd/weiug.jpg"`
	//入学年份
	SchoolYear int `example:"2017"`
	//个人简介
	Describe string `example:"超级帅的前端dalao"`
	//成员类型,老师 - 0;前端 - 1;产品 - 2;后端 - 3;运营 - 4;
	MemberType int `enums:"0,1,2,3,4" example:"0"`
	//如果是老师,则有若干头衔
	TeacherInfo string `example:"创新校长\n制霸西电"`
}
