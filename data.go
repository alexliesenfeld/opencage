package opencage

type Licenses struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Rate struct {
	Limit     int   `json:"limit"`
	Remaining int   `json:"remaining"`
	Reset     int64 `json:"reset"`
}

type DMS struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Mercator struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type NUTSCode struct {
	Code string `json:"code"`
}

type NUTS struct {
	NUTS0 NUTSCode `json:"NUTS0"`
	NUTS1 NUTSCode `json:"NUTS1"`
	NUTS2 NUTSCode `json:"NUTS2"`
	NUTS3 NUTSCode `json:"NUTS3"`
}

type OSM struct {
	EditUrl string `json:"edit_url"`
	NoteUrl string `json:"note_url"`
	Url     string `json:"url"`
}

type Region struct {
	DE            string `json:"DE"`
	EUROPE        string `json:"EUROPE"`
	WESTERNEUROPE string `json:"WESTERN_EUROPE"`
	WORLD         string `json:"WORLD"`
}

type UNM49 struct {
	Regions              Region   `json:"regions"`
	StatisticalGroupings []string `json:"statistical_groupings"`
}

type Currency struct {
	AlternateSymbols     []interface{} `json:"alternate_symbols"`
	DecimalMark          string        `json:"decimal_mark"`
	HtmlEntity           string        `json:"html_entity"`
	IsoCode              string        `json:"iso_code"`
	IsoNumeric           string        `json:"iso_numeric"`
	Name                 string        `json:"name"`
	SmallestDenomination int           `json:"smallest_denomination"`
	Subunit              string        `json:"subunit"`
	SubunitToUnit        int           `json:"subunit_to_unit"`
	Symbol               string        `json:"symbol"`
	SymbolFirst          int           `json:"symbol_first"`
	ThousandsSeparator   string        `json:"thousands_separator"`
}

type RoadInfo struct {
	Road    string `json:"road"`
	DriveOn string `json:"drive_on"`
	SpeedIn string `json:"speed_in"`
}

type SunInfo struct {
	Apparent     int `json:"apparent"`
	Astronomical int `json:"astronomical"`
	Civil        int `json:"civil"`
	Nautical     int `json:"nautical"`
}

type Sun struct {
	Rise SunInfo `json:"rise"`
	Set  SunInfo `json:"set"`
}

type Timezone struct {
	Name         string `json:"name"`
	NowInDst     int    `json:"now_in_dst"`
	OffsetSec    int    `json:"offset_sec"`
	OffsetString string `json:"offset_string"`
	ShortName    string `json:"short_name"`
}

type What3Words struct {
	Words string `json:"words"`
}

type Annotations struct {
	DMS         DMS        `json:"DMS"`
	MGRS        string     `json:"MGRS"`
	Maidenhead  string     `json:"Maidenhead"`
	Mercator    Mercator   `json:"Mercator"`
	NUTS        NUTS       `json:"NUTS"`
	OSM         OSM        `json:"OSM"`
	UNM49       UNM49      `json:"UN_M49"`
	Callingcode int        `json:"callingcode"`
	Currency    Currency   `json:"currency"`
	Flag        string     `json:"flag"`
	Geohash     string     `json:"geohash"`
	Qibla       float64    `json:"qibla"`
	Roadinfo    RoadInfo   `json:"roadinfo"`
	Sun         Sun        `json:"sun"`
	Timezone    Timezone   `json:"timezone"`
	What3Words  What3Words `json:"what3words"`
}

type Bounds struct {
	Northeast Coordinates `json:"northeast"`
	Southwest Coordinates `json:"southwest"`
}

type ResultComponents struct {
	ISO31661Alpha2 string   `json:"ISO_3166-1_alpha-2"`
	ISO31661Alpha3 string   `json:"ISO_3166-1_alpha-3"`
	ISO31662       []string `json:"ISO_3166-2"`
	Category       string   `json:"_category"`
	Type           string   `json:"_type"`
	City           string   `json:"city"`
	CityDistrict   string   `json:"city_district"`
	Continent      string   `json:"continent"`
	Country        string   `json:"country"`
	CountryCode    string   `json:"country_code"`
	County         string   `json:"county"`
	Municipality   string   `json:"municipality"`
	HouseNumber    string   `json:"house_number"`
	Office         string   `json:"office"`
	PoliticalUnion string   `json:"political_union"`
	Postcode       string   `json:"postcode"`
	Road           string   `json:"road"`
	State          string   `json:"state"`
	StateCode      string   `json:"state_code"`
	Suburb         string   `json:"suburb"`
	Town           string   `json:"town"`
	Village        string   `json:"village"`
	Pedestrian     string   `json:"pedestrian"`
	Neighbourhood  string   `json:"neighbourhood"`
	Region         string   `json:"region"`
}

type Geometry struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Result struct {
	Annotations Annotations      `json:"annotations"`
	Bounds      Bounds           `json:"bounds"`
	Components  ResultComponents `json:"components"`
	Confidence  int              `json:"confidence"`
	Formatted   string           `json:"formatted"`
	Geometry    Geometry         `json:"geometry"`
}

type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type StayInformed struct {
	Blog     string `json:"blog"`
	Mastodon string `json:"mastodon"`
}

type Timestamp struct {
	CreatedHttp string `json:"created_http"`
	CreatedUnix int    `json:"created_unix"`
}

type Response struct {
	Documentation string       `json:"documentation"`
	Licenses      []Licenses   `json:"licenses"`
	Rate          Rate         `json:"rate"`
	Results       []Result     `json:"results"`
	Status        Status       `json:"status"`
	StayInformed  StayInformed `json:"stay_informed"`
	Thanks        string       `json:"thanks"`
	Timestamp     Timestamp    `json:"timestamp"`
	TotalResults  int          `json:"total_results"`
}

type GeocodingParams struct {
	// Abbreviate and shorten the formatted string that is returned
	Abbreviate bool

	// AddressOnly will have the following effect: When set to true, the formatted string returned will solely
	// consist of the address, excluding the names of Points of Interest (POIs).
	AddressOnly bool

	// AddRequest will have the following effect: When true, the response includes various request parameters
	// to facilitate ease of debugging.
	AddRequest bool

	// Bounds limits the potential outcomes to a specified bounding box.
	// FORWARD GEOCODING ONLY!
	Bounds []float64

	// CountryCode limits search outcomes to a particular country or set of countries.
	// This code is a two-letter designation according to the ISO 3166-1 Alpha 2 standard, such as 'gb'
	// for the United Kingdom, 'fr' for France, and 'us' for the United States.
	// Country codes that don't conform to the two-letter format will not be considered.
	// FORWARD GEOCODING ONLY!
	CountryCode string

	// JSON wraps the returned data with a function name (AJAX related, see https://en.wikipedia.org/wiki/JSONP).
	JSONPFunctionName string

	// Language is a language code in IETF format (like 'es' for Spanish or 'pt-BR' for Brazilian Portuguese),
	// or 'native', which signals an attempt to provide the response in the local language(s).
	Language string

	// Limit is the upper limit for the number of results to be returned. The default setting is 10,
	// with the highest permissible value being 100.
	Limit int

	// NoAnnotations turns off annotations.
	NoAnnotations bool

	// NoDedupe disables deduplication of results.
	NoDedupe bool

	// NoRecord logging of query contents.
	NoRecord bool

	// Pretty will pretty print the response content. This will most likely only be
	// visible to this library and is transparent to you as a developer.
	Pretty bool

	// Proximity offers a suggestion to the geocoder to prioritize results nearer to a
	// given location. However, it's important to note that this is merely one of several
	// elements used in the internal scoring system for ranking results.
	// FORWARD GEOCODING ONLY!
	Proximity []float64

	// RoadInfo adjusts the geocoder's function is altered to try and locate the nearest road,
	// rather than an address.
	RoadInfo bool
}
