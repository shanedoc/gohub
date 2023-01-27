package user

import "github.com/shanedoc/gohub/pkg/database"

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
