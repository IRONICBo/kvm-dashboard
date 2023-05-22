package param

type VmArgument struct {
	UUID string `json:"uuid" binding:"required"`
}

type VmHistoryArgument struct {
	UUID   string   `json:"uuid" binding:"required"`
	Fields []string `json:"fields" binding:"required"`
}

type VmPageAgrument struct {
	UUID     string `json:"uuid" binding:"required"`
	PageSize int    `json:"pagesize" binding:"required"`
	Page     int    `json:"page" binding:"required"`
}
