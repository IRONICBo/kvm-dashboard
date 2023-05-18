package api

import (
	"kvm-dashboard/model"
	"kvm-dashboard/router/param"
	"kvm-dashboard/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHostBasicInfo(c *gin.Context) {
	svc := services.NewService()
	hostInfo := svc.GetHostInfo()

	// todo: need to reconstruct
	c.JSON(http.StatusOK, model.SuccessWithData(hostInfo))
}

func GetVmBasicInfo(c *gin.Context) {
	param := &param.VmArgument{}
	err := c.ShouldBindQuery(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}

	svc := services.NewService()
	vmInfo := svc.GetVMInfo(param.UUID)

	// todo: need to reconstruct
	c.JSON(http.StatusOK, model.SuccessWithData(vmInfo))
}
