package core

import "errors"

type Arg struct {
	Name  string
	Value interface{}
}

type Wire interface {
	NewWire(...Arg) (interface{}, error)
}

type Initializr interface {
	Initializr() interface{}
}

type WireFunc func(...Arg) (interface{}, error)

func (f WireFunc) NewWire(args ...Arg) (interface{}, error) {
	return f(args...)
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

func Instance(target interface{}, args ...Arg) (interface{}, error) {
	// TODO
	return nil, errors.New("nothing")
}

func buildWith(initializr interface{}) {
	// TODO
}
