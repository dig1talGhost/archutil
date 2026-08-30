package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"charm.land/lipgloss/v2"
	"charm.land/lipgloss/v2/tree"
	"github.com/dig1talGhost/archutil/internal/core"
	"github.com/dig1talGhost/archutil/internal/logger"
	"github.com/dig1talGhost/archutil/internal/styles"
	"github.com/dig1talGhost/archutil/internal/terminal"
)

var (
	mainTitle = "Archutil 󰣇"
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

func RenderHeader() {

	lipgloss.Println(styles.HeaderStyle.Render("", mainTitle, ""))
	makeTree()
}

func main() {

	showVersion := flag.Bool("version", false, "print current version")
	flag.Parse()

	if *showVersion {
		fmt.Println(mainTitle, strings.TrimSpace(version))
		os.Exit(0)
	}

	terminal.CheckEnv()
	terminal.ClearScreen()

	logger.TimeStamped.Info("Initializing...")
	terminal.ClearScreen()

	RenderHeader()
	core.Menu()

	terminal.ClearScreen()
	RenderHeader()
}
