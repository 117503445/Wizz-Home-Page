package serverchan

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"time"
	"wizz-home-page/app/dao"
)

func ResumeRemind() {
	// 7天等于604800000毫秒
	values, err := dao.Resumes.Where("send_time <= ? AND interview_result = 0", gtime.TimestampMilli()-604800000).Array("interviewer_id")
	if err != nil {
		g.Log().Line().Error(err)
	}
	ids := make([]int, 0)
	for _, value := range values {
		ids = append(ids, value.Int())
	}
	for i, id := range ids {
		value, err := dao.Interviewers.Value("serverchan_id", "id", id)
		if err != nil {
			g.Log().Line().Error(err)
		}
		push(value.String(), "test", gconv.String(i)) //to do 内容待完善
		time.Sleep(3 * time.Second)
	}

}
