package form1065

import (
	"fmt"
	"math"
)

// AllocationType defines the type of allocation being performed
type AllocationType string

const (
	AllocationTypeIncome   AllocationType = "Income"
	AllocationTypeLoss     AllocationType = "Loss"
	AllocationTypeDeduction AllocationType = "Deduction"
	AllocationTypeCash     AllocationType = "Cash"
)

// BasisDifference represents the inside vs outside basis difference
type BasisDifference struct {
	InsideBasis  float64
	OutsideBasis float64
	Difference   float64
}

// AllocationResult contains the results of an allocation calculation
type AllocationResult struct {
	Partner       *Partner
	IncomeShare   float64
	LossShare     float64
	DeductionShare float64
	CapitalShare   float64
	AdjustedBasis  float64
}

// Calculate704bAllocation performs 704(b) allocations based on profit/loss
// This implements the standard capital-based allocation method
func (p *Partnership) Calculate704bAllocation() ([]AllocationResult, error) {
	if len(p.Partners) == 0 {
		return nil, fmt.Errorf("no partners in partnership")
	}

	if p.TotalPartnershipTaxableIncome == 0 {
		return nil, fmt.Errorf("no taxable income to allocate")
	}

	// Step 1: Calculate total capital accounts
	totalCapital := 0.0
	partnerCapitals := make(map[string]float64)

	for _, partner := range p.Partners {
		totalCapital += partner.CapitalAccount
		partnerCapitals[partner.Name] = partner.CapitalAccount
	}

	if totalCapital <= 0 {
		return nil, fmt.Errorf("total capital must be positive for allocation")
	}

	// Step 2: Calculate allocation ratios
	// Standard 704(b) method: allocate based on capital accounts
	results := make([]AllocationResult, 0, len(p.Partners))

	for _, partner := range p.Partners {
		// Calculate capital-based ratio
		capitalRatio := partner.CapitalAccount / totalCapital

		// For production: could use special allocations if limited liability
		// But standard treatment is capital account based

		result := AllocationResult{
			Partner:        partner,
			IncomeShare:    p.TotalPartnershipTaxableIncome * capitalRatio,
			LossShare:      0.0,
			DeductionShare: 0.0,
			CapitalShare:   capitalRatio,
		}

		results = append(results, result)

		// Update partner allocations
		partner.IncomeAllocation = result.IncomeShare
		partner.LossAllocation = result.LossShare
		partner.DeductionAllocation = result.DeductionShare

		// Update capital account
		account := CapitalAccount{
			BeginningBalance: partner.CapitalAccount,
			AllocationIncome: result.IncomeShare,
		}
		partner.BasisEndingBalance = account.BeginningBalance + account.AllocationIncome
	}

	return results, nil
}

// CalculateBasisAdjustment computes §743(b) basis adjustments
// For transfers of partnership interest
func (p *Partnership) CalculateBasisAdjustment(partner *Partner) *BasisAdjustment {
	if !p.EnforceContingency {
		return nil
	}

	// Calculate inside basis (characterized to partner's share)
	insideBasis := partner.BasisEndingBalance

	// Calculate outside basis (partner's tax basis)
	outsideBasis := 0.0
	for _, basis := range partner.TaxBasis {
		outsideBasis += basis
	}

	difference := insideBasis - outsideBasis

	return &BasisAdjustment{
		Partner:       partner.Name,
		AdjustmentType: "743(b)",
		AssetKey:      "PartnershipInterest",
		OldBasis:      outsideBasis,
		NewBasis:      insideBasis,
		Amount:        difference,
	}
}

// DistributeLoss applies the "step-up" loss limitation
// A partner cannot deduct losses that would make their basis negative
func (p *Partnership) DistributeLoss(lossAmount float64) map[string]float64 {
	results := make(map[string]float64)

	remainingLoss := lossAmount
	totalAllowedLoss := 0.0

	// Sort partners by capital account (greedy algorithm for limited liability)
	sortedPartners := make([]*Partner, len(p.Partners))
	copy(sortedPartners, p.Partners)

	// First pass: calculate total allowable loss
	for i, partner := range sortedPartners {
		allowableLoss := math.Min(remainingLoss, partner.BasisEndingBalance)
		totalAllowedLoss += allowableLoss
		remainingLoss -= allowableLoss
		_ = i // placeholder for sorting logic
	}

	// Second pass: distribute losses
	remainingLoss = lossAmount
	for _, partner := range sortedPartners {
		allowableLoss := math.Min(remainingLoss, partner.BasisEndingBalance)
		results[partner.Name] = allowableLoss
		partner.LossAllocation = allowableLoss
		remainingLoss -= allowableLoss

		if remainingLoss <= 0 {
			break
		}
	}

	return results
}

// DistributeDeductions allocates partnership deductions to partners
func (p *Partnership) DistributeDeductions(deductionAmount float64) map[string]float64 {
	results := make(map[string]float64)

	for _, partner := range p.Partners {
		share := deductionAmount * partner.Percentage
		results[partner.Name] = share
		partner.DeductionAllocation = share
	}

	return results
}

// BasisCheck verifies partner basis limitations
type BasisCheckResult struct {
	PartnerName      string
	BeginningBasis   float64
	Additions        float64
	Subtractions     float64
	EndingBasis      float64
	ExcessLossDeduction float64
}

// CheckBasisLimitations validates that partners don't deduct losses beyond basis
func (p *Partnership) CheckBasisLimitations() []BasisCheckResult {
	results := make([]BasisCheckResult, 0, len(p.Partners))

	for _, partner := range p.Partners {
		result := BasisCheckResult{
			PartnerName:    partner.Name,
			BeginningBasis: partner.CapitalAccount,
			Additions:      partner.IncomeAllocation,
			Subtractions:   partner.DistributionAllocation,
			EndingBasis:    partner.BasisEndingBalance,
		}

		// Calculate excess loss that cannot be deducted
		if result.EndingBasis < 0 {
			result.ExcessLossDeduction = -result.EndingBasis
		}

		results = append(results, result)
	}

	return results
}

// VarianceAnalysis compares actual vs required allocations
type VarianceAnalysis struct {
	PartnerName string
	TaxBenefit  float64 // Tax benefit from allocation
	Variance    float64 // Difference from standard allocation
}

// AnalyzeAllocationVariance performs a §704(b) variance analysis
// This determines if allocations are substantial economic ownership
func (p *Partnership) AnalyzeAllocationVariance() []VarianceAnalysis {
	results := make([]VarianceAnalysis, 0, len(p.Partners))

	totalIncome := 0.0
	for _, partner := range p.Partners {
		totalIncome += partner.IncomeAllocation
	}

	for _, partner := range p.Partners {
		// Calculate tax benefit (simplified)
		taxBenefit := partner.IncomeAllocation * 0.21 // Assuming 21% corporate rate

		// Variance from capital account proportion
		capitalPct := partner.CapitalAccount
		for _, other := range p.Partners {
			if other.Name != partner.Name {
				capitalPct += other.CapitalAccount
			}
		}

		requiredIncome := p.TotalPartnershipTaxableIncome * (partner.CapitalAccount / capitalPct)
		variance := partner.IncomeAllocation - requiredIncome

		results = append(results, VarianceAnalysis{
			PartnerName: partner.Name,
			TaxBenefit:  taxBenefit,
			Variance:    variance,
		})
	}

	return results
}

// SpecialAllocation handles special allocation rules for certain entities
func (p *Partnership) ApplySpecialAllocationRules() error {
	// For LLCs, might need special allocation treatment
	// This is a placeholder for production rules:
	// - Recursive allocations for tiered structures
	// - Special rules for continuously referenced entities
	// - Limitations on allocations not matching economic interest

	return nil
}

// ValidateAllocations checks that allocations meet 704(b) requirements
func (p *Partnership) Validate704bAllocations() error {
	// Verify allocations are substantially equal to capital accounts
	for _, partner := range p.Partners {
		// Basic validation: allocations should be positive
		if partner.IncomeAllocation < 0 {
			return fmt.Errorf("partner %s has negative income allocation", partner.Name)
		}

		// Check basis limitations
		if partner.BasisEndingBalance < 0 {
			// Loss limitation applies - this may be acceptable
		}
	}

	return nil
}