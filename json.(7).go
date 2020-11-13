package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	type FruitBasket struct {
		Name    string
		Fruit   []string
		Id      int64 `json:"ref"`
		private string
		Created time.Time
	}

	basket := FruitBasket{
		Name:    "Standard",
		Fruit:   []string{"Apple", "Banana", "Orange"},
		Id:      999,
		private: "Second-rate",
		Created: time.Now(),
	}

	var jsonData []byte
	jsonData, err := json.Marshal(basket)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(jsonData))
}
