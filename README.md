# E-Commerce Microservices Platform

A demonstration of a microservices architecture using **Go** and **Node.js**, built with **GraphQL Federation** and **Docker Compose**.

## Architecture

```
┌─────────────────┐     ┌─────────────────┐     ┌─────────────────┐
│  Products       │     │  Users          │     │  Orders         │
│  Service        │     │  Service        │     │  Service        │
│  Port: 4001     │     │  Port: 4002     │     │  Port: 4003     │
│  (Go+GraphQL)   │     │  (Go+GraphQL)   │     │  (Go+GraphQL)   │
└────────┬────────┘     └────────┬────────┘     └────────┬────────┘
         │                       │                       │
         └───────────────────────┼───────────────────────┘
                                 │
                  ┌──────────────┴──────────────┐
                  │     API Gateway             │
                  │     Port: 4000              │
                  │     (Node.js+Apollo Gateway)│
                  │     Unified GraphQL API     │
                  └─────────────────────────────┘
```

## Features

- Microservices architecture with separate services for products, users, and orders
- Apollo Federation-based GraphQL Gateway
- Unified query access across services
- Dockerized setup with health checks
- Cross-service query capability

---

## 🚀 Quick Start (Docker)

### Prerequisites

- [Docker](https://www.docker.com/) installed

### Run Everything

```bash
git clone https://github.com/tagaertner/e-commerce-graphql.git
cd e-commerce-graphql
docker-compose up --build
```

## 🚀 Quick Start (Docker)

### Prerequisites

- [Docker](https://www.docker.com/) installed

### Run Everything

Clone the repo: https://github.com/tagaertner/e-commerce-graphql.git
Navigate to the project: `cd e-commerce-graphql`
Start the services: `docker compose up --build`

🧠 **You will not see terminal logs until you run queries.**

Then, click below to open the GraphQL Playground:
👉 [http://localhost:4000/graphql](http://localhost:4000/graphql)

You can now run live GraphQL queries against the composed services.

### Sample Queries

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

## Service Endpoints

| Service  | Port | GraphQL Playground            | Health Check                 |
| -------- | ---- | ----------------------------- | ---------------------------- |
| Products | 4001 | http://localhost:4001/        | http://localhost:4001/health |
| Users    | 4002 | http://localhost:4002/        | http://localhost:4002/health |
| Orders   | 4003 | http://localhost:4003/        | http://localhost:4003/health |
| Gateway  | 4000 | http://localhost:4000/graphql | http://localhost:4000/health |

## Project Structure

```
e-commerce-graphql/
├── .gitignore
├── docker-compose.yml
├── gateway/                       # Node.js Apollo Federation Gateway
│   ├── gateway.js                 # Gateway entrypoint with subgraph composition
│   ├── package.json               # Gateway dependencies
│   └── Dockerfile
├── services/                      # Go microservices
│   ├── orders/                    # Orders service
│   │   ├── main.go
│   │   ├── schema.graphql
│   │   ├── gqlgen.yml
│   │   ├── generated/
│   │   ├── models/
│   │   ├── resolvers/
│   │   ├── services/              # Business logic (e.g. order_service.go)
│   │   └── Dockerfile
│   ├── products/                  # Product service (same structure)
│   └── users/                     # User service (same structure)
└── README.md
```

## Docker Configuration

The project includes Docker support with:

- **Multi-stage builds** for optimized Go service images
- **Service networking** for inter-service communication
- **Health checks** for container monitoring
- **Volume mounting** for development (optional)

## Future Development

**1. Add data layer/Database Integration**

**2. Add Mutations**

**3. Testing**

**4. Advanced Queries**

**5. Authentication & Authorization**

**6. AWS/GCP Deployment**

## Technical Details

**Backend Services:** Go using `gqlgen` for GraphQL schema and resolver generation  
**API Gateway:** Node.js with **Apollo Gateway** for composing federated GraphQL services  
**Communication:** HTTP and GraphQL between the gateway and Go microservices  
**Containerization:** Docker with multi-stage builds for efficient, production-ready images

## Sample Data

The services include sample data for demonstration:

- 3 products (laptop, smartphone, headphones)
- 3 users (customers and admin)
- 3 orders with different statuses

---

**MIT License Copyright (c) 2025 Tami Gaertner**
