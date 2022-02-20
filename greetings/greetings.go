package greetings

import "fmt"

type Message string

type Greeter struct {
	Message Message
}

type Event struct {
	Greeter Greeter
}

// Returns type message.

func NewMessage() Message {
	return Message("Hello, welcome!")
}

// Instantiates Greeter.

func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

// Returns the message inside the greeter object.

func (g Greeter) Greet() Message {
	return g.Message
}

// Instantiates Event.

func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

// Tells greeter to issue a greeting and then prints the message out.

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
