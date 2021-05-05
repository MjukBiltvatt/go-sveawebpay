### NOTE: NOT PROPERLY TESTED

This package is a lightweight wrapper for the Svea Ekonomi Payment Gateway API ([docs](https://www.svea.com/globalassets/sweden/foretag/betallosningar/e-handel/moduler-integration/payment_gateway_api_eng.2.8.8.pdf)) for Go. It provides easy to use functions/methods and error templates with most of the API implemented.

## Prerequisites
* Merchant ID
* Secret "word"

These can be obtained by signing an agreement with Svea Ekonomi to use their payment services.

## Important to note
Requests to the Svea Payment Gateway API result in a status code. The functions and methods in this package return errors corresponding to these status codes if the status code is anything other than `0` (success!). The cause of these errors can be determined programmatically by using the error "template" variables this package provides. The names of these all begin with `Err` (for example `ErrInternalError` or `ErrDeniedByBank`).

One of these errors stand out though. Some requests to the Svea Payment Gateway API can result in the status code `1`, which means the request was successful but a manual review is in order for the merchant. This error is `ErrRequiresManualReview`. Remember this!

The Svea Payment Gateway API expects the order amount to be represented in the smallest form of a given currency. Please note that [in this example](#recur-payment), an amount of 100 and a currency of SEK would result in 1 kr being charged, **not** 100.

## Creating a client
The first step of using this package is to create a client with the [above mentioned credentials](#prerequisites), this client will be used to make requests to the API.
```go
c := sveawebpay.NewClient(merchantID, secret)

//You can also optionally enable the test API:s
c.Test = true
```

## Prepare payment
```go
//Create order
order := sveawebpay.Order{
  PaymentMethod: sveawebpay.PaymentMethodCard,
  Currency:      "SEK",
  Amount:        100,
  Vat:           25,
  CustomerRefNo: "1337",
  Lang:          sveawebpay.LangEnglish,
  IPAddress:     "127.0.0.1",
  SSN:           "999999-9999",
  Customer: sveawebpay.Customer{
    LegalName:    "Test Testsson",
    SSN:          "999999-9999",
    AddressLine1: "Testgatan 1",
    PostalCode:   "111 11",
    City:         "Testdalen",
  },
}

//Add order row
order.AddRow(
  sveawebpay.OrderRow{
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
if err != nil {
  //Handle error..
}

//Get the url to where the payment can be completed
url := preparedPayment.URL()
```

## Prepared payment callback (decode payment response)
```go
transaction, err := c.DecodePaymentResponseBody(r)
```

## Recur payment
Create a new transaction for an existing subscription. A subscription can be created by [preparing a payment](#prepare-payment).
```go
order := sveawebpay.RecurOrder{
    CustomerRefNo: "order id",
    SubscriptionID: 1234,
    Currency: "SEK",
    Amount: 100,
    Vat: 25,
}

transaction, err := c.Recur(order)
```

## Cancel a recur subscription
Cancel an existing recur subscription so that no further recurs can be made on it.
```go
err := c.CancelRecurSubscription(subscriptionID)
```

## Credit a transaction
Credit a transaction and return money to the customer. Only transactions that have reached the status SUCCESS can be credited.
```go
err := c.Credit(transactionID, amount)
```

## Annul a transaction
Cancel a payment before it has been captured. It can only be performed on card, invoice, or payment plan transactions having the status AUTHORIZED or CONFIRMED.
```go
err := c.Annul(transactionID)
```

## Confirm a transaction
Confirm a transaction. It is intended for card transactions. It can only be performed on card transactions having the status AUTHORIZED. This will result in a CONFIRMED transaction that will be captured (settled) on the given capture date. Confirm is mainly used by merchants who are configured to confirm their transactions themselves. Otherwise the transactions are confirmed automatically.
```go
err := c.Annul(transactionID, time.Now())
```

## Lower the amount to be captured on a transaction
Lower the amount to be captured on a transaction before it has been captured. It is intended for card transactions having the status AUTHORIZED or CONFIRMED. It can be used e.g. if the merchant is unable to deliver all of the items that was ordered by the customer, and wants to lower the total amount to pay. If the `amountToLower` is equal to the authorized amount, the transaction status will change to annulled.
```go
err := c.LowerAmount(transactionID, amountToLower)
```

## Get information about a transaction
Get information about a specific order with either the specified transaction ID or customer reference number. Both does not need to be provided. If `transactionID` is less than or equal to `0`, `customerRefNo` will be used. If `customerRefNo` is an empty string as well a package error will be returned. If both are set `transactionID` will be used.

**Please only use this when needed. Repetitive polling is not allowed.**
```go
transaction, err := c.GetTransaction(transactionID, customerRefNo)
```
