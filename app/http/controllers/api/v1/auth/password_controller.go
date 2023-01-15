package auth

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/shanedoc/gohub/app/http/controllers/api/v1"
	"github.com/shanedoc/gohub/app/models/user"
	"github.com/shanedoc/gohub/app/requests"
	"github.com/shanedoc/gohub/pkg/response"
)

//用户控制器
type PasswordController struct {
	v1.BaseAPIController
}

//使用手机验证重置密码
func (pc *PasswordController) ResetByPhone(c *gin.Context) {
	//表单验证
	request := requests.ResetByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByPhone); !ok {
		return
	}
	//更新密码
	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()
		response.Success(c)
	}
}

//使用邮件重置密码
func (pc *PasswordController) ResetByEmail(c *gin.Context) {
	request := requests.ResetByEmailRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByEmail); !ok {
		return
	}
	userModel := user.GetByEmail(request.Email)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()
		response.Success(c)
	}
}
