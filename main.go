package main

import (
	"http-request/routers"
)

func main() {
	var PORT = ":8080"

	// router.StartServer().Run(PORT)
	routers.StartServer().Run(PORT)
}
