package sveawebpay

import "errors"

//ErrToCode returns the status code that corresponds to the specified error. The
//error must be one of the constants specified by this package or `-1` will be returned
func ErrToCode(err error) int {
	for c, e := range errMap {
		if err == e {
			return c
		}
	}

	return -1
}

//CodeToErr returns the error that corresponds to the specified status code. The
//status code must be one returned by the svea api or `ErrUnknown` will be returned
func CodeToErr(code int) error {
	for c, e := range errMap {
		if code == c {
			return e
		}
	}

	return ErrUnknown
}

var (
	//ErrRequiresManualReview should not be handled like any other error, the
	//request was successful but still for some reason needs to be reviewed manually
	ErrRequiresManualReview = errors.New("request performed successfully but requires manual review by merchant")

	//ErrInternalError - Internal system error
	ErrInternalError = errors.New("internal system error")

	//ErrXMLParseFail - Invalid XML
	ErrXMLParseFail = errors.New("invalid XML")

	//ErrIllegalEncoding - Invalid encoding
	ErrIllegalEncoding = errors.New("invalid encoding")

	//ErrIllegalURL - Invalid URL
	ErrIllegalURL = errors.New("invalid URL")

	//ErrIllegalTransactionStatus - Invalid transaction status
	ErrIllegalTransactionStatus = errors.New("invalid transaction status")

	//ErrExternalError - Failure at third party e.g. the bank
	ErrExternalError = errors.New("failure at third party e.g. the bank")

	//ErrDeniedByBank - Transaction rejected by bank
	ErrDeniedByBank = errors.New("transaction rejected by bank")

	//ErrCancelled - Transaction cancelled
	ErrCancelled = errors.New("transaction cancelled")

	//ErrIllegalTransactionID - Invalid transaction ID
	ErrIllegalTransactionID = errors.New("invalid transaction ID")

	//ErrMerchantNotConfigured - Merchant not configured
	ErrMerchantNotConfigured = errors.New("merchant not configured")

	//ErrMerchantNotConfiguredAtBank - Merchant not configured at bank
	ErrMerchantNotConfiguredAtBank = errors.New("merchant not configured at bank")

	//ErrPaymentMethodNotConfigured - Payment method not configured for merchant
	ErrPaymentMethodNotConfigured = errors.New("payment method not configured for merchant")

	//ErrTimeoutAtBank - Timeout at bank
	ErrTimeoutAtBank = errors.New("timeout at bank")

	//ErrMerchantNotActive - The merchant is inactivated
	ErrMerchantNotActive = errors.New("the merchant is inactivated")

	//ErrPaymentMethodNotActive - The payment method is inactivated
	ErrPaymentMethodNotActive = errors.New("the payment method is inactivated")

	//ErrIllegalAuthorizedAmount - Amount cannot be authorized
	ErrIllegalAuthorizedAmount = errors.New("amount cannot be authorized")

	//ErrIllegalCapturedAmount - Amount cannot be captured
	ErrIllegalCapturedAmount = errors.New("amount cannot be captured")

	//ErrIllegalCreditedAmount - Amount cannot be credited
	ErrIllegalCreditedAmount = errors.New("amount cannot be credited")

	//ErrExceedsAmountLimit - Amount exceeds the limit
	ErrExceedsAmountLimit = errors.New("amount exceeds the limit")

	//ErrTransactionNotBelongingToMerchant - Transaction does not belong to merchant
	ErrTransactionNotBelongingToMerchant = errors.New("transaction does not belong to merchant")

	//ErrCustomerRefNoAlreadyUsed - Customer reference number already
	//used in another transaction
	ErrCustomerRefNoAlreadyUsed = errors.New("customer reference number already used in another transaction")

	//ErrNoSuchTrans - Transaction does not exist
	ErrNoSuchTrans = errors.New("transaction does not exist")

	//ErrDuplicateTransaction - More than one transaction found for
	//the given customerrefno
	ErrDuplicateTransaction = errors.New("more than one transaction found for the given customerrefno")

	//ErrIllegalOperation - Operation not allowed for the given payment method
	ErrIllegalOperation = errors.New("operation not allowed for the given payment method")

	//ErrCompanyNotActive - Company inactivated
	ErrCompanyNotActive = errors.New("company inactivated")

	//ErrSubscriptionNotActive - Subscription not active
	ErrSubscriptionNotActive = errors.New("subscription not active")

	//ErrSubscriptionNotSupported - Payment method doesn’t support subscriptions
	ErrSubscriptionNotSupported = errors.New("payment method doesn’t support subscriptions")

	//ErrIllegalDateFormat - Invalid date format
	ErrIllegalDateFormat = errors.New("invalid date format")

	//ErrIllegalResponseData - Invalid response data
	ErrIllegalResponseData = errors.New("invalid response data")

	//ErrCurrencyNotConfigured - Currency not configured
	ErrCurrencyNotConfigured = errors.New("currency not configured")

	//ErrCurrencyNotActive - Currency not active
	ErrCurrencyNotActive = errors.New("currency not active")

	//ErrCurrencyAlreadyConfigured - Currency is already configured
	ErrCurrencyAlreadyConfigured = errors.New("currency is already configured")

	//ErrNoValidPaymentMethods - No valid payment methods
	ErrNoValidPaymentMethods = errors.New("no valid payment methods")

	//ErrCreditDeniedByBank - Credit denied by bank
	ErrCreditDeniedByBank = errors.New("credit denied by bank")

	//ErrIllegalCreditUsed - User is not allowed to perform credit operation
	ErrIllegalCreditUsed = errors.New("user is not allowed to perform credit operation")

	//ErrCustomerNotFound - Customer not found
	ErrCustomerNotFound = errors.New("customer not found")

	//ErrAgeLimitExceeded - E.g. the transaction is too old
	ErrAgeLimitExceeded = errors.New("age limit exceeded, e.g. the transaction is too old")

	//ErrBrowserNotSupported - The browser version is too old
	ErrBrowserNotSupported = errors.New("the browser version is too old")

	//ErrPending - The request is still being processed
	ErrPending = errors.New("the request is still being processed")

	//ErrCreditPending - The credit could not be handled instantly.
	//It is put in a queue and it will be processed in due time.
	ErrCreditPending = errors.New("the credit could not be handled instantly. It is put in a queue and it will be processed in due time")

	//ErrBadTransactionID - Invalid transaction ID
	ErrBadTransactionID = errors.New("invalid transaction ID")

	//ErrBadMerchantID - Invalid merchant ID
	ErrBadMerchantID = errors.New("invalid merchant ID")

	//ErrBadLang - Invalid language
	ErrBadLang = errors.New("invalid language")

	//ErrBadAmount - Invalid amount
	ErrBadAmount = errors.New("invalid amount")

	//ErrBadCustomerRefNo - Invalid customerrefno
	ErrBadCustomerRefNo = errors.New("invalid customerrefno")

	//ErrBadCurrency - Invalid currency
	ErrBadCurrency = errors.New("invalid currency")

	//ErrBadPaymentMethod - Invalid payment method
	ErrBadPaymentMethod = errors.New("invalid payment method")

	//ErrBadReturnURL - Invalid returnurl
	ErrBadReturnURL = errors.New("invalid returnurl")

	//ErrBadMac - Invalid mac
	ErrBadMac = errors.New("invalid mac")

	//ErrBadCardNumberOrCardTypeNotConfigured - Card type not configured for merchant
	ErrBadCardNumberOrCardTypeNotConfigured = errors.New("card type not configured for merchant")

	//ErrBadSSN - Invalid ssn
	ErrBadSSN = errors.New("invalid ssn")

	//ErrBadVat - Invalid vat
	ErrBadVat = errors.New("invalid vat")

	//ErrBadCaptureDate - Invalid capture date
	ErrBadCaptureDate = errors.New("invalid capture date")

	//ErrBadCampaignCode - Invalid campaign code
	ErrBadCampaignCode = errors.New("invalid campaign code")

	//ErrBadSubscriptionType - Invalid subscription type
	ErrBadSubscriptionType = errors.New("invalid subscription type")

	//ErrBadSubscriptionID - Invalid subscription ID
	ErrBadSubscriptionID = errors.New("invalid subscription ID")

	//ErrBadBase64 - Invalid Base
	ErrBadBase64 = errors.New("invalid Base")

	//ErrBadCallbackURL - Invalid callbackurl
	ErrBadCallbackURL = errors.New("invalid callbackurl")

	//Err3DCheckFailed - 3D check failed
	Err3DCheckFailed = errors.New("3D check failed")

	//ErrCardNotEnrolled - Card not enrolled in 3D Secure
	ErrCardNotEnrolled = errors.New("card not enrolled in 3D Secure")

	//ErrBadIPAddress - Provided IP address is invalid
	ErrBadIPAddress = errors.New("provided IP address is invalid")

	//ErrBadMobile - Invalid mobile phone number
	ErrBadMobile = errors.New("invalid mobile phone number")

	//ErrBadCountry - Invalid country
	ErrBadCountry = errors.New("invalid country")

	//Err3DCheckNotAvailable - Merchant 3D configuration invalid
	Err3DCheckNotAvailable = errors.New("merchant 3D configuration invalid")

	//ErrTimeout - Timeout at Svea
	ErrTimeout = errors.New("timeout at Svea")

	//ErrBadPeriod - The reconciliation period is invalid
	ErrBadPeriod = errors.New("the reconciliation period is invalid")

	//ErrBadAddressID - AddressSelector is not valid for this CountryCode
	ErrBadAddressID = errors.New("addressSelector is not valid for this CountryCode")

	//ErrBadCustomerData - The supplied customer data is invalid
	ErrBadCustomerData = errors.New("the supplied customer data is invalid")

	//ErrBadUnit - Invalid unit
	ErrBadUnit = errors.New("invalid unit")

	//ErrBadExternalPaymentRef - Invalid external payment reference
	ErrBadExternalPaymentRef = errors.New("invalid external payment reference")

	//ErrBadStoredCardAlias - Invalid stored card alias
	ErrBadStoredCardAlias = errors.New("invalid stored card alias")

	//ErrStoredCardAliasNotActive - Stored card alias inactive
	ErrStoredCardAliasNotActive = errors.New("stored card alias inactive")

	//ErrStoredCardsNotEnabled ...
	ErrStoredCardsNotEnabled = errors.New("stored cards not enabled")

	//ErrOnlyDebitCardsAllowed ...
	ErrOnlyDebitCardsAllowed = errors.New("only debitcards allowed")

	//ErrTransactionAlreadyInProgress ...
	ErrTransactionAlreadyInProgress = errors.New("transaction is already in progress")

	//ErrBadCancelURL - Invalid Cancel url
	ErrBadCancelURL = errors.New("invalid Cancel url")

	//ErrBadSimulatorCode - Invalid value for Simulator Code
	ErrBadSimulatorCode = errors.New("invalid value for Simulator Code")

	//ErrBadOrderRow - Invalid format for Order Row
	ErrBadOrderRow = errors.New("invalid format for Order Row")

	//ErrBadPayerAlias - Invalid format for Payer AliasPhone number
	ErrBadPayerAlias = errors.New("invalid format for Payer AliasPhone number")

	//ErrBadShowStoreCardDialog - Invalid value for Show Store Card Dialog
	ErrBadShowStoreCardDialog = errors.New("invalid value for Show Store Card Dialog")

	//ErrAntiFraudCardBinNotAllowed - Antifraud - cardbin not allowed
	ErrAntiFraudCardBinNotAllowed = errors.New("antifraud - cardbin not allowed")

	//ErrAntiFraudIPLocationNotAllowed - Antifraud – iplocation not allowed
	ErrAntiFraudIPLocationNotAllowed = errors.New("antifraud – iplocation not allowed")

	//ErrAntiFraudIPLocationAndBinDoesntMatch - Antifraud – ip-location and bin does not match
	ErrAntiFraudIPLocationAndBinDoesntMatch = errors.New("antifraud – ip-location and bin does not match")

	//ErrAntiFraudMaxAmountPerIPExceeded - Antifraud – max amount per ip exceeded
	ErrAntiFraudMaxAmountPerIPExceeded = errors.New("antifraud – max amount per ip exceeded")

	//ErrAntiFraudMaxTransactionsPerIPExceeded - Antifraud – max transactions per ip exceeded
	ErrAntiFraudMaxTransactionsPerIPExceeded = errors.New("antifraud – max transactions per ip exceeded")

	//ErrAntiFraudMaxTransactionsPerCardNoExceeded - Antifraud – max transactions per card number exceeded
	ErrAntiFraudMaxTransactionsPerCardNoExceeded = errors.New("antifraud – max transactions per card number exceeded")

	//ErrAntiFraudMaxAmountPerCardNoExceeded - Antifraud – max amount per cardnumer exceeded
	ErrAntiFraudMaxAmountPerCardNoExceeded = errors.New("antifraud – max amount per cardnumer exceeded")

	//ErrAntiFraudIPAddressBlocked - Antifraud – IP address blocked.
	ErrAntiFraudIPAddressBlocked = errors.New("antifraud – IP address blocked")

	//ErrSwishNotEnrolled - Payer alias is invalid, or payee not enrolled
	ErrSwishNotEnrolled = errors.New("payer alias is invalid, or payee not enrolled")

	//ErrSwishRefundOrderNotFound - Original Payment not found or original payment is more than
	//months old
	ErrSwishRefundOrderNotFound = errors.New("original Payment not found or original payment is more than months old")

	//ErrSwishRefundPayerError - Payer alias in the refund does not match the payee alias in the
	//original payment
	ErrSwishRefundPayerError = errors.New("payer alias in the refund does not match the payee alias in the original payment")

	//ErrSwishRefundPayerOrgNrError - Payer organization number do not match original payment payee
	//organization number
	ErrSwishRefundPayerOrgNrError = errors.New("payer organization number do not match original payment payee organization number")

	//ErrSwishRefundPayeeSSNError - The Payer SSN in the original payment is not the same as the
	//SSN for the current Payee
	ErrSwishRefundPayeeSSNError = errors.New("the Payer SSN in the original payment is not the same as the SSN for the current Payee")

	//ErrUnknown is given when a status code can not be matched with an error
	ErrUnknown = errors.New("unknown error")
)

var errMap = map[int]error{
	1:   ErrRequiresManualReview,
	100: ErrInternalError,
	101: ErrXMLParseFail,
	102: ErrIllegalEncoding,
	104: ErrIllegalURL,
	105: ErrIllegalTransactionStatus,
	106: ErrExternalError,
	107: ErrDeniedByBank,
	108: ErrCancelled,
	110: ErrIllegalTransactionID,
	111: ErrMerchantNotConfigured,
	112: ErrMerchantNotConfiguredAtBank,
	113: ErrPaymentMethodNotConfigured,
	114: ErrTimeoutAtBank,
	115: ErrMerchantNotActive,
	116: ErrPaymentMethodNotActive,
	117: ErrIllegalAuthorizedAmount,
	118: ErrIllegalCapturedAmount,
	119: ErrIllegalCreditedAmount,
	124: ErrExceedsAmountLimit,
	126: ErrTransactionNotBelongingToMerchant,
	127: ErrCustomerRefNoAlreadyUsed,
	128: ErrNoSuchTrans,
	129: ErrDuplicateTransaction,
	130: ErrIllegalOperation,
	131: ErrCompanyNotActive,
	133: ErrSubscriptionNotActive,
	134: ErrSubscriptionNotSupported,
	135: ErrIllegalDateFormat,
	136: ErrIllegalResponseData,
	138: ErrCurrencyNotConfigured,
	139: ErrCurrencyNotActive,
	140: ErrCurrencyAlreadyConfigured,
	142: ErrNoValidPaymentMethods,
	143: ErrCreditDeniedByBank,
	144: ErrIllegalCreditUsed,
	146: ErrCustomerNotFound,
	147: ErrAgeLimitExceeded,
	148: ErrBrowserNotSupported,
	149: ErrPending,
	150: ErrCreditPending,
	301: ErrBadTransactionID,
	303: ErrBadMerchantID,
	304: ErrBadLang,
	305: ErrBadAmount,
	306: ErrBadCustomerRefNo,
	307: ErrBadCurrency,
	308: ErrBadPaymentMethod,
	309: ErrBadReturnURL,
	311: ErrBadMac,
	316: ErrBadCardNumberOrCardTypeNotConfigured,
	317: ErrBadSSN,
	318: ErrBadVat,
	319: ErrBadCaptureDate,
	320: ErrBadCampaignCode,
	321: ErrBadSubscriptionType,
	322: ErrBadSubscriptionID,
	323: ErrBadBase64,
	325: ErrBadCallbackURL,
	326: Err3DCheckFailed,
	327: ErrCardNotEnrolled,
	328: ErrBadIPAddress,
	329: ErrBadMobile,
	330: ErrBadCountry,
	331: Err3DCheckNotAvailable,
	332: ErrTimeout,
	333: ErrBadPeriod,
	334: ErrBadAddressID,
	335: ErrBadCustomerData,
	336: ErrBadUnit,
	337: ErrBadExternalPaymentRef,
	338: ErrBadStoredCardAlias,
	339: ErrStoredCardAliasNotActive,
	340: ErrStoredCardsNotEnabled,
	341: ErrOnlyDebitCardsAllowed,
	342: ErrTransactionAlreadyInProgress,
	343: ErrBadCancelURL,
	344: ErrBadSimulatorCode,
	345: ErrBadOrderRow,
	346: ErrBadPayerAlias,
	347: ErrBadShowStoreCardDialog,
	500: ErrAntiFraudCardBinNotAllowed,
	501: ErrAntiFraudIPLocationNotAllowed,
	502: ErrAntiFraudIPLocationAndBinDoesntMatch,
	503: ErrAntiFraudMaxAmountPerIPExceeded,
	504: ErrAntiFraudMaxTransactionsPerIPExceeded,
	505: ErrAntiFraudMaxTransactionsPerCardNoExceeded,
	506: ErrAntiFraudMaxAmountPerCardNoExceeded,
	507: ErrAntiFraudIPAddressBlocked,
	600: ErrSwishNotEnrolled,
	601: ErrSwishRefundOrderNotFound,
	602: ErrSwishRefundPayerError,
	603: ErrSwishRefundPayerOrgNrError,
	604: ErrSwishRefundPayeeSSNError,
}
