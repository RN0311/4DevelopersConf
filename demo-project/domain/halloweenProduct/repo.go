package halloweenproduct

import (
	"github.com/RN0311/demo-project/mongo"
	"github.com/globalsign/mgo/bson"
)

type Repository struct {
	mongoClient *mongo.Client
}

func NewRepository(mongoClient *mongo.Client) *Repository {
	repo := Repository{mongoClient}
	return &repo
}

const (
	databaseName   = "halloween-product"
	collectionName = "halloween-shopping"
)

func (r *Repository) GetById(id string) (*HalloweenProduct, error) {
	var session = r.mongoClient.NewSession()
	defer session.Close()

	query := bson.M{"Id": id}

	var com *HalloweenProduct
	err := session.
		DB(databaseName).
		C(collectionName).
		Find(query).
		One(&com)

	if err != nil {
		return nil, err
	}

	return com, nil
}
func (r *Repository) Insert(com *HalloweenProduct) error {
	var session = r.mongoClient.NewSession()
	defer session.Close()

	err := session.
		DB(databaseName).
		C(collectionName).
		Insert(com)

	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(id string) error {
	var session = r.mongoClient.NewSession()
	defer session.Close()

	query := bson.M{"Id": id}

	err := session.
		DB(databaseName).
		C(collectionName).
		Remove(query)

	if err != nil {
		return err
	}
	return nil
}
