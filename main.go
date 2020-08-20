package main

import (
	"be-scraper/recipes"
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Hello!")
	recipes, err := recipes.GetAll()

	if err != nil {
		panic(err)
	}

	result, err := json.Marshal(recipes)
	fmt.Printf("%s\n", result)
}
