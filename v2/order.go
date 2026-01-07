package sveawebpay

import "encoding/xml"

// Available payment methods
const (
	PaymentMethodCard   string = "SVEACARDPAY"
	PaymentMethodSwish  string = "SWISH"
	PaymentMethodPaypal string = "PAYPAL"
)

const (
	//SubscriptionTypeRecurring should be used when referring to the initial order
	//for a recurring subscription when the initial order should not be captured
	SubscriptionTypeRecurring string = "RECURRING"

	//SubscriptionTypeRecurringCapture should be used when referring to the initial order
	//for a recurring subscription when the initial order should be captured
	SubscriptionTypeRecurringCapture string = "RECURRINGCAPTURE"
)

// Available language codes
const (
	LangSwedish   string = "sv"
	LangDanish    string = "da"
	LangGerman    string = "de"
	LangEnglish   string = "en"
	LangFinnish   string = "fi"
	LangFrench    string = "fr"
	LangNorwegian string = "no"
	LangPolish    string = "pl"
	LangSpanish   string = "es"
	LangItalian   string = "it"
	LangDutch     string = "nl"
)

// During the lifetime of a transaction, the status may change many times, most often ending up as
// either SUCCESS or FAILED. The status can be obtained by calling the `GetTransaction` method.
// The range of possible statuses for a specific transaction depends on its payment method.
const (
	TransStatusNew              string = "NEW"
	TransStatusReceived         string = "RECEIVED"
	TransStatusInvalid          string = "INVALID"
	TransStatusValid            string = "VALID"
	TransStatusPending          string = "PENDING"
	TransStatusPendingRecur     string = "PENDINGRECUR"
	TransStatusResponseReceived string = "RESPONSE_RECEIVED"
	TransStatusFailed           string = "FAILED"
	TransStatusCancelled        string = "CANCELLED"
	TransStatusRegistered       string = "REGISTERED"
	TransStatusAuthorized       string = "AUTHORIZED"
	TransStatusConfirmed        string = "CONFIRMED"
	TransStatusAnnulled         string = "ANNULLED"
	TransStatusCapturePending   string = "CAPTPENDING"
	TransStatusCaptureFailed    string = "CAPTFAILED"
	TransStatusSuccess          string = "SUCCESS"
	TransStatusError            string = "ERROR"
)

// Order represents a payment message to the payment gateway api
type Order struct {
	XMLName            xml.Name  `xml:"payment"`
	PaymentMethod      string    `xml:"paymentmethod"`
	Currency           string    `xml:"currency"`
	Amount             int       `xml:"amount"`
	Vat                float64   `xml:"vat"`
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

// AddRow appends an order row the the slice of order rows in the order
func (o *Order) AddRow(row OrderRow) {
	o.Rows.Rows = append(o.Rows.Rows, row)
}

// SetRows sets the order rows slice in the order
func (o *Order) SetRows(rows []OrderRow) {
	o.Rows.Rows = rows
}

// OrderRows represents the xml `orderrows` element in an order
type OrderRows struct {
	Rows []OrderRow `xml:"row"`
}

// OrderRow represents an order row
type OrderRow struct {
	Name          string  `xml:"name"`
	Description   string  `xml:"description"`
	Amount        int     `xml:"amount"`
	Vat           float64 `xml:"vat"`
	Quantity      int     `xml:"quantity"`
	ArticleNumber string  `xml:"sku"`
	Unit          string  `xml:"unit"`
	IPAddress     string  `xml:"ipaddress"`
}

// RecurOrder represents an order to be used to create a recur transaction for
// an existing subscription
type RecurOrder struct {
	XMLName        xml.Name `xml:"recur"`
	CustomerRefNo  string   `xml:"customerrefno"`
	SubscriptionID int      `xml:"subscriptionid"`
	Currency       string   `xml:"currency"`
	Amount         int      `xml:"amount"`
	Vat            float64  `xml:"vat"`
}
