package main

import (
	"fmt"
	"net/http"

	"github.com/1makarov/go-fakeapi-sdk"
)

func main() {
	client := fakeapi.New(fakeapi.Endpoint, http.DefaultClient)

	post, err := client.GetPostByID(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(post.UserID, post.Title)
}
