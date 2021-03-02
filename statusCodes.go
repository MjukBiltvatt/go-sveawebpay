package sveawebpay

//StatusSuccess request performed successfully
const StatusSuccess = 0

//StatusRequiresManualReview request performed successfully but
//requires manual review by merchant
const StatusRequiresManualReview = 1

//StatusInternalError internal system error
const StatusInternalError = 100

//StatusXMLParseFail invalid XML
const StatusXMLParseFail = 101

const StatusIllegalEncoding = 102

const StatusIllegalURL = 104

const StatusIllegalTransactionStatus = 105

const StatusExternalError = 106

const StatusDeniedByBank = 107

const StatusCancelled = 108

const StatusIllegalTransactionID = 110

const StatusMerchantNotConfigured = 111

const StatusMerchantNotConfiguredAtBank = 112

const StatusPaymentMethodNotConfigured = 113

const StatusTimeoutAtBank = 114

const StatusMerchantNotActive = 115

const StatusPaymentMethodNotActive = 116

const StatusIllegalAuthorizedAmount = 117

const StatusIllegalCapturedAmount = 118

const StatusIllegalCreditedAmount = 119

const StatusExceedsAmountLimit = 124

const StatusTransactionNotBelongingToMerchant = 126

const StatusCustomerRefNoAlreadyUsed = 127

const StatusNoSuchTrans = 128

const StatusDuplicateTransaction = 129

const StatusIllegalOperation = 130

const StatusCompanyNotActive = 131

const StatusSubscriptionNotActive = 133

const StatusSubscriptionNotSupported = 134

const StatusIllegalDateFormat = 135

const StatusIllegalResponseData = 136

const StatusCurrencyNotConfigured = 138

const StatusCurrencyNotActive = 139

const StatusCurrencyAlreadyConfigured = 140

const StatusNoValidPaymentMethods = 142

const StatusCreditDeniedByBank = 143

const StatusIllegalCreditUsed = 144

const StatusCustomerNotFound = 146

const StatusAgeLimitExceeded = 147

const StatusBrowserNotSupported = 148

const StatusPending = 149

const StatusCreditPending = 150

const StatusBadTransactionID = 301

const StatusBadMerchantID = 303

const StatusBadLang = 304

const StatusBadAmount = 305

const StatusBadCustomerRefNo = 306

const StatusBadCurrency = 307

const StatusBadPaymentMethod = 308

const StatusBadReturnURL = 309

const StatusBadMac = 311

const StatusBadCardNumberOrCardTypeNotConfigured = 316

const StatusBadSSN = 317

const StatusBadVat = 318

const StatusBadCaptureDate = 319

const StatusBadCampaignCode = 320

const StatusBadSubscriptionType = 321

const StatusBadSubscriptionID = 322

const StatusBadBase64 = 323

const StatusBadCallbackURL = 325

const Status3DCheckFailed = 326

const StatusCardNotEnrolled = 327

const StatusBadIPAddress = 328

const StatusBadMobile = 329

const StatusBadCountry = 330

const Status3DCheckNotAvailable = 331

const StatusTimeout = 332

const StatusBadPeriod = 333

const StatusBadAddressID = 334

const StatusBadCustomerData = 335

const StatusBadUnit = 336

const StatusBadExternalPaymentRef = 337

const StatusBadStoredCardAlias = 338

const StatusStoredCardAliasNotActive = 339

const StatusStoredCardsNotEnabled = 340

const StatusOnlyDebitCardsAllowed = 341

const StatusTransactionAlreadyInProgress = 342

const StatusBadCancelURL = 343

const StatusBadSimulatorCode = 344

const StatusBadOrderRow = 345

const StatusBadPayerAlias = 346

const StatusBadShowStoreCardDialog = 347

const StatusAntiFraudCardBinNotAllowed = 500

const StatusAntiFraudIPLocationNotAllowed = 501

const StatusAntiFraudIPLocationAndBinDoesntMatch = 502

const StatusAntiFraudMaxAmountPerIPExceeded = 503

const StatusAntiFraudMaxTransactionsPerIPExceeded = 504

const StatusAntiFraudMaxTransactionsPerCardNoExceeded = 505

const StatusAntiFraudMaxAmountPerCardNoExceeded = 506

const StatusAntiFraudIPAddressBlocked = 507

const StatusSwishNotEnrolled = 600

const StatusSwishRefundOrderNotFound = 601

const StatusSwishRefundPayerError = 602

const StatusSwishRefundPayerOrgNrError = 603

const StatusSwishRefundPayeeSSNError = 604
