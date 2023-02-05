package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/shanedoc/gohub/app/models/user"
	"github.com/shanedoc/gohub/app/requests"

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
	request := requests.PaginationRequest{}
	//fmt.Println(request)
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}
	data, pager := user.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *UsersController) UserUpdateProfile(c *gin.Context) {
	request := requests.UserUpdateProfileReuqest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateProfile); !ok {
		return
	}
	currentUser := auth.CurrentUser(c)
	currentUser.Name = request.Name
	currentUser.Introduction = request.Introduction
	currentUser.City = request.City
	rowsAffected := currentUser.Save()
	if rowsAffected > 0 {
		response.Data(c, currentUser)
	} else {
		response.Abort500(c, "更新失败!")
	}

}

func (ctrl *UsersController) UpdateEmail(c *gin.Context) {
	request := requests.UserUpdateEmailRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateEmailProfile); !ok {
		return
	}
	currentUser := auth.CurrentUser(c)
	currentUser.Email = request.Email
	rowsAffected := currentUser.Save()
	if rowsAffected > 0 {
		response.Data(c, currentUser)
	} else {
		response.Abort500(c, "更新失败!")
	}
}

func (ctrl *UsersController) UpdatePhone(c *gin.Context) {
	request := requests.UserUpdatePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdatePhoneProfile); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Phone = request.Phone
	rowsAffected := currentUser.Save()

	if rowsAffected > 0 {
		response.Success(c)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *UsersController) UpdatePassword(c *gin.Context) {
	request := requests.UserUpdatePasswordRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdatePasswordProfile); !ok {
		return
	}
	currentUser := auth.CurrentUser(c)
	// 验证原始密码是否正确
	_, err := auth.Attempt(currentUser.Name, request.Password)
	if err != nil {
		// 失败，显示错误提示
		response.Unauthorized(c, "原密码不正确")
	} else {
		// 更新密码为新密码
		currentUser.Password = request.NewPassword
		currentUser.Save()

		response.Success(c)
	}
}
