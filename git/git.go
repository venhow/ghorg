package git

import (
	"os"
	"os/exec"
)

type Gitter struct{}

func NewGit() Gitter {
	return Gitter{}
}

func (g Gitter) Clone(cloneURL string, repoDir string) error {
	args := []string{"clone", cloneURL, repoDir}
	if os.Getenv("GHORG_BACKUP") == "true" {
		args = append(args, "--mirror")
	}

	cmd := exec.Command("git", args...)
	err := cmd.Run()
	return err
}

func (g Gitter) SetOrigin(repoURL string, repoDir string) error {
	// TODO: make configs around remote name
	// we clone with api-key in clone url
	args := []string{"remote", "set-url", "origin", repoURL}
	cmd := exec.Command("git", args...)
	cmd.Dir = repoDir
	return cmd.Run()
}

func (g Gitter) Checkout(branch string, repoDir string) error {
	cmd := exec.Command("git", "checkout", branch)
	cmd.Dir = repoDir
	return cmd.Run()
}

func (g Gitter) Clean(repoDir string) error {
	cmd := exec.Command("git", "clean", "-f", "-d")
	cmd.Dir = repoDir
	return cmd.Run()
}

func (g Gitter) UpdateRemote(repoDir string) error {
	cmd := exec.Command("git", "remote", "update")
	cmd.Dir = repoDir
	return cmd.Run()
}

func (g Gitter) Pull(branch string, repoDir string) error {
	// TODO: handle case where repo was removed, should not give user an error
	cmd := exec.Command("git", "pull", "origin", branch)
	cmd.Dir = repoDir
	return cmd.Run()
}

func (g Gitter) Reset(branch string, repoDir string) error {
	cmd := exec.Command("git", "reset", "--hard", "origin/"+branch)
	cmd.Dir = repoDir
	return cmd.Run()
}
