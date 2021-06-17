package names

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"time"

	"github.com/mmuoDev/commons/httputils"
)

//Person represents names of persons
type Person struct {
	Names []string `json:"names"`
}

func IsValidTribe(tribe string) bool {
	tribes := [3]string{"yoruba", "igbo", "hausa"}
	for _, t := range tribes {
		if tribe == t {
			return true
		}
	}
	return false
}

func GetNames(tribe, gender string) []string {
	var person Person
	if tribe == "igbo" {
		if gender == "female" {
			httputils.FileToStruct(filepath.Join("names", "igbo_female.json"), &person)
		} else {
			httputils.FileToStruct(filepath.Join("names", "igbo_male.json"), &person)
		}
	} else if tribe == "yoruba" {
		if gender == "female" {
			httputils.FileToStruct(filepath.Join("names", "yoruba_female.json"), &person)
		} else {
			httputils.FileToStruct(filepath.Join("names", "yoruba_male.json"), &person)
		}
	} else if tribe == "hausa" {
		if gender == "female" {
			httputils.FileToStruct(filepath.Join("names", "hausa_female.json"), &person)
		} else {
			httputils.FileToStruct(filepath.Join("names", "hausa_male.json"), &person)
		}
	}

	return person.Names
}

func GenerateRandomNames(tribe, gender string, count int) ([]string, error) {
	var res []string
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	isValid := IsValidTribe(tribe)
	if !isValid {
		return res, fmt.Errorf("The tribe, %s is an invalid tribe", tribe)
	}
	names := GetNames(tribe, gender)
	for i := 1; i <= count; i++ {
		n := names[rand.Intn(len(names))]
		res = append(res, n)
	}

	return res, nil
}
