package object

// NewEnvironment initializes an environment.
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// Environment is a scope environment.
type Environment struct {
	store map[string]Object
}

// Get returns identifier object.
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

// Set sets identifier object with associated key.
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
