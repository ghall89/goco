package gitutils

import "github.com/go-git/go-git/v6"

type Adapter struct{ repo *git.Repository }

func NewAdapter(r *git.Repository) *Adapter { return &Adapter{repo: r} }

func (a *Adapter) Stage() error  { return Stage(a.repo) }
func (a *Adapter) Commit() error { _, err := Commit(a.repo); return err }
func (a *Adapter) Push() error   { return Push(a.repo) }
