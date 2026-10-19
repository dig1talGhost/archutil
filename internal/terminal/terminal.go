package terminal

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearScreen() {
	_, err := fmt.Print("\033[2J\033[H")
	if err == nil {
		return
	}

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}
