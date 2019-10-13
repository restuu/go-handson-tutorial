package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://api.github.com/users/defunkt")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	Profile := struct {
		Login     string `json:"login"`
		ID        int    `json:"id"`
		SiteAdmin bool   `json:"site_admin"`
		Bio       string `json:"bio"`
	}{}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &Profile); err != nil {
		panic(err)
	}

	fmt.Println("Login: ", Profile.Login)
	fmt.Println("ID: ", Profile.ID)
	fmt.Println("SiteAdmin: ", Profile.SiteAdmin)
	fmt.Println("Bio: ", Profile.Bio)
}
