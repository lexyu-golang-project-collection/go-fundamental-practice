package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
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
	panic("")
}

func main() {

	cart := LoadCart(1)
	fmt.Printf("%+v\n", cart)

	product := LoadProduct(1)
	fmt.Printf("%+v", product)

	c, ps := LoadCartAndProductsSequential(1)

	slog.Info("cart", "c1", c)

	slog.Info("products", "p1", ps)
}
