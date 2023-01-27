package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/shanedoc/gohub/pkg/auth"
	"github.com/shanedoc/gohub/pkg/response"
)

type UsersController struct {
	BaseAPIController
}

//获取当前登录用户信息
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}
