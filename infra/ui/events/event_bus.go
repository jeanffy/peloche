package events

type EventBus interface {
	Subscribe(id string, fn interface{})
	Publish(id string, args ...interface{})
}
