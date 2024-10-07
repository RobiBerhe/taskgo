package main

import "taskgo/adapters/http"

func main() {
	server := http.NewRouter()
	server.Run(":6060")
}
