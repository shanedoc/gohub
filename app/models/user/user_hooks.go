package user

import (
	"github.com/shanedoc/gohub/pkg/hash"

	"gorm.io/gorm"
)

//BeforeSave GORM模型钩子,创建和更新前调用
func (userModel *User) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(userModel.Password) {
		userModel.Password = hash.BcryptHash(userModel.Password)
	}
	return
}

//校验密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}
