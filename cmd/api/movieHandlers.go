package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/ccontreraso/backend_app/models"
	"github.com/julienschmidt/httprouter"
)

	type jsonResponse struct {
		OK bool `json:"ok"`
		Message string `json:"message"`
	}

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print("Invalid ID Parameter!")
		app.errorJson(w, err)
		return
	}

	app.logger.Println("ID is: ", id)

	movie, err := app.models.DB.Get(id)
	if err != nil {
		app.logger.Print("Error getting movie!")
		app.errorJson(w, err)
		return
	}
	// movie := models.Movie{
	// 	ID: id,
	// 	Title: "The Go Programming Language",
	// 	Year: 2009,
	// 	Description: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.",
	// 	ReleaseDate: time.Date(2009, 8, 01, 01, 0,0,0, time.Local), // "2009-08-01",
	// 	Runtime: 100,
	// 	Rating: 5.0,
	// 	MPAARating: "PG-13",
	// 	CreatedAt: time.Now(),
	// 	UpdatedAt: time.Now(),
	// }

	_ = app.writeJSON(w, http.StatusOK, movie, "movie")
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.models.DB.All()
	if err != nil {
		app.logger.Print("Error getting movies!")
		app.errorJson(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, movies, "movies"); if err != nil {
		app.logger.Print("Error writing JSON!")
		app.errorJson(w, err)
		return
	}
}

func (app *application) getAllGenres(w http.ResponseWriter, r *http.Request) {
	genres, err :=  app.models.DB.AllGenres()

	if err != nil {
		app.errorJson(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, genres, "genres"); if err != nil {
		app.errorJson(w, err)
		return
	}
}

func (app *application) getAllMoviesByGenre(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	genreID, err := strconv.Atoi(params.ByName("genre_id")); if err != nil {
		app.errorJson(w, err)
		return
	}

	movies, err := app.models.DB.All(genreID); if err != nil {
		app.errorJson(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, movies, "movies"); if err != nil {
		app.errorJson(w, err)
		return
	}
}

func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request) {

	// delete the movie
	// return the response

	params := httprouter.ParamsFromContext(r.Context())
	
	id, err := strconv.Atoi(params.ByName("id")); if err != nil {
		app.errorJson(w, err)
		return
	}

	err = app.models.DB.DeleteMovie(id); if err != nil {
		app.errorJson(w, err)
		return
	}

	ok := jsonResponse{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response"); if err != nil {
		app.errorJson(w, err)
		return
	}
	
}

func (app *application) insertMovie(w http.ResponseWriter, r *http.Request) {
	type jsonResponse struct {
		OK bool `json:"ok"`
	}

	ok := jsonResponse{
		OK: true,
	}

	err := app.writeJSON(w, http.StatusOK, ok, "response");
	if err != nil {
		app.logger.Print("Error writing JSON!")
		app.errorJson(w, err)
		return
	}
}

type MoviePayload struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Year string `json:"year"`
	Runtime string `json:"runtime"`
	Rating string `json:"rating"`
	MPAARating string `json:"mpaa_rating"`
	ReleaseDate string `json:"release_date"`
}

func (app *application) updateMovie(w http.ResponseWriter, r *http.Request) {
	var payload MoviePayload

	err := json.NewDecoder(r.Body).Decode(&payload); if err != nil {
		app.logger.Print("Error decoding JSON!")
		app.errorJson(w, err)
		return
	}

	var movie models.Movie

	if (payload.ID != "0") {
		id, _ := strconv.Atoi(payload.ID)
		m, _ := app.models.DB.Get(id)
		movie = *m
		movie.UpdatedAt = time.Now()
	}

	movie.ID, _ = strconv.Atoi(payload.ID)
	movie.Title = payload.Title
	movie.Description = payload.Description
	movie.Year, _ = strconv.Atoi(payload.Year)
	movie.Runtime, _ = strconv.Atoi(payload.Runtime)
	movie.Rating, _ = strconv.ParseFloat(payload.Rating, 64)
	movie.MPAARating = payload.MPAARating
	movie.ReleaseDate, _ = time.Parse("2006-01-02", payload.ReleaseDate)
	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()

	if movie.ID == 0 {
		err = app.models.DB.InsertMovie(movie); if err != nil {
			app.logger.Print("Error inserting movie!")
			app.errorJson(w, err)
			return
		}
	} else {
		err = app.models.DB.UpdateMovie(movie); if err != nil {
			app.logger.Print("Error updating movie!")
			app.errorJson(w, err)
			return
		}
	}




	ok := jsonResponse{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response");
	if err != nil {
		app.logger.Print("Error writing JSON!")
		app.errorJson(w, err)
		return
	}
}

func (app *application) searchMovies(w http.ResponseWriter, r *http.Request) {
	
}

