package adapters

import (
	"reflect"

	"fyne.io/fyne/v2"
)

type SimpleEventsAdapter struct {
	events map[string][]reflect.Value
}

func NewSimpleEventsAdapter() *SimpleEventsAdapter {
	return &SimpleEventsAdapter{
		events: make(map[string][]reflect.Value),
	}
}

func (x *SimpleEventsAdapter) Subscribe(id string, callbackFunc interface{}) {
	subscribers := x.events[id]
	if x.events[id] == nil {
		subscribers = make([]reflect.Value, 0)
	}
	subscribers = append(subscribers, reflect.ValueOf(callbackFunc))
	x.events[id] = subscribers
}

func (x *SimpleEventsAdapter) Publish(id string, arguments ...interface{}) {
	subscribers := x.events[id]
	if subscribers == nil {
		return
	}

	argumentValues := make([]reflect.Value, len(arguments))
	for i, argument := range arguments {
		if argument == nil {
			argumentValues[i] = reflect.ValueOf(nil)
		} else {
			argumentValues[i] = reflect.ValueOf(argument)
		}
	}

	// calling subscribers in a fyne.Do callback to let the Publish caller finish its task before sending events
	// with fyne.Do, subscribers will be called in the next UI tick
	fyne.Do(func() {
		for _, subscriber := range subscribers {
			subscriber.Call(argumentValues)
		}
	})
}
