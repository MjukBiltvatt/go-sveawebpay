package sveawebpay

import "fmt"

//PreparedPayment represents a prepared payment
type PreparedPayment struct {
	test        bool
	ID          int    `xml:"id"`
	DateCreated string `xml:"created"`
}

//URL returns a http url that can be visited to complete the payment
func (p *PreparedPayment) URL() string {
	if p.test {
		return fmt.Sprintf("https://webpaypaymentgatewaystage.svea.com/webpay/preparedpayment/%v", p.ID)
	}

	return fmt.Sprintf("https://webpaypaymentgateway.svea.com/webpay/preparedpayment/%v", p.ID)
}
