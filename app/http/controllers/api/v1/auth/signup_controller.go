package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/shanedoc/gohub/app/http/controllers/api/v1"
	"github.com/shanedoc/gohub/app/models/user"
)

//处理用户身份认证

//SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

//手机号
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	//请求对象
	type PhoneExistequest struct {
		Phone string `json:"phone"`
	}
	request := PhoneExistequest{}

	//解析json请求
	if err := c.ShouldBindJSON(&request); err != nil {
		//失败 返回422
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		//打印错误
		fmt.Println(err.Error())
		return
	}
	//返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
