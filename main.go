package main

import (
	"fmt"

	"github.com/jonlk/go-foodie/food"
)

func main() {

	foodDisplayModels := food.GetFoodDisplayModels(334147)

	for _, f := range foodDisplayModels {
		fmt.Println(f.Name, f.Amount, f.UnitName)
	}

}
