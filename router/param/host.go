package param

type HostArgument struct {
	UUID string `json:"uuid" binding:"required"`
}

type HostHistoryArgument struct {
	UUID   string   `json:"uuid" binding:"required"`
	Fields []string `json:"fields" binding:"required"`
}

type HostPageAgrument struct {
	UUID     string `json:"uuid" binding:"required"`
	PageSize int    `json:"pagesize" binding:"required"`
	Page     int    `json:"page" binding:"required"`
}
