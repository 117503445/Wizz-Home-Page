package models

//Article
type Article struct {
	ID         int    `json:"id" example:"1"`
	Title      string `example:"为之的历程" json:"title"`
	ArticleUrl string `example:"http://weixin.com/wizzpassage" json:"articleUrl"`
}
