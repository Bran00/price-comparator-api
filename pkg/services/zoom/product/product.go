package product

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Product representa a estrutura de um produto retornado pela API do Zoom.
type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Suggestions representa uma lista de produtos sugeridos pela API do Zoom.
type Suggestions struct {
	Products []Product `json:"products"`
}

// PriceHistory representa o histórico de preços de um produto.
type PriceHistory struct {
	ProductID string    `json:"product_id"`
	History   []float64 `json:"history"`
}

// SearchByName busca produtos por nome na API do Zoom.
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

// GetProductIDByName busca um produto pelo nome e retorna o ID do primeiro produto encontrado.
// Esta função reutiliza a função SearchByName e extrai o ID do primeiro produto da lista de sugestões.
func GetProductIDByName(name string, http *http.Client) (string, error) {
	suggestions, err := SearchByName(name, http)
	if err != nil {
		return "", err
	}

	if len(suggestions.Products) == 0 {
		return "", fmt.Errorf("nenhum produto encontrado com o nome: %s", name)
	}

	return suggestions.Products[0].ID, nil
}

// GetPriceHistoryByID busca o histórico de preços de um produto pelo seu ID.
// Esta função faz uma chamada para um endpoint hipotético da API do Zoom para obter o histórico de preços.
func GetPriceHistoryByID(id string, http *http.Client) (*PriceHistory, error) {
	url := fmt.Sprintf("https://api.zoom.com/v2/products/%s/history", id)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var history PriceHistory
	err = json.Unmarshal(body, &history)
	if err != nil {
		return nil, err
	}

	return &history, nil
}
