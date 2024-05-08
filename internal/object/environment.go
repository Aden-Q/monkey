package object

// interface compliance check
var _ Environment = (*environment)(nil)

type Environment interface {
	Get(name string) (Object, bool)
	Set(name string, val Object)
}

type environment struct {
	store map[string]Object
}

func NewEnvironment() Environment {
	return &environment{
		store: make(map[string]Object),
	}
}

func (e *environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

func (e *environment) Set(name string, val Object) {
	e.store[name] = val
}
