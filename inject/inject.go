package inject

import "errors"

var (
	registerInjects = make([]Inject, 0)
)

// Register append injects
func Register(injects ...Inject) {
	registerInjects = append(registerInjects, injects...)
}

// Build start build inject component
func Build(injects ...Inject) {
	registerInjects = append(registerInjects, injects...)
	for _, inject := range registerInjects {
		i := inject.InjectOf()
		if i != nil {
			buildWith(i)
		}
	}
	// TODO
}

func Instance(target interface{}, args ...*Arg) (interface{}, error) {
	// TODO
	return nil, errors.New("nothing")
}

func buildWith(initializr interface{}) {
	// TODO
}
