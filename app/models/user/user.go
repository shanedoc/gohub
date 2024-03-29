package user

import (
	"github.com/shanedoc/gohub/app/models"
	"github.com/shanedoc/gohub/pkg/database"
)

//user 用户模型

type User struct {
	models.BaseModel
	Name string `json:"name,omitempty"`

	City         string `json:"city,omitempty"`
	Introduction string `json:"introduction,omitempty"`
	Avator       string `json:"avator,omitempty"`

	Email    string `json:"-"` //json解析器忽略字段
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

//通过手机号、email、用户名获取用户信息
func GetByMulti(loginId string) (userModel User) {
	database.DB.
		Where("phone=?", loginId).
		Or("email=?", loginId).
		Or("name=?", loginId).
		First(&userModel)
	return userModel
}

//通过手机号获取用户
func GetByPhone(phone string) (userModel User) {
	database.DB.Where("phone=?", phone).First(&userModel)
	return
}

//通过id获取用户信息
func Get(id string) (userModel User) {
	database.DB.Where("id=?", id).First(&userModel)
	return
}

//保存
func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}

//通过邮箱查找用户信息
func GetByEmail(email string) (userModel User) {
	database.DB.Where("email=?", email).First(&userModel)
	return
}
