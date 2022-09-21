package log

// Tags
const (
	// Tags
	TagErrorMessage = "error_message"
	TagRawMessage   = "message"
)

// Events
const (
	// Authorization
	EventBuildAuthorizationMessage   = "integration_build_authorization_message"
	EventSendAuthorizationMessage    = "integration_send_authorization_message"
	EventAuthorizationResponse       = "integration_authorization_response"
	EventCreateAuthorizationResponse = "integration_create_authorization_response"
	EventBrandNotDeterminate         = "brand_not_determinate"

	// Capture
	EventBuildCaptureMessage   = "integration_build_capture_message"
	EventSendCaptureMessage    = "integration_send_capture_message"
	EventCaptureResponse       = "integration_capture_response"
	EventCreateCaptureResponse = "integration_create_capture_response"

	// Refund
	EventBuildRefundMessage   = "integration_build_refund_message"
	EventSendRefundMessage    = "integration_send_refund_message"
	EventRefundResponse       = "integration_refund_response"
	EventCreateRefundResponse = "integration_create_refund_response"

	// Purchase
	EventBuildPurchaseMessage   = "integration_build_purchase_message"
	EventSendPurchaseMessage    = "integration_send_purchase_message"
	EventPurchaseResponse       = "integration_purchase_response"
	EventCreatePurchaseResponse = "integration_create_purchase_response"
)

// Event error
const (
	EventErrorNotSavedReferences          = "integration_error_saving_references"
	EventErrorNotSavedResponse            = "integration_error_saving_response"
	EventErrorNotSavedParsed              = "integration_error_saving_parsed"
	EventErrorBuildAuthorizationMessage   = "integration_error_building_authorization_request"
	EventErrorAcquirerResponse            = "integration_error_in_acquirer_response"
	EventErrorCreateAuthorizationResponse = "integration_error_creating_authorization_response"
	EventErrorBuildCaptureMessage         = "integration_error_building_capture_request"
	EventErrorCreateCaptureResponse       = "integration_error_creating_capture_response"
	EventErrorBuildRefundMessage          = "integration_error_building_refund_request"
	EventErrorCreateRefundResponse        = "integration_error_creating_refund_response"
	EventErrorBuildPurchaseMessage        = "integration_error_building_purchase_request"
	EventErrorCreatePurchaseResponse      = "integration_error_creating_purchase_response"
)

const MessageErrorBrandNoDeterminate = "the brand is not determinate, is credit card flow, purchase"
