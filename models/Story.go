package models

type Story struct {
	ID    uint   `json:"id"`
	TimeStamp uint `json:"time"`
	Name string `json:"name"`
	StoryDescribe string `json:"storyDescribe"`
}
