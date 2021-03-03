package sveawebpay

import "encoding/xml"

//PaymentMethodCard should be used when referring to the card payment method
const PaymentMethodCard = "SVEACARDPAY"

//PaymentMethodSwish should be used when referring to the swish payment method
const PaymentMethodSwish = "SWISH"

//SubscriptionTypeRecurring should be used when referring to the initial order
//for a recurring subscription when the initial order should not be captured
const SubscriptionTypeRecurring = "RECURRING"

//SubscriptionTypeRecurringCapture should be used when referring to the initial order
//for a recurring subscription when the initial order should be captured
const SubscriptionTypeRecurringCapture = "RECURRINGCAPTURE"

//LangSwedish represents the language code for swedish
const LangSwedish = "sv"

//LangDanish represents the language code for danish
const LangDanish = "da"

//LangGerman represents the language code for german
const LangGerman = "de"

//LangEnglish represents the language code for english
const LangEnglish = "en"

//LangFinnish represents the language code for finnish
const LangFinnish = "fi"

//LangFrench represents the language code for french
const LangFrench = "fr"

//LangNorwegian represents the language code for norwegian
const LangNorwegian = "no"

//LangPolish represents the language code for polish
const LangPolish = "pl"

//LangSpanish represents the language code for spanish
const LangSpanish = "es"

//LangItalian represents the language code for italian
const LangItalian = "it"

//LangDutch represents the language code for dutch
const LangDutch = "nl"

//Order represents a payment message to the payment gateway api
type Order struct {
	XMLName            xml.Name  `xml:"payment"`
	PaymentMethod      string    `xml:"paymentmethod"`
	Currency           string    `xml:"currency"`
	Amount             int64     `xml:"amount"`
	Vat                int64     `xml:"vat"`
	CustomerRefNo      string    `xml:"customerrefno"`
	ReturnURL          string    `xml:"returnurl"`
	CancelURL          string    `xml:"cancelurl"`
	CallbackURL        string    `xml:"callbackurl"`
	SimulatorCode      string    `xml:"simulatorcode"`
	Lang               string    `xml:"lang"`
	SubscriptionType   string    `xml:"subscriptiontype"`
	ExternalPaymentRef string    `xml:"externalpaymentref"`
	Customer           Customer  `xml:"customer"`
	IPAddress          string    `xml:"ipaddress"`
	SSN                string    `xml:"ssn"`
	Rows               OrderRows `xml:"orderrows"`
}

//OrderRows represents the xml `orderrows` element in an order
type OrderRows struct {
	Rows []OrderRow `xml:"row"`
}

//OrderRow represents an order row
type OrderRow struct {
	Name          string `xml:"name"`
	Description   string `xml:"description"`
	Amount        int64  `xml:"amount"`
	Vat           int64  `xml:"vat"`
	Quantity      int64  `xml:"quantity"`
	ArticleNumber string `xml:"sku"`
	Unit          string `xml:"unit"`
	IPAddress     string `xml:"ipaddress"`
}

//AddRow appends an order row the the slice of order rows in the order
func (o *Order) AddRow(row OrderRow) {
	o.Rows.Rows = append(o.Rows.Rows, row)
}

//SetRows sets the order rows slice in the order
func (o *Order) SetRows(rows []OrderRow) {
	o.Rows.Rows = rows
}
