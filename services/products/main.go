package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strings"
    "e-commerce/services/products/resolvers"
)

func main() {
    resolver := resolvers.NewResolver()
    
    // GraphQL endpoint that Apollo Gateway expects
    http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        
        if r.Method == "POST" {
            // Handle GraphQL introspection query from Apollo Gateway
            var requestBody map[string]interface{}
            json.NewDecoder(r.Body).Decode(&requestBody)
            
            query, _ := requestBody["query"].(string)
            
            // Handle SDL query (what Apollo Gateway needs)
            if strings.Contains(query, "_service") {
                response := map[string]interface{}{
                    "data": map[string]interface{}{
                        "_service": map[string]interface{}{
                            "sdl": `
                                directive @key(fields: String!) on OBJECT | INTERFACE
                                
                                type Product @key(fields: "id") {
                                    id: ID!
                                    name: String!
                                    price: Float!
                                    description: String
                                    inventory: Int!
                                }
                                
                                extend type Query {
                                    products: [Product!]!
                                    product(id: ID!): Product
                                }
                            `,
                        },
                    },
                }
                json.NewEncoder(w).Encode(response)
                return
            }
            
            // Handle products query
            if strings.Contains(query, "products") {
                products, _ := resolver.Products(r.Context())
                response := map[string]interface{}{
                    "data": map[string]interface{}{
                        "products": products,
                    },
                }
                json.NewEncoder(w).Encode(response)
                return
            }
        }
        
        // Default response
        response := map[string]interface{}{
            "data": map[string]interface{}{
                "message": "Products GraphQL service ready",
            },
        }
        json.NewEncoder(w).Encode(response)
    })
    
    // Health check
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "status": "healthy",
            "service": "products",
        })
    })
    
    log.Println("🛍️ Products service (Federation-ready) at http://localhost:4001/")
    log.Fatal(http.ListenAndServe(":4001", nil))
}
