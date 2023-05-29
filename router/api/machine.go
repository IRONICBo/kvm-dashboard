package api

import (
	"kvm-dashboard/services"

	"github.com/gin-gonic/gin"
)

func GetMachineList(c *gin.Context) {
	svc := services.NewService()
	svc.GetMachineList()
}
