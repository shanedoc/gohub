package user

import "github.com/shanedoc/gohub/app/models"

//user 用户模型

type User struct {
	models.BaseModel
	Name     string `json:"name,omitempty"`
	Email    string `json:"-"` //json解析器忽略字段
	Phone    string `json:"-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}
