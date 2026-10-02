// Package form1065 implements Form 1065 Partnership Return calculations
// Includes 704(b) allocations, Section 754 basis adjustments, and §199A QBI calculations
package form1065

import (
	"fmt"
	"math"
)

// Partnership represents a U.S. partnership for Form 1065 filing
type Partnership struct {
	Name              string
	EIN               string
	Address           Address
	EnforceContingency bool // §754 election flag
	TaxYear           int

	// Financial aggregates
	TotalPartners       int
	TotalAssets         float64
	TotalLiabilities    float64
	TotalPartnershipTaxableIncome float64

	// Members
	Partners []*Partner

	// §754 election adjustments
	Sections754Adjustments []*BasisAdjustment

	// Form data
	K1s        []*K1
	Items      []Item
}

// Address represents partnership location information
type Address struct {
	Street1      string
	Street2      string
	City         string
	State        string
	ZipCode      string
}

// Partner represents a partner in the partnership
type Partner struct {
	Name           string
	SSN            string
	CapitalAccount float64 // Beginning capital account balance
	Percentage     float64 // Ownership percentage (0.0 to 1.0)

	// Tax basis (for §754 adjustments)
	TaxBasis map[string]float64 // Asset key -> basis amount

	// Allocations
	BasisEndingBalance     float64
	IncomeAllocation       float64
	LossAllocation         float64
	DeductionAllocation    float64
	DistributionAllocation float64
	QBIIncome              float64 // §199A Qualified Business Income

	// Partner type
	PartnerType string // "General", "Limited", "Corporate"

	// K-1 generated for this partner
	K1 *K1
}

// K1 represents Schedule K-1 for a partner
type K1 struct {
	PartnerName        string
	PartnerSSN         string
	Item               []K1Item
	Totals             K1Totals
}

// K1Item represents an item on Schedule K-1
type K1Item struct {
	Code         string
	Description  string
	Amount       float64
}

// ValidationError represents a validation failure
type ValidationError struct {
	Field   string
	Message string
}

// K1Totals contains the summary totals for a K-1
type K1Totals struct {
	OrdinaryIncome           float64
	NetTier1Income           float64
	Distributions            float64
	CapitalGainsLosses       float64
	TaxCredits               float64
	Section199AQuotient      float64
	Section199AGINEBonus     float64
	Section199AQBI           float64
	Section199AUBITI         float64
	NetSection199ADeductions float64
}

// Item represents a partnership item (asset/liability for allocation)
type Item struct {
	Type   string
	Key    string
	Value  float64
}

// BasisAdjustment represents a §754 basis adjustment
type BasisAdjustment struct {
	Partner       string
	AdjustmentType string // "Inside", "Outside", "Other"
	AssetKey      string
	OldBasis      float64
	NewBasis      float64
	Amount        float64
}

// 704(b)AllocationParams contains parameters for allocation calculations
type AllocationParams struct {
	CapitalAccountRatio float64
	IncomeRatio         float64
	LossRatio           float64
	DeductionRatio      float64
}

// CapitalAccount tracks partner capital account changes
type CapitalAccount struct {
	BeginningBalance  float64
	Contributions   float64
	AllocationIncome  float64
	AllocationLoss   float64
	Distributions    float64
	EndingBalance    float64
}

// Validate validates the partnership entity
func (p *Partnership) Validate() []ValidationError {
	var errors []ValidationError

	if p.EIN == "" || len(p.EIN) != 10 {
		errors = append(errors, ValidationError{
			Field:   "EIN",
			Message: "Employer Identification Number is required (10 digits)",
		})
	}

	if p.TaxYear < 1954 {
		errors = append(errors, ValidationError{
			Field:   "TaxYear",
			Message: "Invalid tax year",
		})
	}

	if len(p.Partners) == 0 {
		errors = append(errors, ValidationError{
			Field:   "Partners",
			Message: "At least one partner is required",
		})
	}

	for i, partner := range p.Partners {
		if partner.SSN == "" {
			errors = append(errors, ValidationError{
				Field:   fmt.Sprintf("Partner[%d].SSN", i),
				Message: "Partner SSN is required",
			})
		}

		if partner.Percentage <= 0 || partner.Percentage > 1 {
			errors = append(errors, ValidationError{
				Field:   fmt.Sprintf("Partner[%d].Percentage", i),
				Message: "Partner percentage must be between 0 and 1",
			})
		}
	}

	return errors
}

// CalculateAllocations performs 704(b) allocations based on partner capital accounts
func (p *Partnership) CalculateAllocations() error {
	if len(p.Partners) == 0 {
		return fmt.Errorf("no partners to allocate")
	}

	// Verify total ownership sums to 100%
	var totalPct float64
	for _, partner := range p.Partners {
		totalPct += partner.Percentage
	}

	// Allow small rounding errors up to 0.1%
	if math.Abs(totalPct-1.0) > 0.001 {
		return fmt.Errorf("partner percentages must sum to 1.0, got %.4f", totalPct)
	}

	// Calculate allocations based on capital accounts (704(b) standard method)
	totalCapital := 0.0
	for _, partner := range p.Partners {
		totalCapital += partner.CapitalAccount
	}

	if totalCapital <= 0 {
		return fmt.Errorf("total capital account must be positive")
	}

	// Distribute income, loss, and deductions proportionally
	for _, partner := range p.Partners {
		// Capital account ratio for allocation
		capitalRatio := partner.CapitalAccount / totalCapital

		// Use capital ratio as default 704(b) allocation
		// This implements the standard capital-based allocation
		allocation := &AllocationParams{
			CapitalAccountRatio: capitalRatio,
			IncomeRatio:         capitalRatio,
			LossRatio:           capitalRatio,
			DeductionRatio:      capitalRatio,
		}

		// Calculate partner's share of partnership income/loss
		partner.IncomeAllocation = p.TotalPartnershipTaxableIncome * allocation.IncomeRatio
		partner.LossAllocation = 0.0
		partner.DeductionAllocation = 0.0

		// Update capital account (simplified model)
		account := CapitalAccount{
			BeginningBalance: partner.CapitalAccount,
			AllocationIncome: partner.IncomeAllocation,
		}

		partner.BasisEndingBalance = account.BeginningBalance + account.AllocationIncome

		// Generate K-1 for partner
		if partner.K1 == nil {
			partner.K1 = &K1{
				PartnerName: partner.Name,
				PartnerSSN:  partner.SSN,
			}
		}

		partner.K1.Totals.OrdinaryIncome = partner.IncomeAllocation
	}

	return nil
}

// CalculateQBI calculates §199A Qualified Business Income
// QBI is generally the partner's share of partnership income
func (p *Partnership) CalculateQBI() error {
	if len(p.Partners) == 0 {
		return fmt.Errorf("no partners to calculate QBI")
	}

	// QBI excludes:
	// - Capital gains/losses (unless from trading)
	// - Guaranteed payments
	// - Non-recognition gains
	// - Injunctions/reimbursements
	// - Penalties

	for _, partner := range p.Partners {
		// Simplified QBI calculation - partner's share of ordinary business income
		// In a production system, this would filter out non-QBI items
		partner.QBIIncome = partner.IncomeAllocation

		// Ensure QBI doesn't exceed partnership income
		if partner.QBIIncome > p.TotalPartnershipTaxableIncome {
			partner.QBIIncome = p.TotalPartnershipTaxableIncome
		}

		if partner.K1 != nil {
			partner.K1.Totals.Section199AQBI = partner.QBIIncome
		}
	}

	return nil
}

// CalculateSection199A calculates the §199A deduction
// Limit is 20% of QBI, subject to various limitations
func (p *Partnership) CalculateSection199A() error {
	if len(p.Partners) == 0 {
		return fmt.Errorf("no partners to calculate §199A")
	}

	// Calculate tax benefit formula for QBI limitation
	// For OECD partners (not C-corps), minimum of:
	// - 20% of QBI
	// - 20% of taxable income (with modifications)

	for _, partner := range p.Partners {
		// Simplified calculation - in production would calculate taxable income
		// and apply W-2 wage and UBIA limitations

		qbiDeduction := partner.QBIIncome * 0.20

		if partner.K1 != nil {
			partner.K1.Totals.Section199AUBITI = qbiDeduction
		}

		// Apply limitations based on taxpayer type and income
		// For simplicity, we calculate the base 20% deduction
		// Production code would need to consider:
		// - Specified service business limitations
		// - W-2 wage limitations
		// - UBIA (Unrecaptured Basis) limitations
		// - Taxable income thresholds
	}

	return nil
}

// Calculate754BasisAdjustments calculates §754 election basis adjustments
// Uses inside/outside basis difference method under §743(b)
func (p *Partnership) Calculate754BasisAdjustments() error {
	if !p.EnforceContingency {
		return nil
	}

	for _, partner := range p.Partners {
		// Calculate inside basis vs outside basis difference
		// This is a simplified implementation

		// In production:
		// 1. Determine if transfer triggers §743(b) adjustment
		// 2. Calculate inside basis of distributed property
		// 3. Calculate outside basis of transferred interest
		// 4. Difference is the §754(b) adjustment

		// Track basis adjustments for each significant asset
		otherBasis := 0.0
		for _, basis := range partner.TaxBasis {
			otherBasis += basis
		}

		// Basic §754 adjustment calculation
		basisDiff := partner.BasisEndingBalance - otherBasis

		if math.Abs(basisDiff) > 0.01 && len(p.Sections754Adjustments) > 0 {
			adjustment := &BasisAdjustment{
				Partner:       partner.Name,
				AdjustmentType: "Inside",
				AssetKey:      "All",
				OldBasis:      otherBasis,
				NewBasis:      partner.BasisEndingBalance,
				Amount:        basisDiff,
			}
			p.Sections754Adjustments = append(p.Sections754Adjustments, adjustment)
		}
	}

	return nil
}

// GetPartnerBySSN retrieves a partner by SSN
func (p *Partnership) GetPartnerBySSN(ssn string) *Partner {
	for _, partner := range p.Partners {
		if partner.SSN == ssn {
			return partner
		}
	}
	return nil
}

// GetPartnerByName retrieves a partner by name
func (p *Partnership) GetPartnerByName(name string) *Partner {
	for _, partner := range p.Partners {
		if partner.Name == name {
			return partner
		}
	}
	return nil
}

// ValidateAllocations validates that allocations are consistent
func (p *Partnership) ValidateAllocations() []ValidationError {
	var errors []ValidationError

	totalIncome := 0.0
	totalLoss := 0.0
	totalDeduction := 0.0

	for i, partner := range p.Partners {
		totalIncome += partner.IncomeAllocation
		totalLoss += partner.LossAllocation
		totalDeduction += partner.DeductionAllocation

		// Check for negative allocations (may occur with losses)
		if partner.BasisEndingBalance < 0 {
			errors = append(errors, ValidationError{
				Field:   fmt.Sprintf("Partner[%d].BasisEndingBalance", i),
				Message: "Partner cannot have negative basis (loss limitation applies)",
			})
		}
	}

	// Verify total income allocation matches partnership total
	if math.Abs(totalIncome-p.TotalPartnershipTaxableIncome) > 0.01 {
		errors = append(errors, ValidationError{
			Field:   "TotalIncomeAllocation",
			Message: fmt.Sprintf("Income allocations must sum to partnership income: got %.2f, expected %.2f",
				totalIncome, p.TotalPartnershipTaxableIncome),
		})
	}

	return errors
}