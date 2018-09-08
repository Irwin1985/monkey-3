package object

// NewEnclosedEnvironment initializes an environment with outer environment.
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// NewEnvironment initializes an environment.
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// Environment is a scope environment.
type Environment struct {
	store map[string]Object
	outer *Environment
}

// Get returns identifier object.
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set sets identifier object with associated key.
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
