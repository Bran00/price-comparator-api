
# Price Comparator API

This is a simple price comparator API built with Go.

## Features

- Compares prices from Google Shopping.
- Built with a Hexagonal Architecture.

## Endpoints

### `GET /compare`

Compares prices for a given product.

**Query Parameters:**

- `product` (string, required): The name of the product to compare.

**Example:**

```bash
curl "http://localhost:8080/compare?product=playstation+5"
```

**Success Response:**

```json
[
    {
        "source": "PlayStation",
        "price": "$499.99",
        "link": "https://direct.playstation.com/en-us/consoles/console/playstation5-console.3006634",
        "description": "PlayStation 5 Console"
    }
]
```
