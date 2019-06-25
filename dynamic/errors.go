package dynamic

type jsonError struct {
	Type                    jsonErrorType            `json:"type"`
	MissingParameterDetails *missingParameterDetails `json:"missing_parameter_details"`
	BadParameterDetails     *badParameterDetails     `json:"bad_parameter_details"`
}

type jsonErrorType string

const (
	jsonErrorTypeInternalServerError jsonErrorType = "INTERNAL_SERVER_ERROR"
	jsonErrorTypeMissingParameter    jsonErrorType = "MISSING_PARAMETER"
	jsonErrorTypeBadParameter        jsonErrorType = "BAD_PARAMETER"
)

type missingParameterDetails struct {
	ParameterName string
}

type badParameterDetails struct {
	ParameterName string
	SuppliedValue string
	Reason        string
}

func internalServerError() *jsonError {
	return &jsonError{
		Type: jsonErrorTypeInternalServerError,
	}
}

func missingParameterError(parameterName string) *jsonError {
	return &jsonError{
		Type: jsonErrorTypeMissingParameter,
		MissingParameterDetails: &missingParameterDetails{
			ParameterName: parameterName,
		},
	}
}

func badParameterError(parameterName, suppliedValue, reason string) *jsonError {
	return &jsonError{
		Type: jsonErrorTypeBadParameter,
		BadParameterDetails: &badParameterDetails{
			ParameterName: parameterName,
			SuppliedValue: suppliedValue,
			Reason:        reason,
		},
	}
}
