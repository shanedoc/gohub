package models

import (
	"time"

	"github.com/spf13/cast"
)

//basemodel模型基类

type BaseModel struct {
	ID uint64 `gorm:"colum:id;primaryKey:autoIncrement;" json:"id,omitempty"`
}

//CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
}

//获取id的字符串格式
func (a BaseModel) GetStringID() string {
	return cast.ToString(a.ID)
}
