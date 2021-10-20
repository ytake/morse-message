package stream

type Reader interface {
	Proceed(message []byte) error
}

