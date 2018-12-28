package core

type Wire interface {
	NewWire() interface{}
}

type Initializr interface {
	Initializr() interface{}
}

type WireFunc func() interface{}

func (f WireFunc) NewWire() interface{} {
	return f()
}

func Build(initializrs ...Initializr) {
	for _, initializr := range initializrs {
		i := initializr.Initializr()
		if i != nil {
			buildWith(i)
		}
	}
	// TODO
}

func buildWith(initializr interface{}) {
	// TODO
}
