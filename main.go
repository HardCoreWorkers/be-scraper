package main

import (
	"encoding/json"
	"fmt"
	"github.com/HardCoreWorkers/be-scraper/recipes"
)

func main() {
	recipes, err := recipes.GetAllRecipes()

	if err != nil {
		panic(err)
	}

	result, err := json.Marshal(recipes)
	fmt.Printf("%s\n", result)
}
