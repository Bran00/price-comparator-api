package price

import (
	"fmt"
	"io"
	"net/http"
)

func FindHistoryPrice(httpClient *http.Client, productID string) ([]byte, error) {
	url := fmt.Sprint("Here the link to api")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
} 
