package models

//为之介绍
type Passage struct {
	// ID is passage's id
	ID      int    `json:"id" example:"1"`
	Passage string `example:"这是为之介绍" json:"passage"`
}
