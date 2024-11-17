package interfaces

type Queuer interface {
	Publish([]byte) error
}
