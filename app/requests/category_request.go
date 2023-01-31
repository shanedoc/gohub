package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type CategoryRequest struct {
	Name        string `valid:"name" json:"name"`
	Description string `valid:"description" json:"description,omitempty"`
}

func CategorySave(data interface{}, c *gin.Context) map[string][]string {

	//todo::校验分类名是否重名
	rules := govalidator.MapData{
		"name":        []string{"required", "min_cn:2", "max_cn:8"},
		"description": []string{"min_cn:3", "max_cn:255"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:分类名称为必填项",
			"min_cn:分类名称长度需至少 2 个字",
			"max_cn:分类名称长度不能超过 8 个字",
		},
		"description": []string{
			"min_cn:分类描述长度需至少 3 个字",
			"max_cn:分类描述长度不能超过 255 个字",
		},
	}
	return validate(data, rules, messages)
}