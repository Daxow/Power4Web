package main

import (
	"fmt"
	"net/http"
)

type Game struct {
	Board  [6][7]int
	Player int
	Winner int
}

var game = Game{
	Player: 1,
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Puissance 4")
	fmt.Fprintln(w, "Joueur:", game.Player, "Gagnant:", game.Winner)
	for i := 0; i < 6; i++ {
		for j := 0; j < 7; j++ {
			fmt.Fprint(w, game.Board[i][j], " ")
		}
		fmt.Fprintln(w)
	}
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe(":8080", nil)
}
