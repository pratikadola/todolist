package models

import "gorm.io/gorm"

type List struct {
	gorm.Model
	ListID uint   `gorm:"primaryKey" json:"id"`
	Name   string `json:"name"`
	Tasks  []Task `gorm:"foreignKey:ListRefer"`
}
