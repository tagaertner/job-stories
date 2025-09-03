# E-Commerce Microservices Platform

A demonstration of a microservices architecture using **Go** and **Node.js**, built with **GraphQL Federation**, **PostgreSQL**, and **Docker**.

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
                  └─────────────┬───────────────┘
                                │
                  ┌─────────────┴───────────────┐
                  │     PostgreSQL Database     │
                  │     Port: 5432              │
                  │     (Shared across services)│
                  └─────────────────────────────┘
```

## Features

- **Microservices architecture** with separate services for products, users, and orders
- **Apollo Federation** GraphQL Gateway for unified API access
- **PostgreSQL database** with GORM for automatic migrations
- **Automated seeding** with test data
- **Docker** setup with health checks and proper service dependencies
- **Cross-service query** capability through federation

---

## 🚀 Quick Start

### Prerequisites

- [Docker](https://www.docker.com/) and Docker Compose installed

### To Run Everything

1. **Clone the repository:**

   ```bash
   git clone https://github.com/tagaertner/e-commerce-graphql.git
   cd e-commerce-graphql
   ```

2. **Start all services:**

   ```bash
   docker-compose up --build
   ```

3. **Wait for services to initialize:**

   - Database will start and become healthy
   - Go services will connect and create tables via GORM
   - Seed data will be automatically inserted
   - Gateway will start and compose the federated schema

4. **Open GraphQL Playground:**
   👉 [http://localhost:4000](http://localhost:4000)

You can now run live GraphQL queries against the federated services.

### Sample Queries

**Get all products:**

```graphql
query GetProducts {
  products {
    id
    name
    description
    price
    inventory
  }
}
```

**Get user with orders:**

```graphql
query GetUserOrders {
  user(id: "1") {
    name
    email
    role
  }
  orders {
    id
    userId
    totalPrice
    status
  }
}
```

**Cross-service federated query:**

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

| Service  | Port | GraphQL Endpoint            | Container Health |
| -------- | ---- | --------------------------- | ---------------- |
| Products | 4001 | http://localhost:4001/query | ✓ Health checks  |
| Users    | 4002 | http://localhost:4002/query | ✓ Health checks  |
| Orders   | 4003 | http://localhost:4003/query | ✓ Health checks  |
| Gateway  | 4000 | http://localhost:4000       | Federated API    |
| Database | 5432 | PostgreSQL                  | ✓ Health checks  |

## Project Structure

```
e-commerce-graphql/
├── docker-compose.yml             # Service orchestration
├── .env                          # Environment variables (not tracked in git)
├── database/
│   └── init/
│       └── 01-seed-data.sql      # Test data for all services
├── gateway/                      # Node.js Apollo Federation Gateway
│   ├── gateway.js               # Gateway with federation composition
│   ├── package.json
│   └── dockerfile
├── services/                     # Go microservices
│   ├── orders/                  # Orders service
│   │   ├── main.go
│   │   ├── database/            # DB connection & migrations
│   │   ├── models/              # GORM models
│   │   ├── resolvers/           # GraphQL resolvers
│   │   ├── services/            # Business logic
│   │   ├── schema.graphql
│   │   └── dockerfile
│   ├── products/                # Products service (same structure)
│   └── users/                   # Users service (same structure)
└── README.md
```

## Database Features

- **Automatic migrations:** GORM creates tables on service startup
- **Seed data:** Test users, products, and orders inserted automatically
- **Shared database:** All services connect to the same PostgreSQL instance
- **Health monitoring:** Database health checks ensure services start in correct order

## Sample Data

The system includes comprehensive test data:

- **Users:** 10 users (customers and admins, including inactive accounts)
- **Products:** 15 products (Apple ecosystem with realistic pricing)
- **Orders:** 20 orders with various statuses (pending, shipped, completed, cancelled)

## Development

### Environment Variables

Create a `.env` file with:

```bash
# Database
POSTGRES_USER=ecom_user
POSTGRES_PASSWORD=your_password
POSTGRES_DB=ecom_db
DB_HOST=db
DB_PORT=5432

# Service Ports
PORT_PRODUCTS=4001
PORT_USERS=4002
PORT_ORDERS=4003
PORT_GATEWAY=4000
```

### Service Dependencies

The Docker Compose setup ensures proper startup order:

1. PostgreSQL database starts first
2. Go services wait for healthy database, then create tables
3. Seed data is inserted after tables exist
4. Gateway starts after all services are ready

## Technical Details

- **Backend Services:** Go with `gqlgen` for GraphQL schema generation
- **API Gateway:** Node.js with Apollo Gateway for federation
- **Database:** PostgreSQL with GORM for Go services
- **Communication:** Internal Docker networking between services
- **Containerization:** Multi-stage Docker builds for production efficiency

## Future Development

1. **Advanced Mutations** - Create, update, delete operations
2. **Real-time Features** - GraphQL subscriptions
3. **Testing Suite** - Unit and integration tests
4. **Authentication & Authorization** - JWT tokens, role-based access control
5. **Monitoring** - Logging, metrics, tracing
6. **Cloud Deployment** - AWS/GCP with managed databases

---

**MIT License Copyright (c) 2025 Tami Gaertner**
