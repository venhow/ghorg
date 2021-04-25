package git

import (
	"os"
	"os/exec"

	"github.com/gabrie30/ghorg/scm"
)

type GitClient struct{}

func NewGit() GitClient {
	return GitClient{}
}

func (g GitClient) Clone(repo scm.Repo) error {
	args := []string{"clone", repo.CloneURL, repo.HostPath}
	if os.Getenv("GHORG_BACKUP") == "true" {
		args = append(args, "--mirror")
	}

	cmd := exec.Command("git", args...)
	err := cmd.Run()
	return err
}

func (g GitClient) SetOrigin(repo scm.Repo) error {
	// TODO: make configs around remote name
	// we clone with api-key in clone url
	args := []string{"remote", "set-url", "origin", repo.URL}
	cmd := exec.Command("git", args...)
	cmd.Dir = repo.HostPath
	return cmd.Run()
}

func (g GitClient) Checkout(repo scm.Repo) error {
	cmd := exec.Command("git", "checkout", repo.CloneBranch)
	cmd.Dir = repo.HostPath
	return cmd.Run()
}

func (g GitClient) Clean(repo scm.Repo) error {
	cmd := exec.Command("git", "clean", "-f", "-d")
	cmd.Dir = repo.HostPath
	return cmd.Run()
}

func (g GitClient) UpdateRemote(repo scm.Repo) error {
	cmd := exec.Command("git", "remote", "update")
	cmd.Dir = repo.HostPath
	return cmd.Run()
}

func (g GitClient) Pull(repo scm.Repo) error {
	// TODO: handle case where repo was removed, should not give user an error
	cmd := exec.Command("git", "pull", "origin", repo.CloneBranch)
	cmd.Dir = repo.HostPath
	return cmd.Run()
}

func (g GitClient) Reset(repo scm.Repo) error {
	cmd := exec.Command("git", "reset", "--hard", "origin/"+repo.CloneBranch)
	cmd.Dir = repo.HostPath
	return cmd.Run()
}
