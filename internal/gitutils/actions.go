package gitutils

import (
	"fmt"
	"goco/internal/input"
	"os"

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

func Push() {
	// Check if repo has remote
	// If true, ask user if they'd like to Push
	// Push to remote, create remote branch if necessary
}
