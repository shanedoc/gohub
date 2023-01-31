package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type TopicRequest struct {
	// Name        string `valid:"name" json:"name"`
	// Description string `valid:"description" json:"description,omitempty"`
	Title      string `json:"title,omitempty" valid:"title"`
	Body       string `json:"body,omitempty" valid:"body"`
	CategoryID string `json:"category_id,omitempty" valid:"category_id"`
}

func TopicSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"title":       []string{"required", "min_cn:3", "max_cn:40"},
		"body":        []string{"required", "min_cn:10", "max_cn:50000"},
		"category_id": []string{"required", "exists:categories,id"},
	}
	messages := govalidator.MapData{
		"title": []string{
			"required:帖子标题为必填项",
			"min_cn:帖子标题长度需至少 2 个字",
			"max_cn:帖子标题长度不能超过 8 个字",
		},
		"body": []string{
			"required:帖子内容为必填项",
			"min_cn:长度需大于 10",
		},
		"category_id": []string{
			"required:帖子分类为必填项",
			"exists:帖子分类未找到",
		},
	}
	return validate(data, rules, messages)
}
