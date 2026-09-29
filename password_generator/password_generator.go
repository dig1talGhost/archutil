package password_generator

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"charm.land/lipgloss/v2"

	"github.com/dig1talGhost/archutil/internal/styles"
	"github.com/dig1talGhost/archutil/internal/terminal"
)

const defaultChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"

func generatePassword(length int, chars string) ([]byte, error) {

	password := make([]byte, length)

	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return nil, err
		}
		password[i] = chars[num.Int64()]
	}
	return password, nil
}

func Gen() {

	passwordLength := 16
	password, err := generatePassword(passwordLength, defaultChars)

	if err != nil {
		fmt.Println()
		lipgloss.Println(styles.ErrorStyle.Render("Error:"), "Oops, something went wrong...")
		return
	}

	defer func() {

		for i := range password {
			password[i] = 0
		}
	}()

	terminal.ClearScreen()

	fmt.Println()
	lipgloss.Println(styles.LegendStyleBold.Render("Password:"), styles.CommonStyle.Render(string(password)))
}
