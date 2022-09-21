package integration

import (
	channel_http_json "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/channels/http_json"
	"github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/integrations"
)

const (
	_providerName = "genova"
)

func GetCapabilities() integrations.Capabilities {
	return integrations.Capabilities{
		Provider: _providerName,
		Version:  "1.22.6",
		Communication: integrations.Communication{
			Channel: []string{channel_http_json.ChannelName},
			VPN:     false,
		},
		Countries: integrations.Countries{
			integrations.ISOCodeBrazil:    _brazilConfig,
			integrations.ISOCodeArgentina: _argentinaConfig,
			integrations.ISOCodeMexico:    _mexicoConfig,
		}}
}

var _brazilConfig = integrations.Country{
	Acquirers: integrations.Acquirers{
		_providerName: {
			Operations: integrations.Operations{
				Authorization: &integrations.Subtypes{},
				Capture: &integrations.Subtypes{
					integrations.OperationSubtypeTotal,
				},
				Refund: integrations.Refund{
					Capture: &integrations.Subtypes{
						integrations.OperationSubtypeTotal,
						integrations.OperationSubtypePartial,
					},
					Authorization: &integrations.Subtypes{},
				},
				Purchase: &integrations.Subtypes{},
				Query: integrations.ComplementaryOperation{
					Purchase:      false,
					Authorization: false,
					Capture:       false,
					Refund:        false,
				},
				Reverse: integrations.ComplementaryOperation{
					Purchase:      false,
					Authorization: false,
					Capture:       false,
					Refund:        false,
				},
			},
			Parameters: integrations.Parameters{
				Merchant: integrations.Merchant{
					"account": integrations.Description{
						Type:        integrations.TypeString,
						Requirement: integrations.RequirementMandatory,
					},
					"number": integrations.Description{
						Type:        integrations.TypeString,
						Requirement: integrations.RequirementMandatory,
					},
				},
				Provider: integrations.Provider{
					"id": integrations.Description{
						Type:        integrations.TypeString,
						Requirement: integrations.RequirementMandatory,
					},
				},
				Options: integrations.Options{
					"regulation": map[string]interface{}{
						"mcc": integrations.Description{
							Reference:   "MerchantCategoryCode",
							Type:        integrations.TypeString,
							Requirement: integrations.RequirementMandatory,
						},
					},
					"three_ds": map[string]interface{}{
						"cryptogram": integrations.Description{
							Reference:   "AuthenticationValue",
							Type:        integrations.TypeString,
							Requirement: integrations.RequirementOptional,
						},
						"ds_trans_id": integrations.Description{
							Reference:   "TransactionId",
							Type:        integrations.TypeString,
							Requirement: integrations.RequirementOptional,
						},
						"three_ds_version": integrations.Description{
							Reference:   "ThreeDSecureVersion",
							Type:        integrations.TypeString,
							Requirement: integrations.RequirementOptional,
						},
					},
					"tokenization": map[string]integrations.Description{
						"dpan_id": {
							Reference:   "DpanID",
							Type:        integrations.TypeString,
							Requirement: integrations.RequirementOptional,
						},
						"cryptogram": {
							Reference:   "AuthenticationValue",
							Type:        integrations.TypeString,
							Requirement: integrations.RequirementOptional,
						},
						"expiration_month": {
							Reference:   "ExpirationMonth",
							Type:        "int",
							Requirement: integrations.RequirementOptional,
						},
						"expiration_year": {
							Reference:   "ExpirationYear",
							Type:        "int",
							Requirement: integrations.RequirementOptional,
						},
					},
				},
			},
			Cards: integrations.Cards{
				Brands: integrations.Brands{
					"master": integrations.CardTypes{integrations.CardTypeCreditCard, integrations.CardTypeDebitCard},
				},
				EntryModes: integrations.EntryModes{
					integrations.EntryModeManual,
					integrations.EntryModeChip,
					integrations.EntryModeNfc,
					integrations.EntryModeSwipe,
				},
				SecurityCode: integrations.RequirementOptional,
			},
			ProcessingType: integrations.ProcessingType{
				Ecommerce: true,
				Present:   true,
			},
			OperationMode: integrations.OperationMode{
				Recurring:    true,
				Subscription: true,
			},
		},
	},
}

var _argentinaConfig = integrations.Country{
	Acquirers: integrations.Acquirers{
		_providerName: {
			Operations: integrations.Operations{
				Authorization: &integrations.Subtypes{},
				Capture: &integrations.Subtypes{
					integrations.OperationSubtypeTotal,
				},
				Refund: integrations.Refund{
					Capture: &integrations.Subtypes{
						integrations.OperationSubtypeTotal,
						integrations.OperationSubtypePartial,
					},
					Authorization: &integrations.Subtypes{},
				},
				Purchase: &integrations.Subtypes{},
			},
			Parameters: integrations.Parameters{
				Merchant: integrations.Merchant{
					"account": integrations.Description{
						Type:        integrations.TypeString,
						Requirement: integrations.RequirementMandatory,
					},
					"number": integrations.Description{
						Type:        integrations.TypeString,
						Requirement: integrations.RequirementMandatory,
					},
				},
				Provider: integrations.Provider{
					"id": integrations.Description{
						Type:        integrations.TypeString,
						Requirement: integrations.RequirementMandatory,
					},
				},
				Options: integrations.Options{
					"regulation": map[string]interface{}{
						"mcc": integrations.Description{
							Reference:   "MerchantCategoryCode",
							Type:        integrations.TypeString,
							Requirement: integrations.RequirementMandatory,
						},
					},
					"three_ds": map[string]interface{}{
						"cryptogram": integrations.Description{
							Reference:   "AuthenticationValue",
							Type:        integrations.TypeString,
							Requirement: integrations.RequirementOptional,
						},
						"ds_trans_id": integrations.Description{
							Reference:   "TransactionId",
							Type:        integrations.TypeString,
							Requirement: integrations.RequirementOptional,
						},
						"three_ds_version": integrations.Description{
							Reference:   "ThreeDSecureVersion",
							Type:        integrations.TypeString,
							Requirement: integrations.RequirementOptional,
						},
					},
					"tokenization": map[string]integrations.Description{
						"dpan_id": {
							Reference:   "DpanID",
							Type:        integrations.TypeString,
							Requirement: integrations.RequirementOptional,
						},
						"cryptogram": {
							Reference:   "AuthenticationValue",
							Type:        integrations.TypeString,
							Requirement: integrations.RequirementOptional,
						},
						"expiration_month": {
							Reference:   "ExpirationMonth",
							Type:        "int",
							Requirement: integrations.RequirementOptional,
						},
						"expiration_year": {
							Reference:   "ExpirationYear",
							Type:        "int",
							Requirement: integrations.RequirementOptional,
						},
					},
				},
			},
			Cards: integrations.Cards{
				Brands: integrations.Brands{
					"master": integrations.CardTypes{integrations.CardTypeCreditCard, integrations.CardTypeDebitCard},
				},
				EntryModes: integrations.EntryModes{
					integrations.EntryModeManual,
					integrations.EntryModeChip,
					integrations.EntryModeNfc,
					integrations.EntryModeSwipe,
				},
				SecurityCode: integrations.RequirementOptional,
			},
		},
	},
}

var _mexicoConfig = integrations.Country{
	Acquirers: integrations.Acquirers{
		_providerName: {
			Operations: integrations.Operations{
				Authorization: &integrations.Subtypes{},
				Capture: &integrations.Subtypes{
					integrations.OperationSubtypeTotal,
				},
				Refund: integrations.Refund{
					Capture: &integrations.Subtypes{
						integrations.OperationSubtypeTotal,
						integrations.OperationSubtypePartial,
					},
					Authorization: &integrations.Subtypes{},
				},
				Purchase: &integrations.Subtypes{},
			},
			Parameters: integrations.Parameters{
				Merchant: integrations.Merchant{
					"account": integrations.Description{
						Type:        integrations.TypeString,
						Requirement: integrations.RequirementMandatory,
					},
					"number": integrations.Description{
						Type:        integrations.TypeString,
						Requirement: integrations.RequirementMandatory,
					},
				},
				Provider: integrations.Provider{
					"id": integrations.Description{
						Type:        integrations.TypeString,
						Requirement: integrations.RequirementMandatory,
					},
				},
				Options: integrations.Options{
					"regulation": map[string]interface{}{
						"mcc": integrations.Description{
							Reference:   "MerchantCategoryCode",
							Type:        integrations.TypeString,
							Requirement: integrations.RequirementMandatory,
						},
					},
					"three_ds": map[string]interface{}{
						"cryptogram": integrations.Description{
							Reference:   "AuthenticationValue",
							Type:        integrations.TypeString,
							Requirement: integrations.RequirementOptional,
						},
						"ds_trans_id": integrations.Description{
							Reference:   "TransactionId",
							Type:        integrations.TypeString,
							Requirement: integrations.RequirementOptional,
						},
						"three_ds_version": integrations.Description{
							Reference:   "ThreeDSecureVersion",
							Type:        integrations.TypeString,
							Requirement: integrations.RequirementOptional,
						},
					},
					"tokenization": map[string]integrations.Description{
						"dpan_id": {
							Reference:   "DpanID",
							Type:        integrations.TypeString,
							Requirement: integrations.RequirementOptional,
						},
						"cryptogram": {
							Reference:   "AuthenticationValue",
							Type:        integrations.TypeString,
							Requirement: integrations.RequirementOptional,
						},
						"expiration_month": {
							Reference:   "ExpirationMonth",
							Type:        "int",
							Requirement: integrations.RequirementOptional,
						},
						"expiration_year": {
							Reference:   "ExpirationYear",
							Type:        "int",
							Requirement: integrations.RequirementOptional,
						},
					},
				},
			},
			Cards: integrations.Cards{
				Brands: integrations.Brands{
					"master": integrations.CardTypes{integrations.CardTypeCreditCard, integrations.CardTypeDebitCard},
				},
				EntryModes: integrations.EntryModes{
					integrations.EntryModeManual,
					integrations.EntryModeChip,
					integrations.EntryModeNfc,
					integrations.EntryModeSwipe,
				},
				SecurityCode: integrations.RequirementOptional,
			},
		},
	},
}
