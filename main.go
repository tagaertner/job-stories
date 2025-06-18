package main

import (
	"log"
    "net/http"
    "os"

	"github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
    "e-commerce/services/products/generated"
    "e-commerce/services/products/resolvers"
)

const defaultPort = "4001"

func main() {
	port := os.Getenv("PORT")
	if port == ""{
		port = defaultPort
	}

	// Create resolver
	resolver := resolvers.NewResolver()

	// Create GraphQL server using our generated schema and resolver
	srv := handler.New(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
	}))

	// Set up HTTP routes
	http.Handle("/", playground.Handler("Products Service", "/query"))
	http.Handle("/query", srv)

	log.Printf("🛍️ Products service ready at http://localhost:%s/", port)
    log.Printf("📊 GraphQL playground at http://localhost:%s/", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
