package models

type ServerLog struct {
	//log's id
	ID int `json:"id" example:"1"`
	//后端处理时间
	Latency       float64
	//请求的HTTP方法
	RequestMethod string
	//请求的Header
	RequestHeader string
	//请求的Body
	RequestBody	string
	//用户名,只有 POST,PUT 这些需要token的接口才非空
	Username      string
	//请求的URI
	RequestURI    string
	//返回值
	ResponseCode  string
	//返回的Body
	ResponseBody  string
}
