package main

import (
	"fmt"
	"net/http"

	"github.com/1makarov/go-fakeapi-sdk"
)

func main() {
	client := fakeapi.New(fakeapi.Endpoint, http.DefaultClient)

	post, err := client.UpdatePost(100, 1, "title #1", "body #1")
	if err != nil {
		panic(err)
	}

	fmt.Println(post.String())
}
