package controller

import (
	"gin-demo/model/response"

	"github.com/gin-gonic/gin"
)

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [get]
func Login(c *gin.Context) {
	response.OkWithMessage("创建成功", c)
}
