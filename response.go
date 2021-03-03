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

//Transaction represents a transaction receieved in the response message from the svea api
type Transaction struct {
	PaymentMethod    string `xml:"paymentmethod"`
	CustomerRefNo    string `xml:"customerrefno"`
	Amount           int64  `xml:"amount"`
	Currency         string `xml:"currency"`
	SubscriptionID   int    `xml:"subscriptionid"`
	SubscriptionType string `xml:"subscriptiontype"`
}
