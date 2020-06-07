//go:generate wire
package example

import (
	"errors"
	"fmt"
	"os"
	"time"
)

//message
type Message string

func NewMessage() Message {
	return Message("Hi there!")
}

//greeter
type Greeter struct {
	Grumpy  bool
	Message Message // <- adding a Message field
}

func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

//event
type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main() {
	//old main
	// message := NewMessage()
	// greeter := NewGreeter(message)
	// event := NewEvent(greeter)
	// event.Start()

	//new main
	e, err := InitializeEvent()
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
}
