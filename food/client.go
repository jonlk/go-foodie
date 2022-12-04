package food

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GetFoodResult() Food {
	url := fmt.Sprintf(`https://api.nal.usda.gov/fdc/v1/food/334147?api_key=%s`, os.Getenv("USDA_API_KEY"))

	foodresult := Food{}

	if req, err := http.NewRequest("GET", url, nil); err == nil {

		client := &http.Client{}

		if resp, err := client.Do(req); err == nil {

			defer resp.Body.Close()

			if body, err := io.ReadAll(resp.Body); err == nil {
				json.Unmarshal([]byte(body), &foodresult)
			}
		}
	} else {
		println(err.Error())
		panic("Failed")
	}

	return foodresult
}
