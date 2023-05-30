package param

type MachineArgument struct {
	UUID string `json:"uuid" binding:"required"`
}
