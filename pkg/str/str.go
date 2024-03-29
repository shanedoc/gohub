package str

import (
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

//str字符串辅助方法

//字符串转为复数
func Plural(word string) string {
	return pluralize.NewClient().Plural(word)
}

//转成单数
func Singular(word string) string {
	return pluralize.NewClient().Singular(word)
}

//Snake 转为 snake_case，如 TopicComment -> topic_comment
func Snake(s string) string {
	return strcase.ToSnake(s)
}

//Camel 转为 CamelCase，如 topic_comment -> TopicComment
func Camel(s string) string {
	return strcase.ToCamel(s)
}

//LowerCamel 转为 lowerCamelCase，如 TopicComment -> topicComment
func LowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}
