package httpModels

//没有 ID 的 Article
type NoIdArticle struct {
	Title      string `example:"为之的历程" json:"title"`
	ArticleUrl string `example:"http://weixin.com/wizzpassage" json:"articleUrl"`
}
