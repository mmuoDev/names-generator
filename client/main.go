package main

import (
	"fmt"
	"log"
	"math/rand"
	genNames "names-generator"
	"time"
)

func GenerateRandomNames(tribe, gender string, count int) ([]string, error) {
	var res []string
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	isValid := genNames.IsValidTribe(tribe)
	if !isValid {
		return res, fmt.Errorf("The tribe, %s is an invalid tribe", tribe)
	}
	names := genNames.GetNames(tribe, gender)
	for i := 1; i <= count; i++ {
		n := names[rand.Intn(len(names))]
		res = append(res, n)
	}

	return res, nil
}

func main() {
	res, err := GenerateRandomNames("yoruba", "male", 1)
	log.Fatal(err, len(res))
}