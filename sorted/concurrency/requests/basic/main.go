package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

type CartProduct struct {
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}

type Cart struct {
	Id       int           `json:"id"`
	UserId   int           `json:"userId"`
	Date     string        `json:"date"`
	Products []CartProduct `json:"products"`
}

type ProductRating struct {
	Rate  float32 `json:"rate"`
	Count int     `json:"count"`
}

type Product struct {
	Id          int           `json:"id"`
	Price       float32       `json:"price"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Category    string        `json:"category"`
	Image       string        `json:"image"`
	Rating      ProductRating `json:"rating"`
}

func LoadCart(cartId int) Cart {
	url := fmt.Sprintf("https://fakestoreapi.com/carts/%d", cartId)

	resp, err := http.Get(url)
	if err != nil {
		panic(err) // @todo log.Fatalf("Error loading Cart URL "%v", error: %v", url, err)
	}
	defer resp.Body.Close()

	var parsedResp Cart
	decodeErr := json.NewDecoder(resp.Body).Decode(&parsedResp)
	if decodeErr != nil {
		panic(decodeErr) // @todo log.Fatalf("Error decoding LoadCart response %v", err)
	}

	return parsedResp
}

func LoadProduct(id int) Product {
	url := fmt.Sprintf("https://fakestoreapi.com/products/%d", id)

	resp, err := http.Get(url)
	if err != nil {
		panic(err) // @todo log.Fatalf("Error loading Product URL "%w", error: %w", url, err)
	}
	defer resp.Body.Close()

	var parseResp Product
	decodeErr := json.NewDecoder(resp.Body).Decode(&parseResp)
	if decodeErr != nil {
		panic(err) // @todo log.Fatalf("Error decoding LoadProduct response %w", err)
	}
	return parseResp
}

// Naive implementation of loading of products, it is done in a blocking for loop.
func LoadCartAndProductsSequential(cartId int) (Cart, []Product) {
	start := time.Now()

	cartResp := LoadCart(cartId)

	productResponses := make([]Product, 0, len(cartResp.Products))

	for _, product := range cartResp.Products {
		productResp := LoadProduct(product.ProductId)
		productResponses = append(productResponses, productResp)
	}
	end := time.Now()
	duration := end.Sub(start)
	slog.Info("LoadCartAndProductsSequential runtime",
		"duration", duration,
		"cartId", cartResp.Id,
		"len(products)", len(productResponses))

	return cartResp, productResponses
}

// Load products in parallel, since we know the number of calls, we read from the channel that number of times.
func LoadCartAndProductsExhaustChannel(cartId int) (Cart, []Product) {
	start := time.Now()

	cartResp := LoadCart(cartId)

	productRespsesCh := make(chan Product, len(cartResp.Products))

	for _, product := range cartResp.Products {
		go func(product CartProduct) {
			productResp := LoadProduct(product.ProductId)
			productRespsesCh <- productResp
		}(product)
	}

	productResps := make([]Product, 0, len(cartResp.Products))

	for i := 0; i < len(cartResp.Products); i++ {
		comm := <-productRespsesCh
		productResps = append(productResps, comm)
	}

	end := time.Now()
	duration := end.Sub(start)

	slog.Info(
		"LoadCartAndProductsExhaustChannel runtime",
		"duration", duration,
		"cardId", cartResp.Id,
		"len(products)", len(productResps),
	)

	return cartResp, productResps
}

func LoadCartAndProductsWaitGroup(cartId int) (Cart, []Product) {
	start := time.Now()

	cartResp := LoadCart(cartId)

	var wg sync.WaitGroup
	productRespsesCh := make(chan Product, len(cartResp.Products))

	for _, product := range cartResp.Products {
		wg.Add(1)
		go func(product CartProduct) {
			defer wg.Done()
			productResp := LoadProduct(product.ProductId)
			productRespsesCh <- productResp
		}(product)
	}

	wg.Wait()
	close(productRespsesCh)

	productResps := make([]Product, 0, len(cartResp.Products))

	for chValue := range productRespsesCh {
		productResps = append(productResps, chValue)
	}

	end := time.Now()
	duration := end.Sub(start)

	slog.Info("LoadCartAndProductsWaitGroup runtime",
		"duration", duration,
		"cartId", cartResp.Id,
		"len(products)", len(productResps),
	)

	return cartResp, productResps
}

func main() {

	// cart := LoadCart(1)
	// slog.Info("cart", "cart", cart)

	// product := LoadProduct(1)
	// slog.Info("product", "product", product)

	LoadCartAndProductsSequential(1)
	// c, ps := LoadCartAndProductsSequential(1)
	// slog.Info("cart", "c1", c)
	// slog.Info("products", "p1", ps)

	LoadCartAndProductsExhaustChannel(1)
	// c, ps = LoadCartAndProductsExhaustChannel(1)
	// slog.Info("cart", "c1", c)
	// slog.Info("products", "p1", ps)

	LoadCartAndProductsWaitGroup(1)
}
