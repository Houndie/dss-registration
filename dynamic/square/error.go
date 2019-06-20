package square

import "fmt"

type ErrorCategory string

const (
	API_ERROR             ErrorCategory = "API_ERROR"
	AUTHENTICATION_ERROR  ErrorCategory = "AUTHENTICATION_ERROR"
	INVALID_REQUEST_ERROR ErrorCategory = "INVALID_REQUEST_ERROR"
	RATE_LIMIT_ERROR      ErrorCategory = "RATE_LIMIT_ERROR"
	PAYMENT_METHOD_ERROR  ErrorCategory = "PAYMENT_METHOD_ERROR"
	REFUND_ERROR          ErrorCategory = "REFUND_ERROR"
)

type ErrorCode string

const (
	INTERNAL_SERVER_ERROR                               ErrorCode = "INTERNAL_SERVER_ERROR"
	UNAUTHORIZED                                        ErrorCode = "UNAUTHORIZED"
	ACCESS_TOKEN_EXPIRED                                ErrorCode = "ACCESS_TOKEN_EXPIRED"
	ACCESS_TOKEN_REVOKED                                ErrorCode = "ACCESS_TOKEN_REVOKED"
	FORBIDDEN                                           ErrorCode = "FORBIDDEN"
	INSUFFICIENT_SCOPES                                 ErrorCode = "INSUFFICIENT_SCOPES"
	APPLICATION_DISABLED                                ErrorCode = "APPLICATION_DISABLED"
	V1_APPLICATION                                      ErrorCode = "V1_APPLICATION"
	V1_ACCESS_TOKEN                                     ErrorCode = "V1_ACCESS_TOKEN"
	CARD_PROCESSING_NOT_ENABLED                         ErrorCode = "CARD_PROCESSING_NOT_ENABLED"
	BAD_REQUEST                                         ErrorCode = "BAD_REQUEST"
	MISSING_REQUIRED_PARAMETER                          ErrorCode = "MISSING_REQUIRED_PARAMETER"
	INCORRECT_TYPE                                      ErrorCode = "INCORRECT_TYPE"
	INVALID_TIME                                        ErrorCode = "INVALID_TIME"
	INVALID_TIME_RANGE                                  ErrorCode = "INVALID_TIME_RANGE"
	INVALID_VALUE                                       ErrorCode = "INVALID_VALUE"
	INVALID_CURSOR                                      ErrorCode = "INVALID_CURSOR"
	UNKNOWN_QUERY_PARAMETER                             ErrorCode = "UNKNOWN_QUERY_PARAMETER"
	CONFLICTING_PARAMETERS                              ErrorCode = "CONFLICTING_PARAMETERS"
	EXPECTED_JSON_BODY                                  ErrorCode = "EXPECTED_JSON_BODY"
	INVALID_SORT_ORDER                                  ErrorCode = "INVALID_SORT_ORDER"
	VALUE_REGEX_MISMATCH                                ErrorCode = "VALUE_REGEX_MISMATCH"
	VALUE_TOO_SHORT                                     ErrorCode = "VALUE_TOO_SHORT"
	VALUE_TOO_LONG                                      ErrorCode = "VALUE_TOO_LONG"
	VALUE_TOO_LOW                                       ErrorCode = "VALUE_TOO_LOW"
	VALUE_TOO_HIGH                                      ErrorCode = "VALUE_TOO_HIGH"
	VALUE_EMPTY                                         ErrorCode = "VALUE_EMPTY"
	ARRAY_LENGTH_TOO_LONG                               ErrorCode = "ARRAY_LENGTH_TOO_LONG"
	ARRAY_LENGTH_TOO_SHORT                              ErrorCode = "ARRAY_LENGTH_TOO_SHORT"
	ARRAY_EMPTY                                         ErrorCode = "ARRAY_EMPTY"
	EXPECTED_BOOLEAN                                    ErrorCode = "EXPECTED_BOOLEAN"
	EXPECTED_INTEGER                                    ErrorCode = "EXPECTED_INTEGER"
	EXPECTED_FLOAT                                      ErrorCode = "EXPECTED_FLOAT"
	EXPECTED_STRING                                     ErrorCode = "EXPECTED_STRING"
	EXPECTED_OBJECT                                     ErrorCode = "EXPECTED_OBJECT"
	EXPECTED_ARRAY                                      ErrorCode = "EXPECTED_ARRAY"
	EXPECTED_MAP                                        ErrorCode = "EXPECTED_MAP"
	EXPECTED_BASE64_ENCODED_BYTE_ARRAY                  ErrorCode = "EXPECTED_BASE64_ENCODED_BYTE_ARRAY"
	INVALID_ARRAY_VALUE                                 ErrorCode = "INVALID_ARRAY_VALUE"
	INVALID_ENUM_VALUE                                  ErrorCode = "INVALID_ENUM_VALUE"
	INVALID_CONTENT_TYPE                                ErrorCode = "INVALID_CONTENT_TYPE"
	INVALID_FORM_VALUE                                  ErrorCode = "INVALID_FORM_VALUE"
	ONE_INSTRUMENT_EXPECTED                             ErrorCode = "ONE_INSTRUMENT_EXPECTED"
	NO_FIELDS_SET                                       ErrorCode = "NO_FIELDS_SET"
	DEPRECATED_FIELD_SET                                ErrorCode = "DEPRECATED_FIELD_SET"
	RETIRED_FIELD_SET                                   ErrorCode = "RETIRED_FIELD_SET"
	CARD_EXPIRED                                        ErrorCode = "CARD_EXPIRED"
	INVALID_EXPIRATION                                  ErrorCode = "INVALID_EXPIRATION"
	INVALID_EXPIRATION_YEAR                             ErrorCode = "INVALID_EXPIRATION_YEAR"
	INVALID_EXPIRATION_DATE                             ErrorCode = "INVALID_EXPIRATION_DATE"
	UNSUPPORTED_CARD_BRAND                              ErrorCode = "UNSUPPORTED_CARD_BRAND"
	UNSUPPORTED_ENTRY_METHOD                            ErrorCode = "UNSUPPORTED_ENTRY_METHOD"
	INVALID_ENCRYPTED_CARD                              ErrorCode = "INVALID_ENCRYPTED_CARD"
	INVALID_CARD                                        ErrorCode = "INVALID_CARD"
	DELAYED_TRANSACTION_EXPIRED                         ErrorCode = "DELAYED_TRANSACTION_EXPIRED"
	DELAYED_TRANSACTION_CANCELED                        ErrorCode = "DELAYED_TRANSACTION_CANCELED"
	DELAYED_TRANSACTION_CAPTURED                        ErrorCode = "DELAYED_TRANSACTION_CAPTURED"
	DELAYED_TRANSACTION_FAILED                          ErrorCode = "DELAYED_TRANSACTION_FAILED"
	CARD_TOKEN_EXPIRED                                  ErrorCode = "CARD_TOKEN_EXPIRED"
	CARD_TOKEN_USED                                     ErrorCode = "CARD_TOKEN_USED"
	AMOUNT_TOO_HIGH                                     ErrorCode = "AMOUNT_TOO_HIGH"
	UNSUPPORTED_INSTRUMENT_TYPE                         ErrorCode = "UNSUPPORTED_INSTRUMENT_TYPE"
	REFUND_AMOUNT_INVALID                               ErrorCode = "REFUND_AMOUNT_INVALID"
	REFUND_ALREADY_PENDING                              ErrorCode = "REFUND_ALREADY_PENDING"
	PAYMENT_NOT_REFUNDABLE                              ErrorCode = "PAYMENT_NOT_REFUNDABLE"
	INVALID_CARD_DATA                                   ErrorCode = "INVALID_CARD_DATA"
	LOCATION_MISMATCH                                   ErrorCode = "LOCATION_MISMATCH"
	IDEMPOTENCY_KEY_REUSED                              ErrorCode = "IDEMPOTENCY_KEY_REUSED"
	UNEXPECTED_VALUE                                    ErrorCode = "UNEXPECTED_VALUE"
	SANDBOX_NOT_SUPPORTED                               ErrorCode = "SANDBOX_NOT_SUPPORTED"
	INVALID_EMAIL_ADDRESS                               ErrorCode = "INVALID_EMAIL_ADDRESS"
	INVALID_PHONE_NUMBER                                ErrorCode = "INVALID_PHONE_NUMBER"
	CHECKOUT_EXPIRED                                    ErrorCode = "CHECKOUT_EXPIRED"
	BAD_CERTIFICATE                                     ErrorCode = "BAD_CERTIFICATE"
	INVALID_SQUARE_VERSION_FORMAT                       ErrorCode = "INVALID_SQUARE_VERSION_FORMAT"
	API_VERSION_INCOMPATIBLE                            ErrorCode = "API_VERSION_INCOMPATIBLE"
	CARD_DECLINED                                       ErrorCode = "CARD_DECLINED"
	VERIFY_CVV_FAILURE                                  ErrorCode = "VERIFY_CVV_FAILURE"
	VERIFY_AVS_FAILURE                                  ErrorCode = "VERIFY_AVS_FAILURE"
	CARD_DECLINED_CALL_ISSUER                           ErrorCode = "CARD_DECLINED_CALL_ISSUER"
	NOT_FOUND                                           ErrorCode = "NOT_FOUND"
	APPLE_PAYMENT_PROCESSING_CERTIFICATE_HASH_NOT_FOUND ErrorCode = "APPLE_PAYMENT_PROCESSING_CERTIFICATE_HASH_NOT_FOUND"
	METHOD_NOT_ALLOWED                                  ErrorCode = "METHOD_NOT_ALLOWED"
	NOT_ACCEPTABLE                                      ErrorCode = "NOT_ACCEPTABLE"
	REQUEST_TIMEOUT                                     ErrorCode = "REQUEST_TIMEOUT"
	CONFLICT                                            ErrorCode = "CONFLICT"
	REQUEST_ENTITY_TOO_LARGE                            ErrorCode = "REQUEST_ENTITY_TOO_LARGE"
	UNSUPPORTED_MEDIA_TYPE                              ErrorCode = "UNSUPPORTED_MEDIA_TYPE"
	RATE_LIMITED                                        ErrorCode = "RATE_LIMITED"
	NOT_IMPLEMENTED                                     ErrorCode = "NOT_IMPLEMENTED"
	SERVICE_UNAVAILABLE                                 ErrorCode = "SERVICE_UNAVAILABLE"
	GATEWAY_TIMEOUT                                     ErrorCode = "GATEWAY_TIMEOUT"
)

type Error struct {
	Category ErrorCategory `json:"category"`
	Code     ErrorCode     `json:"code"`
	Detail   string        `json:"detail"`
	Field    string        `json:"field"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("Square Error: Category: %s; Code: %s; Detail: %s; Field: %s", e.Category, e.Code, e.Detail, e.Field)
}

type ErrorList struct {
	Errors []*Error
}

func (e *ErrorList) Error() string {
	if len(e.Errors) == 1 {
		return e.Errors[0].Error()
	}

	retVal := "Multiple Square Errors Returned:"
	for _, err := range e.Errors {
		retVal = fmt.Sprintf(retVal+" %s", err.Error())
	}
	return retVal
}
