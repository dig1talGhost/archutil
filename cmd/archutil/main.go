package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"

	"charm.land/lipgloss/v2/tree"
	"github.com/dig1talGhost/archutil/internal/styles"
	"github.com/dig1talGhost/archutil/internal/terminal"
	"github.com/dig1talGhost/archutil/internal/tui"
)

var (
	mainTitle = "󰣇 archutil"
	version   = "dev"
)

func makeTree() {

	t := tree.Root(styles.TreeRootStyle.Render("○ Version")).
		Child(
			tree.New().
				Root(styles.TreeChildStyle.Render(version)),
		)

	fmt.Println(t)
}

func renderHeader() {

	fmt.Println(styles.HeaderStyle.Render("", mainTitle, ""))
	makeTree()
}

func main() {

	showVersion := flag.Bool("version", false, "print current version")

	if runtime.GOOS != "linux" {
		fmt.Println(styles.ErrorStyle.Render("Error:"), "Unsupported operating system detected")
		os.Exit(1)
	}

	flag.Parse()
	if *showVersion {
		fmt.Println(styles.CommonStyle.Render(mainTitle), "-", strings.TrimSpace(version))
		os.Exit(0)
	}
	terminal.ClearScreen()

	renderHeader()
	tui.Menu()
}
