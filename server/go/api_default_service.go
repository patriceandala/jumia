/*
 * Jumia
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"os"
)

// DefaultApiService is a service that implents the logic for the DefaultApiServicer
// This service should implement the business logic for every endpoint for the DefaultApi API.
// Include any external packages or services that will be required by this service.
type DefaultApiService struct {
}

// NewDefaultApiService creates a default api service
func NewDefaultApiService() DefaultApiServicer {
	return &DefaultApiService{}
}

// ProductSkuGet - Get a product by SKU
func (s *DefaultApiService) ProductSkuGet(ctx context.Context, sku string) (ImplResponse, error) {
	// TODO - update ProductSkuGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//todo add postgres logic

	body := Product{
		sku,
		"Some Product",
		"Kenya",
	}

	//body := db.CreateProductParams{}

	return Response(200, body), nil

	//return Response(http.StatusNotImplemented, nil), errors.New("ProductSkuGet method not implemented")
}

// ProductSkuPatch - Consume a product. Checks if the product is available first
func (s *DefaultApiService) ProductSkuPatch(ctx context.Context, sku string) (ImplResponse, error) {
	// TODO - update ProductSkuPatch with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//todo add postgres logic

	body := Product{
		sku,
		"Some Product",
		"Kenya",
	}

	return Response(200, body), nil

	//return Response(http.StatusNotImplemented, nil), errors.New("ProductSkuPatch method not implemented")
}

// ProductsPatch - allows bulk update using CSV file
func (s *DefaultApiService) ProductsPatch(ctx context.Context, file *os.File) (ImplResponse, error) {
	// TODO - update ProductsPatch with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//todo add postgres logic
	body := Product{
		"",
		"Some Product",
		"Kenya",
	}
	return Response(200, body), nil

	//return Response(http.StatusNotImplemented, nil), errors.New("ProductsPatch method not implemented")
}
