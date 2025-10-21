package gitutils

import (
	git "github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/object"
)

func Repository(p string) (*git.Repository, error) {
	repo, err := git.PlainOpen(p)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func Changes(r *git.Repository) (git.Status, error) {
	worktree, err := Worktree(r)
	if err != nil {
		return nil, err
	}

	changes, err := worktree.Status()
	if err != nil {
		return nil, err
	}

	return changes, nil
}

func Worktree(r *git.Repository) (*git.Worktree, error) {
	worktree, err := r.Worktree()
	if err != nil {
		return nil, err
	}

	return worktree, nil
}

func History(r *git.Repository) (object.CommitIter, error) {
	var options = &git.LogOptions{
		All: true,
	}

	history, err := r.Log(options)
	if err != nil {
		return nil, err
	}

	return history, nil
}
