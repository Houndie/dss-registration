package dynamic

type jsonError struct {
	Type                    jsonErrorType            `json:"type,omitempty"`
	MissingParameterDetails *missingParameterDetails `json:"missing_parameter_details,omitempty"`
	BadParameterDetails     *badParameterDetails     `json:"bad_parameter_details,omitempty"`
}

type jsonErrorType string

const (
	jsonErrorTypeInternalServerError jsonErrorType = "INTERNAL_SERVER_ERROR"
	jsonErrorTypeMissingParameter    jsonErrorType = "MISSING_PARAMETER"
	jsonErrorTypeBadParameter        jsonErrorType = "BAD_PARAMETER"
	jsonErrorTypeUnauthorized        jsonErrorType = "UNAUTHORIZED"
)

type missingParameterDetails struct {
	ParameterName string `json:"parameter_name"`
}

type badParameterDetails struct {
	ParameterName string `json:"parameter_name"`
	SuppliedValue string `json:"supplied_value"`
	Reason        string `json:"reason"`
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

func unauthorizedError() *jsonError {
	return &jsonError{
		Type: jsonErrorTypeUnauthorized,
	}
}
