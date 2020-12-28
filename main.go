package main

import (
	"log"
	"os"

	"github.com/zhendong233/Books/internal/book"
)

func main() {
	b, err := book.WireBuild()
	if err != nil {
		log.Print(err)
		os.Exit(-1)
	}
	b.Run()
}
