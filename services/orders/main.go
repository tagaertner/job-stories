package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/tagaertner/e-commerce/services/orders/generated"
	"github.com/99designs/gqlgen/graphql/handler/extension" 
	"github.com/tagaertner/e-commerce/services/orders/resolvers"
	"github.com/99designs/gqlgen/graphql/handler/transport"
)

const defaultPort = "4003"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	resolver := resolvers.NewResolver()

	srv := handler.New(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
	}))

	// Introspection 
	if os.Getenv("ENVIRONMENT") != "production" { 
		srv.Use(extension.Introspection{})
	}

	// Add transports
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.Websocket{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	// Health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status": "healthy", "service": "orders"}`))
	})

	log.Printf("🛍️ Orders service ready at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
