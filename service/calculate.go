package service

import "github.com/thirawatwcprtn/assessment-tax/model"

func CalculateTax(income float64, allowances []model.Allowance) float64 {
	personalAllowance := 60000.0
	totalAllowance := personalAllowance

	for _, allowance := range allowances {
		if allowance.AllowanceType == "donation" {
			donationCap := 100000.0
			if allowance.Amount > donationCap {
				allowance.Amount = donationCap
			}
			totalAllowance += allowance.Amount
		}
	}

	taxableIncome := income - totalAllowance

	tax := 0.0
	if taxableIncome > 0 && taxableIncome <= 150000 {
	} else if taxableIncome > 150000 && taxableIncome <= 500000 {
		tax += (taxableIncome - 150000) * 0.10
	} else if taxableIncome > 500000 && taxableIncome <= 1000000 {
		tax += (350000 * 0.10)
		tax += (taxableIncome - 500000) * 0.15
	} else if taxableIncome > 1000000 && taxableIncome <= 2000000 {
		tax += (350000 * 0.10) + (500000 * 0.15)
		tax += (taxableIncome - 1000000) * 0.20
	} else if taxableIncome > 2000000 {
		tax += (350000 * 0.10) + (500000 * 0.15) + (1000000 * 0.20)
		tax += (taxableIncome - 2000000) * 0.35
	}

	return tax
}
