package model

type File struct {
	ID       int
	File     string
	FolderID int
	Folder   *Folder
}
