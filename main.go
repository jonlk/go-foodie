package main

import (
	"fmt"

	"github.com/jonlk/go-foodie/food"
)

func main() {

	foodDisplayModel := food.GetFoodDisplayModel(334147)

	fmt.Println(`Results for: `, foodDisplayModel.Description)

	for _, f := range foodDisplayModel.FoodDisplayDetails {
		fmt.Println(f.Name, f.Amount, f.UnitName)
	}
}
