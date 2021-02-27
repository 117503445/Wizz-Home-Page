package service

import (
	"wizz-home-page/app/dao"
	"wizz-home-page/app/model"
	"wizz-home-page/library/response"
)

func SortLogs(Page int, pageSize int) (*response.Page, error) {
	n1 := (Page - 1) * pageSize
	n2 := Page * pageSize

	totalSize, err := dao.Logs.Count()
	if err != nil {
		return nil, err
	}

	var logs []model.Logs
	err = dao.Logs.Limit(n1, n2).Structs(&logs)
	if err != nil {
		return nil, err
	}

	p := response.Page{
		PageNum:    Page,
		TotalPages: totalSize/pageSize + 1,
		PageSize:   pageSize,
		TotalSize:  totalSize,
		Content:    logs,
	}
	return &p, nil
}
