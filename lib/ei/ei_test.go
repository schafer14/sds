package ei_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/matryer/is"
	"github.com/rogpeppe/go-internal/testscript"

	"github.com/schafer14/sds/lib/ei"
)

type Input struct {
	Name  string `validate:"required"`
	Email string
}

func TestMain(m *testing.M) {

	os.Exit(testscript.RunMain(m, map[string]func() int{
		"my-edit": Edit,
	}))
}

func TestFetchingInput(t *testing.T) {

	t.Parallel()
	is := is.New(t)
	ctx := context.Background()
	os.Setenv("EDITOR", "my-edit")

	defaultInput := Input{
		Name: "Banner",
	}
	err := ei.Fetch(ctx, &defaultInput)
	is.NoErr(err)

	is.Equal(defaultInput.Name, "Grace")
	is.Equal(defaultInput.Email, "grace@example.com")

}

// WaitFor waits for a close signal and then closes program with a successful
// error code.
func Edit() int {

	var expected Input
	f, err := os.OpenFile(os.Args[1], os.O_RDWR, 0600)
	if err != nil {
		return 2
	}

	if err := json.NewDecoder(f).Decode(&expected); err != nil {
		return 3
	}

	if expected.Name != "Banner" {
		return 4
	}

	if err := f.Truncate(0); err != nil {
		return 5
	}

	myData := []byte(`{"Name": "Grace", "Email": "grace@example.com" }`)
	if _, err := f.Write(myData); err != nil {
		return 6
	}

	return 0

}
