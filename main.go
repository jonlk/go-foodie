package main

import (
	"fmt"

	"github.com/jonlk/go-foodie/food"
)

func main() {

	foodResult := food.GetFoodResult()

	fmt.Println(`Results For`, foodResult.Description)

	// foodDisplayModels := []food.FoodDisplayModel{}
	for _, f := range foodResult.FoodNutrients {
		if f.Amount > 0 {
			//foodDisplayModel := food.FoodDisplayModel{}
			//foodDisplayModels = append(foodDisplayModels, foodDisplayModel)
			fmt.Println(f.Nutrient.Name, f.Amount, f.Nutrient.UnitName)
		}
	}
}
