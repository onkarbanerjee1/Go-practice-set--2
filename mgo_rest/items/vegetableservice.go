package items

import (
	"gopkg.in/mgo.v2"
)

// VegetableService holds the collection represents Vegetables
type VegetableService struct {
	Collection *mgo.Collection
}

// NewVegetableService returns a service to access Vegetables in database
func NewVegetableService(sess *mgo.Session, dbName, collectionName string) *VegetableService {
	collection := sess.DB(dbName).C(collectionName)
	collection.EnsureIndex(vegetableModelIndex())
	return &VegetableService{collection}
}

// Add adds a vegetable
func (vegetableService *VegetableService) Add(vegetable *Vegetable) error {
	vegModel := newVegetableModel(vegetable)
	return vegetableService.Collection.Insert(vegModel)
}
