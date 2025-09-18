package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/tagaertner/job-stories/services/stories/database"
	"github.com/tagaertner/job-stories/services/stories/generated"
	"github.com/tagaertner/job-stories/services/stories/resolvers"
    "github.com/tagaertner/job-stories/services/stories/services"
)

const defaultPort = "4101"

func main() {
    // Only load .env file when not in Docker
    if _, err := os.Stat(".env"); err == nil {
        err := godotenv.Load()
        if err != nil {
            log.Println("⚠️  Failed to load .env file")
        } else {
            log.Println("✅ Loaded .env file")
        }
    } else {
        log.Println("📦 Running in containerized environment, using system environment variables")
    }
    
    // ... rest of your code
    // Flag to check the database connection and exit
	testDB := flag.Bool("test-db", false, "Test DB connection and exit")
	flag.Parse()

	// Connect to the database
	db := database.Connect()

	if *testDB {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("❌ Failed to get sql DB: %v", err)
            fmt.Println("🔍 DB_HOST =", os.Getenv("DB_HOST"))
		}
		if err := sqlDB.Ping(); err != nil {
			log.Fatalf("❌ Database ping failed: %v", err)
		}
		log.Println("✅ Connected to PostgreSQL successfully")
		return // exit after test
	}

    database.RunMigrations(db)


    port := os.Getenv("PORT")
    if port == "" {
        port = defaultPort
    }

  
    // Creates Story services with data
    storyService := services.NewStoryService(db)

    resolver := &resolvers.Resolver{
        StoryService: storyService,
    }

	srv := handler.New(generated.NewExecutableSchema(
		generated.Config{
			Resolvers: resolver,
		},
        ),
    )

    // Just enable introspection (this is what you actually need)
    srv.Use(extension.Introspection{})

    // Add supported transport methods for GraphQL requests:
	// - POST and GET for queries/mutations
	// - WebSocket transport enables live data features like subscriptions
    srv.AddTransport(transport.POST{})
    srv.AddTransport(transport.GET{})    
    srv.AddTransport(transport.Websocket{}) 

    http.Handle("/", playground.Handler("GraphQL playground", "/query"))
    http.Handle("/query", srv)

    // Health check
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"status": "healthy", "service": "stories"}`))
    })

    log.Printf("📓 [stories] service ready at http://stories:%s/query", port)
    log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}

