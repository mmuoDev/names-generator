package main

import (
	"log"
	"github.com/mmuoDev/names-generator/names"
)

func main() {
	n, err := names.GenerateRandomNames("igbo", "male", 2)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(n)
}
