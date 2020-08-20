package recipes

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"math"
	"strconv"
)

type Recipe struct {
	Name        string
	PrepTime    string
	CookTime    string
	ImageURL    string
	Ingredients []string
}

var recipes []Recipe

func GetAll() ([]Recipe, error) {
	url := "https://www.bbc.co.uk/food/recipes/a-z/a/"
	doc, err := goquery.NewDocument(url)

	if err != nil {
		panic(err)
	}

	pagination := doc.Find(".pagination-summary.gel-wrap b.gel-pica-bold").Text()

	numRecipes, err := strconv.ParseFloat(pagination, 10)
	numPages := math.Ceil(numRecipes / 24)
	numPg := int(numPages)

	for i := 1; i < numPg+1; i++ {
		pageNum := strconv.Itoa(i)
		fullURL := (url + pageNum)
		GetRecipes(fullURL)
	}

	return recipes, nil
}

func GetRecipes(url string) {
	doc, err := goquery.NewDocument(url)

	if err != nil {
		panic(err)
	}

	doc.Find(".promo").Each(func(i int, s *goquery.Selection) {
		link, ok := s.Attr("href")
		if ok {
			GetRecipe(link)
		}
	})
}

func GetRecipe(url string) {
	doc, err := goquery.NewDocument("https://www.bbc.co.uk" + url)

	if err != nil {
		panic(err)
	}

	name := doc.Find("h1.gel-trafalgar.content-title__text").Text()

	prepTime := doc.Find("p.recipe-metadata__prep-time").Text()

	cookTime := doc.Find("p.recipe-metadata__cook-time").Text()

	imageURL, ok := doc.Find(".recipe-media__image img").Attr("src")

	if ok {
		fmt.Println(imageURL)
	}

	ingredients := []string{}
	doc.Find("ul.recipe-ingredients__list").Each(func(i int, s *goquery.Selection) {
		s.Find("a.recipe-ingredients__link").Each(func(i int, r *goquery.Selection) {
			ingredients = append(ingredients, (r.Text()))
		})
	})

	r := Recipe{
		Name:        name,
		PrepTime:    prepTime,
		CookTime:    cookTime,
		ImageURL:    imageURL,
		Ingredients: ingredients,
	}

	recipes = append(recipes, r)
}
