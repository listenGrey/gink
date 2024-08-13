package transfer

type Transfer interface {
	Send(filepath string, destinationIndex string) error
	Receive() error
	Stop() error
}
