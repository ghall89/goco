package input

import (
	"errors"
	"log"
	"strings"

	"github.com/charmbracelet/huh"
)

func Commit() (string, string) {
	var (
		message     string
		description string
	)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Message").
				Validate(func(str string) error {
					length := len(strings.Split(str, ""))

					if length < 1 {
						return errors.New("You must enter a commit message")
					}

					if length > 50 {
						return errors.New("Commit message must be less than 50 characters")
					}

					return nil
				}).
				Value(&message),
			huh.NewText().
				Title("Description").
				Value(&description),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	return message, description
}
