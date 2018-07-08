package main

import (
	"fmt"
	"os"

	"github.com/medivo/Go-practice-set--2/mgo_rest/items"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

// // Order holds a sample order
// type Order struct {
// 	OrdererName string   `json:"orderer-name"`
// 	OrderItems  []string `json:"items"`
// }

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println("expected no error, got", err)
		os.Exit(1)
	}
	defer session.Close()

	vegetableService := items.NewVegetableService(session.Copy(), "mgoRest", "Vegetables")

	cauliflower := &items.Vegetable{
		Name:        "Cauliflower",
		Description: "Season's best cauliflowers",
	}

	err = vegetableService.Add(cauliflower)
	if err != nil {
		fmt.Println("expected no error, got", err)
		os.Exit(1)
	}
	fmt.Println("Cauliflower inserted")
}
