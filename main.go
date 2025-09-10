package main

import (
	"fmt"
	"log"
	"net/http"
	"sitcon/alien"
	"sitcon/harry_potter"
	"sitcon/lord_of_the_rings"
	"sitcon/matrix"
	"sitcon/star_wars"
)

func main() {
	// Star Wars themed handlers
	http.HandleFunc("/", star_wars.DroidsHandler)
	http.HandleFunc("/robots.txt", star_wars.RobotsTxtHandler)

	// Harry Potter themed handlers
	http.HandleFunc("/hogwarts", harry_potter.HogwartsHandler)
	http.HandleFunc("/hogwarts/1", harry_potter.GryffindorHandler)
	http.HandleFunc("/hogwarts/2", harry_potter.HufflepuffHandler)
	http.HandleFunc("/hogwarts/3", harry_potter.RavenclawHandler)
	http.HandleFunc("/hogwarts/4", harry_potter.SlytherinHandler)

	// Lord of the Rings themed handlers
	http.HandleFunc("/ring", lord_of_the_rings.LotrHandler)
	http.HandleFunc("/lord_of_the_rings/flag.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "lord_of_the_rings/flag.png")
	})

	// Matrix themed handlers
	http.HandleFunc("/morpheus", matrix.MatrixHandler)
	http.HandleFunc("/matrix/flag.png", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "matrix/flag.png")
	})

	// Alien vs. Predator themed handler
	http.HandleFunc("/alien", alien.AlienHandler)

	fmt.Println("Starting server on http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
