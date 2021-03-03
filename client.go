package sveawebpay

import (
	"bytes"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

//URLTest is the api url used for tests
const URLTest = "https://webpaypaymentgatewaytest.svea.com/webpay/rest/"

//URLProd is the api url used in prod
const URLProd = "https://webpaypaymentgateway.svea.com/webpay/rest/"

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
func (c *Client) post(method string, body interface{}, dst interface{}) error {
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

	//Determine the URL to use
	URL := URLProd
	if c.Test {
		URL = URLTest
	}

	//Do the request
	resp, err := http.DefaultClient.PostForm(URL+method, form)
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

	if dst != nil {
		//Unmarshal the xml response message
		if err := xml.Unmarshal(d, dst); err != nil {
			return fmt.Errorf("package error: failed to unmarshal xml response message: %v", err.Error())
		}
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
func (c *Client) PreparePayment(order Order) (*PreparedPayment, error) {
	order.Amount *= 100

	//Make the post request to the api
	var resp preparedPaymentResponse
	if err := c.post("preparepayment", order, &resp); err != nil {
		return nil, err
	}
	resp.PreparedPayment.test = c.Test

	return &resp.PreparedPayment, nil
}

//DecodePaymentResponseBody decodes the response body of a call to the svea api.
//It can for example be used in the callback of a prepared payment.
func (c *Client) DecodePaymentResponseBody(r io.Reader) (*Transaction, error) {
	//Read the response body
	buf := new(bytes.Buffer)
	if i, err := buf.ReadFrom(r); err != nil {
		return nil, fmt.Errorf("package error: failed to read response body (%v bytes read): %v", i, err.Error())
	}

	//Unmarshal the response body
	var rb responseBody
	if err := xml.Unmarshal(buf.Bytes(), &rb); err != nil {
		return nil, fmt.Errorf("package error: failed to unmarshal xml response body: %v", err.Error())
	}

	//Decode the base64 response message
	d, err := base64.StdEncoding.DecodeString(rb.Message)
	if err != nil {
		return nil, fmt.Errorf("package error: failed to decode base64 response message: %v", err.Error())
	}

	//Unmarshal the xml response message
	var p paymentResponse
	if err := xml.Unmarshal(d, &p); err != nil {
		return nil, fmt.Errorf("package error: failed to unmarshal xml response message: %v", err.Error())
	}
	p.Transaction.Amount /= 100

	return &p.Transaction, nil
}

//Recur calls the api to create a new transaction for an existing subscription
func (c *Client) Recur(order RecurOrder) (Transaction, error) {
	order.Amount *= 100

	//Make the post request to the api
	var resp paymentResponse
	if err := c.post("recur", order, &resp); err != nil {
		return Transaction{}, err
	}

	return Transaction{}, nil
}

//CancelRecurSubscription calls the api to cancel an existing recur subscription so that
//no further recurs can be made on it
func (c *Client) CancelRecurSubscription(subscriptionID int) error {
	//Define the request body
	req := struct {
		XMLName        xml.Name `xml:"cancelrecursubscription"`
		SubscriptionID int      `xml:"subscriptionid"`
	}{
		SubscriptionID: subscriptionID,
	}

	//Make the post request to the api
	var resp paymentResponse
	if err := c.post("cancelrecursubscription", req, &resp); err != nil {
		return err
	}

	return nil
}

//Credit calls the api to return money to the customer. Only transactions that have reached
//the status SUCCESS can be credited
func (c *Client) Credit(transactionID int, amount float64) error {
	//Define the request body
	req := struct {
		XMLName        xml.Name `xml:"credit"`
		TransactionID  int      `xml:"transactionid"`
		AmountToCredit float64  `xml:"amounttocredit"`
	}{
		TransactionID:  transactionID,
		AmountToCredit: amount * 100,
	}

	//Make the post request to the api
	if err := c.post("credit", req, nil); err != nil {
		return err
	}

	return nil
}

//Annul calls the api to cancel a payment before it has been captured. It can only be
//performed on card, invoice, or payment plan transactions having the status AUTHORIZED or CONFIRMED
func (c *Client) Annul(transactionID int) error {
	//Define the request body
	req := struct {
		XMLName       xml.Name `xml:"annul"`
		TransactionID int      `xml:"transactionid"`
	}{
		TransactionID: transactionID,
	}

	//Make the post request to the api
	if err := c.post("annul", req, nil); err != nil {
		return err
	}

	return nil
}

//Confirm calls the api to confirm a transaction. It is intended for card transactions.
//It can only be performed on card transactions having the status AUTHORIZED. This will
//result in a CONFIRMED transaction that will be captured (settled) on the given capture
//date. Confirm is mainly used by merchants who are configured to confirm their transactions
//themselves. Otherwise the transactions are confirmed automatically.
func (c *Client) Confirm(transactionID int, captureDate time.Time) error {
	//Define the request body
	req := struct {
		XMLName       xml.Name `xml:"confirm"`
		TransactionID int      `xml:"transactionid"`
		CaptureDate   string   `xml:"capturedate"`
	}{
		TransactionID: transactionID,
		CaptureDate:   captureDate.Format("2006-01-02"),
	}

	//Make the post request to the api
	if err := c.post("confirm", req, nil); err != nil {
		return err
	}

	return nil
}

//LowerAmount calls the api to lower the amount on a transaction before it has been captured.
//It is intended for card transactions having the status AUTHORIZED or CONFIRMED. It can be
//used e.g. if the merchant is unable to deliver all of the items that was ordered by the customer,
//and wants to lower the total amount to pay. If the `amountToLower` is equal to the authorized
//amount, the transaction status will change to annulled.
func (c *Client) LowerAmount(transactionID int, amountToLower float64) error {
	//Define the request body
	req := struct {
		XMLName       xml.Name `xml:"loweramount"`
		TransactionID int      `xml:"transactionid"`
		AmountToLower float64  `xml:"amounttolower"`
	}{
		TransactionID: transactionID,
		AmountToLower: amountToLower * 100,
	}

	//Make the post request to the api
	if err := c.post("loweramount", req, nil); err != nil {
		return err
	}

	return nil
}

//GetTransaction calls the api to get information about a specific order with either the specified
//transaction-id or customer reference number. Both does not need to be provided. If `transactionID <= 0`,
//`customerRefNo` will be used. If `customerRefNo` is an empty string as well a package error will
//be returned. If both are set `transactionID` will be used.
//
//**Please only use this when needed. Repetitive polling is not allowed.**
func (c *Client) GetTransaction(transactionID int, customerRefNo string) (*Transaction, error) {
	var req interface{}
	var method string

	//Define the request body
	if transactionID <= 0 {
		//Transaction id is not provided, use the customer reference number
		method = "querycustomerrefno"
		req = struct {
			XMLName       xml.Name `xml:"loweramount"`
			CustomerRefNo string   `xml:"customerrefno"`
		}{
			CustomerRefNo: customerRefNo,
		}
	} else if len(customerRefNo) > 0 {
		//Use the customer reference number if the transaction id is not provided
		method = "querytransactionid"
		req = struct {
			XMLName       xml.Name `xml:"loweramount"`
			TransactionID int      `xml:"transactionid"`
		}{
			TransactionID: transactionID,
		}
	} else {
		return nil, errors.New("package error: neither a transaction id or customer reference number is provided")
	}

	//Make the post request to the api
	var resp paymentResponse
	if err := c.post(method, req, &resp); err != nil {
		return nil, err
	}

	return &resp.Transaction, nil
}
