package response

//Page 返回分页 data
type Page struct {
	PageNum    int         `json:"pageNum"`    // 现在是第几页
	TotalPages int         `json:"totalPages"` // 总共有几页
	PageSize   int         `json:"pageSize"`   // 满的一页有几个对象
	TotalSize  int         `json:"totalSize"`  // 总共有几个对象
	Content    interface{} `json:"content"`    // 这页的对象数组
}
