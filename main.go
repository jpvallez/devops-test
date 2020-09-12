package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type GitSha struct {
	Sha string `json:"sha"`
}

type Response struct {
	LastCommitSha string
	Version       string
	Description   string
}

var version = "undefined"

func getGithubCommit() (string, error) {
	resp, err := http.Get("https://api.github.com/repos/jpvallez/devops-test/commits/master")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var sha GitSha
	json.Unmarshal(body, &sha)

	return sha.Sha, err
}

func getApplicationVersion() string {
	return version
}

func serviceResponse(w http.ResponseWriter, r *http.Request) {
	// Get the latest github commit for our repo (hardcoded).
	latestCommit, _ := getGithubCommit()
	response := Response{
		LastCommitSha: latestCommit,
		Version:       getApplicationVersion(),
		Description:   "This is a pre-interview technical test",
	}

	json.NewEncoder(w).Encode(response)
}

func handleRequests() {
	http.HandleFunc("/version", serviceResponse)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
