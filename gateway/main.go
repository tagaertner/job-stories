package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// GraphQL request struct

type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

// GraphQL response structure  
type GraphQLResponse struct {
	Data   interface{}      `json:"data,omitempty"`
	Errors []GraphQLError   `json:"errors,omitempty"`
}

type GraphQLError struct {
	Message string `json:"message"`
}

// handle Graphql request
func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received GraphQL request:", r.Method)

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the GraphQL request
	var req GraphQLRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received query: %s\n", req.Query)

	// Determine which service to route to
	serviceURL := routeQuery(req.Query)
	if serviceURL == "unknown" {
		http.Error(w, "Unable to route query", http.StatusBadRequest)
		return
	}

	fmt.Printf("Forwarding to: %s\n", serviceURL)

	// Forward request to the service
	responseBytes, err := forwardRequest(serviceURL, req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Service error: %v", err), http.StatusInternalServerError)
		return
	}

	// Return the response from the service
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBytes)
}

// this func will serve a simpkle GraphQL playground/test page
func playgroundHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Serving playground")
	
	html := `
<!DOCTYPE html>
<html>
<head>
    <title>GraphQL Gateway</title>
</head>
<body>
    <h1>GraphQL Gateway</h1>
    <p>Gateway is running!</p>
    <p>Send POST requests to <code>/query</code> for GraphQL queries</p>
    <p>Example: <code>curl -X POST http://localhost:8080/query -H "Content-Type: application/json" -d '{"query":"{ hello }"}'</code></p>
</body>
</html>`
	
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func routeQuery(query string) string {
	query = strings.ToLower(query)
	
	if strings.Contains(query, "getallproducts") || strings.Contains(query, "getproduct") {
		return "http://localhost:8081/query"
	}
	if strings.Contains(query, "getallusers") || strings.Contains(query, "getuser") {
		return "http://localhost:8082/query"
	}
	
	return "unknown"
}

func forwardRequest(serviceURL string, req GraphQLRequest) ([]byte, error) {
	// Convert request back to JSON
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	// Create HTTP request to the service
	httpReq, err := http.NewRequest("POST", serviceURL, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")

	// Send request to the service
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response from the service
	return io.ReadAll(resp.Body)
}

func main(){
	fmt.Println("Starting gateway...")

	// Route for GraphQL queries (post requests)
	http.HandleFunc("/query", graphqlHandler)

	// Route for the playground/homepage (get request)
	http.HandleFunc("/", playgroundHandler)

	fmt.Println("Gateway running on:")
	fmt.Println("  - Playground: http://localhost:8080")
	fmt.Println("  - GraphQL:    http://localhost:8080/query")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
