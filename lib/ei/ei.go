package ei

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func Fetch(ctx context.Context, input interface{}) error {

	text, err := json.MarshalIndent(input, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling default : %v", err)
	}

	f, err := os.CreateTemp("", "*.json")
	if err != nil {
		return fmt.Errorf("creating temp director : %v", err)
	}
	defer os.Remove(f.Name()) // clean up

	if _, err := f.Write(text); err != nil {
		return err
	}

	if err := os.Chmod(f.Name(), 0600); err != nil {
		return err
	}

	editor := os.Getenv("EDITOR")
	path, err := exec.LookPath(editor)
	if err != nil {
		return fmt.Errorf("editor not found")
	}

	cmd := exec.CommandContext(ctx, path, f.Name())
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()

	nf, err := os.Open(f.Name())
	if err != nil {
		return err
	}
	defer nf.Close()

	if err := json.NewDecoder(nf).Decode(input); err != nil {
		return fmt.Errorf("decoding : %v", err)
	}

	return nil
}
