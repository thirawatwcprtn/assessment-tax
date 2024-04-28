package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/thirawatwcprtn/assessment-tax/model"
	"github.com/thirawatwcprtn/assessment-tax/service"
)

func CalculateTaxHandler(c echo.Context) error {
	var request model.TaxRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	tax := service.CalculateTax(request.TotalIncome, request.Allowances)

	response := model.TaxResponse{
		Tax: tax,
	}
	return c.JSON(http.StatusOK, response)
}
