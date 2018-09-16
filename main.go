package main

func assignRoutes(handler *Handler) {
	firstRoutes(handler)
}


func main() {
	handler := SetupServer()
	assignRoutes(handler)
	StartServer(handler)
}

