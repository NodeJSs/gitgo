package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//GithubUser defines the sturcture of each user
type GithubUser struct {
	Name        string `json:"name"`
	Location    string `json:"location"`
	Bio         string `json:"bio"`
	PublicRepos int    `json:"public_repos"`
	Followers   int    `json:"followers"`
	Following   int    `json:"following"`
}

const (
	apiURL       = "https://api.github.com"
	userEndpoint = "/users/"
)

//GetUsersData queries the github REST API for the user's data
func GetUsersData(user string) GithubUser {
	var result GithubUser
	fmt.Println("fetching user's data")

	response, err := http.Get(apiURL + userEndpoint + user)

	if err != nil {
		log.Fatalf("Error fetching user's data %s\n", err)
	}
	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal(data, &result)
	return result
}
