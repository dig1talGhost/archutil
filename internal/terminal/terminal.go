package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/dig1talGhost/archutil/internal/logger"
)

func CheckEnv() {

	if runtime.GOOS != "linux" {
		logger.Standard.Fatal("Error:", "unsupported operating system", runtime.GOOS)
	}
}

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
