package router

import (
	"crudApi/controllers"
	"github.com/gorilla/mux"
)

func AddAppRoute(r *mux.Router){

	r.HandleFunc("/movies", controllers.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", controllers.GetMovie).Methods("GET")
	r.HandleFunc("/movies", controllers.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", controllers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", controllers.DeleteMovie).Methods("DELETE")

}