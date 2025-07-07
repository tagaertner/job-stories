package main

import (
    "log"
    "net/http"
    "os"

    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
    "e-commerce/services/products/generated"
    "e-commerce/services/products/resolvers"
    "github.com/99designs/gqlgen/graphql/handler/transport" 
    "github.com/99designs/gqlgen/graphql/handler/extension"
    // Remove this import - you don't need models anymore
    // "e-commerce/services/products/models"
)

const defaultPort = "4001"

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = defaultPort
    }

    // NEW WAY - Use the service pattern
    resolver := resolvers.NewResolver()  // ✅ This creates ProductService with data
    
    srv := handler.New(generated.NewExecutableSchema(generated.Config{
        Resolvers: resolver,
    }))


    // Enable introspection in non-production environments
    if os.Getenv("ENVIRONMENT") != "production" {
        srv.Use(extension.Introspection{})
    }

    // Add Post transport
    srv.AddTransport(transport.POST{})
    srv.AddTransport(transport.GET{})    
    srv.AddTransport(transport.Websocket{}) 

    http.Handle("/", playground.Handler("GraphQL playground", "/query"))
    http.Handle("/query", srv)

    // Health check
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"status": "healthy", "service": "products"}`))
    })

    log.Printf("🛍️ Products service ready at http://localhost:%s/", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}

