package param

type HostArgument struct {
	UUID string `json:"uuid" binding:"required"`
}
