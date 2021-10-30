package main

import (
	"fmt"
	"github.com/1makarov/go-fakeapi-client"
)

func main() {
	fake := fakeapi.New()

	posts, err := fake.GetAllPosts()
	if err != nil {
		panic(err)
	}

	for _, p := range posts {
		fmt.Println(p.String())
	}
}
