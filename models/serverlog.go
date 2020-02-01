package models

type ServerLog struct {
	//log's id
	ID int `json:"id" example:"1"`
	//收到请求的时间,精准到秒的格林威治时间戳
	TimeStamp	int64
	//请求的HTTP方法
	RequestMethod string
	//用户名,只有 POST,PUT 这些需要token的接口才非空
	Username      string
	//请求的URI
	RequestURI    string
	//返回状态码
	ResponseCode  string
	//请求POST,PUT or DELETE 的Model中的name字段
	ModelName	string
}
