package main

import (
	"fmt"

	"github.com/daniel/code-review-bot/internal/git"
	"github.com/daniel/code-review-bot/internal/gpt"
)

func main() {
	commitHash := "HEAD"
	diff, err := git.GetDiff(commitHash)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	analysis, err := gpt.AnalyzeDiff(diff)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("AI Analysis:", analysis)
}
