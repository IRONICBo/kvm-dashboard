package services

import (
	"kvm-dashboard/consts"
	"kvm-dashboard/model"
)

func (svc *Service) GetAlertInfo(uuid string, pagesize, page int) *model.Page {

	// get alert info
	pageData := svc.Dao.ReadAlertData(uuid, consts.PERIOD_30D, pagesize, page)

	return pageData
}
