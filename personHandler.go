package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// All retrieves all persons from the database
func All() ([]Person, error) {

	persons := []Person{}
	for key, element := range PersonMap {
		persons = append(persons, element)
		fmt.Println("Key:", key, "=>", "Element:", element)
	}

	return persons, nil
}

// One returns a single user record from the database
func One(name string) (*Person, error) {

	p := PersonMap[name]
	return &p, nil
}

func bodyToPerson(r *http.Request, p *Person) error {
	if r == nil {
		return errors.New("a request is required")
	}
	if r.Body == nil {
		return errors.New("request body is empty")
	}
	if p == nil {
		return errors.New("a person is required")
	}
	bd, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bd, p)
}

func personsGetAll(w http.ResponseWriter, r *http.Request) {

	persons, err := All()
	if err != nil {

		postError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}

	postBodyResponse(w, http.StatusOK, jsonResponse{"persons": persons})
}

func personsPostOne(w http.ResponseWriter, r *http.Request) {
	p := new(Person)
	err := bodyToPerson(r, p)

	if err != nil {
		postError(w, http.StatusBadRequest)
		return
	}
	PersonMap[p.Name] = *p

	w.WriteHeader(http.StatusCreated)
}

func personsGetOne(w http.ResponseWriter, r *http.Request, name string) {

	p, err := One(name)
	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}
	if r.Method == http.MethodHead {
		postBodyResponse(w, http.StatusOK, jsonResponse{})
		return
	}

	postBodyResponse(w, http.StatusOK, jsonResponse{"person": p})
}
