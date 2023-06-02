package db

type GenericFileDB interface {
	AppendFile(string, []byte) error
}
