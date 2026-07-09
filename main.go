package main

import "go1f/pkg/server"

func main() {
	err := server.Run()
	if err != nil {
		panic(err)
	}
}
