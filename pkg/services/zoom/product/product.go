package zoom

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
    "net/url"
    "strings"

    "github.com/PuerkitoBio/goquery"
)

// --- Models ---

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

type SuggestionRequest struct {
    Pagination               PaginationRequest `json:"pagination"`
    Query                    string            `json:"query"`
    Order                    string            `json:"order"`
    DisableRefinementsStats  bool              `json:"disableRefinementsStats"`
    EnableCategorySuggestion bool              `json:"enableCategorySuggestion"`
    Refinements              []interface{}     `json:"refinements"`
}

type PaginationRequest struct {
    HitsPerPage int `json:"hitsPerPage"`
    Page        int `json:"page"`
}

type SuggestionResponse struct {
    Query       string             `json:"query"`
    Index       string             `json:"index"`
    Pagination  PaginationResponse `json:"pagination"`
    Hits        []Hit              `json:"hits"`
    IsIndexable bool               `json:"isIndexable"`
}

type PaginationResponse struct {
    Page        int `json:"page"`
    HitsPerPage int `json:"hitsPerPage"`
    NbHits      int `json:"nbHits"`
    NbPages     int `json:"nbPages"`
}

type Hit struct {
    ObjectID           string             `json:"objectID"`
    Position           int                `json:"position"`
    Query              string             `json:"query"`
    CategorySuggestion CategorySuggestion `json:"categorySuggestion"`
}

type CategorySuggestion struct {
    BestOption Option   `json:"bestOption"`
    AllOptions []Option `json:"allOptions"`
}

type Option struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

// --- API Functions ---

func SearchByName(name string, client *http.Client) (*Suggestions, error) {
    url := fmt.Sprintf("https://api.zoom.com/v2/products/search?name=%s", name)
    resp, err := client.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    log.Printf("Resposta bruta:\n%s", string(body))

    var suggestions Suggestions
    if err := json.Unmarshal(body, &suggestions); err != nil {
        return nil, err
    }

    return &suggestions, nil
}

func SuggestionsOfProducts(product string, client *http.Client) (*SuggestionResponse, error) {
    baseURL := "https://api-v1.zoom.com.br/search/suggestions"

    log.Println("product" + product)
    requestBody := SuggestionRequest{
        Pagination: PaginationRequest{
            HitsPerPage: 5,
            Page:        0,
        },
        Query:                    product,
        Order:                    "default",
        DisableRefinementsStats:  true,
        EnableCategorySuggestion: true,
        Refinements:              []interface{}{},
    }

    bodyBytes, err := json.Marshal(requestBody)
    if err != nil {
        return nil, fmt.Errorf("erro ao serializar corpo da requisição: %w", err)
    }

    req, err := http.NewRequest("POST", baseURL, strings.NewReader(string(bodyBytes)))
    if err != nil {
        return nil, fmt.Errorf("erro ao criar requisição: %w", err)
    }
    req.Header.Set("Content-Type", "application/json")

    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("erro ao fazer requisição: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("erro ao ler corpo da resposta: %w", err)
    }

    log.Println("Resposta da API:")
    var pretty bytes.Buffer
    if err := json.Indent(&pretty, body, "", "  "); err == nil {
        log.Println(pretty.String())
    } else {
        log.Println(string(body))
    }

    var suggestionResponse SuggestionResponse
    if err := json.Unmarshal(body, &suggestionResponse); err != nil {
        return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
    }

    return &suggestionResponse, nil
}

// --- Web Scraping Functions ---

func ExtractProductIDsFromSearchPage(pageURL string) ([]string, error) {
    client := &http.Client{}
    req, err := http.NewRequest("GET", pageURL, nil)
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

