package testutil

import (
	"os/exec"
	"strings"
)

var RepoDir string

func init() {
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		panic(err)
	}
	RepoDir = strings.TrimSpace(string(out))
}
