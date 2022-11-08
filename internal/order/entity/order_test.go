package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAndEmptyID_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{}

	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenAndEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123"}

	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenAndEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAnError(t *testing.T) {
	order := Order{ID: "123", Price: 10}

	assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenValidParams_WhenCreateANewOrder_ThenSouldReceiveACreatedOrderWithAllParams(t *testing.T) {
	order := Order{
		ID:    "123",
		Price: 100.0,
		Tax:   1.2,
	}

	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 100.0, order.Price)
	assert.Equal(t, 1.2, order.Tax)
	assert.Nil(t, order.IsValid())
}

func TestGivenValidParams_WhenCallNewOrderFunc_ThenSouldReceiveACreatedOrderWithAllParams(t *testing.T) {
	order, err := NewOrder("123", 100.0, 1.2)

	assert.Nil(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 100.0, order.Price)
	assert.Equal(t, 1.2, order.Tax)
}

func TestGivenAnPriceAndText_WhenCallCalculateFinalPrice_ThenShouldSetFinalPrice(t *testing.T) {
	order, err := NewOrder("123", 100.0, 1.2)

	assert.Nil(t, err)

	err = order.CalculateFinalPrice()
	assert.Nil(t, err)

	assert.Equal(t, 120.0, order.FinalPrice)
}
