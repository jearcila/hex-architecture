package model

import "reflect"

type Options struct {
	Regulation   *Regulation   `json:"regulation"`
	Tokenization *Tokenization `json:"tokenization"`
	ThreeDS      *ThreeDS      `json:"three_ds"`
	Subscription *Subscription `json:"subscription"`
	CollectorID  uint64        `json:"collector_id"`
	Plan         *Plan         `json:"plan"`
	WalletID     string        `json:"wallet_id"`
}

type Plan struct {
	ID string `json:"id"`
}

type Regulation struct {
	MCC               *string `json:"mcc"`
	LegalName         string  `json:"legal_name"`
	ZIPCode           string  `json:"zip"`
	City              string  `json:"city"`
	Country           string  `json:"country,omitempty"`
	AddressStreet     string  `json:"address_street,omitempty"`
	AddressDoorNumber uint64  `json:"address_door_number,omitempty"`
	RegionCode        string  `json:"region_code,omitempty"`
	RegionCodeIso     string  `json:"region_code_iso,omitempty"`
	DocumentNumber    string  `json:"document_number,omitempty"`
	DocumentType      string  `json:"document_type,omitempty"`
	FiscalCondition   string  `json:"fiscal_condition,omitempty"`
}

func (tds ThreeDS) IsEmpty() bool {
	return reflect.DeepEqual(tds, ThreeDS{})
}

type Tokenization struct {
	ExpirationMonth uint32 `json:"expiration_month"`
	ExpirationYear  uint32 `json:"expiration_year"`
	DPANID          string `json:"dpan_id"`
	Cryptogram      string `json:"cryptogram"`
}

func (t Tokenization) IsEmpty() bool {
	return reflect.DeepEqual(t, Tokenization{})
}

type Subscription struct {
	SubscriptionID       string                `json:"subscription_id"`
	SubscriptionSequence *SubscriptionSequence `json:"subscription_sequence,omitempty"`
	FirstTimeUse         bool                  `json:"first_time_use"`
	BillingDate          *string               `json:"billing_date"`
}

func (s Subscription) IsEmpty() bool {
	return reflect.DeepEqual(s, Subscription{})
}

type SubscriptionSequence struct {
	Number *int `json:"number,omitempty"`
	Total  *int `json:"total,omitempty"`
}
