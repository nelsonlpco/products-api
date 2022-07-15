package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nelsonlpco/products-api/src/adapter/postgres"
	"github.com/nelsonlpco/products-api/src/di"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()

	postgres.RunMigrations()
	productService := di.ConfigureProductDI(conn)

	router := mux.NewRouter()
	router.Handle("/product", http.HandlerFunc(productService.Create)).Methods(http.MethodPost)
	router.Handle("/product", http.HandlerFunc(productService.Fetch)).
		Queries(
			"page", "{page}",
			"itemsPerPage", "{itemsPerPage}",
			"descending", "{descending}",
			"sort", "{sort}",
			"search", "{search}",
		).Methods("GET")

	port := viper.GetString("server.port")
	log.Printf("LISTEN ON PORT: %v", port)

	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
