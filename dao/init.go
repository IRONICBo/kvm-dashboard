package dao

import (
	"errors"
	"fmt"
	"kvm-dashboard/consts"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// timezone: Asia/Shanghai

type InfluxDB struct {
	// for query
	Org    string
	Bucket string

	write api.WriteAPIBlocking
	query api.QueryAPI
}

var client *InfluxDB

func Init(severURL, authToken, org, bucket string) error {
	influxDBClient := influxdb2.NewClient(severURL, authToken)
	if influxDBClient == nil {
		return errors.New("can not connect to influxdb")
	}

	client = &InfluxDB{
		Org:    org,
		Bucket: bucket,
		write:  influxDBClient.WriteAPIBlocking(org, bucket),
		query:  influxDBClient.QueryAPI(org),
	}

	return nil
}

func GetInfluxDBClient() *InfluxDB {
	return client
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

func buildQueryWithPage(bucket, period, measurement, uuid string,
	pagesize, page int) string {
	// use group to get the all data
	return fmt.Sprintf(`
		from(bucket: "%s")
		|> range(start: -%s)
		|> filter(fn: (r) => r._measurement == "%s")
		|> filter(fn: (r) => r.uuid == "%s")
		|> group(columns: ["*"])
		|> sort(columns: ["_time"], desc: true)
		|> limit(n: %d, offset: %d)
		|> yield(name: "results")`, bucket, period, measurement,
		uuid, pagesize, (page-1)*pagesize)
}

func buildQueryWithCount(bucket, period, measurement, uuid string) string {
	// use group to get the all data
	return fmt.Sprintf(`
		from(bucket: "%s")
		|> range(start: -%s)
		|> filter(fn: (r) => r._measurement == "%s")
		|> filter(fn: (r) => r.uuid == "%s")
		|> group(columns: ["*"])	
		|> count()`, bucket, period, measurement, uuid)
}
