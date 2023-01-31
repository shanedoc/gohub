package link

import (
	"github.com/shanedoc/gohub/app/models"
	"github.com/shanedoc/gohub/pkg/database"
)

type Link struct {
	models.BaseModel

	// Put fields in here
	Name string `json:"name,optiempty"`
	URL  string `json:"url,optiempty"`
	models.CommonTimestampsField
}

func (link *Link) Create() {
	database.DB.Create(&link)
}

func (link *Link) Save() (rowsAffected int64) {
	result := database.DB.Save(&link)
	return result.RowsAffected
}

func (link *Link) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&link)
	return result.RowsAffected
}
