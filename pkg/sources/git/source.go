package git

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

type GitSource struct {
	Url        string
	Branch     string
	Repository *git.Repository
	path       string
}

func (g GitSource) Clean(path string) {

}

func (g *GitSource) GetRepositoryName() string {
	var name string

	parts := strings.Split(g.Url, "/")
	name = parts[len(parts)-1]

	name = strings.Replace(name, ".git", "", -1)

	return name
}

func (g *GitSource) SetPath(BasePath string) string {
	gitPath := g.GetRepositoryName()
	g.path = filepath.Join(BasePath, gitPath)
	return g.path
}

func (g GitSource) GetPath() string {
	return g.path
}

func (g *GitSource) Download() (bool, error) {
	gs, err := git.PlainClone(g.path, false, &git.CloneOptions{
		URL:      g.Url,
		Progress: os.Stdout,
		// SingleBranch:  true,
		// ReferenceName: plumbing.ReferenceName(g.Branch),
	})

	g.Repository = gs
	if err != nil {
		log.Printf(err.Error())
		if err == git.ErrRepositoryAlreadyExists {
			return true, nil
		}
		return false, err
	}

	return true, err
}
