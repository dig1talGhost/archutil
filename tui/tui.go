package tui

import (
	"fmt"
	"os"

	"charm.land/huh/v2"
	"charm.land/lipgloss/v2"
	"github.com/dig1talGhost/archutil/getarch"
	ghosttp "github.com/dig1talGhost/archutil/ghosttp/cmd"
	"github.com/dig1talGhost/archutil/glyphs"
	"github.com/dig1talGhost/archutil/logger"
	"github.com/dig1talGhost/archutil/password_generator"
	"github.com/dig1talGhost/archutil/styles"
	"github.com/dig1talGhost/archutil/terminal"
)

func Menu() {

	fmt.Println()
	var choice string

	for {
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Main Menu").
					Description("Choose an option to execute, or exit the program.").
					Options(
						huh.NewOption("Initiate local HTTP server", "opt1"),
						huh.NewOption("Get latest archiso and sig", "opt2"),
						huh.NewOption("Glyphs menu", "opt3"),
						huh.NewOption("Password generator", "opt4"),
						huh.NewOption("Exit", "exit"),
					).
					Value(&choice),
			),
		)

		err := form.Run()

		if err != nil {
			logger.Standard.Fatalf("Error running form: %v", err)
		}

		switch choice {
		case "opt1":
			ghosttp.Serve()
		case "opt2":
			getarch.Latest()
		case "opt3":
			glyphs.Pager()
		case "opt4":
			password_generator.Gen()
		case "exit":
			lipgloss.Println(styles.CommonStyle.Render("\nExiting..."))
			os.Exit(0)
		}

		fmt.Print("\nPress Enter to return to the menu...")
		fmt.Scanln()

		terminal.ClearScreen()
	}
}
