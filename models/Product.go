package models

type Product struct {
	ID    			uint   `json:"id"`
	UrlAvatar    	string
	LittleDescribe 	string
	Describe		string
	Partner			string
	UrlBackground	string
	UrlScreenshot	string
	ProjectType		string
}