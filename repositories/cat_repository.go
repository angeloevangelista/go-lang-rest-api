package cat_repositories

import "github.com/angeloevangelista/go-rest-api/models"

var cats []models.Cat

func GetCats() []models.Cat {
	return cats
}

func SetCats(newCats []models.Cat) {
	cats = newCats
}
