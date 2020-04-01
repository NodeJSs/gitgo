package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type githubUser struct {
	name        interface{}
	location    interface{}
	bio         interface{}
	publicRepos interface{}
	followers   interface{}
	following   interface{}
}

var endpointString string = "https://api.github.com/users/"

//GetUsersData queries the github REST API for the user's data
func GetUsersData(users []string) {
	var result []githubUser
	dataMap := make(map[string]interface{})
	fmt.Println("fetching user's data")
	for index, value := range users {
		response, err := http.Get(endpointString + value)
		if err != nil {
			fmt.Printf("Error fetching user's data %s\n", err)
		} else {
			data, _ := ioutil.ReadAll(response.Body)
			err = json.Unmarshal(data, &dataMap)
			result = append(result, githubUser{
				name:        dataMap["name"],
				location:    dataMap["location"],
				bio:         dataMap["bio"],
				publicRepos: dataMap["public_repos"],
				followers:   dataMap["followers"],
				following:   dataMap["following"],
			})

			fmt.Println("#", index+1, "Github user")
			fmt.Println("Name: ", result[index].name)
			fmt.Println("Location: ", result[index].location)
			fmt.Println("Bio: ", result[index].bio)
			fmt.Println("Number of Public Repos: ", result[index].publicRepos)
			fmt.Println("Followers: ", result[index].followers)
			fmt.Println("Following: ", result[index].following)

		}
	}

}
