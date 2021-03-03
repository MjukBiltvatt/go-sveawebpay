package sveawebpay

//StatusSuccess - Request performed successfully
const StatusSuccess = 0

//StatusRequiresManualReview - Request performed successfully but
//requires manual review by merchant
const StatusRequiresManualReview = 1

//StatusInternalError - Internal system error
const StatusInternalError = 100

//StatusXMLParseFail - Invalid XML
const StatusXMLParseFail = 101

//StatusIllegalEncoding - Invalid encoding
const StatusIllegalEncoding = 102

//StatusIllegalURL - Invalid URL
const StatusIllegalURL = 104

//StatusIllegalTransactionStatus - Invalid transaction status
const StatusIllegalTransactionStatus = 105

//StatusExternalError - Failure at third party e.g. the bank
const StatusExternalError = 106

//StatusDeniedByBank - Transaction rejected by bank
const StatusDeniedByBank = 107

//StatusCancelled - Transaction cancelled
const StatusCancelled = 108

//StatusIllegalTransactionID - Invalid transaction ID
const StatusIllegalTransactionID = 110

//StatusMerchantNotConfigured - Merchant not configured
const StatusMerchantNotConfigured = 111

//StatusMerchantNotConfiguredAtBank - Merchant not configured at bank
const StatusMerchantNotConfiguredAtBank = 112

//StatusPaymentMethodNotConfigured - Payment method not configured for merchant
const StatusPaymentMethodNotConfigured = 113

//StatusTimeoutAtBank - Timeout at bank
const StatusTimeoutAtBank = 114

//StatusMerchantNotActive - The merchant is inactivated
const StatusMerchantNotActive = 115

//StatusPaymentMethodNotActive - The payment method is inactivated
const StatusPaymentMethodNotActive = 116

//StatusIllegalAuthorizedAmount - Amount cannot be authorized
const StatusIllegalAuthorizedAmount = 117

//StatusIllegalCapturedAmount - Amount cannot be captured
const StatusIllegalCapturedAmount = 118

//StatusIllegalCreditedAmount - Amount cannot be credited
const StatusIllegalCreditedAmount = 119

//StatusExceedsAmountLimit - Amount exceeds the limit
const StatusExceedsAmountLimit = 124

//StatusTransactionNotBelongingToMerchant - Transaction does not belong to merchant
const StatusTransactionNotBelongingToMerchant = 126

//StatusCustomerRefNoAlreadyUsed - Customer reference number already
//used in another transaction
const StatusCustomerRefNoAlreadyUsed = 127

//StatusNoSuchTrans - Transaction does not exist
const StatusNoSuchTrans = 128

//StatusDuplicateTransaction - More than one transaction found for
//the given customerrefno
const StatusDuplicateTransaction = 129

//StatusIllegalOperation - Operation not allowed for the given
//payment method
const StatusIllegalOperation = 130

//StatusCompanyNotActive - Company inactivated
const StatusCompanyNotActive = 131

//StatusSubscriptionNotActive - Subscription not active
const StatusSubscriptionNotActive = 133

//StatusSubscriptionNotSupported - Payment method doesn’t support subscriptions
const StatusSubscriptionNotSupported = 134

//StatusIllegalDateFormat - Invalid date format
const StatusIllegalDateFormat = 135

//StatusIllegalResponseData - Invalid response data
const StatusIllegalResponseData = 136

//StatusCurrencyNotConfigured - Currency not configured
const StatusCurrencyNotConfigured = 138

//StatusCurrencyNotActive - Currency not active
const StatusCurrencyNotActive = 139

//StatusCurrencyAlreadyConfigured - Currency is already configured
const StatusCurrencyAlreadyConfigured = 140

//StatusNoValidPaymentMethods - No valid payment methods
const StatusNoValidPaymentMethods = 142

//StatusCreditDeniedByBank - Credit denied by bank
const StatusCreditDeniedByBank = 143

//StatusIllegalCreditUsed - User is not allowed to perform credit operation
const StatusIllegalCreditUsed = 144

//StatusCustomerNotFound - Customer not found
const StatusCustomerNotFound = 146

//StatusAgeLimitExceeded - E.g. the transaction is too old
const StatusAgeLimitExceeded = 147

//StatusBrowserNotSupported - The browser version is too old
const StatusBrowserNotSupported = 148

//StatusPending - The request is still being processed
const StatusPending = 149

//StatusCreditPending - The credit could not be handled instantly.
//It is put in a queue it will be processed in due time.
const StatusCreditPending = 150

//StatusBadTransactionID - Invalid transaction ID
const StatusBadTransactionID = 301

//StatusBadMerchantID - Invalid merchant ID
const StatusBadMerchantID = 303

//StatusBadLang - Invalid language
const StatusBadLang = 304

//StatusBadAmount - Invalid amount
const StatusBadAmount = 305

//StatusBadCustomerRefNo - Invalid customerrefno
const StatusBadCustomerRefNo = 306

//StatusBadCurrency - Invalid currency
const StatusBadCurrency = 307

//StatusBadPaymentMethod - Invalid payment method
const StatusBadPaymentMethod = 308

//StatusBadReturnURL - Invalid returnurl
const StatusBadReturnURL = 309

//StatusBadMac - Invalid mac
const StatusBadMac = 311

//StatusBadCardNumberOrCardTypeNotConfigured - Card type not configured for merchant
const StatusBadCardNumberOrCardTypeNotConfigured = 316

//StatusBadSSN - Invalid ssn
const StatusBadSSN = 317

//StatusBadVat - Invalid vat
const StatusBadVat = 318

//StatusBadCaptureDate - Invalid capture date
const StatusBadCaptureDate = 319

//StatusBadCampaignCode - Invalid campaign code
const StatusBadCampaignCode = 320

//StatusBadSubscriptionType - Invalid subscription type
const StatusBadSubscriptionType = 321

//StatusBadSubscriptionID - Invalid subscription ID
const StatusBadSubscriptionID = 322

//StatusBadBase64 - Invalid Base64
const StatusBadBase64 = 323

//StatusBadCallbackURL - Invalid callbackurl
const StatusBadCallbackURL = 325

//Status3DCheckFailed - 3D check failed
const Status3DCheckFailed = 326

//StatusCardNotEnrolled - Card not enrolled in 3D Secure
const StatusCardNotEnrolled = 327

//StatusBadIPAddress - Provided IP address is invalid
const StatusBadIPAddress = 328

//StatusBadMobile - Invalid mobile phone number
const StatusBadMobile = 329

//StatusBadCountry - Invalid country
const StatusBadCountry = 330

//Status3DCheckNotAvailable - Merchant 3D configuration invalid
const Status3DCheckNotAvailable = 331

//StatusTimeout - Timeout at Svea
const StatusTimeout = 332

//StatusBadPeriod - The reconciliation period is invalid
const StatusBadPeriod = 333

//StatusBadAddressID - AddressSelector is not valid for this CountryCode
const StatusBadAddressID = 334

//StatusBadCustomerData - The supplied customer data is invalid
const StatusBadCustomerData = 335

//StatusBadUnit - Invalid unit
const StatusBadUnit = 336

//StatusBadExternalPaymentRef - Invalid external payment reference
const StatusBadExternalPaymentRef = 337

//StatusBadStoredCardAlias - Invalid stored card alias
const StatusBadStoredCardAlias = 338

//StatusStoredCardAliasNotActive - Stored card alias inactive
const StatusStoredCardAliasNotActive = 339

//StatusStoredCardsNotEnabled ...
const StatusStoredCardsNotEnabled = 340

//StatusOnlyDebitCardsAllowed ...
const StatusOnlyDebitCardsAllowed = 341

//StatusTransactionAlreadyInProgress ...
const StatusTransactionAlreadyInProgress = 342

//StatusBadCancelURL - Invalid Cancel url
const StatusBadCancelURL = 343

//StatusBadSimulatorCode - Invalid value for Simulator Code
const StatusBadSimulatorCode = 344

//StatusBadOrderRow - Invalid format for Order Row
const StatusBadOrderRow = 345

//StatusBadPayerAlias - Invalid format for Payer AliasPhone number
const StatusBadPayerAlias = 346

//StatusBadShowStoreCardDialog - Invalid value for Show Store Card Dialog
const StatusBadShowStoreCardDialog = 347

//StatusAntiFraudCardBinNotAllowed - Antifraud - cardbin not allowed
const StatusAntiFraudCardBinNotAllowed = 500

//StatusAntiFraudIPLocationNotAllowed - Antifraud – iplocation not allowed
const StatusAntiFraudIPLocationNotAllowed = 501

//StatusAntiFraudIPLocationAndBinDoesntMatch - Antifraud – ip-location and bin does not match
const StatusAntiFraudIPLocationAndBinDoesntMatch = 502

//StatusAntiFraudMaxAmountPerIPExceeded - Antofraud – max amount per ip exceeded
const StatusAntiFraudMaxAmountPerIPExceeded = 503

//StatusAntiFraudMaxTransactionsPerIPExceeded - Antifraud – max transactions per ip exceeded
const StatusAntiFraudMaxTransactionsPerIPExceeded = 504

//StatusAntiFraudMaxTransactionsPerCardNoExceeded - Antifraud – max transactions per card number exceeded
const StatusAntiFraudMaxTransactionsPerCardNoExceeded = 505

//StatusAntiFraudMaxAmountPerCardNoExceeded - Antifraud – max amount per cardnumer exceeded
const StatusAntiFraudMaxAmountPerCardNoExceeded = 506

//StatusAntiFraudIPAddressBlocked - Antifraud – IP address blocked.
const StatusAntiFraudIPAddressBlocked = 507

//StatusSwishNotEnrolled - Payer alias is invalid, or payee not enrolled
const StatusSwishNotEnrolled = 600

//StatusSwishRefundOrderNotFound - Original Payment not found or original payment is more than 13
//months old
const StatusSwishRefundOrderNotFound = 601

//StatusSwishRefundPayerError - Payer alias in the refund does not match the payee alias in the
//original payment
const StatusSwishRefundPayerError = 602

//StatusSwishRefundPayerOrgNrError - Payer organization number do not match original payment payee
//organization number
const StatusSwishRefundPayerOrgNrError = 603

//StatusSwishRefundPayeeSSNError - The Payer SSN in the original payment is not the same as the
//SSN for the current Payee
const StatusSwishRefundPayeeSSNError = 604
