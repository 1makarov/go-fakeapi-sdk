package main

import (
	"fmt"
	"github.com/1makarov/go-fakeapi-client"
)

func main() {
	fake := fakeapi.New()

	post, err := fake.CreatePost(fakeapi.PostCreateInput{
		Title:  "no title #1",
		Body:   "no body #1",
		UserID: 001,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(post.String())
}
