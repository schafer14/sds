package examples

import (
	"fmt"
	"log"

	"github.com/schafer14/sds/lib/sm"
)

// Setup some local types.
// The event interface limits the types that our state machine will accpet
// as inputs.
// We could do this with state as well to switch between different types of
// state...
// Notice each event type is public but the event interface is private, allowing
// us to expose our event types without sacraficing type safty of our machine.
type event interface {
	private()
}

type SimpleEventAdd int

func (SimpleEventAdd) private() {}

type SimpleEventSubtract int

func (SimpleEventSubtract) private() {}

type state = int

// These helper functions are just ortimental. They make running the machine
// a little clearer.
func add(i int) event {
	return SimpleEventAdd(i)
}
func sub(i int) event {
	return SimpleEventSubtract(i)
}

// Setup and run a simple state machine.
func Simple() int {

	def := sm.StateOnly(func(s state, e event) (state, error) {
		switch event := e.(type) {
		case SimpleEventAdd:
			return s + int(event), nil
		case SimpleEventSubtract:
			return s - int(event), nil
		default:
			return s, fmt.Errorf("unknown event")
		}
	})

	state, err := def.Zeroed().Exec(add(4), sub(1), sub(2), add(5))
	if err != nil {
		log.Println(err)
		return 1
	}

	log.Printf("state : %d\n", state)
	return 0
}
