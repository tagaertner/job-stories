package federation

import (
    "context"
    "fmt"
)

// Entity represents a federated entity
type Entity interface{}

// EntityResolver handles entity resolution for federation
type EntityResolver interface {
    FindEntityByRepresentation(ctx context.Context, representation map[string]interface{}) (Entity, error)
}

// Service holds federation metadata
type Service struct {
    SDL string
}

// FederationConfig contains federation configuration
type FederationConfig struct {
    EntityResolvers map[string]EntityResolver
    ServiceSDL      string
}

// NewFederationConfig creates a new federation config
func NewFederationConfig(sdl string) *FederationConfig {
    return &FederationConfig{
        EntityResolvers: make(map[string]EntityResolver),
        ServiceSDL:      sdl,
    }
}

// AddEntityResolver adds an entity resolver for a specific type
func (fc *FederationConfig) AddEntityResolver(typeName string, resolver EntityResolver) {
    fc.EntityResolvers[typeName] = resolver
}

// ResolveEntities resolves multiple entities from representations
func (fc *FederationConfig) ResolveEntities(ctx context.Context, representations []map[string]interface{}) ([]Entity, error) {
    entities := make([]Entity, len(representations))
    
    for i, representation := range representations {
        typename, ok := representation["__typename"].(string)
        if !ok {
            return nil, fmt.Errorf("missing __typename in representation")
        }
        
        resolver, exists := fc.EntityResolvers[typename]
        if !exists {
            return nil, fmt.Errorf("no entity resolver for type %s", typename)
        }
        
        entity, err := resolver.FindEntityByRepresentation(ctx, representation)
        if err != nil {
            return nil, fmt.Errorf("failed to resolve entity %s: %w", typename, err)
        }
        
        entities[i] = entity
    }
    
    return entities, nil
}

// GetServiceSDL returns the service SDL
func (fc *FederationConfig) GetServiceSDL() string {
    return fc.ServiceSDL
}
