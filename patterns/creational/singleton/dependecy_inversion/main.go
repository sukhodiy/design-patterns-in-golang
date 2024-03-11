package main

import (
	"fmt"
	"sync"
)

type Database interface {
	GetPopulation(name string) int
}

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

	result := map[string]int{
		"Tokyo": 5000000,
		"Kyiv":  3000000,
	}

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

func GetTotalPopulationEx(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	}
	return result
}

// for testing reasons
type DummyDatabse struct {
	dummyData map[string]int
}

func (d *DummyDatabse) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3,
		}
	}

	return d.dummyData[name]
}

func main() {
	names := []string{"alpha", "gamma"}
	tp := GetTotalPopulationEx(&DummyDatabse{}, names)
	fmt.Println(tp == 4)
}
