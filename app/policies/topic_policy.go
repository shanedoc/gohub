package policies

import (
	"github.com/gin-gonic/gin"
	"github.com/shanedoc/gohub/app/models/topic"
	"github.com/shanedoc/gohub/pkg/auth"
)

//授权策略:数据权限限制

func CanModifyTopic(c *gin.Context, _topic topic.Topic) bool {
	return auth.CurrentUID(c) == _topic.UserID
}
