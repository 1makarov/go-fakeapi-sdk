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

	post, err := client.GetPostByID(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(post.UserID, post.Title)
}
