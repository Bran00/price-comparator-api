package product

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Suggestions struct {
	Products []Product `json:"products"`
}

type PriceHistory struct {
	ProductID string    `json:"product_id"`
	History   []float64 `json:"history"`
}

func SearchByName(name string, http *http.Client) (*Suggestions, error) {
	url := fmt.Sprintf("https://api.zoom.com/v2/products/search?name=%s", name)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
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

func ExtractProductIDsFromSearchPage(url string) ([]string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var productIDs []string
	doc.Find("button[data-testid='save-product']").Each(func(i int, s *goquery.Selection) {
		if id, exists := s.Attr("id"); exists && strings.HasPrefix(id, "save-product-") {
			productID := strings.TrimPrefix(id, "save-product-")
			productIDs = append(productIDs, productID)
		}
	})

	return productIDs, nil
}

func SearchProductIDsByTerm(term string) ([]string, error) {
	baseURL := "https://www.zoom.com.br/search"
	query := url.Values{}
	query.Set("q", term)

	searchURL := fmt.Sprintf("%s?%s", baseURL, query.Encode())
	return ExtractProductIDsFromSearchPage(searchURL)
}
