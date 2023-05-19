package model

import "time"

type Metric struct {
	Name      string        `json:"name"`
	Unit      string        `json:"unit"`
	Data      []interface{} `json:"data"`
	Timestamp []int64       `json:"timestamp"`
}

// return a new metric with single data
func NewSingleMetric(name, unit string, data interface{}) *Metric {
	return &Metric{
		Name:      name,
		Unit:      unit,
		Data:      []interface{}{data},
		Timestamp: []int64{time.Now().Unix()},
	}
}
