package gitutils

import (
	"fmt"
	"goco/internal/input"
	"log"
	"os"
	"os/exec"

	"github.com/charmbracelet/huh/spinner"
	git "github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
)

func Stage(r *git.Repository) error {
	changes, err := Changes(r)
	if err != nil {
		return err
	}

	if changes.IsClean() == true {
		fmt.Println("Clean working tree.")
		os.Exit(0)
	}

	// TODO: Add ability to select changes they'd like to stage

	tree, err := Worktree(r)
	if err != nil {
		return err
	}

	tree.Add(".")

	return nil
}

func Commit(r *git.Repository) (plumbing.Hash, error) {
	msg, dsc := input.Commit()

	tree, err := Worktree(r)
	if err != nil {
		return plumbing.NewHash(""), err
	}

	body := msg + "\n\n" + dsc

	hash, err := tree.Commit(body, &git.CommitOptions{
		AllowEmptyCommits: false,
	})
	if err != nil {
		return plumbing.NewHash(""), err
	}

	return hash, nil
}

func Push(r *git.Repository) error {
	var push bool = false

	remotes, err := r.Remotes()
	if err != nil {
		return err
	}

	input.Push(&push)

	if len(remotes) == 0 {
		return nil
	}

	if push == true {
		loadingMsg := "Pushing to " + remotes[0].String()

		// get users public SSH keys
		// auth := PublicKeys()

		err := spinner.New().
			Title(loadingMsg).
			Action(func() {
				// err := r.Push(&git.PushOptions{
				// 	Auth: auth,
				// })

				// temporary workaround until I can look more into auth issues
				cmd := exec.Command("git", "push")

				err := cmd.Run()
				if err != nil {
					log.Fatal(err)
				}
			}).
			Run()

		if err != nil {
			return err
		}

		fmt.Println("✔ Successfully pushed changes")
	}

	return nil
}
