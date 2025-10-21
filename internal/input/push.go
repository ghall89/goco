package input

import (
	"log"

	"github.com/charmbracelet/huh"
)

func Push(push *bool) {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Push to remote?").
				Value(push),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
}
