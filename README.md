<div align="center">
    <h1>opencage</h1>
</div>

<p align="center">
An API client implementation for the OpenCage Geocoder API.
</p>
<div align="center">

[![Build](https://github.com/alexliesenfeld/opencage/actions/workflows/go.yml/badge.svg)](https://github.com/alexliesenfeld/health/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/alexliesenfeld/opencage/branch/main/graph/badge.svg?token=V2mVh8RvYE)](https://codecov.io/gh/alexliesenfeld/health)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexliesenfeld/opencage)](https://goreportcard.com/report/github.com/alexliesenfeld/health)
[![GolangCI](https://golangci.com/badges/github.com/alexliesenfeld/opencage.svg)](https://golangci.com/r/github.com/alexliesenfeld/health)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Falexliesenfeld%2Fopencage.svg?type=small)](https://app.fossa.com/projects/git%2Bgithub.com%2Falexliesenfeld%2Fopencage?ref=badge_small)
</div>

<p align="center">
    <a href="https://pkg.go.dev/github.com/alexliesenfeld/opencage">Documentation</a>
    ·
    <a href="https://github.com/alexliesenfeld/opencage/issues">Report Bug</a>
    ·
    <a href="https://github.com/alexliesenfeld/opencage/issues">Request Feature</a>
</p>


An API client implementation for the OpenCage Geocoder API. It provides a reverse (latitude/longitude to text) and 
forward (text to latitude/longitude) geocoding API. Please visit https://opencagedata.com/api for 
more information.

## How to install
```shell
go get github.com/alexliesenfeld/opencage
```

## Usage
You can import this package into your application as follows:
```go
import "github.com/alexliesenfeld/opencage"
```

Use it as follows:

```go
client := New("my-api-key")

// Can be used to control timeouts (e.g., using context.WithDeadline), cancellation, etc.
ctx := context.Background()

// Call the API with default parameters.
// Also shows how to use forward geocoding API using a location.
response, err := client.Geocode(ctx, "Berlin, Germany", nil)

// Or set your own API parameters.
// Also shows how to use reverse geocoding API using latitude and longitude.
response, err := client.Geocode(ctx, "52.3877830 9.7334394", &GeocodingParams{
    RoadInfo:  true,
    Proximity: []float32{1.0, -1.0},
    Language:  "de",
})
```

**Please refer to [this file](https://github.com/alexliesenfeld/opencage/blob/main/data.go) to see more about how the 
response looks like and what parameters you are able to set.**

## Legal

Please note that the OpenCage API is property of OpenCage GmbH. 
Head over to https://opencagedata.com for more information.

### License

This library is free software: you can redistribute it and/or modify it under the terms of the MIT Public License.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied
warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the MIT Public License for more details.

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Falexliesenfeld%2Fopencage.svg?type=large&issueType=license)]