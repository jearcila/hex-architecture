package constants

// References keys
const (
	IccRelatedData             = "icc_related_data"
	IsoResponseCode            = "iso_response_code"
	MerchantOperationReference = "merchant_operation_reference"
	ReconciliationTicketID     = "ticket_id"
)

const (
	AcquirerTransactionID     = "acquirer_transaction_id"
	DefaultDeliveryTime       = "360" // default delivery time in minutes
	GenovaIntegration         = "genova"
	_acqMastercardIntegration = "acq-mastercard"
)

const (
	CreditCard = "credit_card"
	DebitCard  = "debit_card"

	// Card present entry modes
	ManualEntryMode           = "0"
	SwipeEntryMode            = "1"
	ChipEntryMode             = "2"
	ContactlessSwipeEntryMode = "5"
	ContactlessEntryMode      = "6"

	// Recurring billing date layout
	BillingDateLayout = "2006-01-02"

	// Possible lengths for hard descriptor
	Len3  = 3
	Len7  = 7
	Len12 = 12
)
