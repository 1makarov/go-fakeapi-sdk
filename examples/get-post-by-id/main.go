package main

import (
	"fmt"
	"github.com/1makarov/go-fakeapi-client"
)

func main() {
	fake := fakeapi.New()

	post, err := fake.GetPostByID(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(post.String())
}
