package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/thirawatwcprtn/assessment-tax/handler"
	"github.com/thirawatwcprtn/assessment-tax/model"
)

func TestCalculateTaxHandler_ValidRequest(t *testing.T) {
	// สร้าง Echo instance เพื่อทดสอบฮันเดลเลอร์
	e := echo.New()

	// สร้างคำขอที่ถูกต้อง
	request := model.TaxRequest{
		TotalIncome: 500000,
		WHT:         10000,
		Allowances:  []model.Allowance{{AllowanceType: "donation", Amount: 50000}},
	}
	requestJSON, _ := json.Marshal(request)

	// สร้าง httptest request
	req := httptest.NewRequest(http.MethodPost, "/calculate-tax", bytes.NewReader(requestJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	// สร้าง httptest recorder
	rec := httptest.NewRecorder()

	// สร้าง context ใน Echo
	c := e.NewContext(req, rec)

	// เรียกฮันเดลเลอร์
	if assert.NoError(t, handler.CalculateTaxHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code, "Response status should be 200")

		var response model.TaxResponse
		err := json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err, "Response body should be valid JSON")
		assert.Greater(t, response.Tax, 0.0, "Calculated tax should be greater than zero")
	}
}

func TestCalculateTaxHandler_InvalidRequest(t *testing.T) {
	e := echo.New()

	// สร้างคำขอที่ไม่ถูกต้อง (JSON ไม่ถูกต้อง)
	req := httptest.NewRequest(http.MethodPost, "/calculate-tax", bytes.NewReader([]byte("invalid json")))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.CalculateTaxHandler(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code, "Response status should be 400 for invalid request")
		assert.Contains(t, rec.Body.String(), "Invalid request", "Error message should indicate invalid request")
	}
}

func TestCalculateTaxHandler_NoContentType(t *testing.T) {
	e := echo.New()

	// สร้างคำขอโดยไม่มี Content-Type
	request := model.TaxRequest{
		TotalIncome: 500000,
		WHT:         10000,
		Allowances:  []model.Allowance{{AllowanceType: "donation", Amount: 50000}},
	}
	requestJSON, _ := json.Marshal(request)

	req := httptest.NewRequest(http.MethodPost, "/calculate-tax", bytes.NewReader(requestJSON))
	// ไม่ได้ตั้งค่า Content-Type

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler.CalculateTaxHandler(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code, "Response status should be 400 when Content-Type is not set")
		assert.Contains(t, rec.Body.String(), "Invalid request", "Error message should indicate invalid request")
	}
}
