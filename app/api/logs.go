package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
	"wizz-home-page/library/response"
)

var Log = new(logsApi)

type logsApi struct{}

// @Summary 获取所有日志
// @Tags 日志
// @Accept  json
// @Produce  json
// @Success 200 {array} model.LogsApiRep
// @Router /api/logs [get]
func (*logsApi) ReadAll(r *ghttp.Request) {
	g.Log().Debug("GetAll")
	var logs []model.Logs
	if err := dao.Logs.Structs(&logs); err != nil {
		response.JsonOld(r, 500, "")
	}
	if len(logs) == 0 {
		r.Response.Write("[]")
		r.Exit()
	} else {
		var logsRsp []model.LogsApiRep
		if err := gconv.Structs(logs, &logsRsp); err != nil {
			g.Log().Line().Error(err)
		}
		response.JsonOld(r, 200, logsRsp)
	}
}
