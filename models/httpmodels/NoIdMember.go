package httpModels

//没有 ID 的 Member
type NoIdMember struct {
	//成员姓名
	Name string `example:"tengye"`
	//成员头像图片的url
	UrlAvatar string `example:"http://wuygewfuyd/weiug.jpg"`
	//入学年份
	SchoolYear int `example:"2017"`
	//个人简介
	Describe string `example:"超级帅的前端dalao"`
	//成员类型,0 - 学生,1 - 导师
	MemberType int `enums:"0,1" example:"0"`
	//如果是老师,则有若干头衔
	TeacherInfo string `example:"创新校长\n制霸西电"`
}
