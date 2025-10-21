package gitutils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v6/plumbing/transport/ssh"
)

func PublicKeys() *ssh.PublicKeys {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	publicKeys, err := ssh.NewPublicKeysFromFile("git", filepath.Join(homedir, ".ssh/id_rsa"), "")

	if err != nil {
		log.Fatal(err)
	}

	return publicKeys
}
