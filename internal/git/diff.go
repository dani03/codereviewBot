package git

import (
	"os/exec"
)

func GetDiff(commitHash string) (string, error) {
	cmd := exec.Command("git", "diff", commitHash+"^!", commitHash)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
