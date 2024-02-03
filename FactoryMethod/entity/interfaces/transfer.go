package interfaces

type Transfer interface {
	Send() string
	Processing() string
	Flow()
}

var NewID uint

func Increment() uint {
	NewID++
	return NewID
}
