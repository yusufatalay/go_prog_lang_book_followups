package main

import ("fmt"
		"JSON"
)
type Movie struct {
	Title	string
	Year	int	`json:"released"`
	Color	bool  `json:"color,omitempty"`
	Actors	[]string
}

var movies = []Movie{
	{Title: "Casablanca",Year: 1942,Color:false,Actors: []string{"Humphrey Bogart","Ingrid Bergman"}},
	{Title: "Cool Hand Luke",Year: 1967, Color: true,Actors: []string{"Steve McQueen","Jacqueline Bisset"}}
}

func main(){

	data , _ := json.Marshal(movies)

	fmt.Printf("%s\n",data)
}
