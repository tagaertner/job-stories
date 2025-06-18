package generated

import (
    "context"
    "e-commerce/services/products/models"
    "github.com/99designs/gqlgen/graphql"
)

type Config struct {
    Resolvers ResolverRoot
}

type ResolverRoot interface {
    Query() QueryResolver
}

type QueryResolver interface {
    Products(ctx context.Context) ([]*models.Product, error)
    Product(ctx context.Context, id string) (*models.Product, error)
}

// Simple implementation that returns a basic executable schema
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
    return &executableSchema{
        resolvers: cfg.Resolvers,
    }
}

type executableSchema struct {
    resolvers ResolverRoot
}

func (e *executableSchema) Schema() *graphql.Schema {
    return nil // We'll implement this later
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
    return 1, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
    return func(ctx context.Context) *graphql.Response {
        return &graphql.Response{Data: []byte(`{"data":{}}`)}
    }
}
