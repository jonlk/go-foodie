package food

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getFoodResult(fdcId int) Food {

	url := fmt.Sprintf(`https://api.nal.usda.gov/fdc/v1/food/%v?api_key=%s`,
		fdcId,
		os.Getenv("USDA_API_KEY"))

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
