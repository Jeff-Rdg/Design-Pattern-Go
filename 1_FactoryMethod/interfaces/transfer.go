package interfaces

type Transfer interface {
	Create()
}

var NewID uint

func Increment() uint {
	NewID++
	return NewID
}
