package main

import (
	"net/http"
	"os"
	"strconv"
	"fmt"
)

type Product struct {
	Name     string
	Category string
	Price    float64
}

var products = []Product{}


func main() {

	addExampleData()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		write(w, " <a href=\"/product_form\">Add product</a>")
		write(w, " <a href=\"/products\">Show products</a>")

	})

	http.HandleFunc("/add_product", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		name := r.URL.Query().Get("name")
		category := r.URL.Query().Get("category")
		price, _ := strconv.ParseFloat(r.URL.Query().Get("price"), 64)


		p := Product{Name: name, Category: category, Price: price}

		products = append(products, p)
		http.Redirect(w, r, "/", 303)

	})

	http.HandleFunc("/product_form", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		write(w, "<form action='add_product'>PRODUCT<input name='name'>CATEGORY<input name='category'>PRICE<input name='price'><input type='submit' value='Add'></form>")
	})

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		write(w, "<table>")
		for _, product := range products {

			write(w, fmt.Sprintf("<tr><td>%v</td><td>%v</td><td>%v</td><td><form action='Put_in'><input type='hidden' name='name' value='%v'><input type='submit' value='Put'></form></td></tr>", product.Name,product.Category, product.Price, product.Name))

		}
		write(w, "</table>")

	})
	http.HandleFunc("/Put_in", func(w http.ResponseWriter, r *http.Request) {

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

func addExampleData() {
	products = append(products, Product{"Carrot", "Vegetables", 123})
	products = append(products, Product{"Apple", "Fruits", 666})
}