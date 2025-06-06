package utils

type naiveDI struct {
	objects map[string]interface{}
}

var diContainer *naiveDI

func GetNaiveDI() *naiveDI {
	if diContainer != nil {
		return diContainer
	}
	diContainer = &naiveDI{
		objects: make(map[string]interface{}),
	}
	return diContainer
}

func (x *naiveDI) Provide(token string, obj interface{}) {
	x.objects[token] = obj
}

func (x *naiveDI) Resolve(token string) interface{} {
	return x.objects[token]
}
