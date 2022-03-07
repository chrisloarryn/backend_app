package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	// movies
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMovies)

	// movie get and post
	router.HandlerFunc(http.MethodPost, "/v1/admin/editmovie", app.updateMovie)
	router.HandlerFunc(http.MethodGet, "/v1/admin/deletemovie/:id", app.deleteMovie)
	router.HandlerFunc(http.MethodGet, "/v1/movie/:id", app.getOneMovie)

	// movies by genre
	router.HandlerFunc(http.MethodGet, "/v1/movies/:genre_id", app.getAllMoviesByGenre)

	// genres
	router.HandlerFunc(http.MethodGet, "/v1/genres", app.getAllGenres)

	return app.enableCORS(router)
}