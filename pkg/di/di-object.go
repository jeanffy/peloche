package di

import (
	"reflect"
)

type DIInstanceType int8

const (
	Singleton = iota
)

type DIObject struct {
	token        string
	generator    reflect.Value
	instanceType DIInstanceType
	value        interface{}
}

func (x *DIObject) AsSingleton() {
	x.instanceType = Singleton
}
