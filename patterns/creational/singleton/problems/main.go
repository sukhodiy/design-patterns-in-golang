package main

import (
	"fmt"
	"sync"
)

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// sync.Once, init() - thread safety
// laziness

var once sync.Once
var instance *singletonDatabase

func readData(path string) (map[string]int, error) {
	// dummy to read data from file

	result := make(map[string]int)
	result["Tokyo"] = 5000000
	result["Kyiv"] = 3000000
	return result, nil
}

func GetSingletonDatabse() *singletonDatabase {
	once.Do(func() {
		caps, err := readData("./capitals.txt")
		db := singletonDatabase{caps}
		if err == nil {
			db.capitals = caps
		}
		instance = &db
	})
	return instance
}

// the problem is we break Dependency Inversion Principle
// we don't use interface for DB, that's why it would be hard for testins and improving

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabse().GetPopulation(city)
	}
	return result
}

func main() {
	cities := []string{"Tokyo", "Kyiv"}
	tp := GetTotalPopulation(cities)
	ok := tp == 8000000
	fmt.Println(ok)
}
