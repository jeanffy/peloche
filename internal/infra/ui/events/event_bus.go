package events

var EVENT_BUS_TOKEN = "EventBus"

type EventBus interface {
	Subscribe(id string, fn interface{})
	Publish(id string, args ...interface{})
}
