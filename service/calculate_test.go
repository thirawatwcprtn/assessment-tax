package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thirawatwcprtn/assessment-tax/model"
)

func TestCalculateTax(t *testing.T) {
	tests := []struct {
		name        string
		income      float64
		allowances  []model.Allowance
		expectedTax float64
	}{
		{
			name:        "No income, no tax",
			income:      0,
			allowances:  []model.Allowance{},
			expectedTax: 0,
		},
		{
			name:        "Income under personal allowance",
			income:      50000,
			allowances:  []model.Allowance{},
			expectedTax: 0,
		},
		{
			name:        "Income over personal allowance, under first bracket",
			income:      100000,
			allowances:  []model.Allowance{},
			expectedTax: 0,
		},
		{
			name:        "Income in first taxable bracket",
			income:      200000,
			allowances:  []model.Allowance{},
			expectedTax: (200000 - 150000 - 60000) * 0.10,
		},
		{
			name:        "Income in second taxable bracket",
			income:      700000,
			allowances:  []model.Allowance{},
			expectedTax: (350000 * 0.10) + (700000-500000-60000)*0.15,
		},
		{
			name:        "Income in third taxable bracket",
			income:      1500000,
			allowances:  []model.Allowance{},
			expectedTax: (350000 * 0.10) + (500000 * 0.15) + (1500000-1000000-60000)*0.20,
		},
		{
			name:        "Income in top bracket",
			income:      3000000,
			allowances:  []model.Allowance{},
			expectedTax: (350000 * 0.10) + (500000 * 0.15) + (1000000 * 0.20) + (3000000-2000000-60000)*0.35,
		},
		{
			name:   "Income with donation cap",
			income: 1000000,
			allowances: []model.Allowance{
				{AllowanceType: "donation", Amount: 150000},
			},
			expectedTax: (350000 * 0.10) + ((1000000 - 60000 - 100000 - 150000) * 0.15),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tax := CalculateTax(test.income, test.allowances)
			assert.Equal(t, test.expectedTax, tax)
		})
	}
}
