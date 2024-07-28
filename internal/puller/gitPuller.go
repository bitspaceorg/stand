package puller

import (
	"log"

	"gopkg.in/src-d/go-git.v4"
)

// this encapsulates all the pull logic of the git folder
type GitPuller struct {
	RepoLink string
	Path     string
}

func (p *GitPuller) Pull() error {
	// Clone the repository
	_, err := git.PlainClone(p.Path, false, &git.CloneOptions{
		URL: p.RepoLink,
		// Auth: &http.BasicAuth{
		// 	Username: "your-username", // Username can be anything except an empty string
		// 	Password: "your-access-token", // Personal access token
		// },
	})
	if err != nil {
		log.Fatalf("Failed to clone the repository: %v", err)
	}
	return nil
}
