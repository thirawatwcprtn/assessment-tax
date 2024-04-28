package service

import (
	"net/http"

	"github.com/labstack/echo"
)

type PersonalDeductionRequest struct {
	Amount float64 `json:"amount"`
}

type PersonalDeductionResponse struct {
	PersonalDeduction float64 `json:"personalDeduction"`
}

type KReceiptDeductionRequest struct {
	Amount float64 `json:"amount"`
}

// Response body structure
type KReceiptDeductionResponse struct {
	KReceipt float64 `json:"kReceipt"`
}

func NewSetPersonalDeduction() echo.HandlerFunc {
	return SetPersonalDeduction
}

func SetPersonalDeduction(c echo.Context) error {
	var request PersonalDeductionRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request data")
	}

	// Validation logic
	if request.Amount < 10000 || request.Amount > 100000 {
		return c.JSON(http.StatusBadRequest, "Amount must be between 10,000 and 100,000")
	}

	response := PersonalDeductionResponse{
		PersonalDeduction: request.Amount,
	}

	return c.JSON(http.StatusOK, response)
}

func SetKReceiptDeduction(c echo.Context) error {
	var request KReceiptDeductionRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request data")
	}
	if request.Amount <= 0 || request.Amount > 100000 {
		return c.JSON(http.StatusBadRequest, "Amount must be between 0 and 100000")
	}

	response := KReceiptDeductionResponse{
		KReceipt: request.Amount,
	}

	return c.JSON(http.StatusOK, response)
}
