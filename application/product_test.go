package application_test

import (
	"github.com/chjoaquim/go-hexagonal-arch/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_EnableValid(t *testing.T) {
	product := application.Product{}
	product.Name = "Test Enable"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)
}

func TestProduct_EnableInValid(t *testing.T) {
	product := application.Product{}
	product.Name = "Test Enable"
	product.Status = application.DISABLED
	product.Price = 0

	err := product.Enable()
	require.Equal(t, "price must be greater than zero to enable product", err.Error())
}

func TestProduct_DisableValid(t *testing.T) {
	product := application.Product{}
	product.Name = "Test Disabled"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)
}

func TestProduct_DisableInValid(t *testing.T) {
	product := application.Product{}
	product.Name = "Test Disabled"
	product.Status = application.ENABLED
	product.Price = 10

	err := product.Disable()
	require.Equal(t, "price must be zero to enable product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.Name = "Test IsValid"
	product.Status = application.DISABLED
	product.Price = 10
	product.ID = uuid.NewV4().String()

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Status = ""
	_, err = product.IsValid()
	require.Nil(t, err)
}

func TestProduct_IsValidWithStatusInvalid(t *testing.T) {
	product := application.Product{}
	product.Name = "Test IsValid"
	product.Status = "INVALID STATUS"
	product.Price = 10
	product.ID = uuid.NewV4().String()

	_, err := product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())
}

func TestProduct_IsValidWithPriceInvalid(t *testing.T) {
	product := application.Product{}
	product.Name = "Test IsValid"
	product.Status = application.ENABLED
	product.Price = -2
	product.ID = uuid.NewV4().String()

	_, err := product.IsValid()
	require.Equal(t, "the price must be greater than zero", err.Error())
}

func TestProduct_GetID(t *testing.T) {
	id := uuid.NewV4().String()
	product := application.Product{}
	product.Name = "Test IsValid"
	product.Status = application.DISABLED
	product.Price = 10
	product.ID = id

	require.Equal(t, id, product.GetID())
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	product.Name = "Test GetName"
	product.Status = application.DISABLED
	product.Price = 10
	product.ID = uuid.NewV4().String()

	require.Equal(t, "Test GetName", product.GetName())
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
	product.Name = "Test GetName"
	product.Status = application.DISABLED
	product.Price = 10
	product.ID = uuid.NewV4().String()

	require.Equal(t, application.DISABLED, product.GetStatus())
	product.Status = application.ENABLED
	require.Equal(t, application.ENABLED, product.GetStatus())
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{}
	product.Name = "Test GetName"
	product.Status = application.DISABLED
	product.Price = 10
	product.ID = uuid.NewV4().String()

	require.Equal(t, float64(10), product.GetPrice())
}

func TestProduct_IsValidByValidation(t *testing.T) {
	product := application.Product{}
	product.Name = ""
	product.Status = application.DISABLED
	product.Price = 10
	product.ID = uuid.NewV4().String()

	_, err := product.IsValid()
	require.Error(t, err)
}
