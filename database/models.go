package database

type Movie struct {
	Id       string   `json:"id"`
	Isbn     string   `json:"isbn"`
	Title    string   `json:"title"`
	Director Director `json:"director"`
	Actors   []Actor  `json:"actors"`
}

type Director struct {
	Person
}

type Actor struct {
	Person
}

type Person struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
