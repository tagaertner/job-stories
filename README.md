# E-Commerce Microservices Platform

A demonstration of microservices architecture using **Go** and **Node.js**, implementing an e-commerce system with GraphQL APIs.

## Architecture

```
┌──────────────────┐     ┌──────────────────┐     ┌──────────────────┐
│  🛍️ Products      │     │  👥 Users         │     │  📦 Orders        │
│  Service          │     │  Service          │     │  Service          │
│  Port: 4001       │     │  Port: 4002       │     │  Port: 4003       │
│  (Go + GraphQL)   │     │  (Go + GraphQL)   │     │  (Go + GraphQL)   │
└─────────┬────────┘     └─────────┬────────┘     └─────────┬────────┘
          │                        │                        │
          └────────────────────────┼────────────────────────┘
                                   │
                     ┌─────────────┴─────────────┐
                     │    🚀 API Gateway         │
                     │    Port: 4000             │
                     │    (Node.js + Express)    │
                     │    Unified GraphQL API    │
                     └───────────────────────────┘
```

## Features

- **Microservices Architecture**: Three independent Go services
- **API Gateway**: Service composition and unified client interface
- **Cross-Service Queries**: Data fetching across multiple services
- **Health Monitoring**: Status endpoints for all services

## Quick Start

### Prerequisites

- Go 1.24+
- Node.js 18+

### Setup & Installation

**1. Clone and navigate:**

```bash
git clone https://github.com/tagaertner/e-commerce-graphql.git
cd e-commerce
```

**2. Install gateway dependencies:**

```bash
cd gateway
npm install
```

**3. Install Go dependencies for each service:**

```bash
# Products
cd services/products
go mod tidy
go install github.com/99designs/gqlgen@latest

# Users
cd ../users
go mod tidy

# Orders
cd ../orders
go mod tidy
```

### Running the Services

Start each service in a separate terminal:

**Products Service:**

```bash
cd services/products
go run main.go
# Running on http://localhost:4001
```

**Users Service:**

```bash
cd services/users
go run main.go
# Running on http://localhost:4002
```

**Orders Service:**

```bash
cd services/orders
go run main.go
# Running on http://localhost:4003
```

**API Gateway:**

```bash
cd gateway
npm run dev
# Running on http://localhost:4000
```

### Testing

**GraphQL Playground:** http://localhost:4000/graphql

## API Examples

### Basic Queries

**Get products:**

```graphql
query {
  products {
    id
    name
    price
    inventory
  }
}
```

**Get user orders:**

```graphql
query {
  user(id: "1") {
    name
    email
  }
  ordersByUser(userId: "1") {
    id
    totalPrice
    status
  }
}
```

**Cross-service data:**

```graphql
query {
  products {
    id
    name
    price
  }
  users {
    id
    name
    email
  }
  orders {
    id
    userId
    productId
    totalPrice
  }
}
```

### Service Endpoints

| Service  | Port | GraphQL Playground            | Health Check                 |
| -------- | ---- | ----------------------------- | ---------------------------- |
| Products | 4001 | http://localhost:4001/        | http://localhost:4001/health |
| Users    | 4002 | http://localhost:4002/        | http://localhost:4002/health |
| Orders   | 4003 | http://localhost:4003/        | http://localhost:4003/health |
| Gateway  | 4000 | http://localhost:4000/graphql | http://localhost:4000/health |

## Project Structure

```
e-commerce/
├── gateway/                   # Node.js API Gateway
│   ├── gateway.js            # Service composition logic
│   └── package.json
├── services/                 # Go microservices
│   ├── products/
│   │   ├── main.go          # HTTP server setup
│   │   ├── schema.graphql   # GraphQL schema definition
│   │   ├── gqlgen.yml       # Code generation config
│   │   ├── generated/       # Auto-generated GraphQL code
│   │   ├── models/          # Data models
│   │   └── resolvers/       # Business logic
│   ├── users/               # Same structure
│   └── orders/              # Same structure
└── README.md
```

## Development

### Adding New Features

**1. Update GraphQL schema:**

```bash
cd services/[service-name]
# Edit schema.graphql
gqlgen generate
```

**2. Implement resolvers:**

```bash
# Edit resolvers/resolver.go
# Add business logic
```

**3. Update gateway:**

```bash
cd gateway
# Edit gateway.js to proxy new queries
```

## Technical Details

**Backend Services:** Go with gqlgen for GraphQL server generation  
**API Gateway:** Node.js with Express and node-fetch for service composition  
**Communication:** HTTP/GraphQL between gateway and services

## Sample Data

The services include sample data for demonstration:

- 3 products (laptop, smartphone, headphones)
- 3 users (customers and admin)
- 3 orders with different statuses

---

**MIT License Copyright (c) 2025 Tami Gaertner**
