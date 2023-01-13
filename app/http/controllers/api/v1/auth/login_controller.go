package auth

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/shanedoc/gohub/app/http/controllers/api/v1"
	"github.com/shanedoc/gohub/app/requests"
	"github.com/shanedoc/gohub/pkg/auth"
	"github.com/shanedoc/gohub/pkg/jwt"
	"github.com/shanedoc/gohub/pkg/response"
)

type LoginController struct {
	v1.BaseAPIController
}

//手机号登录
func (lc *LoginController) LoginByPhone(c *gin.Context) {
	//验证表单信息
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}

	//尝试登录
	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		//失败
		response.Error(c, err, "账号不存在")
	} else {
		//成功
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		response.JSON(c, gin.H{
			"token": token,
		})
	}

}
