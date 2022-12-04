package food

import "sort"

func GetFoodDisplayModels(fdcId int) []FoodDisplayModel {

	foodResult := getFoodResult(fdcId)
	foodDisplayModels := []FoodDisplayModel{}

	for _, f := range foodResult.FoodNutrients {

		if f.Amount > 0 {

			foodDisplayModel := FoodDisplayModel{
				Name:     f.Nutrient.Name,
				Amount:   f.Amount,
				UnitName: f.Nutrient.UnitName,
			}

			foodDisplayModels = append(foodDisplayModels, foodDisplayModel)
		}
	}

	//we want to show highest to lowest grams of ingredients
	sort.Slice(foodDisplayModels,
		func(i, j int) bool {
			return foodDisplayModels[j].Amount < foodDisplayModels[i].Amount
		})

	return foodDisplayModels
}
