package routes

import (
	"github.com/gorilla/mux"
	"github.com/saltchang/church-music-api/models"
)

var (
	db = models.DB
)

// Routers struct
type Routers struct{}

// InitRouters function
func (router *Routers) InitRouters() *mux.Router {
	mainRouter := mux.NewRouter()

	// Route Handlers and Endpoints

	// Route: Home
	mainRouter.HandleFunc("/", GetIndex).Methods("GET")

	// Route: Get all songs
	mainRouter.HandleFunc("/api/songs", GetAllSongs).Methods("GET")

	// Route: Get a song by its SID
	mainRouter.HandleFunc("/api/songs/sid/{sid}", GetSongBySID).Methods("GET")

	// Route: Get a song by searching title
	mainRouter.HandleFunc("/api/songs/search", GetSongBySearch).Queries("title", "{title}", "lang", "{lang}", "c", "{c}", "to", "{to}").Methods("GET")

	// Route: Create a song
	mainRouter.HandleFunc("/api/songs", CreateSong).Methods("POST")

	// Route: Update a song by its SID
	mainRouter.HandleFunc("/api/songs/sid/{sid}", UpdateSong).Methods("PUT")

	// Route: Delete a song by its SID
	mainRouter.HandleFunc("/api/songs/{sid}", DeleteSong).Methods("DELETE")

	return mainRouter
}