package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

var (
	database = make(map[string]Product)
)

func SetJSONResp(res http.ResponseWriter, message []byte, httpCode int) {
	res.Header().Set("Content-type", "application/json")
	res.WriteHeader(httpCode)
	res.Write(message)
}

func main() {

	//init db
	database["001"] = Product{ID: "001", Name: "Realme Pro 2", Quantity: 10}
	database["002"] = Product{ID: "002", Name: "Iphone 7", Quantity: 5}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		message := []byte(`{"message": "ready to go"}`)
		SetJSONResp(res, message, http.StatusOK)
	})

	http.HandleFunc("/get-products", func(res http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			message := []byte(`{"message": "Invalid http method"}`)
			SetJSONResp(res, message, http.StatusMethodNotAllowed)
			return
		}

		var products []Product

		for _, product := range database {
			products = append(products, product)
		}

		productJSON, err := json.Marshal(&products)
		if err != nil {
			message := []byte(`{"message": "Error when parsing data"}`)
			SetJSONResp(res, message, http.StatusInternalServerError)
			return
		}

		SetJSONResp(res, productJSON, http.StatusOK)
	})

	err := http.ListenAndServe("localhost:9000", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
