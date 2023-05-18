package model

type Metric struct {
	Name      string        `json:"name"`
	Unit      string        `json:"unit"`
	Data      []interface{} `json:"data"`
	Timestamp []int64       `json:"timestamp"`
}
