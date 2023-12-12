//go:build !coverage
// +build !coverage

package constants

const CL = "CL"
const PE = "PE"

const ERROR = "ERROR"

// SHARED ARCHETYPE CONSTANTS
const PROCESS_NAME = "PROCESS_NAME"

// ARCHETYPE SUBSCRIPTION CONTANTS
const (
	SUBSCRIPTION_NAME          = "SUBSCRIPTION_NAME"
	SUSBCRIPTION_SIGNAL_BROKEN = "SUSBCRIPTION_SIGNAL_BROKEN"
	RAW_EVENT                  = "RAW_EVENT"
	UNMARSHALED_EVENT          = "UNMARSHALED_EVENT"
	ERROR_UNMARSHALING_EVENT   = "ERROR_UNMARSHALING_EVENT"
)

// ARCHETYPE RESTY CONSTATS
const (
	REQUEST_BODY  = "REQUEST_BODY"
	RESPONSE_BODY = "RESPONSE_BODY"
	STATUS_CODE   = "STATUS_CODE"
)

// ARCHETYPE DATADOG LOGS CONSTANTS
const (
	STRUCT          = "struct"
	METHOD          = "method"
	DD_SPAN_ID      = "dd.span_id"
	DD_TRACE_ID     = "dd.trace_id"
	SERVICE_NAME    = "service.name"
	SERVICE_VERSION = "service.version"
)

// HTTP CLIENT HEADERS CONSTANTS
const (
	CONTENT_TYPE_HEADER_KEY  = "Content-Type"
	X_COMMERCE_HEADER_KEY    = "x-commerce"
	X_COUNTRY_HEADER_KEY     = "x-country"
	X_CHREF_HEADER_KEY       = "x-chref"
	X_CMREF_HEADER_KEY       = "x-cmref"
	X_TXREF_HEADER_KEY       = "x-txref"
	X_ENVIRONMENT_HEADER_KEY = "x-environment"
)
