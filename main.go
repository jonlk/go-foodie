package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jonlk/go-foodie/food"
)

const (
	API_KEY = `<your api key goes here>`
)

func main() {

	//TODO: pass in fdcid arg
	url := fmt.Sprintf(`https://api.nal.usda.gov/fdc/v1/food/334147?api_key=%s`, API_KEY)

	req, _ := http.NewRequest("GET", url, nil)

	client := &http.Client{}

	if resp, err := client.Do(req); err == nil {

		defer resp.Body.Close()

		if body, err := io.ReadAll(resp.Body); err == nil {

			food := food.Food{}

			json.Unmarshal([]byte(body), &food)

			fmt.Println(`Results For`, food.Description)

			for _, f := range food.FoodNutrients {
				if f.Amount > 0 {
					fmt.Println(f.Nutrient.Name, f.Amount, f.Nutrient.UnitName)
				}
			}

		}
	}
}
