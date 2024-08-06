package transfer

type Transfer interface {
	StartListener() error
	SendFile(filename string, destinationIndex string) error
}
