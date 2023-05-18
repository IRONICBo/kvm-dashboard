package services

import "kvm-dashboard/dao"

type Service struct {
	Dao *dao.InfluxDB
}

func NewService() *Service {
	return &Service{
		Dao: dao.GetInfluxDBClient(),
	}
}
