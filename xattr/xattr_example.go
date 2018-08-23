package main

import (
	"fmt"
	"log"

	"github.com/pkg/xattr"
)

func main() {

	path := "abc.txt"
	prefix := "user."

	fmt.Println("Begin listing xattr")
	var list []string
	list, err := xattr.List(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("list is ", list)
	fmt.Println("Finish listing xattr")

	fmt.Println("Begin setting xattr")
	if err := xattr.Set(path, prefix+"metadata-attr", []byte("metadata-value")); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Finished setting xattr")

	fmt.Println("Begin listing xattr")
	list, err = xattr.List(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("list is ", list)
	fmt.Println("Finish listing xattr")

	fmt.Println("Begin getting xattr")
	var data []byte
	data, err = xattr.Get(path, prefix+"metadata-attr")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data got is", string(data))
	fmt.Println("Finished getting xattr")

	fmt.Println("Begin removing xattr")
	if err = xattr.Remove(path, prefix+"metadata-attr"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Finished removing xattr")

	fmt.Println("Begin listing xattr")
	list, err = xattr.List(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("list is ", list)
	fmt.Println("Finish listing xattr")
}
