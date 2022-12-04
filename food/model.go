package food

import "sort"

func GetFoodDisplayModel(fdcId int) FoodDisplayModel {

	foodResult := getFoodResult(fdcId)

	foodDisplayModel := FoodDisplayModel{
		Description: foodResult.Description,
	}

	for _, f := range foodResult.FoodNutrients {

		if f.Amount > 0 {

			foodDisplayDetail := FoodDisplayDetail{
				Name:     f.Nutrient.Name,
				Amount:   f.Amount,
				UnitName: f.Nutrient.UnitName,
			}

			foodDisplayModel.FoodDisplayDetails =
				append(foodDisplayModel.FoodDisplayDetails,
					foodDisplayDetail)
		}
	}

	//we want to show highest to lowest grams of ingredients
	sort.Slice(foodDisplayModel.FoodDisplayDetails,
		func(i, j int) bool {
			return foodDisplayModel.FoodDisplayDetails[j].Amount < foodDisplayModel.FoodDisplayDetails[i].Amount
		})

	return foodDisplayModel
}
