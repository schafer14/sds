package sm

import "fmt"

type ErrorList struct {
	list []error
}

func (e *ErrorList) Error() string {
	return fmt.Sprintf("%d errors processing state machine : %v", len(e.list), e.list)
}

func (e *ErrorList) add(err error) {
	if e.list == nil {
		e.list = []error{}
	}

	e.list = append(e.list, err)
}

func (e *ErrorList) err() error {
	if e.list == nil {
		return nil
	}

	return e
}
