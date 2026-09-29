package main

import (
	"flag"
	"os"
	"runtime"
	"strings"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/tree"
	"github.com/dig1talGhost/archutil/styles"
	"github.com/dig1talGhost/archutil/terminal"
	"github.com/dig1talGhost/archutil/tui"
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
	lipgloss.Println(t)
}

func renderHeader() {

	lipgloss.Println(styles.HeaderStyle.Render("", mainTitle, ""))
	makeTree()
}

func main() {

	showVersion := flag.Bool("version", false, "print current version")

	if runtime.GOOS != "linux" {
		lipgloss.Println(styles.ErrorStyle.Render("Error:"), "Unsupported operating system detected")
		os.Exit(1)
	}

	flag.Parse()
	if *showVersion {
		lipgloss.Println(styles.CommonStyle.Render(mainTitle), "-", strings.TrimSpace(version))
		os.Exit(0)
	}
	terminal.ClearScreen()

	renderHeader()
	tui.Menu()
}
