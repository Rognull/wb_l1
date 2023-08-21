package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Address string
}

type Book struct {
	Title  string
	Author string
}


type JSONAdapter interface {
	ToJson() (string, error)
}

type PersonAdapter struct {
	Person
}

func AdapterToJson(ja JSONAdapter)(string,error){
	data, err := json.Marshal(ja)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (pa *PersonAdapter) ToJson() (string, error) {
	data, err := json.Marshal(pa)
	if err != nil {
		return "", err
	}
	return string(data), nil
}


type BookAdapter struct {
	Book
}

func (ba *BookAdapter) ToJson() (string, error) {
	data, err := json.Marshal(ba)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

 


func main() {
	person := &PersonAdapter{Person{Name: "kirill", Age: 29, Address: "Russia Moskow"}}
	book := &BookAdapter{Book{Title: "Кладбище жЫвотных", Author: "С.Кинг"}}

	personJson, err := person.ToJson()

	if err != nil {
		panic(err)
	}
	bookJson, err := book.ToJson()
	if err != nil {
		panic(err)
	}

 
	fmt.Println("Person JSON:", personJson)
	fmt.Println("Book JSON:", bookJson)

	newString, err :=AdapterToJson(person)
	
	if err != nil {
		panic(err)
	}

	fmt.Println("Person JSON:", newString)
}