package examples_test

import (
	"testing"

	"github.com/schafer14/sds/lib/sm/examples"
)

func TestSimpl(t *testing.T) {

	returnCode := examples.Simple()
	if returnCode != 0 {
		t.Error("non zero return code")
	}
}
