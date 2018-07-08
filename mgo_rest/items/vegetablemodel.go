package items

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// vegetableModel models a Vegetable
type vegetableModel struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string
	Description string
}

// vegetableModelIndex returns an index for Vegetable
func vegetableModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"Name"},
		Unique:     true,
		DropDups:   true,
		Sparse:     true,
		Background: true,
	}
}

func newVegetableModel(v *Vegetable) *vegetableModel {
	return &vegetableModel{
		Name:        v.Name,
		Description: v.Description,
	}
}

func toVegetable(model *vegetableModel) *Vegetable {
	return &Vegetable{
		model.ID.Hex(),
		model.Name,
		model.Description,
	}
}
