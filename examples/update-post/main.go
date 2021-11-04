package main

import (
	"fmt"
	"github.com/1makarov/go-fakeapi-sdk"
)

func main() {
	fake := fakeapi.New()

	post, err := fake.UpdatePost(fakeapi.PostUpdateInput{
		Title:  "no title #1",
		Body:   "no body #1",
		UserID: 5,
		PostID: 6,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(post.String())
}
