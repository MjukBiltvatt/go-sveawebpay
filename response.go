package sveawebpay

// The raw response from the Svea Payment Gateway API
type ResponseBody struct {
	Message string `xml:"message"`
}

// The response from the Svea Payment Gateway API with the status code isolated
type Response struct {
	StatusCode int `xml:"statuscode"`
}

// Prepared payment response
type PreparedPaymentResponse struct {
	PreparedPayment PreparedPayment `xml:"preparedpayment"`
	Response
}

// Payment response containing the affected transaction object
type PaymentResponse struct {
	Transaction Transaction `xml:"transaction"`
	Response
}

// A card transaction as represented in the responses from the API
type CardTransaction struct {
	CardType     string `xml:"cardtype"`
	MaskedCardNo string `xml:"maskedcardno"`
	ExpiryMonth  string `xml:"expirymonth"`
	ExpiryYear   string `xml:"expiryyear"`
	AuthCode     string `xml:"authcode"`
}

// A payment transaction as represented in the responses from the API
type Transaction struct {
	PaymentMethod    string `xml:"paymentmethod"`
	CustomerRefNo    string `xml:"customerrefno"`
	Amount           int64  `xml:"amount"`
	Currency         string `xml:"currency"`
	SubscriptionID   int    `xml:"subscriptionid"`
	SubscriptionType string `xml:"subscriptiontype"`
	ID               string `xml:"id,attr"`
	CardTransaction
}
