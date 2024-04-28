package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestSetPersonalDeduction(t *testing.T) {
	e := echo.New()

	// Test valid request
	t.Run("Valid request", func(t *testing.T) {
		request := PersonalDeductionRequest{Amount: 20000}
		requestJSON, _ := json.Marshal(request)
		req := httptest.NewRequest(http.MethodPost, "/personal-deduction", bytes.NewReader(requestJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, SetPersonalDeduction(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			var response PersonalDeductionResponse
			json.Unmarshal(rec.Body.Bytes(), &response)
			assert.Equal(t, request.Amount, response.PersonalDeduction)
		}
	})

	// Test invalid request (amount too low)
	t.Run("Invalid request (Amount too low)", func(t *testing.T) {
		request := PersonalDeductionRequest{Amount: 5000}
		requestJSON, _ := json.Marshal(request)
		req := httptest.NewRequest(http.MethodPost, "/personal-deduction", bytes.NewReader(requestJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, SetPersonalDeduction(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Contains(t, rec.Body.String(), "Amount must be between 10,000 and 100,000")
		}
	})

	// Test invalid request (amount too high)
	t.Run("Invalid request (Amount too high)", func(t *testing.T) {
		request := PersonalDeductionRequest{Amount: 150000}
		requestJSON, _ := json.Marshal(request)
		req := httptest.NewRequest(http.MethodPost, "/personal-deduction", bytes.NewReader(requestJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, SetPersonalDeduction(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Contains(t, rec.Body.String(), "Amount must be between 10,000 and 100,000")
		}
	})
}

func TestSetKReceiptDeduction(t *testing.T) {
	e := echo.New()

	// Test valid request
	t.Run("Valid request", func(t *testing.T) {
		request := KReceiptDeductionRequest{Amount: 25000}
		requestJSON, _ := json.Marshal(request)
		req := httptest.NewRequest(http.MethodPost, "/kreceipt-deduction", bytes.NewReader(requestJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, SetKReceiptDeduction(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			var response KReceiptDeductionResponse
			json.Unmarshal(rec.Body.Bytes(), &response)
			assert.Equal(t, request.Amount, response.KReceipt)
		}
	})

	// Test invalid request (negative amount)
	t.Run("Invalid request (negative amount)", func(t *testing.T) {
		request := KReceiptDeductionRequest{Amount: -5000}
		requestJSON, _ := json.Marshal(request)
		req := httptest.NewRequest(http.MethodPost, "/kreceipt-deduction", bytes.NewReader(requestJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, SetKReceiptDeduction(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Contains(t, rec.Body.String(), "Amount must be between 0 and 100000")
		}
	})

	// Test invalid request (amount too high)
	t.Run("Invalid request (Amount too high)", func(t *testing.T) {
		request := KReceiptDeductionRequest{Amount: 150000}
		requestJSON, _ := json.Marshal(request)
		req := httptest.NewRequest(http.MethodPost, "/kreceipt-deduction", bytes.NewReader(requestJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, SetKReceiptDeduction(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Contains(t, rec.Body.String(), "Amount must be between 0 and 100000")
		}
	})
}

func TestSetPersonalDeduction_Additional(t *testing.T) {
	e := echo.New()

	// Test empty body
	t.Run("Empty body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/personal-deduction", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, SetPersonalDeduction(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Contains(t, rec.Body.String(), "Invalid request data")
		}
	})

	// Test invalid JSON
	t.Run("Invalid JSON", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/personal-deduction", bytes.NewReader([]byte("invalid json")))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, SetPersonalDeduction(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Contains(t, rec.Body.String(), "Invalid request data")
		}
	})

	// Test Content-Type not set
	t.Run("Content-Type not set", func(t *testing.T) {
		request := PersonalDeductionRequest{Amount: 20000}
		requestJSON, _ := json.Marshal(request)
		req := httptest.NewRequest(http.MethodPost, "/personal-deduction", bytes.NewReader(requestJSON))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, SetPersonalDeduction(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Contains(t, rec.Body.String(), "Invalid request data")
		}
	})
}

func TestSetKReceiptDeduction_Additional(t *testing.T) {
	e := echo.New()

	// Test empty body
	t.Run("Empty body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/kreceipt-deduction", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, SetKReceiptDeduction(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Contains(t, rec.Body.String(), "Invalid request data")
		}
	})

	// Test invalid JSON
	t.Run("Invalid JSON", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/kreceipt-deduction", bytes.NewReader([]byte("invalid json")))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, SetKReceiptDeduction(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Contains(t, rec.Body.String(), "Invalid request data")
		}
	})

	// Test Content-Type not set
	t.Run("Content-Type not set", func(t *testing.T) {
		request := KReceiptDeductionRequest{Amount: 25000}
		requestJSON, _ := json.Marshal(request)
		req := httptest.NewRequest(http.MethodPost, "/kreceipt-deduction", bytes.NewReader(requestJSON))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, SetKReceiptDeduction(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.Contains(t, rec.Body.String(), "Invalid request data")
		}
	})
}
