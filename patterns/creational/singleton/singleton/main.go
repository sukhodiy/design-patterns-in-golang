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

func main() {
	db := GetSingletonDatabse()
	pop := db.GetPopulation("Kyiv")

	fmt.Println("Pop of Kyiv =", pop)
}
