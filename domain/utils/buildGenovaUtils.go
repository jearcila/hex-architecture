package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jearcila/hex-architecture/application/appcfg"
	"github.com/jearcila/hex-architecture/domain/constants"
	"github.com/jearcila/hex-architecture/domain/model"
	"github.com/jearcila/hex-architecture/domain/utils/format"
	transactions_context "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/context"
	transactions_models "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/transactions/models"
	"github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/utils/currencies"
	descriptor "github.com/mercadolibre/fury_gateway-kit/pkg/g2/framework/utils/descriptor"
	"github.com/shopspring/decimal"
)

// only digits
var digitCheck = regexp.MustCompile(`^\d+$`)

func ParseContext(ctx transactions_context.Context, req *model.FirstOperationRequest) error {
	req.MerchantOperationReference = BuildMerchantOperationReference(ctx)
	req.ProcessorID = appcfg.GetString(ctx, appcfg.GenovaProcessorID)
	floatAmount := currencies.CentToFloat(ctx.Transaction.Operation.Amount)
	req.Amount.Amount = decimal.NewFromFloat(floatAmount)
	req.Amount.Currency = ctx.Transaction.Operation.Currency
	req.Installments = ctx.Transaction.Operation.Installments

	if err := parseCard(ctx.Transaction, req); err != nil {
		return err
	}

	option, err := BuildOptions(*ctx.Transaction)
	if err != nil {
		return err
	}

	if option.WalletID != "" {
		req.WalletID = option.WalletID
	}

	if option.Plan != nil && option.Plan.ID != "" {
		req.InstallmentPlanType = option.Plan.ID
	}

	if err := parseMerchant(ctx.Transaction, option, req); err != nil {
		return err
	}

	parseAuthentication(option, req)

	parseTokenization(option, req)

	if err := parseSubscription(option, req); err != nil {
		return err
	}

	return nil
}

func parseCard(trx *transactions_models.Transaction, req *model.FirstOperationRequest) error {
	if trx.Operation.Card == nil {
		return errors.New("card information not found")
	}
	req.Card.NumberID = trx.Operation.Card.Number
	req.Card.ExpirationMonth = trx.Operation.Card.ExpirationMonth
	req.Card.ExpirationYear = trx.Operation.Card.ExpirationYear

	cardType, err := getCardType(trx.Operation.Card.Type)
	if err != nil {
		return err
	}
	req.CardProcessingMode = cardType

	// Fills the security code only if it's present
	if trx.Operation.Card.SecurityCode != nil {
		req.Card.SecurityCodeID = trx.Operation.Card.SecurityCode
	}

	// Fills the card on file info
	req.Card.CredentialOnFile = trx.Operation.Card.Id != ""

	// Fills the card holder info
	if trx.Operation.Card.Holder != nil {
		req.Card.Holder.Name = trx.Operation.Card.Holder.Name
		if trx.Operation.Card.Holder.Identification != nil {
			req.Card.Holder.ID.Type = trx.Operation.Card.Holder.Identification.Type
			req.Card.Holder.ID.Number = trx.Operation.Card.Holder.Identification.Number
		}
	}

	// Fills the present data only if it's present in case of first operation
	// Note: Refunds need to be processed as ecommerce. Card present info in refunds must be discarded.
	if trx.Operation.Card.Present != nil {
		err := parsePresent(trx.Operation.Card.Present, req)
		if err != nil {
			return err
		}
	}
	return nil
}

func parsePresent(trxPresent *transactions_models.Present, req *model.FirstOperationRequest) error {
	cardDataEntryMode, err := getCardEntryMode(trxPresent.Meta.DataEntryMode)
	if err != nil {
		return err
	}
	req.Card.EntryModeDetails = &model.EntryModeDetails{
		EntryMode:         cardDataEntryMode,
		CardPresentID:     trxPresent.Id,
		PINBlock:          trxPresent.Meta.PinBlock,
		FallbackIndicator: trxPresent.Meta.FallbackIndicator,
		ICCRelatedData:    trxPresent.Meta.IccRelatedData,
		ICCSequenceNumber: trxPresent.Meta.SequenceNumber,
	}
	req.PointOfInteraction = &model.PointOfInteraction{
		ID:        trxPresent.Meta.Poi,
		Type:      trxPresent.Meta.PoiType,
		Signature: trxPresent.Meta.PoiSignature,
	}
	return nil
}

func getCardEntryMode(cardEntryMode string) (string, error) {
	switch cardEntryMode {
	case constants.ManualEntryMode:
		return "manual", nil
	case constants.SwipeEntryMode:
		return "swipe", nil
	case constants.ChipEntryMode:
		return "chip", nil
	case constants.ContactlessSwipeEntryMode:
		return "contactless_swipe", nil
	case constants.ContactlessEntryMode:
		return "contactless_chip", nil
	}
	return "", errors.New("invalid card entry mode")
}

func BuildOptions(trx transactions_models.Transaction) (*model.Options, error) {
	var option model.Options
	err := json.Unmarshal(trx.Options, &option)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal options: %w", err)
	}
	return &option, nil

}

func parseMerchant(trx *transactions_models.Transaction, option *model.Options, req *model.FirstOperationRequest) error {
	if req.MerchantID = trx.Merchant.Account; req.MerchantID == "" {
		return errors.New("merchant account is empty")
	}
	// Send descriptor only if it's present
	if trx.Merchant.Descriptor != nil {
		sanitizedsoftDescriptor := descriptor.SetDescriptorValues(descriptor.Properties{Value: trx.Merchant.Descriptor.Product})
		sanitizedHardDescriptor := descriptor.SetDescriptorValues(descriptor.Properties{Value: trx.Merchant.Descriptor.MerchantName})

		length := len(sanitizedHardDescriptor)
		var finalHardDescriptor string

		switch {
		case length < constants.Len3:
			finalHardDescriptor = format.PadRight(sanitizedHardDescriptor, " ", constants.Len3)
		case length < constants.Len7:
			finalHardDescriptor = format.PadRight(sanitizedHardDescriptor, " ", constants.Len7)
		case length < constants.Len12:
			finalHardDescriptor = format.PadRight(sanitizedHardDescriptor, " ", constants.Len12)
		default:
			finalHardDescriptor = sanitizedHardDescriptor[:constants.Len12]
		}
		req.SoftDescriptor = sanitizedsoftDescriptor
		req.HardDescriptor = finalHardDescriptor
	}

	if option.Regulation == nil {
		return errors.New("regulations info not received")
	}

	sanitizedCity, err := format.GetNormalizedASCII(option.Regulation.City)
	if err != nil {
		return errors.New("invalid city format")
	}
	req.SubMerchant.Location.City = sanitizedCity

	req.MCC = ""
	if option.Regulation.MCC != nil && *option.Regulation.MCC != "" {
		req.MCC = *option.Regulation.MCC
	}

	sanitizedZip, err := format.GetNormalizedASCII(option.Regulation.ZIPCode)
	if err != nil {
		return errors.New("invalid zip format")
	}
	req.SubMerchant.Location.ZipCode = sanitizedZip

	sanitizedLegalName, err := format.GetNormalizedASCII(option.Regulation.LegalName)
	if err != nil {
		return errors.New("invalid legal name format")
	}
	req.SubMerchant.LegalName = sanitizedLegalName

	req.SubMerchant.Location.CountryCode = option.Regulation.Country

	sanitizedAddress, err := format.GetNormalizedASCII(option.Regulation.AddressStreet)
	if err != nil {
		return errors.New("invalid address street format")
	}
	req.SubMerchant.Location.Address = sanitizedAddress

	req.SubMerchant.Location.AddressDoorNumber = option.Regulation.AddressDoorNumber
	req.SubMerchant.Location.Region = option.Regulation.RegionCodeIso
	req.SubMerchant.TaxID.Number = option.Regulation.DocumentNumber
	req.SubMerchant.TaxID.Type = option.Regulation.DocumentType
	req.SubMerchant.FiscalCondition = option.Regulation.FiscalCondition

	if option.CollectorID > 0 {
		req.SubMerchant.ID = strconv.FormatUint(option.CollectorID, 10)
	}

	return nil
}

func parseAuthentication(option *model.Options, req *model.FirstOperationRequest) {
	if option.ThreeDS != nil && !option.ThreeDS.IsEmpty() {
		req.Authentication = &model.Authentication{
			ThreeDS: &model.ThreeDS{
				Cryptogram:         option.ThreeDS.Cryptogram,
				ServerTransID:      option.ThreeDS.ThreeDSServerTransID,
				ACSReferenceNumber: option.ThreeDS.ACSReferenceNumber,
				Eci:                option.ThreeDS.Eci,
				DSTransID:          option.ThreeDS.DSTransID,
				ACSTransID:         option.ThreeDS.ACSTransID,
				ThreeDSVersion:     getMajorVersion(option.ThreeDS.ThreeDSVersion, ""),
			}}
	}
}

// getMajorVersion returns the major version value from version. For example,
// getMajorVersion("2.1.0", "2") == "2". If version is an empty string, getMajorVersion returns defaultVersion.
func getMajorVersion(version, defaultVersion string) string {
	if version != "" {
		return strings.Split(version, ".")[0]
	}

	return defaultVersion
}

func parseTokenization(option *model.Options, req *model.FirstOperationRequest) {
	if option.Tokenization != nil && !option.Tokenization.IsEmpty() {
		req.Card.Tokenization = &model.TokenizationInfo{
			Cryptogram:      option.Tokenization.Cryptogram,
			ExpirationMonth: option.Tokenization.ExpirationMonth,
			ExpirationYear:  option.Tokenization.ExpirationYear,
			DPANID:          option.Tokenization.DPANID,
		}
		if digitCheck.MatchString(option.Tokenization.Cryptogram) {
			// If the cryptogram value is numeric then it is a cryptogram id
			req.Card.Tokenization.Cryptogram = ""
			req.Card.Tokenization.CryptogramID = option.Tokenization.Cryptogram
		}
	}
}

func parseSubscription(option *model.Options, req *model.FirstOperationRequest) error {
	if option.Subscription != nil && !option.Subscription.IsEmpty() {
		req.Recurring = &model.Recurring{
			ID:                 option.Subscription.SubscriptionID,
			FirstPayment:       option.Subscription.FirstTimeUse,
			InvoicePeriodMonth: int(time.Now().Month()),
			InvoicePeriodYear:  time.Now().Year(),
		}

		if option.Subscription.BillingDate != nil {
			billingDate, err := time.Parse(constants.BillingDateLayout, *option.Subscription.BillingDate)
			if err != nil {
				return errors.New("cannot parse subscription billing date")
			}

			year, month, _ := billingDate.Date()
			req.Recurring.InvoicePeriodMonth = int(month)
			req.Recurring.InvoicePeriodYear = year
		}
	}

	return nil
}

func getCardType(cardType string) (string, error) {
	switch cardType {
	case constants.DebitCard:
		return "debit", nil
	case constants.CreditCard:
		return "credit", nil
	}
	return "", errors.New("invalid card type")
}
