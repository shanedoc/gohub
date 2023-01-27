package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/shanedoc/gohub/app/models/user"
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

//index所有用户
func (ctrl *UsersController) Index(c *gin.Context) {
	data, pager := user.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}
