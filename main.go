package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	models "github.com/angeloevangelista/go-rest-api/models"
	repositories "github.com/angeloevangelista/go-rest-api/repositories"
	cat_routes "github.com/angeloevangelista/go-rest-api/routes"
)

func ResponseHeadersMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Add("Content-Type", "application/json")

		handler.ServeHTTP(responseWriter, request)
	})
}

func main() {
	repositories.SetCats(append(
		repositories.GetCats(),
		models.Cat{
			Id:   1,
			Name: "Luna",
			Owner: &models.Person{
				Name: "James",
				Age:  22,
			},
		},
		models.Cat{
			Id:   2,
			Name: "Milo",
			Owner: &models.Person{
				Name: "Robert",
				Age:  36,
			},
		},
		models.Cat{
			Id:   3,
			Name: "Jessie",
			Owner: &models.Person{
				Name: "Fernanda",
				Age:  24,
			},
		},
	))

	router := mux.NewRouter()

	router.HandleFunc("/cats", cat_routes.GetCats).Methods("GET")
	router.HandleFunc("/cats/{id}", cat_routes.GetCat).Methods("GET")
	router.HandleFunc("/cats", cat_routes.CreateCat).Methods("POST")
	router.HandleFunc("/cats/{id}", cat_routes.DeleteCat).Methods("DELETE")

	router.Use(ResponseHeadersMiddleware)

	log.Fatal(http.ListenAndServe(":3333", router))
}
