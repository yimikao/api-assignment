package main

import (
	"github.com/yimikao/api-assignment/api/app"
	"github.com/yimikao/api-assignment/api/routes"
)

func StartServer(a *app.Server) {
	s, port := a.InitializeRouter()
	routes.RunServer(s, port)
}
func main() {
	StartServer(&app.Server{})
}
