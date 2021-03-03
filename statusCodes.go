package sveawebpay

//StatusSuccess - Request performed successfully
const StatusSuccess int = 0

//StatusRequiresManualReview - Request performed successfully but
//requires manual review by merchant
const StatusRequiresManualReview int = 1

//StatusInternalError - Internal system error
const StatusInternalError int = 100

//StatusXMLParseFail - Invalid XML
const StatusXMLParseFail int = 101

//StatusIllegalEncoding - Invalid encoding
const StatusIllegalEncoding int = 102

//StatusIllegalURL - Invalid URL
const StatusIllegalURL int = 104

//StatusIllegalTransactionStatus - Invalid transaction status
const StatusIllegalTransactionStatus int = 105

//StatusExternalError - Failure at third party e.g. the bank
const StatusExternalError int = 106

//StatusDeniedByBank - Transaction rejected by bank
const StatusDeniedByBank int = 107

//StatusCancelled - Transaction cancelled
const StatusCancelled int = 108

//StatusIllegalTransactionID - Invalid transaction ID
const StatusIllegalTransactionID int = 110

//StatusMerchantNotConfigured - Merchant not configured
const StatusMerchantNotConfigured int = 111

//StatusMerchantNotConfiguredAtBank - Merchant not configured at bank
const StatusMerchantNotConfiguredAtBank int = 112

//StatusPaymentMethodNotConfigured - Payment method not configured for merchant
const StatusPaymentMethodNotConfigured int = 113

//StatusTimeoutAtBank - Timeout at bank
const StatusTimeoutAtBank int = 114

//StatusMerchantNotActive - The merchant is inactivated
const StatusMerchantNotActive int = 115

//StatusPaymentMethodNotActive - The payment method is inactivated
const StatusPaymentMethodNotActive int = 116

//StatusIllegalAuthorizedAmount - Amount cannot be authorized
const StatusIllegalAuthorizedAmount int = 117

//StatusIllegalCapturedAmount - Amount cannot be captured
const StatusIllegalCapturedAmount int = 118

//StatusIllegalCreditedAmount - Amount cannot be credited
const StatusIllegalCreditedAmount int = 119

//StatusExceedsAmountLimit - Amount exceeds the limit
const StatusExceedsAmountLimit int = 124

//StatusTransactionNotBelongingToMerchant - Transaction does not belong to merchant
const StatusTransactionNotBelongingToMerchant int = 126

//StatusCustomerRefNoAlreadyUsed - Customer reference number already
//used in another transaction
const StatusCustomerRefNoAlreadyUsed int = 127

//StatusNoSuchTrans - Transaction does not exist
const StatusNoSuchTrans int = 128

//StatusDuplicateTransaction - More than one transaction found for
//the given customerrefno
const StatusDuplicateTransaction int = 129

//StatusIllegalOperation - Operation not allowed for the given
//payment method
const StatusIllegalOperation int = 130

//StatusCompanyNotActive - Company inactivated
const StatusCompanyNotActive int = 131

//StatusSubscriptionNotActive - Subscription not active
const StatusSubscriptionNotActive int = 133

//StatusSubscriptionNotSupported - Payment method doesn’t support subscriptions
const StatusSubscriptionNotSupported int = 134

//StatusIllegalDateFormat - Invalid date format
const StatusIllegalDateFormat int = 135

//StatusIllegalResponseData - Invalid response data
const StatusIllegalResponseData int = 136

//StatusCurrencyNotConfigured - Currency not configured
const StatusCurrencyNotConfigured int = 138

//StatusCurrencyNotActive - Currency not active
const StatusCurrencyNotActive int = 139

//StatusCurrencyAlreadyConfigured - Currency is already configured
const StatusCurrencyAlreadyConfigured int = 140

//StatusNoValidPaymentMethods - No valid payment methods
const StatusNoValidPaymentMethods int = 142

//StatusCreditDeniedByBank - Credit denied by bank
const StatusCreditDeniedByBank int = 143

//StatusIllegalCreditUsed - User is not allowed to perform credit operation
const StatusIllegalCreditUsed int = 144

//StatusCustomerNotFound - Customer not found
const StatusCustomerNotFound int = 146

//StatusAgeLimitExceeded - E.g. the transaction is too old
const StatusAgeLimitExceeded int = 147

//StatusBrowserNotSupported - The browser version is too old
const StatusBrowserNotSupported int = 148

//StatusPending - The request is still being processed
const StatusPending int = 149

//StatusCreditPending - The credit could not be handled instantly.
//It is put in a queue it will be processed in due time.
const StatusCreditPending int = 150

//StatusBadTransactionID - Invalid transaction ID
const StatusBadTransactionID int = 301

//StatusBadMerchantID - Invalid merchant ID
const StatusBadMerchantID int = 303

//StatusBadLang - Invalid language
const StatusBadLang int = 304

//StatusBadAmount - Invalid amount
const StatusBadAmount int = 305

//StatusBadCustomerRefNo - Invalid customerrefno
const StatusBadCustomerRefNo int = 306

//StatusBadCurrency - Invalid currency
const StatusBadCurrency int = 307

//StatusBadPaymentMethod - Invalid payment method
const StatusBadPaymentMethod int = 308

//StatusBadReturnURL - Invalid returnurl
const StatusBadReturnURL int = 309

//StatusBadMac - Invalid mac
const StatusBadMac int = 311

//StatusBadCardNumberOrCardTypeNotConfigured - Card type not configured for merchant
const StatusBadCardNumberOrCardTypeNotConfigured int = 316

//StatusBadSSN - Invalid ssn
const StatusBadSSN int = 317

//StatusBadVat - Invalid vat
const StatusBadVat int = 318

//StatusBadCaptureDate - Invalid capture date
const StatusBadCaptureDate int = 319

//StatusBadCampaignCode - Invalid campaign code
const StatusBadCampaignCode int = 320

//StatusBadSubscriptionType - Invalid subscription type
const StatusBadSubscriptionType int = 321

//StatusBadSubscriptionID - Invalid subscription ID
const StatusBadSubscriptionID int = 322

//StatusBadBase64 - Invalid Base64
const StatusBadBase64 int = 323

//StatusBadCallbackURL - Invalid callbackurl
const StatusBadCallbackURL int = 325

//Status3DCheckFailed - 3D check failed
const Status3DCheckFailed int = 326

//StatusCardNotEnrolled - Card not enrolled in 3D Secure
const StatusCardNotEnrolled int = 327

//StatusBadIPAddress - Provided IP address is invalid
const StatusBadIPAddress int = 328

//StatusBadMobile - Invalid mobile phone number
const StatusBadMobile int = 329

//StatusBadCountry - Invalid country
const StatusBadCountry int = 330

//Status3DCheckNotAvailable - Merchant 3D configuration invalid
const Status3DCheckNotAvailable int = 331

//StatusTimeout - Timeout at Svea
const StatusTimeout int = 332

//StatusBadPeriod - The reconciliation period is invalid
const StatusBadPeriod int = 333

//StatusBadAddressID - AddressSelector is not valid for this CountryCode
const StatusBadAddressID int = 334

//StatusBadCustomerData - The supplied customer data is invalid
const StatusBadCustomerData int = 335

//StatusBadUnit - Invalid unit
const StatusBadUnit int = 336

//StatusBadExternalPaymentRef - Invalid external payment reference
const StatusBadExternalPaymentRef int = 337

//StatusBadStoredCardAlias - Invalid stored card alias
const StatusBadStoredCardAlias int = 338

//StatusStoredCardAliasNotActive - Stored card alias inactive
const StatusStoredCardAliasNotActive int = 339

//StatusStoredCardsNotEnabled ...
const StatusStoredCardsNotEnabled int = 340

//StatusOnlyDebitCardsAllowed ...
const StatusOnlyDebitCardsAllowed int = 341

//StatusTransactionAlreadyInProgress ...
const StatusTransactionAlreadyInProgress int = 342

//StatusBadCancelURL - Invalid Cancel url
const StatusBadCancelURL int = 343

//StatusBadSimulatorCode - Invalid value for Simulator Code
const StatusBadSimulatorCode int = 344

//StatusBadOrderRow - Invalid format for Order Row
const StatusBadOrderRow int = 345

//StatusBadPayerAlias - Invalid format for Payer AliasPhone number
const StatusBadPayerAlias int = 346

//StatusBadShowStoreCardDialog - Invalid value for Show Store Card Dialog
const StatusBadShowStoreCardDialog int = 347

//StatusAntiFraudCardBinNotAllowed - Antifraud - cardbin not allowed
const StatusAntiFraudCardBinNotAllowed int = 500

//StatusAntiFraudIPLocationNotAllowed - Antifraud – iplocation not allowed
const StatusAntiFraudIPLocationNotAllowed int = 501

//StatusAntiFraudIPLocationAndBinDoesntMatch - Antifraud – ip-location and bin does not match
const StatusAntiFraudIPLocationAndBinDoesntMatch int = 502

//StatusAntiFraudMaxAmountPerIPExceeded - Antofraud – max amount per ip exceeded
const StatusAntiFraudMaxAmountPerIPExceeded int = 503

//StatusAntiFraudMaxTransactionsPerIPExceeded - Antifraud – max transactions per ip exceeded
const StatusAntiFraudMaxTransactionsPerIPExceeded int = 504

//StatusAntiFraudMaxTransactionsPerCardNoExceeded - Antifraud – max transactions per card number exceeded
const StatusAntiFraudMaxTransactionsPerCardNoExceeded int = 505

//StatusAntiFraudMaxAmountPerCardNoExceeded - Antifraud – max amount per cardnumer exceeded
const StatusAntiFraudMaxAmountPerCardNoExceeded int = 506

//StatusAntiFraudIPAddressBlocked - Antifraud – IP address blocked.
const StatusAntiFraudIPAddressBlocked int = 507

//StatusSwishNotEnrolled - Payer alias is invalid, or payee not enrolled
const StatusSwishNotEnrolled int = 600

//StatusSwishRefundOrderNotFound - Original Payment not found or original payment is more than 13
//months old
const StatusSwishRefundOrderNotFound int = 601

//StatusSwishRefundPayerError - Payer alias in the refund does not match the payee alias in the
//original payment
const StatusSwishRefundPayerError int = 602

//StatusSwishRefundPayerOrgNrError - Payer organization number do not match original payment payee
//organization number
const StatusSwishRefundPayerOrgNrError int = 603

//StatusSwishRefundPayeeSSNError - The Payer SSN in the original payment is not the same as the
//SSN for the current Payee
const StatusSwishRefundPayeeSSNError int = 604
