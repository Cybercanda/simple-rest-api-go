package main

import "simple-rest-api-go/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
