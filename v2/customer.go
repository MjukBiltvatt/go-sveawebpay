package sveawebpay

// Customer represents the customer
type Customer struct {
	LegalName    string `xml:"legalname"`
	SSN          string `xml:"ssn"`
	AddressLine1 string `xml:"addressline1"`
	AddressLine2 string `xml:"addressline2"`
	PostalCode   string `xml:"postcode"`
	City         string `xml:"postarea"`
}
