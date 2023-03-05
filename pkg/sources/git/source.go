package git

import (
	"log"
	"os"

	"github.com/go-git/go-git/v5"
)

type GitSource struct {
	Url    string
	Branch string
}

func (g GitSource) Clean(path string) {

}

func (g GitSource) Download(path string) (bool, error) {
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      g.Url,
		Progress: os.Stdout,
		// SingleBranch:  true,
		// ReferenceName: plumbing.ReferenceName(g.Branch),
	})

	if err != nil {
		log.Printf(err.Error())
		if err == git.ErrRepositoryAlreadyExists {
			return true, nil
		}
		return false, err
	}

	return true, err
}
