package main

import (
	"encoding/json"

	"github.com/mailru/easyjson"
)

type Person struct {
	FirstName string
	LastName  string
}

//easyjson:json
type PersonGenerated struct {
	FirstName string
	LastName  string
}

var (
	person = &Person{
		FirstName: "John",
		LastName:  "Snow",
	}
	personGenerated = &PersonGenerated{
		FirstName: "John",
		LastName:  "Snow",
	}
)

func defaultJsonSerializer() string {
	bytes, _ := json.Marshal(person)
	return string(bytes)
}

func easyJson() string {
	bytes, _ := easyjson.Marshal(personGenerated)
	return string(bytes)
}
