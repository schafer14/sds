package sm_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/schafer14/sds/lib/sm"
)

type event = int
type state = int

func TestRunningAStateMachine(t *testing.T) {

	t.Parallel()
	is := is.New(t)

	def := sm.Def(func(s state, i event) (int, state, error) {
		ns := s + i
		return ns, ns, nil
	})

	smInst := def.Init(1)
	out, err := smInst.Run(2, 3, 4, 5)
	is.NoErr(err)
	is.Equal(out, []int{3, 6, 10, 15})
	is.Equal(smInst.State(), 15)

	ns, err := smInst.Exec(6, 7, 8, 9)
	is.NoErr(err)
	is.Equal(ns, 45)

	zeroedSM := def.Zeroed()
	is.Equal(zeroedSM.State(), 0)
}
