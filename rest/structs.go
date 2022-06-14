package main

// album represents data about a record album.
type Person struct {
    ID     string  `json:"id"`
    Name     string  `json:"name"`
    Gender  string  `json:"gender"`
    Homeworld string  `json:"homeworld"`
    BirthYear string  `json:"birthYear"`
}

type Planet struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Gravity string `json:"gravity"`
	Climate string `json:"climate"`
}
