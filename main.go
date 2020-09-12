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

func getGithubCommit() string {
	resp, err := http.Get("https://api.github.com/repos/jpvallez/learn-docker/commits/master")
	if err != nil {
		// TODO: Handle this properly
		return "Unable to get github commit sha."
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// TODO: Handle this properly
		return "Unable to get github commit sha."
	}

	var sha GitSha
	json.Unmarshal(body, &sha)

	return sha.Sha
}

func getApplicationVersion() string {
	// Will need to understand this properly
	return "1.0"
}

func serviceResponse(w http.ResponseWriter, r *http.Request) {
	response := Response{
		LastCommitSha: getGithubCommit(),
		Version:       "1.0",
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
