package model

type File struct {
	ID       int     `gorm:"primaryKey"`
	FolderID int     `gorm:"not null"`
	Folder   *Folder `gorm:"not null;references:FolderID;foreignKey:ID"`
	File     string  `gorm:"not null"`
}
