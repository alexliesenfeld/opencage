# opencage
An API client implementation for the OpenCage Geocoder API. It provides a reverse (latitude/longitude to text) and 
forward (text to latitude/longitude) geocoding API. Please visit https://opencagedata.com/api#quickstart for 
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

// Call the API with default parameters (also shows how to use forward geocoding API using a location)
response, err := client.Geocode(context.Background(), "Berlin, Germany", nil)

// Or set your own API parameters (also shows how to use reverse geocoding API using coordingates)
response, err := client.Geocode(context.Background(), "52.3877830 9.7334394", &GeocodingParams{
    RoadInfo:  true,
    Proximity: []float32{1.0, -1.0},
    Language:  "de",
})
```
## Legal

Please note that the OpenCage API and is property of OpenCage GmbH. Head over to https://opencagedata.com for more
information.

### License

This software is free software: you can redistribute it and/or modify it under the terms of the MIT Public License.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied
warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the MIT Public License for more details.