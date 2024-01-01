package opencage

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"
)

func TestClientForwardGeocodingRoundTrip(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("roadinfo") != "1" {
			t.Fatalf("roadinfo should be 1")
		}

		if r.URL.Query().Get("proximity") != "1.0000000,-1.0000000" {
			t.Fatalf("proximity should be 1.0000000,-1.0000000")
		}

		if r.URL.Query().Get("language") != "de" {
			t.Fatalf("language should be de")
		}

		body, err := ioutil.ReadFile(filepath.Join("testdata", "reverse_geocoding_response.json"))
		if err != nil {
			t.Fatal("error reading source file:", err)
		}

		w.WriteHeader(200)
		_, _ = w.Write(body)
	}))

	client := New("my-api-key", WithEndpoint(server.URL), WithGlobalTimeout(10*time.Second))

	response, err := client.Geocode(context.Background(), "52.3877830 9.7334394", &GeocodingParams{
		RoadInfo:  true,
		Proximity: []float32{1.0, -1.0},
		Language:  "de",
	})

	if err != nil {
		t.Fatalf("there should be no error but was: %s", err.Error())
	}

	if response.Status.Code != 200 {
		t.Fatalf("status code is not 200")
	}

	if response.Status.Message != "OK" {
		t.Fatalf("status code is not OK")
	}
}

func TestClientBackwardGeocodingRoundTrip(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadFile(filepath.Join("testdata", "forward_geocoding_response.json"))
		if err != nil {
			t.Fatal("error reading source file:", err)
		}

		if r.URL.Query().Get("roadinfo") != "1" {
			t.Fatalf("roadinfo should be 1")
		}

		if r.URL.Query().Get("bounds") != "1.0000000,-1.0000000" {
			t.Fatalf("proximity should be 1.0000000,-1.0000000")
		}

		if r.URL.Query().Get("language") != "de" {
			t.Fatalf("language should be de")
		}

		w.WriteHeader(200)
		_, _ = w.Write(body)
	}))

	client := New("my-api-key", WithEndpoint(server.URL), WithGlobalTimeout(10*time.Second))

	response, err := client.Geocode(context.Background(), "Berlin, Germany", &GeocodingParams{
		RoadInfo: true,
		Bounds:   []float32{1.0, -1.0},
		Language: "de",
	})

	if err != nil {
		t.Fatalf("there should be no error but was: %s", err.Error())
	}

	if response.Status.Code != 200 {
		t.Fatalf("status code is not 200")
	}

	if response.Status.Message != "OK" {
		t.Fatalf("status code is not OK")
	}
}

func TestClientError(t *testing.T) {
	testData := []struct {
		statusCode    int
		expectedError error
	}{
		{statusCode: 400, expectedError: ErrInvalidRequest},
		{statusCode: 401, expectedError: ErrAuthFailure},
		{statusCode: 402, expectedError: ErrQuotaExceeded},
		{statusCode: 403, expectedError: ErrForbidden},
		{statusCode: 404, expectedError: ErrInvalidEndpoint},
		{statusCode: 405, expectedError: ErrMethodNotAllowed},
		{statusCode: 408, expectedError: ErrTimeout},
		{statusCode: 410, expectedError: ErrRequestTooLong},
		{statusCode: 426, expectedError: ErrUpgradeRequired},
		{statusCode: 429, expectedError: ErrRateLimitExceeded},
		{statusCode: 500, expectedError: ErrInternalServerError},
	}

	for _, data := range testData {
		t.Run(fmt.Sprintf("Status code: %d", data.statusCode), func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(data.statusCode)
			}))

			client := New("my-api-key", WithEndpoint(server.URL), WithGlobalTimeout(10*time.Second))

			_, err := client.Geocode(context.Background(), "Berlin, Germany", nil)
			if err == nil {
				t.Fatalf("there should be an error")
			}

			if !errors.Is(err, data.expectedError) {
				t.Fatalf("there should be an ErrInvalidRequest")
			}
		})
	}

}
