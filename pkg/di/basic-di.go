package di

import (
	"log"
	"reflect"
)

// singleton pattern
var diContainer *basicDI

// ---------------------------------------------------------------------------
// #region definition

type basicDI struct {
	diObjects map[string]*DIObject
}

// #endregion

// ---------------------------------------------------------------------------
// #region constructor

func GetBasicDI() *basicDI {
	if diContainer != nil {
		return diContainer
	}
	diContainer = &basicDI{
		diObjects: make(map[string]*DIObject),
	}
	return diContainer
}

// #endregion

// ---------------------------------------------------------------------------
// #region public

func (x *basicDI) Provide(token string, object interface{}) *DIObject {
	diOject := &DIObject{
		token:        token,
		generator:    reflect.ValueOf(object),
		instanceType: Singleton,
		value:        nil,
	}
	x.diObjects[token] = diOject
	return diOject
}

func (x *basicDI) Resolve(token string) interface{} {
	diObject, found := x.diObjects[token]
	if !found {
		log.Fatalf("No DI object for token %s", token)
		return nil
	}

	if diObject.instanceType == Singleton && diObject.value != nil {
		return diObject.value
	}

	if diObject.generator.Kind() != reflect.Func {
		diObject.value = diObject.generator.Interface()
	} else {
		results := diObject.generator.Call([]reflect.Value{})
		if len(results) == 0 {
			log.Fatalf("DI object for token %s returns nothing", token)
			return nil
		}
		diObject.value = results[0].Interface()
	}

	return diObject.value
}

// #endregion

// ---------------------------------------------------------------------------
// #region private

// #endregion
