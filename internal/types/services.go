package types

type GitService interface {
	Stage() error
	Commit() error
	Push() error
}
