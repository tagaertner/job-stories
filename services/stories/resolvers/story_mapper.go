package resolvers

// import (
// 	"github.com/tagaertner/job-stories/services/stories/generated"
// 	"github.com/tagaertner/job-stories/services/stories/models"
// )

// func ToGraphQLStory(p *models.Story) *generated.Product {
// 	return &generated.Story{
// 		ID:          p.ID,
// 		Name:        p.Name,
// 		Price:       p.Price,
// 		Description: p.Description,
// 		Inventory:   p.Inventory,
// 		Available:   p.Available,
// 	}
// }

// func ToGraphQLProductList(products []*models.Story) []*generated.Story{
// 	var gqlStories []*generated.Product
// 	for _, p := range products {
// 		gqlStories = append(gqlStories, ToGraphQLProduct(p))
// 	}
// 	return gqlStories
// }
