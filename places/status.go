package places

type Status string

const (
	// StatusOK indicates that no errors occurred; the place was successfully detected and at least one result was returned.
	StatusOK Status = "OK"
	// StatusUnknown indicates a server-side error; trying again may be successful.
	StatusUnknown Status = "UNKNOWN"
	// StatusZeroResults indicates that the reference was valid but no longer refers to a valid result. This may occur if the establishment is no longer in business.
	StatusZeroResults Status = "ZERO_RESULTS"
	// StatusOverQueryLimit indicates that you are over your quota.
	StatusOverQueryLimit Status = "OVER_QUERY_LIMIT"
	// StatusRequestDenied indicates that your request was denied, generally because of lack of an invalid key parameter.
	StatusRequestDenied Status = "REQUEST_DENIED"
	// StatusInvalidRequest generally indicates that the query (reference) is missing.
	StatusInvalidRequest Status = "INVALID_REQUEST"
	// StatusNotFound indicates that the referenced location was not found in the Places database.
	StatusNotFound Status = "NOT_FOUND"
)
