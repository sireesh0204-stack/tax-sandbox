package form1065

import (
	"testing"
)

func TestCalculate704bAllocation(t *testing.T) {
	tests := []struct {
		name          string
		partners      []*Partner
		taxableIncome float64
		wantErr       bool
	}{
		{
			name: "Simple allocation",
			partners: []*Partner{
				{Name: "Partner A", CapitalAccount: 50.0, Percentage: 0.5, SSN: "111-11-1111"},
				{Name: "Partner B", CapitalAccount: 50.0, Percentage: 0.5, SSN: "222-22-2222"},
			},
			taxableIncome: 100.0,
			wantErr:       false,
		},
		{
			name: "Proportional allocation",
			partners: []*Partner{
				{Name: "Partner A", CapitalAccount: 25.0, Percentage: 0.25, SSN: "111-11-1111"},
				{Name: "Partner B", CapitalAccount: 75.0, Percentage: 0.75, SSN: "222-22-2222"},
			},
			taxableIncome: 100.0,
			wantErr:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Partnership{
				Partners:                      tt.partners,
				TotalPartnershipTaxableIncome: tt.taxableIncome,
			}

			results, err := p.Calculate704bAllocation()
			if (err != nil) != tt.wantErr {
				t.Fatalf("Calculate704bAllocation() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				for _, res := range results {
					expectedIncome := tt.taxableIncome * (res.Partner.CapitalAccount / (tt.partners[0].CapitalAccount + tt.partners[1].CapitalAccount))
					if res.IncomeShare != expectedIncome {
						t.Errorf("Partner %s: got income %f, want %f", res.Partner.Name, res.IncomeShare, expectedIncome)
					}
				}
			}
		})
	}
}
