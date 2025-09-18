const { ApolloGateway, IntrospectAndCompose } = require("@apollo/gateway");
const { ApolloServer } = require("@apollo/server");
const { startStandaloneServer } = require("@apollo/server/standalone");
const { ApolloServerPluginLandingPageLocalDefault } = require("@apollo/server/plugin/landingPage/default");

async function startServer() {
  try {
    console.log("🔄 Starting job-stories Federation Gateway...");

    // Create the federation gateway
    const gateway = new ApolloGateway({
      supergraphSdl: new IntrospectAndCompose({
        subgraphs: [
          { name: "stories", url: "http://stories:4101/query" },
          { name: "users", url: "http://users:4102/query" },
          // { name: "orders", url: "http://orders:4003/query" },
        ],
      }),
      // Poll every 10 seconds for schema changes
      pollIntervalInMs: 10000,
    });

    // Create Apollo Server with the gateway
    const server = new ApolloServer({
      gateway,
      introspection: true, // Enable introspection for development
      csrfPrevention: {
        requestHeaders: ["apollo-required-preflight"],
      },
      plugins: [
        // Enable GraphQL Playground
        ApolloServerPluginLandingPageLocalDefault({
          embed: true,
          settings: {
            "editor.theme": "dark",
            "editor.fontSize": 14,
          },
        }),

        // Custom plugin for health checks and logging
        {
          requestDidStart() {
            return {
              didResolveOperation(requestContext) {
                console.log(`📊 Query: ${requestContext.request.operationName || "Anonymous"}`);
              },
              didEncounterErrors(requestContext) {
                console.error("❌ GraphQL errors:", requestContext.errors);
              },
            };
          },
        },
      ],
      // Custom error formatting
      formatError: (error) => {
        console.error("🚨 Gateway Error:", error);
        return {
          message: error.message,
          code: error.extensions?.code,
          path: error.path,
        };
      },
    });

    // Start the server
    const { url } = await startStandaloneServer(server, {
      listen: { port: 4100 },
    });

    // Success logging
    console.log("✅ Federation Gateway Successfully Started!");
    console.log("");
    console.log(`🚀 Gateway ready at ${url}`);
    console.log(`🎮 GraphQL Playground: ${url}`);
    console.log("");
    console.log("📋 Connected Services:");
    console.log("  📓 Stories Service: http://localhost:4101/query");
    console.log("  👥 Users Service: http://localhost:4102/query");
    // console.log("  📦 Orders Service: http://localhost:4003/query");
    console.log("");
    console.log("🔗 Example Federated Query:");
    console.log(`
    query ExampleCrossServiceQuery {
    // todo need to fix the example incorrect obj
      user(id: "1") {
        name
        email
        # This will automatically resolve across services
      }
      stories {
        title
        category
      }
      // orders {
      //   id
      //   status
      // }
    }`);
  } catch (error) {
    console.error("💥 Failed to start federation gateway:", error);

    // Error messages
    if (error.message.includes("ECONNREFUSED")) {
      console.error("");
      console.error("🚨 Connection Error: Make sure all services are running:");
      console.error("  - Stories service on port 4101");
      console.error("  - Users service on port 4102");
      // console.error("  - Orders service on port 4103");
      console.error("");
      console.error("💡 Start each service with: go run main.go");
    }

    process.exit(1);
  }
}

// Graceful shutdown handling
process.on("SIGINT", () => {
  console.log("\n🛑 Shutting down gateway gracefully...");
  process.exit(0);
});

process.on("SIGTERM", () => {
  console.log("\n🛑 Shutting down gateway gracefully...");
  process.exit(0);
});

// Start the gateway
startServer().catch((error) => {
  console.error("💥 Unhandled gateway startup error:", error);
  process.exit(1);
});

// Export for testing purposes
module.exports = { startServer };
