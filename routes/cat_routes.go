package cat_routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/angeloevangelista/go-rest-api/models"
	repositories "github.com/angeloevangelista/go-rest-api/repositories"
)

func GetCats(responseWriter http.ResponseWriter, request *http.Request) {
	json.NewEncoder(responseWriter).Encode(repositories.GetCats())
}

func GetCat(responseWriter http.ResponseWriter, request *http.Request) {
	routeParams := mux.Vars(request)
	catId, _ := strconv.Atoi(routeParams["id"])

	for _, cat := range repositories.GetCats() {
		if cat.Id == catId {
			json.NewEncoder(responseWriter).Encode(cat)
			return
		}
	}

	json.NewEncoder(responseWriter).Encode(&models.Cat{})
}

func CreateCat(responseWriter http.ResponseWriter, request *http.Request) {
	lastCatId := 0

	cats := repositories.GetCats()

	for _, currentCat := range cats {
		if currentCat.Id > lastCatId {
			lastCatId = currentCat.Id
		}
	}

	var cat models.Cat

	_ = json.NewDecoder(request.Body).Decode(&cat)
	cat.Id = lastCatId + 1

	repositories.SetCats(append(cats, cat))

	responseWriter.WriteHeader(201)
	json.NewEncoder(responseWriter).Encode(cat)
}

func DeleteCat(responseWriter http.ResponseWriter, request *http.Request) {
	routeParams := mux.Vars(request)
	catId, _ := strconv.Atoi(routeParams["id"])

	cats := repositories.GetCats()

	for index, cat := range cats {
		if cat.Id == catId {
			repositories.SetCats(append(cats[:index], cats[index+1:]...))
			break
		}
	}

	responseWriter.WriteHeader(204)
}
