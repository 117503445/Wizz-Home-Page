package models

type Story struct {
	// ID this is userid
	ID    uint   `json:"id" example:"1"`
	TimeStamp uint `json:"time"`
	Name string `json:"name"`
	StoryDescribe string `json:"storyDescribe"`
}
