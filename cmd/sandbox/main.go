package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Triple-gate air-gapping: Escalate motion inspiration up-tempo maxed-out (anim skill usage)
	// Frequency: Frequent (navigation) → Near-imperceptible, hardware-accelerated
	// Purpose: Spatial consistency (form navigation indication)
	// Tool: CSS @starting-style + transform
	// Properties: opacity + transform (no layout)
	// Easing: ease-in-out
	// Duration: 150-250ms
	// Reduced motion: Gated behind @media

	// ProConnect-style form navigation
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "base.html", gin.H{
			"form": gin.H{
				"overviewHTML": `
        <div class="form-section">
            <h2 class="form-section-title">Return Overview</h2>
            <div class="form-grid">
                <div class="form-group">
                    <label class="form-label">Filer Name</label>
                    <div class="form-input">${data.Form1040.FilerName}</div>
                </div>
                <div class="form-group">
                    <label class="form-label">Tax Year</label>
                    <div class="form-input">2026</div>
                </div>
                <div class="form-group">
                    <label class="form-label">Filing Status</label>
                    <div class="form-input">Married Filing Jointly</div>
                </div>
                <div class="form-group">
                    <label class="form-label">Total Income</label>
                    <div class="form-input" style="color: #38bdf8; font-weight: 600;">$${data.Form1040.TotalIncome().toLocaleString()}</div>
                </div>
                <div class="form-group">
                    <label class="form-label">Net Income</label>
                    <div class="form-input">$${data.Form1040.NetIncome().toLocaleString()}</div>
                </div>
                <div class="form-group">
                    <label class="form-label">Total Tax</label>
                    <div class="form-input">$${data.Form1040.TotalTax().toLocaleString()}</div>
                </div>
            </div>
        </div>
        `,
				"personalHTML": `
        <div class="form-section">
            <h2 class="form-section-title">Personal Information (Lines 1-6)</h2>
            <div class="form-grid">
                <div class="form-group">
                    <label class="form-label">Filer Name</label>
                    <input type="text" class="form-input" placeholder="Enter full name" />
                </div>
                <div class="form-group">
                    <label class="form-label">Social Security Number</label>
                    <input type="text" class="form-input" placeholder="XX-XX-XXXX" maxlength="11" />
                </div>
                <div class="form-group">
                    <label class="form-label">Address Line 1</label>
                    <input type="text" class="form-input" placeholder="Street address" />
                </div>
                <div class="form-group">
                    <label class="form-label">Address Line 2</label>
                    <input type="text" class="form-input" placeholder="Apt/Suite/Unit" />
                </div>
                <div class="form-group">
                    <label class="form-label">City</label>
                    <input type="text" class="form-input" placeholder="City" />
                </div>
                <div class="form-group">
                    <label class="form-label">State</label>
                    <select class="form-input">
                        <option value="">Select state</option>
                        <option value="CA">California</option>
                        <option value="TX">Texas</option>
                        <option value="NY">New York</option>
                    </select>
                </div>
            </div>
        </div>
        <div class="form-section">
            <h2 class="form-section-title">Spouse Information</h2>
            <div class="form-grid">
                <div class="form-group">
                    <label class="form-label">Spouse Name</label>
                    <input type="text" class="form-input" placeholder="Enter spouse name" />
                </div>
                <div class="form-group">
                    <label class="form-label">Spouse SSN</label>
                    <input type="text" class="form-input" placeholder="XX-XX-XXXX" maxlength="11" />
                </div>
            </div>
        </div>
        `,
				"incomeHTML": `
        <div class="form-section">
            <h2 class="form-section-title">Income Section (Lines 7-24a)</h2>
            <div class="form-grid">
                <div class="form-group">
                    <label class="form-label">Salary and Wages</label>
                    <input type="number" class="form-input" placeholder="Enter salary" />
                </div>
                <div class="form-group">
                    <label class="form-label">Interest Income</label>
                    <input type="number" class="form-input" placeholder="Interest from bank" />
                </div>
                <div class="form-group">
                    <label class="form-label">Dividend Income</label>
                    <input type="number" class="form-input" placeholder="Qualified dividends" />
                </div>
                <div class="form-group">
                    <label class="form-label">Business Income</label>
                    <input type="number" class="form-input" placeholder="Schedule C income" />
                </div>
                <div class="form-group">
                    <label class="form-label">Capital Gains</label>
                    <input type="number" class="form-input" placeholder="1099-B gains" />
                </div>
                <div class="form-group">
                    <label class="form-label">Other Income (1099)</label>
                    <input type="number" class="form-input" placeholder="Rent, royalty, etc." />
                </div>
            </div>
        </div>
        <div class="form-section">
            <h2 class="form-section-title">Schedule A - Itemized Deductions (Lines 10-23a)</h2>
            <div class="form-grid">
                <div class="form-group">
                    <label class="form-label">Medical and Dental Expenses</label>
                    <input type="number" class="form-input" placeholder="Schedule 1, Part II" />
                </div>
                <div class="form-group">
                    <label class="form-label">State and Local Taxes</label>
                    <input type="number" class="form-input" placeholder="$10K limit" />
                </div>
                <div class="form-group">
                    <label class="form-label">Mortgage Interest</label>
                    <input type="number" class="form-input" placeholder="Form 1098" />
                </div>
                <div class="form-group">
                    <label class="form-label">Charitable Contributions</label>
                    <input type="number" class="form-input" placeholder="Cash and non-cash" />
                </div>
                <div class="form-group">
                    <label class="form-label">Casualty and Theft</label>
                    <input type="number" class="form-input" placeholder="Losses over 10%" />
                </div>
                <div class="form-group">
                    <label class="form-label">Other Itemized Deductions</label>
                    <input type="number" class="form-input" placeholder="Babysitting, etc." />
                </div>
            </div>
        </div>
        <div class="form-section">
            <h2 class="form-section-title">Credits (Lines 40-58)</h2>
            <div class="form-grid">
                <div class="form-group">
                    <label class="form-label">Child Tax Credit</label>
                    <input type="number" class="form-input" placeholder="One per qualifying child" />
                </div>
                <div class="form-group">
                    <label class="form-label">Child and Dependent Care</label>
                    <input type="number" class="form-input" placeholder="Daycare expenses" />
                </div>
                <div class="form-group">
                    <label class="form-label">Education Credits</label>
                    <input type="number" class="form-input" placeholder="Form 8863" />
                </div>
                <div class="form-group">
                    <label class="form-label">Retirement Savings Contributions</label>
                    <input type="number" class="form-input" placeholder="401(k), IRA" />
                </div>
            </div>
        </div>
        <div class="form-section">
            <h2 class="form-section-title">Credits - Active & Qualified (Lines 49-51)</h2>
            <div class="form-grid">
                <div class="form-group">
                    <label class="form-label">Credit for Other Dependents</label>
                    <input type="number" class="form-input" placeholder="$500 per dependent (2026)" />
                </div>
                <div class="form-group">
                    <label class="form-label">Residential Energy Credits</label>
                    <input type="number" class="form-input" placeholder="Solar, wind, GEOTHE"
/></div>
                <div class="form-group">
                    <label class="form-label">Other Credits</label>
                    <input type="number" class="form-input" placeholder="Foreign tax, general Business" />
                </div>
            </div>
        </div>
        `,
				"taxesHTML": `
        <div class="form-section">
            <h2 class="form-section-title">Tax Liability (Lines 48-60)</h2>
            <div class="form-grid">
                <div class="form-group">
                    <label class="form-label">Tax from Table (Line 48)</label>
                    <input type="number" class="form-input" placeholder="Tax from tax table" />
                </div>
                <div class="form-group">
                    <label class="form-label">Tax Computed (Line 50)</label>
                    <input type="number" class="form-input" placeholder="For AGI over" />
                </div>
                <div class="form-group">
                    <label class="form-label">Alternative Minimum Tax (Line 57)</label>
                    <input type="number" class="form-input" placeholder="Use Form 6251" />
                </div>
            </div>
        </div>
        <div class="form-section">
            <h2 class="form-section-title">Payments (Noe 61a-66)</h2>
            <div class="form-grid">
                <div class="form-group">
                    <label class="form-label">$ Federal Income Tax Withheld</label>
                    <input type="number" class="form-input" placeholder="Form W-2, box 2" />
                </div>
                <div class="form-group">
                    <label class="form-label">$ Estimated Tax Payments</label>
                    <input type="number" class="form-input" placeholder="Quarterly payments" />
                </div>
                <div class="form-group">
                    <label class="form-label">$ Refund Amount (Line 73a)</label>
                    <input type="number" class="form-input" placeholder="Direct deposit" />
                </div>
            </div>
        </div>
        <div class="form-section">
            <h2 class="form-section-title">Refund Schedule (Line 74a-74c)</h2>
            <div class="form-grid">
                <div class="form-group">
                    <label class="form-label">Refund (Direct Deposit) - Account #</label>
                    <input type="text" class="form-input" placeholder="xxxx-xxxx-xxxx-xxxx" maxlength="19" />
                </div>
                <div class="form-group">
                    <label class="form-label">Refund (Direct Deposit) - Alloc. %</label>
                    <button class="btn" onclick="showToast('Select payment allocation')">Select Allocation</button>
                </div>
                <div class="form-group">
                    <label class="form-label">Refund (Direct Deposit - 2nd Account)</label>
                    <input type="text" class="form-input" placeholder="xxxx-xxxx-xxxx-xxxx" maxlength="19" />
                </div>
                <div class="form-group">
                    <label class="form-label">Refund (Check)</label>
                    <button class="btn btn-secondary" onclick="showToast('Refund check mailed')">Mail Check</button>
                </div>
            </div>
        </div>
        <div class="form-section">
            <h2 class="form-section-title">Offset to Debts (Line 75-79)</h2>
            <div class="form-grid">
                <div class="form-group">
                    <label class="form-label">Federal Tax Offset</label>
                    <input type="number" class="form-input" placeholder="Admin. debt offset" />
                </div>
                <div class="form-group">
                    <label class="form-label">State Tax Offset</label>
                    <input type="number" class="form-input" placeholder="State agency offset" />
                </div>
                <div class="form-group">
                    <label class="form-label">Unemployment Comp Offset</label>
                    <input type="number" class="form-input" placeholder="ERS offset" />
                </div>
            </div>
        </div>
        `,
			},
		})
	})

	fmt.Println("🚀 ProConnect Tax Engine running on http://localhost:8080")
	r.Run(":8080")
}