package main

import (
    "log"
    "net/http"
    "os"

    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
    "e-commerce/services/users/generated"
    "e-commerce/services/users/resolvers"
    "github.com/99designs/gqlgen/graphql/handler/transport" 
    "github.com/99designs/gqlgen/graphql/handler/extension"
    
)

const defaultPort = "4002"

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = defaultPort
    }

    resolver := resolvers.NewResolver()
    
    srv := handler.New(generated.NewExecutableSchema(generated.Config{
        Resolvers: resolver,
    }))

    // Enable introspection
    if os.Getenv("ENVIRONMENT") != "production" { 
		srv.Use(extension.Introspection{})
	}


    // Post transport
    srv.AddTransport((transport.POST{}))
    srv.AddTransport(transport.GET{})   
    srv.AddTransport(transport.Websocket{}) 

    http.Handle("/", playground.Handler("GraphQL playground", "/query"))
    http.Handle("/query", srv)

    // Health check
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"status": "healthy", "service": "users"}`))
    })

    log.Printf("🛍️ Users service ready at http://localhost:%s/", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}