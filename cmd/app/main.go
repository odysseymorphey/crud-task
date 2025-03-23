package main

import "github.com/odysseymorphey/crud-task/internal/server"

func main() {
	app := server.NewServer()

	app.Run()
}