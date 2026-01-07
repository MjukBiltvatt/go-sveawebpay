package sveawebpay

import (
	"errors"
	"fmt"
)

// ErrToCode returns the status code that corresponds to the specified error. The
// error must be one of the constants specified by this package or `-1` will be returned
func ErrToCode(err error) int {
	for c, e := range errMap {
		if err == e {
			return c
		}
	}

	return -1
}

// CodeToErr returns the error that corresponds to the specified status code. The
// status code must be one returned by the svea api or "unknown error: <code>" will be returned
func CodeToErr(code int) error {
	if code == 0 {
		return nil
	}

	if e, ok := errMap[code]; ok {
		return e
	}

	return fmt.Errorf("unknown error: %d", code)
}

// ErrRequiresManualReview should not be handled like any other error, the
// request was successful but still for some reason needs to be reviewed manually
var ErrRequiresManualReview = errors.New("request performed successfully but requires manual review by merchant")

// Possible errors returned by the svea api
var (
	ErrInternalError                             = errors.New("internal system error")
	ErrXMLParseFail                              = errors.New("invalid XML")
	ErrIllegalEncoding                           = errors.New("invalid encoding")
	ErrIllegalURL                                = errors.New("invalid URL")
	ErrIllegalTransactionStatus                  = errors.New("invalid transaction status")
	ErrExternalError                             = errors.New("failure at third party e.g. the bank")
	ErrDeniedByBank                              = errors.New("transaction rejected by bank")
	ErrCancelled                                 = errors.New("transaction cancelled")
	ErrIllegalTransactionID                      = errors.New("invalid transaction ID")
	ErrMerchantNotConfigured                     = errors.New("merchant not configured")
	ErrMerchantNotConfiguredAtBank               = errors.New("merchant not configured at bank")
	ErrPaymentMethodNotConfigured                = errors.New("payment method not configured for merchant")
	ErrTimeoutAtBank                             = errors.New("timeout at bank")
	ErrMerchantNotActive                         = errors.New("the merchant is inactivated")
	ErrPaymentMethodNotActive                    = errors.New("the payment method is inactivated")
	ErrIllegalAuthorizedAmount                   = errors.New("amount cannot be authorized")
	ErrIllegalCapturedAmount                     = errors.New("amount cannot be captured")
	ErrIllegalCreditedAmount                     = errors.New("amount cannot be credited")
	ErrExceedsAmountLimit                        = errors.New("amount exceeds the limit")
	ErrTransactionNotBelongingToMerchant         = errors.New("transaction does not belong to merchant")
	ErrCustomerRefNoAlreadyUsed                  = errors.New("customer reference number already used in another transaction")
	ErrNoSuchTrans                               = errors.New("transaction does not exist")
	ErrDuplicateTransaction                      = errors.New("more than one transaction found for the given customerrefno")
	ErrIllegalOperation                          = errors.New("operation not allowed for the given payment method")
	ErrCompanyNotActive                          = errors.New("company inactivated")
	ErrSubscriptionNotActive                     = errors.New("subscription not active")
	ErrSubscriptionNotSupported                  = errors.New("payment method doesn’t support subscriptions")
	ErrIllegalDateFormat                         = errors.New("invalid date format")
	ErrIllegalResponseData                       = errors.New("invalid response data")
	ErrCurrencyNotConfigured                     = errors.New("currency not configured")
	ErrCurrencyNotActive                         = errors.New("currency not active")
	ErrCurrencyAlreadyConfigured                 = errors.New("currency is already configured")
	ErrNoValidPaymentMethods                     = errors.New("no valid payment methods")
	ErrCreditDeniedByBank                        = errors.New("credit denied by bank")
	ErrIllegalCreditUsed                         = errors.New("user is not allowed to perform credit operation")
	ErrCustomerNotFound                          = errors.New("customer not found")
	ErrAgeLimitExceeded                          = errors.New("age limit exceeded, e.g. the transaction is too old")
	ErrBrowserNotSupported                       = errors.New("the browser version is too old")
	ErrPending                                   = errors.New("the request is still being processed")
	ErrCreditPending                             = errors.New("the credit could not be handled instantly. It is put in a queue and it will be processed in due time")
	ErrCardClosed                                = errors.New("card is closed")
	ErrBadTransactionID                          = errors.New("invalid transaction ID")
	ErrBadMerchantID                             = errors.New("invalid merchant ID")
	ErrBadLang                                   = errors.New("invalid language")
	ErrBadAmount                                 = errors.New("invalid amount")
	ErrBadCustomerRefNo                          = errors.New("invalid customerrefno")
	ErrBadCurrency                               = errors.New("invalid currency")
	ErrBadPaymentMethod                          = errors.New("invalid payment method")
	ErrBadReturnURL                              = errors.New("invalid returnurl")
	ErrBadMac                                    = errors.New("invalid mac")
	ErrBadCardNumberOrCardTypeNotConfigured      = errors.New("card type not configured for merchant")
	ErrBadSSN                                    = errors.New("invalid ssn")
	ErrBadVat                                    = errors.New("invalid vat")
	ErrBadCaptureDate                            = errors.New("invalid capture date")
	ErrBadCampaignCode                           = errors.New("invalid campaign code")
	ErrBadSubscriptionType                       = errors.New("invalid subscription type")
	ErrBadSubscriptionID                         = errors.New("invalid subscription ID")
	ErrBadBase64                                 = errors.New("invalid Base")
	ErrBadCallbackURL                            = errors.New("invalid callbackurl")
	Err3DCheckFailed                             = errors.New("3D check failed")
	ErrCardNotEnrolled                           = errors.New("card not enrolled in 3D Secure")
	ErrBadIPAddress                              = errors.New("provided IP address is invalid")
	ErrBadMobile                                 = errors.New("invalid mobile phone number")
	ErrBadCountry                                = errors.New("invalid country")
	Err3DCheckNotAvailable                       = errors.New("merchant 3D configuration invalid")
	ErrTimeout                                   = errors.New("timeout at Svea")
	ErrBadPeriod                                 = errors.New("the reconciliation period is invalid")
	ErrBadAddressID                              = errors.New("addressSelector is not valid for this CountryCode")
	ErrBadCustomerData                           = errors.New("the supplied customer data is invalid")
	ErrBadUnit                                   = errors.New("invalid unit")
	ErrBadExternalPaymentRef                     = errors.New("invalid external payment reference")
	ErrBadStoredCardAlias                        = errors.New("invalid stored card alias")
	ErrStoredCardAliasNotActive                  = errors.New("stored card alias inactive")
	ErrStoredCardsNotEnabled                     = errors.New("stored cards not enabled")
	ErrOnlyDebitCardsAllowed                     = errors.New("only debitcards allowed")
	ErrTransactionAlreadyInProgress              = errors.New("transaction is already in progress")
	ErrBadCancelURL                              = errors.New("invalid Cancel url")
	ErrBadSimulatorCode                          = errors.New("invalid value for Simulator Code")
	ErrBadOrderRow                               = errors.New("invalid format for Order Row")
	ErrBadPayerAlias                             = errors.New("invalid format for Payer AliasPhone number")
	ErrBadShowStoreCardDialog                    = errors.New("invalid value for Show Store Card Dialog")
	ErrServiceUnavailableTryLater                = errors.New("temporary problem, try later")
	ErrAntiFraudCardBinNotAllowed                = errors.New("antifraud - cardbin not allowed")
	ErrAntiFraudIPLocationNotAllowed             = errors.New("antifraud – iplocation not allowed")
	ErrAntiFraudIPLocationAndBinDoesntMatch      = errors.New("antifraud – ip-location and bin does not match")
	ErrAntiFraudMaxAmountPerIPExceeded           = errors.New("antifraud – max amount per ip exceeded")
	ErrAntiFraudMaxTransactionsPerIPExceeded     = errors.New("antifraud – max transactions per ip exceeded")
	ErrAntiFraudMaxTransactionsPerCardNoExceeded = errors.New("antifraud – max transactions per card number exceeded")
	ErrAntiFraudMaxAmountPerCardNoExceeded       = errors.New("antifraud – max amount per cardnumer exceeded")
	ErrAntiFraudIPAddressBlocked                 = errors.New("antifraud – IP address blocked")
	ErrSwishNotEnrolled                          = errors.New("payer alias is invalid, or payee not enrolled")
	ErrSwishRefundOrderNotFound                  = errors.New("original Payment not found or original payment is more than months old")
	ErrSwishRefundPayerError                     = errors.New("payer alias in the refund does not match the payee alias in the original payment")
	ErrSwishRefundPayerOrgNrError                = errors.New("payer organization number do not match original payment payee organization number")
	ErrSwishRefundPayeeSSNError                  = errors.New("the Payer SSN in the original payment is not the same as the SSN for the current Payee")
)

// Map status codes to errors
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
	152: ErrCardClosed,
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
	373: ErrServiceUnavailableTryLater,
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
