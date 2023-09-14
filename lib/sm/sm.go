package sm

// The definition of a state machine contains instructions on how to move from
// state to state.
//
// It is parameterirsed on a state, inputs and outputs.
type Definition[S any, I any, O any] struct {
	fn func(S, I) (O, S, error)
}

// SM is an instance of a state machine.
type SM[S any, I any, O any] struct {
	state S
	def   Definition[S, I, O]
}

// Def creates a new template for a state machine. The function provided
// is used to move between states.
func Def[S any, I any, O any](f func(S, I) (O, S, error)) Definition[S, I, O] {
	return Definition[S, I, O]{f}
}

// Empty is a useless type we use as the output of state machines that do
// not care about there outputs.
type Empty interface {
	empty()
}

// StateOnly provides a way to define a state machine if you do not care about
// outputs at individual steps.
func StateOnly[S any, I any](f func(S, I) (S, error)) Definition[S, I, Empty] {
	fn := func(s S, i I) (Empty, S, error) {
		ns, err := f(s, i)
		return nil, ns, err
	}

	return Definition[S, I, Empty]{fn}
}

// Init creates an instance of a state machine from a Definition and an intial
// state.
func (d Definition[State, I, O]) Init(s State) *SM[State, I, O] {
	return &SM[State, I, O]{
		state: s,
		def:   d,
	}
}

// Zeroed initialises a state machine with a zero state.
func (d Definition[State, I, O]) Zeroed() *SM[State, I, O] {
	var s State
	return d.Init(s)
}

// Run processes a slice of events through an instance of a state machine.
func (sm *SM[S, I, O]) Run(inputs ...I) ([]O, error) {

	outputs := []O{}
	errs := ErrorList{}

	for _, input := range inputs {
		out, ns, err := sm.def.fn(sm.state, input)
		if err != nil {
			errs.add(err)
			continue
		}

		sm.state = ns
		outputs = append(outputs, out)
	}

	return outputs, errs.err()
}

// State returns the current state of the machine.
func (sm *SM[State, _, _]) State() State {
	return sm.state
}

// Exec runs a slice of events and returns the new state and ignores the
// output of each iteration.
func (sm *SM[State, I, _]) Exec(inputs ...I) (State, error) {

	_, err := sm.Run(inputs...)
	return sm.state, err
}
