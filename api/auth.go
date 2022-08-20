package api

import (
	"github.com/gin-gonic/gin"
	"inwind-blog-server-v3/common"
	"inwind-blog-server-v3/service"
	"net/http"
)

//服务入口，负责处理路由，参数校验，请求转发

func Login(c *gin.Context) {
	var loginParams common.LoginParams
	//校验参数
	if err := c.ShouldBind(&loginParams); err != nil {
		common.ResponseError(c, http.StatusUnprocessableEntity, 422, err.Error())
		return
	}

	service.Login(c)
}