// Import required dependencies
const { ApolloServer } = require("apollo-server-express");
const express = require("express");
const fetch = require("node-fetch");

// Main proxy gateway that combines all three services
async function startGateway() {
  console.log("🔄 Starting E-Commerce Gateway...");

  // GraphQL schema - defines the API struct that combines all services
  const typeDefs = `
    type Product {
      id: ID!
      name: String!
      price: Float!
      description: String
      inventory: Int!
    }

    type User {
      id: ID!
      name: String!
      email: String!
      role: String!
      active: Boolean!
    }

    type Order {
      id: ID!
      userId: ID!
      productId: ID!
      quantity: Int!
      totalPrice: Float!
      status: String!
      createdAt: String!
    }

    type Query {
      # Product queries
      products: [Product!]!
      product(id: ID!): Product

      # User queries  
      users: [User!]!
      user(id: ID!): User

      # Order queries
      orders: [Order!]!
      order(id: ID!): Order
      ordersByUser(userId: ID!): [Order!]!
    }
  `;

  // Resolvers - func that fetch data when GraphQL queries are made
  const resolvers = {
    Query: {
      // Get all products
      products: async () => {
        try {
          const response = await fetch("http://localhost:4001/query", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
              query: "{ products { id name price description inventory } }",
            }),
          });
          const data = await response.json();
          return data.data.products;
        } catch (error) {
          console.error("Products service error:", error);
          return [];
        }
      },

      // Get single product by ID from Products services
      product: async (_, { id }) => {
        try {
          const response = await fetch("http://localhost:4001/query", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
              query: `{ product(id: "${id}") { id name price description inventory } }`,
            }),
          });
          const data = await response.json();
          return data.data.product;
        } catch (error) {
          console.error("Product service error:", error);
          return null;
        }
      },

      // Get all users
      users: async () => {
        try {
          const response = await fetch("http://localhost:4002/query", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
              query: "{ users { id name email role active } }",
            }),
          });
          const data = await response.json();
          return data.data.users;
        } catch (error) {
          console.error("Users service error:", error);
          return [];
        }
      },

      // Get single user by ID
      user: async (_, { id }) => {
        try {
          const response = await fetch("http://localhost:4002/query", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
              query: `{ user(id: "${id}") { id name email role active } }`,
            }),
          });
          const data = await response.json();
          return data.data.user;
        } catch (error) {
          console.error("User service error:", error);
          return null;
        }
      },

      // Get all orders
      orders: async () => {
        try {
          const response = await fetch("http://localhost:4003/query", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
              query: "{ orders { id userId productId quantity totalPrice status createdAt } }",
            }),
          });
          const data = await response.json();
          return data.data.orders;
        } catch (error) {
          console.error("Orders service error:", error);
          return [];
        }
      },

      // Get single order by ID
      order: async (_, { id }) => {
        try {
          const response = await fetch("http://localhost:4003/query", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
              query: `{ order(id: "${id}") { id userId productId quantity totalPrice status createdAt } }`,
            }),
          });
          const data = await response.json();
          return data.data.order;
        } catch (error) {
          console.error("Order service error:", error);
          return null;
        }
      },

      // Get all orders for specific user
      ordersByUser: async (_, { userId }) => {
        try {
          const response = await fetch("http://localhost:4003/query", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
              query: `{ ordersByUser(userId: "${userId}") { id userId productId quantity totalPrice status createdAt } }`,
            }),
          });
          const data = await response.json();
          return data.data.ordersByUser;
        } catch (error) {
          console.error("OrdersByUser service error:", error);
          return [];
        }
      },
    },
  };

  // Create Apollo GraphQL server with schema and resolver
  const server = new ApolloServer({
    typeDefs,
    resolvers,
    introspection: true,
    playground: true,
  });

  // Create Express web server
  const app = express();

  // Start Apollo server and connect it to Express
  await server.start();
  server.applyMiddleware({ app, path: "/graphql" });

  // Health check endpoint: Monitor and Debug
  app.get("/health", (req, res) => {
    res.json({
      status: "healthy",
      gateway: "e-commerce-proxy",
      services: {
        products: "http://localhost:4001",
        users: "http://localhost:4002",
        orders: "http://localhost:4003",
      },
    });
  });

  // Start server
  const port = 4000;

  app.listen(port, () => {
    console.log("🚀 E-Commerce Gateway ready at http://localhost:4000/graphql");
    console.log("🏥 Health check at http://localhost:4000/health");
    console.log("");
    console.log("📋 Available services:");
    console.log("  🛍️  Products: http://localhost:4001");
    console.log("  👥 Users: http://localhost:4002");
    console.log("  📦 Orders: http://localhost:4003");
  });
}

// Start gateway and handle startup errors
startGateway().catch(console.error);
