package sveawebpay

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestClientPreparePayment(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		t.Errorf("failed to load .env: %v", err.Error())
	}

	c := NewClient(os.Getenv("MERCHANT_ID_TEST"), os.Getenv("SECRET_TEST"))
	c.Test = true

	preparedPayment, statusCode, err := c.PreparePayment(Order{
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
		OrderRows: OrderRows{
			Rows: []OrderRow{
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
			},
		},
	})
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("PreparedPayment: %v", preparedPayment)
	t.Logf("URL: %v", preparedPayment.URL())
	t.Logf("Statuscode: %v", statusCode)
}
