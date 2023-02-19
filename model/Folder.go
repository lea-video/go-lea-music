package model

type Folder struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"not null"`
}
