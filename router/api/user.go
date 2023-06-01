package api

import (
	"kvm-dashboard/model"
	"kvm-dashboard/router/param"
	"kvm-dashboard/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	param := &param.LoginParam{}
	err := c.ShouldBind(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.FailedWithMsg(err.Error()))
		return
	}

	svc := services.NewService()
	if svc.Login(param.Username, param.Password, c) {
		c.JSON(http.StatusOK, model.Success())
	} else {
		c.JSON(http.StatusOK, model.Failed())
	}
}
