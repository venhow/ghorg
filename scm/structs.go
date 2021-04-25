package scm

// Repo represents an SCM repo, Path can be optionally set by the user, HostPath is the full final path to find the repo in the users machine. CloneURL includes authentication while URL does not, the remote will be reset to the URL after cloning.
type Repo struct {
	Name        string
	HostPath    string
	Path        string
	URL         string
	CloneURL    string
	CloneBranch string
}
