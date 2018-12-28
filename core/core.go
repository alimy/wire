package core

type Args map[string]interface{}

type Wire interface {
	NewWire(Args) (interface{}, error)
}

type Initializr interface {
	Initializr() interface{}
}

type WireFunc func(Args) (interface{}, error)

func (f WireFunc) NewWire(args Args) (interface{},error) {
	return f(args)
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
