package sveawebpay

import (
	"fmt"
)

//URLTest is the url that will be used for tests
const URLTest = "https://webpaypaymentgatewaytest.svea.com/webpay/payment"

//URLProd is the url that will be used in prod
const URLProd = "https://webpaypaymentgateway.svea.com/webpay/payment"

//Client holds the necessary credentials to make requests to the svea apis
type Client struct {
	merchantID string
	secret     string
	Test       bool
}

type request struct {
	Payment Order `xml:"payment"`
}

//NewClient creates a new client used for calls to the svea payment gateway api
func NewClient(merchantID string, secret string) Client {
	return Client{
		merchantID: merchantID,
		secret:     secret,
	}
}

//getURL returns the url to the payment gateway api depending on whether or not
//the client has been set for testing or not
func (c *Client) getURL() string {
	if c.Test {
		return URLTest
	}

	return URLProd
}

//post makes a http post request to the specified url with the specified body and
//unmarshals the response from xml to the specified response interface
func (c *Client) post(url string, body interface{}, response interface{}) error {
	return nil
}

//PreparePayment calls the api to prepare a payment and if successful
//gets a url that can be visited to complete the payment
func (c *Client) PreparePayment(order Order) (url string, err error) {
	type response struct {
		PreparedPayment struct {
			ID          int    `xml:"id"`
			DateCreated string `xml:"created"`
		} `xml:"preparedpayment"`
		StatusCode int `xml:"statuscode"`
	}

	//Determine the request url to use
	reqURL := "https://webpaypaymentgateway.svea.com/webpay/rest/preparepayment"
	if c.Test {
		reqURL = "https://webpaypaymentgatewaytest.svea.com/webpay/rest/preparepayment"
	}

	//Make the post request to the api
	var resp response
	if err := c.post(reqURL, request{order}, &resp); err != nil {
		return "", err
	}

	//Format the url to return
	if c.Test {
		url = fmt.Sprintf("https://webpaypaymentgatewaystage.svea.com/webpay/preparedpayment/%v", resp.PreparedPayment.ID)
	} else {
		url = fmt.Sprintf("https://webpaypaymentgateway.svea.com/webpay/preparedpayment/%v", resp.PreparedPayment.ID)
	}

	return url, nil
}
