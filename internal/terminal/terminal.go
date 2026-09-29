package terminal

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearScreen() {

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("\033[2J\033[H")
		return
	}
}
