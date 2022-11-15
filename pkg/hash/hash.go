//hash操作类
package hash

import (
	"github.com/shanedoc/gohub/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

//使用hash加密
func BcryptHash(password string) string {
	//第二个参数 cost,建议大于12,数值越大耗费时间越长
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogIf(err)
	return string(bytes)
}

//对比明文密码和数据库密码
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

//判断字符串是否hash过的数据
func BcryptIsHashed(str string) bool {
	//bcrypt加密后字符串长度为60
	return len(str) == 60
}
