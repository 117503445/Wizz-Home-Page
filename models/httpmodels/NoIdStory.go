package httpModels

//没有 ID 的 Story
type NoIdStory struct {
	TimeStamp     int64  `example:"1580397149" json:"timeStamp"`
	Name          string `example:"为之诞生" json:"name"`
	StoryDescribe string `example:"11月,前身 TgClub 诞生" json:"storyDescribe"`
}
