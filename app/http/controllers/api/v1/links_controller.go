package v1

import (
	"github.com/shanedoc/gohub/app/models/link"
	"github.com/shanedoc/gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type LinksController struct {
	BaseAPIController
}

func (ctrl *LinksController) Index(c *gin.Context) {
	//links := link.All()
	response.Data(c, link.AllCached())
}

func (ctrl *LinksController) Show(c *gin.Context) {
	linkModel := link.Get(c.Param("id"))
	if linkModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, linkModel)
}
