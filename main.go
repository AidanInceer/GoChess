package main

import (
	"chess/chess"
)

func main() {
	game := chess.Game{}

	game.Setup()
	game.Play()

}
