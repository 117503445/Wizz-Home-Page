package models
//成员
type Member struct {
	// ID is member's id
	ID			uint   `json:"id"`
	Name		string
	UrlAvatar	string
	SchoolYear	int
	Describe	string
	MemberType	int
	TeacherInfo	[]string `gorm:"type:string[]"`
}
