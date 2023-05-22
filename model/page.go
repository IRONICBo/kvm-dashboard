package model

type Page struct {
	Data     []*PageItem `json:"data"`
	PageSize int         `json:"pagesize"`
	Page     int         `json:"page"`
	Total    int         `json:"total"`
}

func NewPage(data []*PageItem, pagesize, page, total int) *Page {
	return &Page{
		Data:     data,
		Page:     page,
		PageSize: pagesize,
		Total:    total,
	}
}

type PageItem struct {
	Metric    string      `json:"metric"`
	Value     interface{} `json:"value"`
	Timestamp int64       `json:"timestamp"`
}

func NewPageItem(metric string, value interface{}, timestamp int64) *PageItem {
	return &PageItem{
		Metric:    metric,
		Value:     value,
		Timestamp: timestamp,
	}
}
