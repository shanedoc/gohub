package topic

import (
	"github.com/shanedoc/gohub/app/models"
	"github.com/shanedoc/gohub/app/models/category"
	"github.com/shanedoc/gohub/app/models/user"
	"github.com/shanedoc/gohub/pkg/database"
)

type Topic struct {
	models.BaseModel

	Title      string `json:"title,omitempty" `
	Body       string `json:"body,omitempty" `
	UserID     string `json:"user_id,omitempty"`
	CategoryID string `json:"category_id,omitempty"`

	//通过user_id关联user表
	User user.User `json:"user"`
	//通过category_id关联category表
	Category category.Category `json:"category"`

	models.CommonTimestampsField
}

func (topic *Topic) Create() {
	database.DB.Create(&topic)
}

func (topic *Topic) Save() (rowsAffected int64) {
	result := database.DB.Save(&topic)
	return result.RowsAffected
}

func (topic *Topic) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&topic)
	return result.RowsAffected
}
