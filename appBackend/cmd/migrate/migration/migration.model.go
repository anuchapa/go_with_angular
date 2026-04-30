package migration

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ProductCode string         `gorm:"column:product_code;size:50;index:uix_code,unique,where:deleted_at IS NULL" json:"product_code"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at"`
}