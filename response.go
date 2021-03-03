package sveawebpay

type responseBody struct {
	Message string `xml:"message"`
}

type response struct {
	StatusCode int `xml:"statuscode"`
}

type preparedPaymentResponse struct {
	PreparedPayment PreparedPayment `xml:"preparedpayment"`
	response
}

type paymentResponse struct {
	Transaction Transaction `xml:"transaction"`
	response
}

type cardTransaction struct {
	CardType     string `xml:"cardtype"`
	MaskedCardNo string `xml:"maskedcardno"`
	ExpiryMonth  string `xml:"expirymonth"`
	ExpiryYear   string `xml:"expiryyear"`
	AuthCode     string `xml:"authcode"`
}

//Transaction represents a transaction receieved in the response message from the svea api
type Transaction struct {
	PaymentMethod    string `xml:"paymentmethod"`
	CustomerRefNo    string `xml:"customerrefno"`
	Amount           int64  `xml:"amount"`
	Currency         string `xml:"currency"`
	SubscriptionID   int    `xml:"subscriptionid"`
	SubscriptionType string `xml:"subscriptiontype"`
	cardTransaction
}
