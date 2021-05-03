package code

type Code string

const (
	Disconnected     = Code("DISCONNECTED")
	Interrupt        = Code("INTERRUPT")
	CloseError       = Code("CLOSE_ERROR")
	ReadMessageError = Code("READ_MESSAGE_ERROR")
	DialError        = Code("DIAL_ERROR")
	SocketCloseError = Code("SOCKET_CLOSE_ERROR")
)

func (c Code) IsDisconnected() bool {
	return c == Disconnected
}

func (c Code) IsInterrupt() bool {
	return c == Interrupt
}

func (c Code) IsCloseError() bool {
	return c == CloseError
}

func (c Code) IsReadMessageError() bool {
	return c == ReadMessageError
}

func (c Code) IsDialError() bool {
	return c == DialError
}

func (c Code) IsSocketCloseError() bool {
	return c == SocketCloseError
}

func (c Code) ToString() string {
	return string(c)
}
