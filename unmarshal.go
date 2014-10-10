package main

//
// Simple demo of unmarshalling nested structs in Go
//

import (
	"fmt"
	"encoding/json"
	)

// Struct that the json is going to pupulate
type Parents struct {
	Mother string
	Father string
	MothersParents *Parents
	FathersParents *Parents
}

func main(){
	// The json encoded data, a deeply nested (but incomplete) family
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
	fmt.Print("JSON encoded data:\n", json_encoded, "\n")
	fmt.Println("Unmarshalling...")
	parents := &Parents{}
	err := json.Unmarshal([]byte(json_encoded), &parents)
    if err != nil {
        panic(err)
    }
	fmt.Println("Created struct:")
	fmt.Printf("Mother: %s\n", parents.Mother)
	fmt.Printf("Father: %s\n", parents.Father)
	fmt.Printf("Alice's mother: %s\n", parents.MothersParents.Mother)
	fmt.Printf("Alice's father: %s\n", parents.MothersParents.Father)
	fmt.Printf("Bob's mother: %s\n", parents.FathersParents.Mother)
	fmt.Printf("Bob's father: %s\n", parents.FathersParents.Father)
	fmt.Printf("Nancy's mother: %s\n", parents.MothersParents.MothersParents.Mother)
	fmt.Printf("Nancy's father: %s\n", parents.MothersParents.MothersParents.Father)
	fmt.Printf("Dave's mother: %s\n", parents.MothersParents.FathersParents.Mother)
	fmt.Printf("Dave's father: %s\n", parents.MothersParents.FathersParents.Father)
}
