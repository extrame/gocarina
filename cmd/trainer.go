package main

import (
	"log"

	"github.com/armhold/gocarina"
)

func main() {
	m := gocarina.ProcessGameBoards()

	tile := m['A']
	tileWidth := tile.Bounds().Dx()
	tileHeight := tile.Bounds().Dy()
	n := gocarina.NewNetwork(tileWidth, tileHeight)

	for i := 0; i < 100; i++ {
		for r, tile := range m {
			log.Printf("training: %c\n", r)
			n.Train(tile, r)
		}
	}

	n.Save("trained_network.out")

	log.Printf("success")
}
