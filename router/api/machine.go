package api

import (
	"kvm-dashboard/model"
	"kvm-dashboard/router/param"
	"kvm-dashboard/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMachineList(c *gin.Context) {
	svc := services.NewService()
	machineInfo, err := svc.GetMachineList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}

	c.JSON(http.StatusOK, model.SuccessWithData(machineInfo))
}

func GetMachineThreshold(c *gin.Context) {
	// get host or vm threshold, and init threshold in memory
	param := &param.MachineArgument{}
	err := c.ShouldBindQuery(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
		return
	}

	svc := services.NewService()
	thresholdInfo := svc.GetMachineThreshold(param.UUID) // get or return a default one
	c.JSON(http.StatusOK, model.SuccessWithData(thresholdInfo))
}
