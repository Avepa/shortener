package mongodb

import (
	"context"

	"github.com/Avepa/shortener/pkg"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type URLmongoDB struct {
	coll *mongo.Collection
}

func NewURLmongoDB(coll *mongo.Collection) *URLmongoDB {
	return &URLmongoDB{coll: coll}
}

type URLData struct {
	ID   primitive.ObjectID `bson:"_id"`
	URL  string             `bson:"url"`
	PATH string             `bson:"path"`
}

func (db *URLmongoDB) Add(url, path string) error {
	data := &URLData{
		ID:   primitive.NewObjectID(),
		URL:  url,
		PATH: path,
	}

	_, err := db.coll.InsertOne(context.Background(), data)
	if err != nil {
		mwe := "multiple write errors: [{write errors: [{E11000 duplicate key error collection:"
		n := len(mwe)
		if len(err.Error()) >= n {
			e := err.Error()[:n]
			if e == mwe {
				return pkg.ErrorPathIsBusy
			}
		}
		return err
	}

	return nil
}

func (db *URLmongoDB) Get(path string) (string, error) {
	data := new(URLData)
	find := db.coll.FindOne(context.Background(), bson.M{"path": path})
	if err := find.Decode(data); err != nil {
		if err.Error() == "mongo: no documents in result" {
			return "", pkg.ErrorDcumentNotFound
		} else {
			return "", err
		}
	}
	return data.URL, nil
}
