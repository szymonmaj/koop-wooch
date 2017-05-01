package main

import (
	"net/http"
	"os"
	"strconv"
	"fmt"
)
type Product struct {
	Name string
	Price float64

}
var products = []Product{}


func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		write(w, " <a href=\"/product_form\">Add product</a>")
		write(w, " <a href=\"/products\">Show products</a>")

	})
	http.HandleFunc("/add_product", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		name := r.URL.Query().Get("name")
		price, _ := strconv.ParseFloat(r.URL.Query().Get("price"), 64)

		p := Product{Name: name, Price: price}

		products = append(products, p)
	})
	http.HandleFunc("/product_form", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		write(w, "<form action='add_product'><input name='name'>+<input name='price'><input type='submit' value='Add'></form>")
	})
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		write(w, "<table>")
		for _, product := range products {

			write(w, fmt.Sprintf("<tr><td>%v</td><td>%v</td></tr>", product.Name, product.Price))

		}
		write(w, "</table>")
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "1234"
	}

	http.ListenAndServe("0.0.0.0:"+port, nil)
}
func write(w http.ResponseWriter, text string) {
	w.Write([]byte(text))
}