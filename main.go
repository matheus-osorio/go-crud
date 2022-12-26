package main

import "gocrud/server"

func main() {
	server.Setup()
	server.Run("8080")
}
