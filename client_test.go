package sveawebpay

import (
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

// TestClientPreparePayment is an integration test that validates PreparePayment
// with real API calls. Requires the following environment variables:
//   - MERCHANT_ID_TEST: Test merchant ID
//   - SECRET_TEST: Test secret
func TestClientPreparePayment(t *testing.T) {
	//Load .env file
	if err := godotenv.Load(); err != nil {
		t.Errorf("failed to load .env: %v", err.Error())
	}

	//Create client
	c := NewClient(os.Getenv("MERCHANT_ID_TEST"), os.Getenv("SECRET_TEST"))
	c.Test = true

	//Create order
	order := Order{
		PaymentMethod: PaymentMethodCard,
		Currency:      "SEK",
		Amount:        100,
		Vat:           25,
		CustomerRefNo: "1337",
		Lang:          LangEnglish,
		IPAddress:     "127.0.0.1",
		SSN:           "999999-9999",
		Customer: Customer{
			LegalName:    "Test Testsson",
			SSN:          "999999-9999",
			AddressLine1: "Testgatan 1",
			PostalCode:   "111 11",
			City:         "Testdalen",
		},
	}

	//Add order row
	order.AddRow(
		OrderRow{
			Name:          "Test",
			Description:   "Some test",
			Amount:        100,
			Vat:           25,
			Quantity:      1,
			ArticleNumber: "1",
			Unit:          "st",
			IPAddress:     "127.0.0.1",
		},
	)

	//Prepare the payment
	preparedPayment, err := c.PreparePayment(order)
	if err != nil && ErrToCode(err) < 0 {
		t.Error(err)
		return
	}

	t.Logf("PreparedPayment: %v", preparedPayment)
	if preparedPayment != nil {
		t.Logf("URL: %v", preparedPayment.URL())
	}
	t.Logf("Statuscode: %v", ErrToCode(err))
}

// TestClientGetTransaction is an integration test that validates GetTransaction
// with real API calls. Requires the following environment variables:
//   - MERCHANT_ID_TEST: Test merchant ID
//   - SECRET_TEST: Test secret
//   - CUSTOMER_REF_NO: Valid customer reference number
//   - TRANSACTION_ID: Valid transaction ID (numeric)
func TestClientGetTransaction(t *testing.T) {
	//Load .env file
	if err := godotenv.Load(); err != nil {
		t.Errorf("failed to load .env: %v", err.Error())
	}

	//Create client
	c := NewClient(os.Getenv("MERCHANT_ID_TEST"), os.Getenv("SECRET_TEST"))
	c.Test = true

	customerRefNo := os.Getenv("CUSTOMER_REF_NO")
	transactionID, err := strconv.Atoi(os.Getenv("TRANSACTION_ID"))
	if err != nil {
		t.Errorf("Failed to parse TRANSACTION_ID env var: %v", err)
		return
	}

	tests := []struct {
		name          string
		transactionID int
		customerRefNo string
		wantErr       bool
	}{
		{
			name:          "Get by Transaction ID",
			transactionID: transactionID,
			customerRefNo: "",
			wantErr:       false,
		},
		{
			name:          "Get by Customer Ref No",
			transactionID: 0,
			customerRefNo: customerRefNo,
			wantErr:       false,
		},
		{
			name:          "Both ID and Ref No (ID priority)",
			transactionID: transactionID,
			customerRefNo: customerRefNo,
			wantErr:       false,
		},
		{
			name:          "Neither",
			transactionID: 0,
			customerRefNo: "",
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transaction, err := c.GetTransaction(tt.transactionID, tt.customerRefNo)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && transaction == nil {
				t.Error("GetTransaction() returned nil transaction without error")
				return
			}
		})
	}
}
