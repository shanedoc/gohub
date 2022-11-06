package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/shanedoc/gohub/app/http/controllers/api/v1"
	"github.com/shanedoc/gohub/app/models/user"
	"github.com/shanedoc/gohub/app/requests"
	"github.com/shanedoc/gohub/pkg/response"
)

//处理用户身份认证

//SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

//手机号
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	request := requests.SignupPhoneExistRequest{}
	//解析json请求
	if err := requests.Validate(c, &request, requests.ValidateSignupPhoneExist); !err {
		return
	}
	//表单验证
	errs := requests.ValidateSignupPhoneExist(&request, c)
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": errs,
		})
		return
	}
	//返回响应
	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

//邮箱
func (sc *SignupController) IsEmailExist(c *gin.Context) {
	//初始化请求
	request := requests.SignupEmailExistRequest{}

	//解析json请求
	if err := requests.Validate(c, &request, requests.ValidateSignupEmailExist); !err {
		return
	}

	//返回响应
	response.JSON(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})

}
