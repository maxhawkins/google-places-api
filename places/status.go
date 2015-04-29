package places

import "fmt"

type apiError struct {
	Status  string
	Message string
}

func (e *apiError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("%s: %s", e.Status, e.Message)
	}
	return e.Status
}

// IsUnknown returns true if the error indicates a server-side error and trying again may be successful.
func IsUnknown(err error) bool {
	if e, ok := err.(*apiError); ok {
		return e.Status == "UNKNOWN"
	}
	return false
}

// IsZeroResults returns true if the error indicates that the search was successful but returned no results. This may occur if the search was passed a latlng in a remote location.
func IsZeroResults(err error) bool {
	if e, ok := err.(*apiError); ok {
		return e.Status == "ZERO_RESULTS"
	}
	return false
}

// IsOverQueryLimit returns true if the error indicates that you are over your quota.
func IsOverQueryLimit(err error) bool {
	if e, ok := err.(*apiError); ok {
		return e.Status == "OVER_QUERY_LIMIT"
	}
	return false
}

// IsRequestDenied returns true if the error indicates that your request was denied, generally because of lack of an invalid key parameter.
func IsRequestDenied(err error) bool {
	if e, ok := err.(*apiError); ok {
		return e.Status == "REQUEST_DENIED"
	}
	return false
}

// IsInvalidRequest returns true if the error indicates that the request is invalid. Generally this means that the query (reference) is missing.
func IsInvalidRequest(err error) bool {
	if e, ok := err.(*apiError); ok {
		return e.Status == "INVALID_REQUEST"
	}
	return false
}

// IsNotFound returns true if the error indicates that the referenced location was not found in the Places database.
func IsNotFound(err error) bool {
	if e, ok := err.(*apiError); ok {
		return e.Status == "NOT_FOUND"
	}
	return false
}
