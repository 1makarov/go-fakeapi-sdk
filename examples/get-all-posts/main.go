package main

import (
	"fmt"
	"net/http"

	"github.com/1makarov/go-fakeapi-sdk"
)

func main() {
	client, err := fakeapi.New(fakeapi.Endpoint, http.DefaultClient)
	if err != nil {
		panic(err)
	}

	posts, err := client.GetAllPosts()
	if err != nil {
		panic(err)
	}

	for _, p := range posts {
		fmt.Println(p.UserID, p.Title)
	}
}
