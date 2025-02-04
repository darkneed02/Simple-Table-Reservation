package models

type Table struct {
	ID          uint   `gorm:"primaryKey"`
	TableNumber int    `gorm:"unique;not null"`
	Status      string `gorm:"default:available"`
}
