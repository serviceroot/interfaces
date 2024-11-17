package interfaces

type EventManager interface {
	LineUp() error
	Publish(Queuer) error
}
