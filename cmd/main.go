package main

import "github.com/RiadMefti/go-api-boilerplate/cmd/api"

func main() {
	server := api.NewApiServer(":3000")

	api.Run(server)
}
