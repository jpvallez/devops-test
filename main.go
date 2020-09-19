package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GitSha struct {
	Sha string `json:"id"`
}

type Response struct {
	LastCommitSha string
	Version       string
	Description   string
}

var version = "undefined"

func getGitLabCommit() (string, error) {
	resp, err := http.Get("https://gitlab.com/api/v4/projects/21089254/repository/commits/master")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var sha GitSha
	err = json.Unmarshal(body, &sha)
	if err != nil {
		return "", err
	}

	return sha.Sha, err
}

func getApplicationVersion() string {
	return version
}

func serviceResponse(w http.ResponseWriter, r *http.Request) {
	// Get the latest github/lab commit for our repo
	latestCommit, _ := getGitLabCommit()

	// Create response object
	response := Response{
		LastCommitSha: latestCommit,
		Version:       getApplicationVersion(),
		Description:   "This is a pre-interview technical test",
	}

	// Respond with our nice json
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		// Handle issue with /version endpoint. Right now I'll throw a 500
		// if this ever happens...
		fmt.Println("Issue encoding json in response object")
		w.WriteHeader(500)
	}
}

func handleRequests() {
	http.HandleFunc("/version", serviceResponse)
	fmt.Println("Starting web service...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
