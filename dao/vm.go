package dao

import (
	"context"
	"fmt"
	"kvm-dashboard/consts"
	"kvm-dashboard/model"
	"kvm-dashboard/utils"
	"kvm-dashboard/vm/data"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func (i *InfluxDB) WriteVMData(vmData *data.LibvirtData, tags map[string]string) error {
	vmFields, err := utils.StructToMapWithJSONTag(vmData)
	if err != nil {
		utils.Log.Error("Can not convert vm data to map", err)
		return err
	}
	vmFields_flattened := utils.FlattenMap(vmFields, ".")

	vmPoint := influxdb2.NewPoint(
		"vm",
		tags,
		vmFields_flattened,
		time.Now(), // set timezone
	)

	return i.write.WritePoint(context.Background(), vmPoint)
}

func (i *InfluxDB) ReadVMData(uuid, field, period, agg string) *model.Metric {
	bucket := i.Bucket
	measurement := "vm"
	count := consts.VISION_COUNT
	metric := &model.Metric{}

	// query data
	query := buildQuery(bucket, period, measurement, field, uuid, agg, count)
	result, err := i.query.Query(context.Background(), query)
	if err != nil {
		utils.Log.Error("Can not query data", err)
		return metric
	}

	// process result to map
	fieldList := make([]interface{}, 0)
	timeList := make([]int64, 0)
	for result.Next() {
		// process data
		fieldList = append(fieldList, result.Record().Value())
		timeList = append(timeList, result.Record().Time().Unix())
	}

	metric.Name = field
	metric.Unit = consts.UNIT[field]
	metric.Data = fieldList
	metric.Timestamp = timeList

	return metric
}

func buildQuery(bucket, period, measurement, field, agg, uuid string,
	count int) string {
	switch period {
	case consts.PERIOD_5M:
		return fmt.Sprintf(`
		from(bucket: "%s")
		|> range(start: -%s)
		|> filter(fn: (r) => r._measurement == "%s")
		|> filter(fn: (r) => r._field == "%s")	
		|> filter(fn: (r) => r.uuid == "%s")
		|> aggregateWindow(every: %ds, fn: %s, createEmpty: false)
		|> limit(n: %d)
		|> yield(name: "results")`, bucket, period, measurement,
			field, uuid, consts.VM_DATA_NEARLY_5M_INTERVAL, agg, count)
	case consts.PERIOD_1H:
		return fmt.Sprintf(`
		from(bucket: "%s")
		|> range(start: -%s)
		|> filter(fn: (r) => r._measurement == "%s")
		|> filter(fn: (r) => r._field == "%s")	
		|> filter(fn: (r) => r.uuid == "%s")
		|> aggregateWindow(every: %ds, fn: %s, createEmpty: false)
		|> limit(n: %d)
		|> yield(name: "results")`, bucket, period, measurement,
			field, uuid, consts.VM_DATA_NEARLY_1H_INTERVAL, agg, count)
	default: // 1m
		return fmt.Sprintf(`
			from(bucket: "%s")
			|> range(start: -%s)
			|> filter(fn: (r) => r._measurement == "%s")
			|> filter(fn: (r) => r._field == "%s")	
			|> filter(fn: (r) => r.uuid == "%s")	
			|> limit(n: %d)
			|> yield(name: "results")`, bucket, period, measurement,
			field, uuid, count)
	}
}

// deprecated
func (i *InfluxDB) ReadVMDataRealtime(uuid, field string) *model.Metric {
	bucket := i.Bucket
	period := consts.PERIOD_1M
	_measurement := "vm"
	count := consts.VISION_COUNT
	metric := &model.Metric{}

	query := fmt.Sprintf(`
	from(bucket: "%s")
	|> range(start: -%s)
	|> filter(fn: (r) => r._measurement == "%s")
	|> filter(fn: (r) => r._field == "%s")
	|> filter(fn: (r) => r.uuid == "%s")	
	|> limit(n: %d)
	|> yield(name: "results")`, bucket, period, _measurement, field, uuid, count)

	result, err := i.query.Query(context.Background(), query)
	if err != nil {
		utils.Log.Error("Can not query data", err)
		return metric
	}

	// process result to map
	fieldList := make([]interface{}, 0)
	timeList := make([]int64, 0)
	for result.Next() {
		// process data
		fieldList = append(fieldList, result.Record().Value())
		timeList = append(timeList, result.Record().Time().Unix())
	}

	metric.Name = field
	metric.Unit = consts.UNIT[field]
	metric.Data = fieldList
	metric.Timestamp = timeList

	return metric
}
