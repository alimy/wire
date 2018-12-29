package inject

type Arg struct {
	Name  string
	Value interface{}
}

type Wire interface {
	NewWire(...*Arg) (interface{}, error)
}

type Inject interface {
	InjectOf() interface{}
}

type WireFunc func(...*Arg) (interface{}, error)

func (f WireFunc) NewWire(args ...*Arg) (interface{}, error) {
	return f(args...)
}

