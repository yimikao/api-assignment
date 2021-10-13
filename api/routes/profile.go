package routes

import (
	"github.com/yimikao/api-assignment/api/app"
	"github.com/yimikao/api-assignment/api/handlers"
)

func RunServer(a *app.Server, port string) {
	a.Router.HandleFunc("/api/profiles", handlers.CreateProfile).Methods("POST")
	a.Router.HandleFunc("/api/profiles", handlers.GetPausedProfiles).
		Queries(
			"status", "{paused}",
		).
		Methods("GET")
	a.Router.HandleFunc("/api/profiles", handlers.GetAllProfiles).Methods("GET")
	a.Router.HandleFunc("/api/profiles/{id}/pause", handlers.PauseProfile).Methods("PUT")
	a.Router.HandleFunc("/api/profiles/{id}/unpause", handlers.UnPauseProfile).Methods("PUT")
	a.Router.HandleFunc("/api/profiles/{id}/delete", handlers.DeleteProfile).Methods("DELETE")

	a.Run(":" + port)
}
