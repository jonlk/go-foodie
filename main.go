package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jonlk/go-foodie/food"
)

func main() {
	//334147

	if len(os.Args) == 0 || len(os.Args) > 2 {
		fmt.Println("USAGE: Only one argument as a numeric FDC ID")
		os.Exit(1)
	}

	if search, err := strconv.Atoi(os.Args[1]); err == nil {

		foodDisplayModel := food.GetFoodDisplayModel(search)

		if foodDisplayModel.Description == "" {
			fmt.Println("Nothing found")
		} else {

			fmt.Println(`Results for: `, foodDisplayModel.Description)

			for _, f := range foodDisplayModel.FoodDisplayDetails {
				fmt.Println(f.Name, f.Amount, f.UnitName)
			}
		}

	} else {
		fmt.Println("USAGE: Argument was not a numeric FDC ID")
	}
}
