package dao

import (
	"context"
	"fmt"
	"kvm-dashboard/consts"
	"kvm-dashboard/model"
	"kvm-dashboard/utils"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func (i *InfluxDB) WriteAlertData(data map[string]interface{}, tags map[string]string) error {
	vmPoint := influxdb2.NewPoint(
		consts.ALERT_MEASUREMENT,
		tags,
		data,
		time.Now(), // set timezone
	)

	return i.write.WritePoint(context.Background(), vmPoint)
}

func (i *InfluxDB) ReadAlertData(uuid, period string, pagesize, page int) *model.Page {
	bucket := i.Bucket
	measurement := consts.ALERT_MEASUREMENT

	// query data
	query := buildQueryWithPage(bucket, period, measurement, uuid, pagesize, page)
	result, err := i.query.Query(context.Background(), query)
	if err != nil {
		utils.Log.Error("Can not query data", err)
		return &model.Page{}
	}
	fmt.Println(query)

	// process result to map
	pageItems := make([]*model.PageItem, 0)
	for result.Next() {
		// process data
		pageItem := model.NewPageItem(
			result.Record().Field(),
			result.Record().Value(),
			result.Record().Time().Unix(),
		)
		pageItems = append(pageItems, pageItem)
	}

	// query count
	var total int
	query = buildQueryWithCount(bucket, period, measurement, uuid)
	result, err = i.query.Query(context.Background(), query)
	if err != nil {
		utils.Log.Error("Can not query data", err)
		return &model.Page{}
	}

	for result.Next() {
		// process data
		total = int(result.Record().Value().(int64))
	}

	pageData := model.NewPage(pageItems, pagesize, page, total)

	return pageData
}
