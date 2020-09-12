package main

import (
	"fmt"
	"testing"
)

func TestGetGithubCommit(t *testing.T) {
	latestCommitSha, err := getGithubCommit()
	fmt.Println("Begin unit testing getGithubCommit function")

	if err != nil {
		t.Errorf("getGithubCommit function returned an error. %s.", err)
		t.Fail()
	}

	// Have we received a 64 character sha?
	if len(latestCommitSha) == 64 {
		t.Errorf("getGithubCommit response is not a sha: %s.", latestCommitSha)
		t.Fail()
	}
}
