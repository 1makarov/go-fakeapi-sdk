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

	if err := client.DeletePostByID(1); err != nil {
		panic(err)
	}

	fmt.Println("successfully deleted")
}