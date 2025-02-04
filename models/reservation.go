package models

type Reservation struct {
	ID           uint `gorm:"primaryKey"`
	TableID      uint
	CustomerName string
}
