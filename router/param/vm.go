package param

type VmArgument struct {
	UUID string `json:"uuid" binding:"required"`
}
