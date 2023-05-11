package request_param

type VmArgument struct {
	Uid string `uri:"uid" binding:"required"`
}
