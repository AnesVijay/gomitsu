package main

import (
	"os"
)

func getToken() string {
	data, err := os.ReadFile("token")
	if err != nil {
		panic(err)
	}

	return string(data)
}
