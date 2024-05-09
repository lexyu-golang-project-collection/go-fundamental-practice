package main

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
	panic("")
}

func LoadProduct(id int) Product {
	panic("")
}

func LoadCartAndProductsSequential(cartId int) (Cart, []Product) {
	panic("")
}

func LoadCartAndProductsExhaustChannel(cartId int) (Cart, []Product) {
	panic("")
}

func main() {

}
