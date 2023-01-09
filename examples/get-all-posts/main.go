package main

import (
	"fmt"
	"net/http"

	"github.com/1makarov/go-fakeapi-sdk"
)

func main() {
	client := fakeapi.New(fakeapi.Endpoint, http.DefaultClient)

	posts, err := client.GetAllPosts()
	if err != nil {
		panic(err)
	}

	for _, p := range posts {
		fmt.Println(p.UserID, p.Title)
	}
}
