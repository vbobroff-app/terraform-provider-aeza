// internal/interfaces/client.go
package interfaces

import (
	"context"

	"github.com/vbobroff-app/terraform-provider-aeza/internal/models"
)

// DataClient интерфейс для data sources
type DataClient interface {
	ListServices(ctx context.Context) ([]models.TerraformService, error)
	ListProducts(ctx context.Context) ([]models.Product, error)
	ListServiceTypes(ctx context.Context) ([]models.ServiceType, error)
	ListServiceGroups(ctx context.Context, serviceType string) ([]models.ServiceGroup, error)
	ListOS(ctx context.Context) ([]models.OperatingSystem, error)
}

// ResourceClient интерфейс для resources
type ResourceClient interface {
	DataClient // Включаем все методы data sources

	CreateService(ctx context.Context, req models.ServiceCreateRequest) (*models.ServiceCreateResponse, error)
	GetService(ctx context.Context, id int64) (*models.Service, error)
	UpdateService(ctx context.Context, id int64, req models.ServiceUpdateRequest) error
	DeleteService(ctx context.Context, id int64) error

	ProlongService(ctx context.Context, serviceID int64, req models.ServiceProlongRequest) (*models.ServiceProlongResponse, error)
}
