package app

import "github.com/habibiahmada/shopwithgolang/app/controllers"

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}