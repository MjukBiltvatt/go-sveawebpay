package sveawebpay

import (
	"bytes"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

//Client holds the necessary credentials to make requests to the svea apis
type Client struct {
	merchantID string
	secret     string
	Test       bool
}

//NewClient creates a new client used for calls to the svea payment gateway api
func NewClient(merchantID string, secret string) Client {
	return Client{
		merchantID: merchantID,
		secret:     secret,
	}
}

//post makes a http post request to the specified url with the specified body and
//unmarshals the response from xml to the specified response interface
func (c *Client) post(URL string, body interface{}, dst interface{}) error {
	//Marshal xml body
	b, err := xml.Marshal(body)
	if err != nil {
		return fmt.Errorf("package error: failed to marshal xml body: %v", err.Error())
	}
	b = []byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" + string(b))

	//Create the mac
	h := sha512.New()
	if _, err := h.Write([]byte(base64.StdEncoding.EncodeToString(b) + c.secret)); err != nil {
		return fmt.Errorf("package error: failed to hash mac: %v", err.Error())
	}
	mac := hex.EncodeToString(h.Sum(nil))

	//Create the post form
	form := url.Values{}
	form.Add("merchantid", c.merchantID)
	form.Add("message", base64.StdEncoding.EncodeToString(b))
	form.Add("mac", mac)

	//Do the request
	resp, err := http.DefaultClient.PostForm(URL, form)
	if err != nil {
		return fmt.Errorf("package error: request failed: %v", err.Error())
	}
	defer resp.Body.Close()

	//Read the response body
	buf := new(bytes.Buffer)
	if i, err := buf.ReadFrom(resp.Body); err != nil {
		return fmt.Errorf("package error: failed to read response body (%v bytes read): %v", i, err.Error())
	}

	//Unmarshal the response body
	var x responseBody
	if err := xml.Unmarshal(buf.Bytes(), &x); err != nil {
		return fmt.Errorf("package error: failed to unmarshal xml response body: %v", err.Error())
	}

	//Decode the base64 response message
	d, err := base64.StdEncoding.DecodeString(x.Message)
	if err != nil {
		return fmt.Errorf("package error: failed to decode base64 response message: %v", err.Error())
	}

	//Unmarshal the xml response message
	if err := xml.Unmarshal(d, dst); err != nil {
		return fmt.Errorf("package error: failed to unmarshal xml response message: %v", err.Error())
	}

	//Unmarshal the xml response to get the status code
	var r response
	if err := xml.Unmarshal(d, &r); err != nil {
		return fmt.Errorf("package error: failed to unmarshal xml response message status code: %v", err.Error())
	}

	//Check the status code for error
	err = CodeToErr(r.StatusCode)
	if err != ErrUnknown {
		return err
	}

	return nil
}

//PreparePayment calls the api to prepare a payment and if successful
//gets a url that can be visited to complete the payment. If an error occurs
//ahead of the request, status code -1 and a non nil error is returned.
//Otherwise a nil error is returned along with the statuscode returned by the api.
//Always check the error and status code before using the prepared payment.
func (c *Client) PreparePayment(order Order) (PreparedPayment, error) {
	//Determine the request url to use
	reqURL := "https://webpaypaymentgateway.svea.com/webpay/rest/preparepayment"
	if c.Test {
		reqURL = "https://webpaypaymentgatewaytest.svea.com/webpay/rest/preparepayment"
	}

	//Create a copy of the order prepared for the request
	o := order
	o.Amount *= 100

	//Make the post request to the api
	var resp preparedPaymentResponse
	if err := c.post(reqURL, o, &resp); err != nil && err != ErrRequiresManualReview {
		return PreparedPayment{}, err
	}
	resp.PreparedPayment.test = c.Test

	return resp.PreparedPayment, nil
}

//DecodePaymentResponseBody decodes the response body of a call to the svea api.
//It can for example be used in the callback of a prepared payment.
func (c *Client) DecodePaymentResponseBody(r io.Reader) (Transaction, error) {
	//Read the response body
	buf := new(bytes.Buffer)
	if i, err := buf.ReadFrom(r); err != nil {
		return Transaction{}, fmt.Errorf("package error: failed to read response body (%v bytes read): %v", i, err.Error())
	}

	//Unmarshal the response body
	var rb responseBody
	if err := xml.Unmarshal(buf.Bytes(), &rb); err != nil {
		return Transaction{}, fmt.Errorf("package error: failed to unmarshal xml response body: %v", err.Error())
	}

	//Decode the base64 response message
	d, err := base64.StdEncoding.DecodeString(rb.Message)
	if err != nil {
		return Transaction{}, fmt.Errorf("package error: failed to decode base64 response message: %v", err.Error())
	}

	//Unmarshal the xml response message
	var p paymentResponse
	if err := xml.Unmarshal(d, &p); err != nil {
		return Transaction{}, fmt.Errorf("package error: failed to unmarshal xml response message: %v", err.Error())
	}
	p.Transaction.Amount /= 100

	return p.Transaction, nil
}
