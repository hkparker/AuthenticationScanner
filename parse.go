package main

import (
	"fmt"
	"encoding/json"
	"errors"
	)

type Parents struct {
	Mother string
	Father string
	MothersParents *Parents
	FathersParents *Parents
}

func (parents Parents) GrandmotherOnMotherSide() (name string, err error) {
	defer func() {
        if r := recover(); r != nil {
            err = errors.New("Error: mother name not present")
        }
    }()
    name = parents.MothersParents.Mother
    return
}

func PrintStructParent(json_encoded string) {
	fmt.Print("JSON encoded data:\n", json_encoded, "\n")
	fmt.Println("Unmarshalling...")
	parents := &Parents{}
	err := json.Unmarshal([]byte(json_encoded), &parents)
	if err != nil {
		panic(err)
	}
	fmt.Println("Created struct.")
	fmt.Println("Printing Alice's grandmother on her mother's side")
	aname, err := parents.MothersParents.GrandmotherOnMotherSide()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(aname)
	}
	fmt.Println("Printing Dave's grandmother on her mother's side")
	dname, err := parents.FathersParents.GrandmotherOnMotherSide()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(dname)
	}
	
}

func main(){
	json_encoded := `{
	"Mother": "Alice",
	"Father": "Bob",
	"MothersParents": {
		"Mother": "Nancy",
		"Father": "Dave",
		"MothersParents": {
			"Mother": "Anna",
			"Father": "Dan"
		},
		"FathersParents": {
			"Mother": "Paula",
			"Father": "Jake"
		}
	},
	"FathersParents": {
		"Mother": "Eve",
		"Father": "Drew"
	}
}`
	PrintStructParent(json_encoded)
}

