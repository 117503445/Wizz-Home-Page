package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"wizz-home-page/app/model"
	"wizz-home-page/app/service"
	"wizz-home-page/library/response"
)

var Log = new(logsApi)

type logsApi struct{}

// @Summary 获取所有日志
// @Tags 日志
// @Accept  json
// @Produce  json
// @Param   pageSize	query	int true  "一页几条"
// @Param   pageIndex	query	int true  "页数"
// @Success 200 {array} model.LogsApiRep
// @Router /api/logs [get]
func (*logsApi) ReadAll(r *ghttp.Request) {
	pageSize := r.GetInt("pageSize")
	pageIndex := r.GetInt("pageIndex")
	responsePage, err := service.SortLogs(pageIndex, pageSize)

	if err != nil {
		response.JsonOld(r, 500, err)
	} else {
		var logRsp []model.LogsApiRep
		if err = gconv.Structs(responsePage.Content, &logRsp); err != nil {
			g.Log().Line().Error(err)
			response.JsonOld(r, 500, err)
		} else {
			responsePage.Content = logRsp
			response.JsonOld(r, 200, responsePage)
		}

	}
}
