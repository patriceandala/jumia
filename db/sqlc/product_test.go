package db

import (
	"context"
	"github.com/patriceandala/jumia/util"
	"github.com/stretchr/testify/require"
	"testing"

	_ "github.com/stretchr/testify"
)

func TestCreateProduct(t *testing.T) {
	createRandomProduct(t)
}

func createRandomProduct(t *testing.T) Product {
	arg := CreateProductParams{
		Sku:     util.RandomSKu(),
		Name:    util.RandomProductName(),
		Country: util.RandomCountry(),
	}

	product, err := testQueries.CreateProduct(context.Background(), arg)
	if err != nil {
		t.Errorf("could not create product %v", err)
		return Product{}
	}

	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Sku, product.Sku)
	require.Equal(t, arg.Name, product.Name)
	require.Equal(t, arg.Country, product.Country)

	require.Zero(t, product.Stock)

	return product
}

func TestGetProduct(t *testing.T) {
	product := createRandomProduct(t)
	product2, err := testQueries.GetProduct(context.Background(), product.Sku)
	if err != nil {
		t.Errorf("could not get product %v", err)
	}

	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product.Sku, product2.Sku)
	require.Equal(t, product.Country, product2.Country)
	require.Equal(t, product.Stock, product2.Stock)
}
