package main

import (
	"github.com/zhendong233/Books/internal/book"
)

func main() {
	b, err := book.WireBuild()
	if err != nil {
		panic(err)
	}
	b.Run()
}
