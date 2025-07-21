package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"price-comparator-api/internal/searchengineer/domain"
)

type SerpAPIRepository struct{}

func NewSerpAPIRepository() *SerpAPIRepository {
	return &SerpAPIRepository{}
}

type SerpAPIResponse struct {
	ShoppingResults []domain.ProductPrice `json:"shopping_results"`
}

func (r *SerpAPIRepository) FindPrices(productName string) ([]domain.ProductPrice, error) {
	apiKey := os.Getenv("SERPAPI_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("SERPAPI_KEY environment variable not set")
	}

	url := fmt.Sprintf("https://serpapi.com/search.json?engine=google_shopping&q=%s&api_key=%s", productName, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data from SerpApi: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var serpApiResponse SerpAPIResponse
	err = json.Unmarshal(body, &serpApiResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return serpApiResponse.ShoppingResults, nil
}
