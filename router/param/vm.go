package param

type VmArgument struct {
	UUID string `json:"uuid" binding:"required"`
}

type VmHistoryArgument struct {
	UUID   string   `json:"uuid" binding:"required"`
	Fields []string `json:"fields" binding:"required"`
}
