package requests

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shanedoc/gohub/pkg/response"
	"github.com/thedevsaddam/govalidator"
)

//validator 公用方法封装

// ValidatorFunc 验证函数类型
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

//校验表单参数
func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	//解析request请求,支持json数据、表单、和url request
	if err := c.ShouldBindJSON(obj); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。",
			"error":   err.Error(),
		})
		//打印错误
		//fmt.Println(err.Error())
		return false
	}
	//表单验证
	err := handler(obj, c)

	//验证是否通过
	if len(err) > 0 {
		response.ValidationError(c, err)
		return false
	}
	return true
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}

	//validate
	return govalidator.New(opts).ValidateStruct()
}

func validateFile(c *gin.Context, data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Request:       c.Request,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}
	//验证文件
	return govalidator.New(opts).Validate()
}
