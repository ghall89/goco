package main

import (
	"errors"
	"fmt"
	"goco/internal/fileutils"
	"goco/internal/gitutils"
	"goco/internal/styles"
	"goco/internal/types"
	"os"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(styles.Error.Render(err.Error()))
		os.Exit(1)
	}

}

func run() error {
	path := "./"

	if err := validateDirectory(path); err != nil {
		return err
	}

	repo, err := gitutils.Repository(path)
	if err != nil {
		return errors.New("Not a git repo.")
	}

	svc := gitutils.NewAdapter(repo)
	return commitAndPush(svc)
}

func validateDirectory(p string) error {
	exists, err := fileutils.Exists(p)
	if err != nil {
		return err
	}
	if exists == false {
		return errors.New("Not a valid directory.")
	}

	return nil
}

func commitAndPush(svc types.GitService) error {
	if err := svc.Stage(); err != nil {
		return fmt.Errorf("stage: %w", err)
	}
	if err := svc.Commit(); err != nil {
		return fmt.Errorf("commit: %w", err)
	}
	if err := svc.Push(); err != nil {
		return fmt.Errorf("push: %w", err)
	}
	return nil
}
