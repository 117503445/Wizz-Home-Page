package serverchan

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
	"time"
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
)

func ResumeRemind() {
	overtime := g.Cfg().GetInt64("remind.overtime")
	var resumes []model.Resumes
	err := dao.Resumes.Where("send_time <= ? AND interview_result = 0", gtime.TimestampMilli()-overtime*86400000).Structs(&resumes)
	if err != nil {
		g.Log().Line().Error(err)
	}
	var interviewer model.Interviewers
	for _, resume := range resumes {
		err = dao.Interviewers.Where("Id", resume.InterviewerId).Struct(&interviewer)
		if err != nil {
			g.Log().Line().Error(err)
		}
		//experienceStr := ""
		//if resume.Experience == 0 {
		//experienceStr = "无项目经历"
		//} else {
		//experienceStr = "有项目经历"
		//}
		url := GetUrl(resume.Id)
		title := fmt.Sprintf("未提交%v的面评", resume.Name)
		fileUrl := fmt.Sprintf("\n\n<a href=\"%v\">简历下载链接</a>", resume.FileUrl)
		if resume.FileUrl == "" {
			fileUrl = ""
		}
		content := fmt.Sprintf("距离投递者投递简历已经过去一周啦~请尽快到此url填写投递者的面试结果以及面试评价，否则将会每天提醒哦\n\n（ps:不需要面试也需要选择是否通过，填写理由哦）\n\n<a href=\"%v\">url</a>\n\n---\n\n【快速回顾%v的简历】\n\n%v %v级 %v\n\n联系电话：%v\n微信：%v\nqq：%v\n\n%v%v", url, resume.Name, resume.Name, resume.Grade, resume.CollegeMajor, resume.TelephoneNumber, resume.WechatNumber, resume.QqNumber, resume.Describe, fileUrl)
		Push(interviewer.ServerchanId, title, content)
		var message model.Messages
		err = dao.Messages.Where("resume_id", resume.Id).Struct(&message)
		if err != nil {
			message.Content = fmt.Sprintf("面试官%v未及时处理简历", interviewer.Name)
			message.ResumeId = resume.Id
			message.CreateTime = gtime.TimestampMilli()
			message.ReadStatus = 0
			if _, err = dao.Messages.Insert(message); err != nil {
				g.Log().Line().Error(err)
			}
		} else {
			message.CreateTime = gtime.TimestampMilli()
			if _, err = dao.Messages.Data(message).Where("id", message.Id).Update(); err != nil {
				g.Log().Line().Error(err)
			}
		}
		time.Sleep(3 * time.Second)
	}

}

func GetUrl(id int) string {
	username := "interviewer"
	password := gfile.GetContents("./tmp/password/admin.txt")
	response := g.Client().ContentJson().PostContent("http://localhost:80/api/auth/login", g.Map{
		"username": username,
		"password": password,
	})
	token := ""
	if js, err := gjson.DecodeToJson(response); err != nil {
		g.Log().Line().Error(err)
	} else {
		token = js.GetString("token")
	}

	url := fmt.Sprintf("https://wizzstudio.com/#/pass?id=%v&jwt=%v", id, token)
	g.Log().Line().Debug(url)
	return url
}
