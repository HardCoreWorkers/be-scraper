package main

import (
	"encoding/json"
	"fmt"
	"web-scraper/recipes"
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
