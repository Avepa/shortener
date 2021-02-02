package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func NewMongoDB(cfg *Config) (*mongo.Database, error) {
	clientOpt := options.Client().ApplyURI("mongodb://" +
		cfg.Host + ":" +
		cfg.Port +
		"/?readPreference=primary&appname=MongoDB%20Compass%20Community&ssl=false",
	)

	client, err := mongo.Connect(context.TODO(), clientOpt)
	if err != nil {
		return nil, err
	}

	db := client.Database(cfg.DBName)
	return db, nil
}

func NewMongoDBCollection(collName string, db *mongo.Database) *mongo.Collection {
	coll := db.Collection(collName)
	return coll
}
