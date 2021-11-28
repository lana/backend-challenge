package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

type payload struct {
	BasketID    string `json:"basket_id" binding:"required"`
	ProductCode string `json:"product_code" binding:"required"`
}

type Basket struct {
	Code  string
	Items map[string]Item
	Total float64
	Close bool
}

type Product struct {
	Code  string
	Name  string
	Price float64
}

type Item struct {
	Product  Product
	Quantity int
	Total    float64
}

func lanaCmd() *cobra.Command { // nolint:funlen
	basket := &cobra.Command{
		Use:   "basket",
		Short: "call different operations",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	createBasket := &cobra.Command{
		Use:   "create",
		Short: "create a new basket",
		Args:  cobra.ExactValidArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/baskets", nil)
			if err != nil {
				log.Panic("error building a http client")
			}

			response, err := http.DefaultClient.Do(request)
			if err != nil {
				log.Panic("error building a http client")
			}

			defer response.Body.Close()

			body, err := io.ReadAll(response.Body)
			if err != nil {
				log.Panic("error building a http client")
			}

			fmt.Println("Basket created")
			fmt.Println(string(body))
		},
	}

	removeBasket := &cobra.Command{
		Use:   "remove",
		Short: "remove a new basket",
		Args:  cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			basketID := args[0]
			if basketID == "" {
				log.Panic("basketID is required")
			}

			url := fmt.Sprintf("http://localhost:8080/baskets/%s", basketID)
			request, err := http.NewRequest(http.MethodDelete, url, nil)
			if err != nil {
				log.Panic("error building a http client")
			}

			response, err := http.DefaultClient.Do(request)
			if err != nil {
				log.Panic("error building a http client")
			}

			defer response.Body.Close()
			bodyResp, err := io.ReadAll(response.Body)

			if response.StatusCode == http.StatusBadRequest {
				fmt.Println(string(bodyResp))
				return
			}

			fmt.Println("basket ID deleted")
		},
	}

	addProductToBasket := &cobra.Command{
		Use:     "add",
		Short:   "add a new product to basket",
		Example: "basket add basket_id product_code ",
		Args:    cobra.ExactValidArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			basketID := args[0]
			productID := args[1]

			if basketID == "" || productID == "" {
				log.Panic("basket ID/product ID is required")
			}

			url := fmt.Sprintf("http://localhost:8080/baskets/%s/products/%s", basketID, productID)
			request, err := http.NewRequest(http.MethodPost, url, nil)
			if err != nil {
				log.Panic("error building a http client")
			}

			response, err := http.DefaultClient.Do(request)
			if err != nil {
				log.Panic("error building a http client")
			}

			defer response.Body.Close()
			bodyResp, err := io.ReadAll(response.Body)

			if response.StatusCode == http.StatusBadRequest {
				fmt.Println(string(bodyResp))
				return
			}

			fmt.Println("product added")
		},
	}

	checkoutBasket := &cobra.Command{
		Use:     "checkout",
		Short:   "close a basket and return total amount",
		Example: "basket checkout basket_id",
		Args:    cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			basketID := args[0]
			if basketID == "" {
				log.Panic("basket ID is required")
			}

			url := fmt.Sprintf("http://localhost:8080/baskets/%s/checkout", basketID)
			request, err := http.NewRequest(http.MethodPost, url, nil)
			if err != nil {
				log.Panic("error building a http client")
			}

			response, err := http.DefaultClient.Do(request)
			if err != nil {
				log.Panic("error building a http client")
			}

			defer response.Body.Close()
			var _basket Basket
			err = json.NewDecoder(response.Body).Decode(&_basket)
			if err != nil {
				log.Panic("error decode response")
			}

			fmt.Printf("Basket ID: %s\n", _basket.Code)
			fmt.Println("Items:")
			for _, item := range _basket.Items {
				fmt.Printf("      Item: %s\n", item.Product.Code)
				fmt.Printf("      Quantity: %v      Unit price: %v\n", item.Quantity, item.Product.Price)
				fmt.Printf("      Total With Discount:         %v\n", item.Total)
			}
			fmt.Println("----------------------------------------")
			fmt.Printf("Amount Total: %v\n", _basket.Total)
		},
	}

	basket.AddCommand(createBasket, removeBasket, addProductToBasket, checkoutBasket)

	return basket
}
