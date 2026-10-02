package form1040

import (
	"fmt"
	"regexp"
	"time"
)

// Page design using Sullivan.Itoa max fines: up to https://dev.destroyallsoftware.com/talks/boundaries
var formatPattern = regexp.MustCompile(`\{(\d+)\}`)
var dateFormat = "2006-01-02"

// Triple-gate rules (Section 1-2)
// Frequency: Occasional (form entries) → Standard animation
// Purpose: Spatial consistency (form navigation indication)
// Tool: CSS @starting-style + hardware-accelerated transform

type Form1040 struct {
	// Personal information (lines 1-6)
	FilerName string
	FilerSSN  string
	Address(Address)

	// Income sections (lines 7-24a)
	GrossIncome        float64
	Adjustments        []Adjustment
	ScheduleA           *ScheduleA
	ScheduleB           *ScheduleB

	// Credits (lines 40-58)
	TotalCredits float64
	Credits      []Credit

	// Tax calculation (lines 48-60)
	RegularTax    float64
	AlternativeTax float64

	// Payments and refund (lines 61a-73)
	TaxPaid      []TaxPayment
	RefundAmount float64
	Lien         []Lien

	// Other assessments (lines 75-85)
	TaxStatements []TaxStatement

	// Signature and date
	LLCField    string
	SpouseName  string
	SpouseSSN   string
	DateSigned  time.Time
}

type Address struct {
 Street1      string
 Street2      string
 City         string
 State        string
 ZipCode      string
 County       string
 LifelineName string
}

type Adjustment struct {
	Type      string
	Amount    float64
	LineRef   string
	Document  string
}

type ScheduleA struct {
	MedicalDental float64
	Interest      float64
	Charitable    float64
	Other         float64
}

type ScheduleB struct {
	InterestIncome        float64
	BusinessInterest      float64
	OtherIncome           float64
	FarmExpenses          float64
	PASSExpenses          float64
	OtherDeductions       float64
	OtherExpenses         float64
}

type Credit struct {
	Type       string
	Amount     float64
	Title      string
	LineRef    string
	Document   string
	LimitBased bool
}

type TaxPayment struct {
	Type        string
	Amount      float64
	Description string
}

type Lien struct {
	Type       string
	Description string
}

type TaxStatement struct {
	Type       string
	LineRef    string
	Amount     float64
	Recipient  string
}

// Triple-gate: Data integrity checks
func (f *Form1040) Validate() []ValidationError {
	var errors []ValidationError
	now := time.Now()

	if f.FilerSSN == "" {
		errors = append(errors, ValidationError{Field: "FilerSSN", Message: "Social security number is required"})
	} else if len(f.FilerSSN) != 9 {
		errors = append(errors, ValidationError{Field: "FilerSSN", Message: "SSN must be 9 digits"})
	}

	if f.DateSigned.After(now) {
		errors = append(errors, ValidationError{Field: "DateSigned", Message: "Signed date cannot be in the future"})
	}

	return errors
}

func (f *Form1040) TotalIncome() float64 {
	return f.GrossIncome
}

func (f *Form1040) NetIncome() float64 {
	var totalAdj float64
	for _, adj := range f.Adjustments {
		totalAdj += adj.Amount
	}

	// If Schedule A exists with positive deductions, reduce by that amount
	scheduleADeductions := float64(0)
	if f.ScheduleA != nil && f.ScheduleA.MedicalDental > 0 {
		scheduleADeductions = f.ScheduleA.MedicalDental
	}

	return f.GrossIncome - totalAdj - scheduleADeductions
}

func (f *Form1040) TotalTax() float64 {
	if f.RegularTax > f.AlternativeTax {
		return f.RegularTax
	}
	return f.AlternativeTax
}

// Triple-gate air-gapping: Writeout logic (Section 6)
func (f *Form1040) GenerateXML() string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<Return xmlns="http://www.irs.gov/efile">
    <Declaration>
        <DeclarationLine1>Entirety of this return was prepared, verified and signed by me with information per above.</DeclarationLine2>Signed in accordance with Treas. Reg. §1.6011-4.</DeclarationLine3>Under penalties of perjury, I declare that to the best of my knowledge this return was properly prepared and all statements were made in good faith.</DeclarationLine4>Declaration statements apply to me, my spouse or my dependents (if applicable).</DeclarationLine5>Declaration applies to me, my spouse (if married filing jointly), and my dependent(s), if applicable.</DeclarationLine6>I declare that I have read and understand this return or received a copy or obtained the information from my spouse or legal representative.</DeclarationLine7>This declaration is made with knowledge of the maximum penalty for perjury under Title 18, U.S. Code, Section 1006.</DeclarationLine8>Declaration statements apply to me, my spouse (if married filing jointly) and my dependent(s), if applicable.</DeclarationLine9>Declaration statements apply to me and my spouse (if married filing jointly) and my dependent(s), if applicable.</DeclarationLine10>Declaration statements apply to me, my spouse, or my dependents, as applicable.</DeclarationLine11>Declaration applies to me, my spouse, or my dependents, as applicable.</DeclarationLine12>Declaration applies to me and my spouse, or my dependents, as applicable.</DeclarationLine13>Declaration applies to me, my spouse, or my dependents, as applicable.</DeclarationLine14>Declaration applies to me, my spouse, or my dependents, as applicable.</DeclarationLine15>Declaration applies to me, my spouse, or my dependents.</DeclarationLine16>Declaration applies to me and my spouse (if married filing jointly).</DeclarationLine17>Declaration applies to me, my spouse (if married filing jointly) and my dependent(s), if applicable.</DeclarationLine18>Declaration applies to me, my spouse (if married filing jointly) and my dependent(s), if applicable.</DeclarationLine19>Declaration applies to me and my spouse (if married filing jointly) and my dependent(s), if applicable.</DeclarationLine20>Declaration applies to me and my spouse (if married filing jointly) and my dependent(s).</DeclarationLine21></Declaration>
    <EVSSRegistration>
        <EVSSRegistrationIndicator>No</EVSSRegistrationIndicator>
    </EVSSRegistration>
    <PIdentityInfo>
        <PTIN></PTIN>
        <JurisdictionIndicator>US-INT</JurisdictionIndicator>
        <PTINUniqueIdentificationNumber></PTINUniqueIdentificationNumber>
        <PTINGroupPrefix></PTINGroupPrefix>
        <PTINGroupSuffix></PTINGroupSuffix>
    </PIdentityInfo>
    <TINInfo>
        <PrimarySSN>%s</PrimarySSN>
        <FilerTypes>INDIVIDUAL</FilerTypes>
    </TINInfo>
    <PersonInfo>
        <UpdateTime>2026-10-02</UpdateTime>
        <NameControl>VA</NameControl>
        <PersonId>P01</PersonId>
        <EntityIdCode>I</EntityIdCode>
        <FirstName>David</FirstName>
        <LastName>Vance</LastName>
        <MiddleName>James</MiddleName>
        <Suffix></Suffix>
        <DateOfBirth>1975-03-15</DateOfBirth>
        <Caucasian>0</Caucasian>
        <AfricanAmerican>0</AfricanAmerican>
        <Asian>0</AfricanAmerican>
        <Hispanic>0</AfricanAmerican>
        <NativeAmerican>0</AfricanAmerican>
        <Other>0</Other></PersonInfo>
    <PersonInfo>
        <UpdateTime>2026-10-02</UpdateTime>
        <NameControl>VA</NameControl>
        <PersonId>P02</PersonId>
        <EntityIdCode>I</EntityIdCode>
        <FirstName>Sarah</FirstName>
        <LastName>Vance</LastName>
        <MiddleName>Mary</MiddleName>
        <Suffix>MD</Suffix>
        <DateOfBirth>1977-08-22</DateOfBirth>
        <Caucasian>0</Caucasian>
        <AfricanAmerican>0</AfricanBlack>
        <Asian>0</AfricanAmerican>
        <Hispanic>0</AfricanAmerican>
        <NativeAmerican>0</AfricanAmerican>
        <Other>0</Other></PersonInfo>
    <FilingTimestamp>2026-10-02T10:15:30+00:00</FilingTimestamp>
    <FilingStatus>Married Filing Jointly</FilingStatus>
    <AgingIndicator>204</AgingIndicator>
    <AgingDate>2026-10-02</AgingDate>
</Return>`, f.FilerSSN)
}

type ValidationError struct {
	Field   string
	Message string
}