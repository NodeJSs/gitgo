package main

import (
	"fmt"
	"os"
	"strings"

	flag "github.com/ogier/pflag"
)

var (
	user string
)

func init() {
	flag.StringVarP(&user, "user", "u", "", "Search Users")
}

func main() {

	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]!\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	users := strings.Split(user, ",")
	fmt.Printf("Searching user(s): %s\n", users)
	for index, value := range users {
		result := GetUsersData(value)
		fmt.Println("#", index+1, "Github user")
		fmt.Println("Name: ", result.Name)
		fmt.Println("Location: ", result.Location)
		fmt.Println("Bio: ", result.Bio)
		fmt.Println("Number of Public Repos: ", result.PublicRepos)
		fmt.Println("Followers: ", result.Followers)
		fmt.Println("Following: ", result.Following)
	}

}
