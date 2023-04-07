package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mpragliola/pipe/pkg/pipe"
)

type Product struct {
	Id    int     `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

type ProductList struct {
	Products []Product `json:"products"`
	Total    int       `json:"total"`
	Skip     int       `json:"skip"`
	Limit    int       `json:"limit"`
}

// We will set up a pipeline to process the data of some example API con the web (dummyjson.com).
// We will read a list of products that is paginated, extract the products from each page,
// flatten the products into a single stream, then we will get the sum of all the prices that are
// valued at least 30 or more.
func main() {
	ctx, cancel := context.WithCancel(context.Background())

	i := 0

	// Beginning of pipeline - here we fetch the API pages with the products
	// into various ProductList structs
	webStream := pipe.OfFunc(
		ctx,
		func() ProductList {
			fmt.Printf("Fetching page %d\n", i+1)

			pl := fetchPage(i)
			if pl.Total < (i+1)*30 {
				cancel()
			}
			i++
			return pl
		},
	)

	// Pagination means that the products come in a field and in separate arrays, one
	// for each page; we extract that field and we spread the arrays in a continuous stream
	getProduct := func(pl ProductList) []Product { return pl.Products }
	products := pipe.Spread(pipe.P(getProduct, webStream))

	// The products are now all streaming on the same level, we can map to a stream of prices
	getPrice := func(p Product) int { return int(p.Price) }
	prices := pipe.P(getPrice, products)

	// We remove the prices that are less than 30
	notTooCheapFilter := func(price int) bool { return price >= 30 }
	poshPrices := pipe.Filter(notTooCheapFilter, prices)

	// And we sum them
	reduceSum := func(a, b int) int { return a + b }
	sum := pipe.Reduce(reduceSum, 0, poshPrices)

	// this will actually consume the data
	all := pipe.Scoop(sum)

	total := all[len(all)-1]

	fmt.Println(total)
}

func fetchPage(page int) (pl ProductList) {
	resp, err := http.Get("https://dummyjson.com/products?skip=" + fmt.Sprint(30*page))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	pl = ProductList{}
	err = json.Unmarshal(body, &pl)
	if err != nil {
		log.Fatalln(err)
	}

	return
}
