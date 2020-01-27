package models

type Member struct {
	ID			uint   `json:"id"`
	Name		string
	UrlAvatar	string
	SchoolYear	int
	Describe	string
	MemberType	int
	TeacherInfo	string
}
