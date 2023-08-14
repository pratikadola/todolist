package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	TaskID         uint   `gorm:"primaryKey"`
	Name           string `json:"name"`
	ScheduleInDays int
	LastOccurred   time.Time
	Total          int
	Completed      int
	ListRefer      uint
}
