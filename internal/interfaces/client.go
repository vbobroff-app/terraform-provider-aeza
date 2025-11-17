// internal/interfaces/client.go
package interfaces

import (
	"context"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/source-models"
)

// DataClient интерфейс для data sources
type DataClient interface {
	ListServices(ctx context.Context) ([]source_models.Service, error)
	ListProducts(ctx context.Context) ([]source_models.Product, error)
	ListServiceTypes(ctx context.Context) ([]source_models.ServiceType, error)
}

// ResourceClient интерфейс для resources
type ResourceClient interface {
	DataClient // Включаем все методы data sources

	CreateService(ctx context.Context, req source_models.ServiceCreateRequest) (*source_models.ServiceCreateResponse, error)
	GetService(ctx context.Context, id int64) (*source_models.ServiceGetResponse, error)
	UpdateService(ctx context.Context, id int64, req source_models.ServiceCreateRequest) error
	DeleteService(ctx context.Context, id int64) error
}
