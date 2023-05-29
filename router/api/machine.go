package api

import (
	"kvm-dashboard/model"
	"kvm-dashboard/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMachineList(c *gin.Context) {
	svc := services.NewService()
	machineInfo, err := svc.GetMachineList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FailedWithMsg(err.Error()))
	}

	c.JSON(http.StatusOK, model.SuccessWithData(machineInfo))
}
