package product

import (
	"encoding/json"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Suggestions struct {
	Products []Product `json:"products"`
}

func SearchByName(name string) (*Suggestions, error) {
	url := fmt.Sprintf("https://api.zoom.com/v2/products/search?name=%s", name)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var suggestions Suggestions
	err = json.Unmarshal(body, &suggestions)
	if err != nil {
		return nil, err
	}

	return &suggestions, nil
}
