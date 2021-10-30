package main

import (
	"fmt"
	"github.com/1makarov/go-fakeapi-client"
)

func main() {
	fake := fakeapi.New()

	if err := fake.DeletePostByID(1); err != nil {
		panic(err)
	}

	fmt.Println("successfully deleted")
}