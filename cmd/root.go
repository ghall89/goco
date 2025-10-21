package main

import (
	"fmt"
	"goco/internal/fileutils"
	"goco/internal/gitutils"
	"log"
	"os"

	"github.com/go-git/go-git/v6"
)

func main() {
	path := "./"

	exists, err := fileutils.Exists(path)
	if err != nil {
		log.Fatal(err)
	}
	if exists == false {
		fmt.Println("Not a valid directory.")
		os.Exit(1)
	}

	repo, err := gitutils.Repository(path)
	if err != nil {
		fmt.Println("Not a git repo.")
		os.Exit(1)
	}

	userInput(repo)
}

func userInput(r *git.Repository) {
	stageErr := gitutils.Stage(r)
	if stageErr != nil {
		log.Fatal(stageErr)
	}

	_, commitErr := gitutils.Commit(r)
	if commitErr != nil {
		log.Fatal(commitErr)
	}

	pushErr := gitutils.Push(r)
	if pushErr != nil {
		log.Fatal(commitErr)
	}
}
