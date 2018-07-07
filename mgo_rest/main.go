package main

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
)

// Order holds a sample order
type Order struct {
	OrdererName string   `json:"orderer-name"`
	OrderItems  []string `json:"items"`
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println("expected no error, got", err)
		os.Exit(1)
	}
	defer session.Close()

	sampleOrder := Order{"Onkar", []string{"Potatoes", "Tomatoes", "Onions"}}
	fmt.Println("Sample order is", sampleOrder)

	c := session.DB("mgoRest").C("Order")
	err = c.Insert(sampleOrder)
	if err != nil {
		fmt.Println("expected no error, got", err)
		os.Exit(1)
	}
}
