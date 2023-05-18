package dao

import (
	"errors"

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
