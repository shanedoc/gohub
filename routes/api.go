//注册路由
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	//v1路由组,把路由全部放到该路由组下
	v1 := r.Group("v1")
	{
		//注册一个路由
		v1.GET("/", func(c *gin.Context) {
			//json格式响应
			c.JSON(http.StatusOK, gin.H{
				"hello": "world",
			})
		})
	}

}
