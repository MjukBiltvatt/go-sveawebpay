package sveawebpay

type paymentResponse struct {
	Transaction transaction `xml:"transaction"`
	StatusCode  int         `xml:"statuscode"`
}

type transaction struct {
	PaymentMethod    string `xml:"paymentmethod"`
	MerchantID       string `xml:"merchantid"`
	CustomerRefNo    string `xml:"customerrefno"`
	Amount           int64  `xml:"amount"`
	Currency         string `xml:"currency"`
	SubscriptionID   int    `xml:"subscriptionid"`
	SubscriptionType string `xml:"subscriptiontype"`
}
