package main

import (
	"crudApi/models"
	"crudApi/router"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	models.Movies = append(models.Movies, models.Movie{ID: "1", Isbn: "34253", Title: "Tiger", Director: &models.Director{Firstname: "Jhon", Lastname: "Doe"}})
	models.Movies = append(models.Movies, models.Movie{ID: "2", Isbn: "64632", Title: "Jawan", Director: &models.Director{Firstname: "Rony", Lastname: "Roy"}})
	
    router.AddAppRoute(r)

	fmt.Printf("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
