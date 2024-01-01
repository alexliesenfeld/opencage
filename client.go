package opencage

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Client represents a client interface to the OpenCage Geocoder API.
type Client struct {
	endpoint   string
	apiKey     string
	httpClient *http.Client
}

// New creates a new OpenCage client. Please have a look at https://opencagedata.com/api for more information.
// It requires an API key to function, which you can also get at https://opencagedata.com/api.
// The options parameter allows for some adjustments on how the client works internally,
// but should not be required in most cases.
func New(apiKey string, options ...Option) *Client {
	client := Client{
		apiKey:     apiKey,
		httpClient: &http.Client{},
		endpoint:   "https://api.opencagedata.com/geocode/v1",
	}

	for idx := range options {
		options[idx](&client)
	}

	return &client
}

// Geocode will call the remote API endpoint with the provided data. Please have a look at
// https://opencagedata.com/api for more information.
func (c *Client) Geocode(ctx context.Context, query string, params *GeocodingParams) (ReverseGeocodeResponse, error) {
	requestURL := c.createURL(query, params)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return ReverseGeocodeResponse{}, fmt.Errorf("cannot create HTTP request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ReverseGeocodeResponse{}, fmt.Errorf("failed to send the request: %w", err)
	}
	defer resp.Body.Close()

	if err := statusCodeToError(resp.StatusCode); err != nil {
		return ReverseGeocodeResponse{}, err
	}

	var result ReverseGeocodeResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return ReverseGeocodeResponse{}, fmt.Errorf("failed to parse the response: %w", err)
	}

	if err := statusCodeToError(result.Status.Code); err != nil {
		return ReverseGeocodeResponse{}, err
	}

	return result, nil
}

func (c *Client) createURL(query string, params *GeocodingParams) string {
	u, _ := url.Parse(fmt.Sprintf("%s/json", c.endpoint))

	q := u.Query()
	q.Set("q", query)
	q.Set("key", c.apiKey)

	if params != nil {
		// Common params
		if params.Abbreviate {
			q.Set("abbrv", "1")
		}

		if params.AddressOnly {
			q.Set("address_only", "1")
		}

		if params.AddRequest {
			q.Set("add_request", "1")
		}

		if params.JSONPFunctionName != "" {
			q.Set("jsonp", params.JSONPFunctionName)
		}

		if params.Language != "" {
			q.Set("language", params.Language)
		}

		if params.Limit > 0 {
			q.Set("limit", strconv.FormatInt(int64(params.Limit), 32))
		}

		if params.NoAnnotations {
			q.Set("no_annotations", "1")
		}

		if params.NoDedupe {
			q.Set("no_dedupe", "1")
		}

		if params.NoRecord {
			q.Set("no_record", "1")
		}

		if params.Pretty {
			q.Set("pretty", "1")
		}

		if params.RoadInfo {
			q.Set("roadinfo", "1")
		}

		// Only forward geocoding
		if len(params.Bounds) > 0 {
			q.Set("bounds", strings.Join(formatFloat32Slice(params.Bounds), ","))
		}

		if params.CountryCode != "" {
			q.Set("countrycode", params.CountryCode)
		}

		if len(params.Proximity) > 0 {
			q.Set("proximity", strings.Join(formatFloat32Slice(params.Proximity), ","))
		}

	}

	u.RawQuery = q.Encode()

	return u.String()
}

func statusCodeToError(statusCode int) error {
	switch statusCode {
	case 400:
		return ErrInvalidRequest
	case 401:
		return ErrAuthFailure
	case 402:
		return ErrQuotaExceeded
	case 403:
		return ErrForbidden
	case 404:
		return ErrInvalidEndpoint
	case 405:
		return ErrMethodNotAllowed
	case 408:
		return ErrTimeout
	case 410:
		return ErrRequestTooLong
	case 426:
		return ErrUpgradeRequired
	case 429:
		return ErrRateLimitExceeded
	case 500:
		return ErrInternalServerError
	}

	if statusCode >= 200 && statusCode < 300 {
		return nil
	}

	return fmt.Errorf("unexpected status code %d", statusCode)
}
