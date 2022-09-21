package model

import (
	"github.com/shopspring/decimal"
)

type FirstOperationRequest struct {
	MerchantOperationReference string              `json:"merchant_operation_reference"`
	ProcessorID                string              `json:"processor_id"`
	MerchantID                 string              `json:"merchant_id"`
	Card                       Card                `json:"card"`
	CardProcessingMode         string              `json:"card_processing_mode"`
	Amount                     Amount              `json:"amount"`
	Installments               uint32              `json:"installments"`
	SoftDescriptor             string              `json:"soft_descriptor"`
	HardDescriptor             string              `json:"hard_descriptor"`
	MCC                        string              `json:"mcc"`
	SubMerchant                SubMerchant         `json:"sub_merchant"`
	PointOfInteraction         *PointOfInteraction `json:"point_of_interaction"`
	Authentication             *Authentication     `json:"authentication"`
	Recurring                  *Recurring          `json:"recurring"`
	InstallmentPlanType        string              `json:"installment_plan_type"`
	WalletID                   string              `json:"wallet_id"`
}

type Amount struct {
	Amount   decimal.Decimal `json:"amount"`
	Currency string          `json:"currency"`
}

type SubMerchant struct {
	ID              string   `json:"id"`
	LegalName       string   `json:"legal_name"`
	FiscalCondition string   `json:"fiscal_condition"`
	TaxID           TaxID    `json:"tax_id"`
	Location        Location `json:"location"`
}

type TaxID struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type Location struct {
	Address           string `json:"address"`
	AddressDoorNumber uint64 `json:"address_door_number"`
	City              string `json:"city"`
	ZipCode           string `json:"zip_code"`
	Region            string `json:"region"`
	CountryCode       string `json:"country_code"`
	Latitude          string `json:"latitude,omitempty"`
	Longitude         string `json:"longitude,omitempty"`
}

type Card struct {
	NumberID         string            `json:"number_id"`
	ExpirationMonth  uint32            `json:"expiration_month"`
	ExpirationYear   uint32            `json:"expiration_year"`
	Holder           Holder            `json:"holder"`
	CredentialOnFile bool              `json:"credential_on_file"`
	SecurityCodeID   *string           `json:"security_code_id"`
	EntryModeDetails *EntryModeDetails `json:"entry_mode_details"`
	Tokenization     *TokenizationInfo `json:"tokenization"`
}

type Holder struct {
	Name string   `json:"name"`
	ID   HolderID `json:"id"`
}

type HolderID struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

type PointOfInteraction struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Signature string `json:"signature"`
}

type TokenizationInfo struct {
	ExpirationMonth uint32 `json:"expiration_month"`
	ExpirationYear  uint32 `json:"expiration_year"`
	DPANID          string `json:"dpan_id"`
	Cryptogram      string `json:"cryptogram"`
	CryptogramID    string `json:"cryptogram_id"`
}

type EntryModeDetails struct {
	CardPresentID     string `json:"card_present_id"`
	EntryMode         string `json:"entry_mode"`
	PINBlock          bool   `json:"pin_block"`
	FallbackIndicator bool   `json:"fallback_indicator"`
	ICCRelatedData    string `json:"icc_related_data"`
	ICCSequenceNumber string `json:"icc_sequence_number"`
}

type Authentication struct {
	ThreeDS *ThreeDS `json:"three_ds"`
}

type ThreeDS struct {
	Cryptogram           string `json:"cryptogram"`
	ServerTransID        string `json:"server_trans_id"`
	ThreeDSServerTransID string `json:"three_ds_server_trans_id"`
	ACSReferenceNumber   string `json:"acs_reference_number"`
	Eci                  string `json:"eci"`
	DSTransID            string `json:"ds_trans_id"`
	ACSTransID           string `json:"acs_trans_id"`
	ThreeDSVersion       string `json:"three_ds_version"`
}

type FirstOperationResponse struct {
	AcquirerTransactionID string `json:"acquirer_transaction_id"`
	ResponseCode          string `json:"response_code"`
	ResponseMessage       string `json:"response_message"`
	AuthorizationCode     string `json:"authorization_code,omitempty"`
	ICCRelatedData        string `json:"icc_related_data,omitempty"`
}

type CaptureRequest struct {
	AcquirerTransactionID      string `json:"acquirer_transaction_id"`
	MerchantOperationReference string `json:"merchant_operation_reference"`
	Amount                     Amount `json:"amount"`
}

type CaptureResponse struct {
	AcquirerTransactionID string `json:"acquirer_transaction_id"`
	ResponseCode          string `json:"response_code"`
	ResponseMessage       string `json:"response_message"`
	AuthorizationCode     string `json:"authorization_code,omitempty"`
}

type CancelRequest struct {
	AcquirerTransactionID      string `json:"acquirer_transaction_id"`
	MerchantOperationReference string `json:"merchant_operation_reference"`
	Amount                     Amount `json:"amount"`
}

type CancelResponse struct {
	AcquirerTransactionID string `json:"acquirer_transaction_id"`
	ResponseCode          string `json:"response_code"`
	ResponseMessage       string `json:"response_message"`
	AuthorizationCode     string `json:"authorization_code,omitempty"`
}

type Recurring struct {
	ID                 string `json:"id"`
	FirstPayment       bool   `json:"first_payment"`
	InvoicePeriodMonth int    `json:"invoice_period_month"`
	InvoicePeriodYear  int    `json:"invoice_period_year"`
}
