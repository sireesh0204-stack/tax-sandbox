package form1065

import (
	"fmt"
	"time"
)

// ScheduleK1 represents Schedule K-1 details for Form 1065
// This captures all the line items that flow to partners
type ScheduleK1 struct {
	PartnerName string
	PartnerID   string // SSN orEIN

	// Part I - Gross Receipts and Sales
	CodeA GrossReceiptsCodeA
}

// GrossReceiptsCodeA represents Part I, Code A: Gross receipts or sales
type GrossReceiptsCodeA struct {
	GrossReceiptsOrSales float64 // Line 1a
	lessReturns          float64 // Line 1b
	NetGrossReceipts     float64 // Line 1c = 1a - 1b
	Narrative            string  // Line 1d
}

// K1LineItem represents a single line item on Schedule K-1
type K1LineItem struct {
	Line   int
	Code   string
	Desc   string
	Amount float64
	TaxTreatment string // "Ordinary", "Capital", "Other"
}

// GenerateK1 creates a Schedule K-1 for a partner
func (p *Partnership) GenerateK1(partner *Partner) *K1 {
	k1 := &K1{
		PartnerName: partner.Name,
		PartnerSSN:  partner.SSN,
		Item:        make([]K1Item, 0),
	}

	// Ordinary business income (Line 1)
	if partner.IncomeAllocation > 0 {
		k1.Item = append(k1.Item, K1Item{
			Code:        "1",
			Description: "Ordinary business income (loss)",
			Amount:      partner.IncomeAllocation,
		})
	}

	// Other information lines
	// Line 2 - Guaranteed payments
	// Line 3 - Other deductions
	// Line 4 - Cost of goods sold
	// Line 5 - Page 2, Part II - Other items

	// §199A QBI information
	k1.Totals.Section199AQBI = partner.QBIIncome

	// Calculate §199A limitation amount
	k1.Totals.Section199AGINEBonus = calculateSection199ALimit(partner.QBIIncome, p.TotalAssets)

	// Net Section 199A deduction
	k1.Totals.NetSection199ADeductions = k1.Totals.Section199AGINEBonus

	return k1
}

// calculateSection199ALimit calculates the §199A deduction limit
func calculateSection199ALimit(qbiIncome float64, totalAssets float64) float64 {
	// Simplified calculation
	// In production: would apply W-2 wage limitations, UBIA limitations
	// and threshold phase-outs based on taxable income

	// Base deduction is 20% of QBI
	baseDeduction := qbiIncome * 0.20

	// Could be limited by:
	// - 50% of W-2 wages (for specified service businesses)
	// - 25% of W-2 wages + 2.5% of UBIA (for farmers, etc.)
	// - 100% of QBI for certain entities

	return baseDeduction
}

// GenerateAllK1s generates K-1s for all partners
func (p *Partnership) GenerateAllK1s() error {
	if len(p.Partners) == 0 {
		return fmt.Errorf("no partners found")
	}

	for _, partner := range p.Partners {
		partner.K1 = p.GenerateK1(partner)
	}

	return nil
}

// K1Error represents an error in K-1 generation
type K1Error struct {
	Partner string
	Field   string
	Message string
}

func (e K1Error) Error() string {
	return fmt.Sprintf("K-1 error for %s: %s - %s", e.Partner, e.Field, e.Message)
}

// ValidateK1 validates a partner's K-1 data
func (k1 *K1) Validate() []ValidationError {
	var errors []ValidationError

	if k1.PartnerName == "" {
		errors = append(errors, ValidationError{
			Field:   "PartnerName",
			Message: "Partner name is required",
		})
	}

	if k1.PartnerSSN == "" || len(k1.PartnerSSN) != 9 {
		errors = append(errors, ValidationError{
			Field:   "PartnerSSN",
			Message: "Valid partner SSN is required",
		})
	}

	// Validate totals
	totals := k1.Totals

	if totals.OrdinaryIncome < 0 {
		errors = append(errors, ValidationError{
			Field:   "Totals.OrdinaryIncome",
			Message: "Ordinary income cannot be negative",
		})
	}

	if totals.Section199AQBI < 0 {
		errors = append(errors, ValidationError{
			Field:   "Totals.Section199AQBI",
			Message: "QBI cannot be negative",
		})
	}

	return errors
}

// GenerateK1XML generates XML representation of Schedule K-1 for MeF filing
func (k1 *K1) GenerateK1XML() string {
	// This is a simplified XML structure
	// Production MeF requires full IRS specification
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<K1 xmlns="http://www.irs.gov/efile">
    <PartnerName>%s</PartnerName>
    <PartnerSSN>%s</PartnerSSN>
    <K1Totals>
        <OrdinaryIncome>%.2f</OrdinaryIncome>
        <NetTier1Income>%.2f</NetTier1Income>
        <Section199AQBI>%.2f</Section199AQBI>
        <Section199ADeduction>%.2f</Section199ADeduction>
    </K1Totals>
</K1>`,
		k1.PartnerName,
		k1.PartnerSSN,
		k1.Totals.OrdinaryIncome,
		k1.Totals.NetTier1Income,
		k1.Totals.Section199AQBI,
		k1.Totals.NetSection199ADeductions,
	)
}

// UpdateK1WithAllocation populates K-1 from allocation results
func (k1 *K1) UpdateK1WithAllocation(partner *Partner) {
	k1.PartnerName = partner.Name
	k1.PartnerSSN = partner.SSN

	k1.Totals.OrdinaryIncome = partner.IncomeAllocation
	k1.Totals.NetTier1Income = partner.IncomeAllocation
	k1.Totals.Section199AQBI = partner.QBIIncome
}

// Timestamp for K-1 generation
func (k1 *K1) Timestamp() time.Time {
	return time.Now()
}