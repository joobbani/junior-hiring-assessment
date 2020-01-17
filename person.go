package main

// Person model holds all information relevant to the API surface involving /people and /people/
type Person struct {
	Name     string `json:"name"`
	Age      int64
	FavColor string `json:"favorite_color"`
}
