package models
//成员
type Member struct {
	//member's id
	ID			int   `json:"id"`
	//成员姓名
	Name		string
	//成员头像图片的url
	UrlAvatar	string
	//入学年份
	SchoolYear	int
	//个人简介
	Describe	string
	//成员类型,0 - 学生,1 - 导师
	MemberType	int	`enums:"0,1"`
	//如果是老师,则有若干头衔
	TeacherInfo	[]string `gorm:"type:string[]" json:"omitempty"`
}
