//+build wireinject

package wire

import (
	"errors"
	. "github.com/alimy/wire/core"
	"github.com/alimy/wire/module/example/room"
	"time"
)

type RoomWire struct {
	Message Wire `wire:"Args(phrase string)"`
	Greeter Wire `wire:"Args(msg github.com/alimy/wire/module/example/room.Message)"`
	Event   Wire `wire:"Args(greeter .../room.Greeter)"`
}

// message creates a default Message.
func (w *RoomWire) message(args ...*Arg) (interface{}, error) {
	phrase, _ := args[0].Value.(string)
	return room.Message(phrase), nil
}

// greeter initializes a Greeter. If the current epoch time is an even
// number, greeter will create a grumpy Greeter.
func (w *RoomWire) greeter(args ...*Arg) (interface{}, error) {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	msg, _ := args[0].Value.(room.Message)
	return &room.Greeter{Message: msg, Grumpy: grumpy}, nil
}

// NewEvent creates an event with the specified greeter.
func (w *RoomWire) event(args ...*Arg) (interface{}, error) {
	g, _ := args[0].Value.(room.Greeter)
	if g.Grumpy {
		return &room.Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return &room.Event{Greeter: g}, nil
}

func (w *RoomWire) Initializr() interface{} {
	w.Message = WireFunc(w.message)
	w.Greeter = WireFunc(w.greeter)
	w.Event = WireFunc(w.event)

	return w
}
