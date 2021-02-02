package repository

import (
	"github.com/Avepa/shortener/pkg/repository/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

type URL interface {
	Add(url, path string) error
	Get(path string) (url string, err error)
}

type Repository struct {
	URL
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		URL: mongodb.NewURLmongoDB(
			mongodb.NewMongoDBCollection("link", db),
		),
	}
}
