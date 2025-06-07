package ui

var EVENTS_PORT_TOKEN = "EventsPort"

type EventsPort interface {
	Subscribe(id string, fn interface{})
	Publish(id string, args ...interface{})
}
