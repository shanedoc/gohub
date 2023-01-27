package user

import (
	"github.com/gin-gonic/gin"
	"github.com/shanedoc/gohub/pkg/app"
	"github.com/shanedoc/gohub/pkg/database"
	"github.com/shanedoc/gohub/pkg/paginator"
)

//校验:邮箱是否存在
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(&User{}).Where("email=?", email).Count(&count)
	return count > 0
}

//校验:手机号是否存在
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(&User{}).Where("phone=?", phone).Count(&count)
	return count > 0
}

func All() (users []User) {
	database.DB.Find(&users)
	return
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (users []User, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(User{}),
		&users,
		app.V1URL(database.TableName(&User{})),
		perPage,
	)
	return
}
