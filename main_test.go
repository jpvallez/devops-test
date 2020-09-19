package main

import (
	"fmt"
	"testing"
)

func TestGetGitLabCommit(t *testing.T) {
	latestCommitSha, err := getGitLabCommit()
	fmt.Println("Begin unit test getGitLabCommit function")

	if err != nil {
		t.Errorf("getGitLabCommit function returned an error. %s.", err)
		t.Fail()
	}

	// Have we received a 40 character sha?
	if len(latestCommitSha) != 40 {
		t.Errorf("getGitLabCommit response is not a sha: %s.", latestCommitSha)
		t.Fail()
	}
}

func TestGetVersion(t *testing.T) {
	version := getApplicationVersion()
	fmt.Println("Begin unit test getVersion function")
	if (version != "undefined") {
		t.Errorf("Cannot reach version. Got: %s, Want: %s", version, "undefined")
		t.Fail()
	}
}

// Negative testing..?