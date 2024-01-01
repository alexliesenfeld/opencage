package opencage

import "errors"

var ErrInvalidRequest = errors.New("invalid request - a required parameter is missing, invalid coordinates, invalid version, or invalid format")
var ErrAuthFailure = errors.New("unable to authenticate - missing, invalid, or unknown API key")
var ErrQuotaExceeded = errors.New("valid request but quota exceeded (payment required)")
var ErrForbidden = errors.New("forbidden - API key disabled or IP address rejected")
var ErrInvalidEndpoint = errors.New("invalid API endpoint")
var ErrMethodNotAllowed = errors.New("method not allowed - non-GET request")
var ErrTimeout = errors.New("timeout - you can try again")
var ErrRequestTooLong = errors.New("request too long")
var ErrUpgradeRequired = errors.New("upgrade required - unsupported TLS")
var ErrRateLimitExceeded = errors.New("too many requests - rate limiting in effect")
var ErrInternalServerError = errors.New("internal server error")
