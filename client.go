package osm

import (
	"context"
	"fmt"
	"net/url"
)

const (
	OsmTypeRelation = "R"
	OsmTypeNode     = "N"
)

// Reverse performs a reverse geocoding request to convert latitude and longitude into a human-readable address.
// It takes a context, latitude, longitude, and optional functions for additional configurations.
// It returns a ReverseResult and an error if the request fails.
func Reverse(ctx context.Context, lat float64, long float64, opts ...optFunc) (*ReverseResult, error) {
	urlVals := url.Values{}
	urlVals.Set("format", "json")
	urlVals.Set("lat", fmt.Sprintf("%f", lat))
	urlVals.Set("lon", fmt.Sprintf("%f", long))
	return runRequest[ReverseResult](ctx, apiUrl+"reverse?"+urlVals.Encode(), opts...)
}

// Details retrieves detailed information about a specific OpenStreetMap object.
// It takes a context, OSM type, OSM ID, and optional functions for additional configurations.
// It returns a DetailsResult and an error if the request fails.
func Details(ctx context.Context, osmType string, osmID int, opts ...optFunc) (*DetailsResult, error) {
	urlVals := url.Values{}
	urlVals.Set("format", "json")
	urlVals.Set("osmtype", osmType)
	urlVals.Set("osmid", fmt.Sprintf("%d", osmID))
	return runRequest[DetailsResult](ctx, apiUrl+"details?"+urlVals.Encode(), opts...)
}

// DetailsWithPlaceID retrieves detailed information about a specific place using its place ID.
// It takes a context, place ID, and optional functions for additional configurations.
// It returns a DetailsResult and an error if the request fails.
func DetailsWithPlaceID(ctx context.Context, placeID int, opts ...optFunc) (*DetailsResult, error) {
	urlVals := url.Values{}
	urlVals.Set("format", "json")
	urlVals.Set("place_id", fmt.Sprintf("%d", placeID))
	return runRequest[DetailsResult](ctx, apiUrl+"details?"+urlVals.Encode(), opts...)
}

// Search performs a search query to find places matching the given query string.
// It takes a context, query string, and optional functions for additional configurations.
// It returns a slice of SearchResult and an error if the request fails.
func Search(ctx context.Context, q string, opts ...optFunc) ([]SearchResult, error) {
	urlVals := url.Values{}
	urlVals.Set("format", "json")
	urlVals.Set("q", q)
	res, err := runRequest[[]SearchResult](ctx, apiUrl+"search?"+urlVals.Encode(), opts...)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return []SearchResult{}, nil
	}
	return *res, nil
}

// Lookup retrieves information about a specific OpenStreetMap object using its type and ID.
// It takes a context, OSM type, OSM ID, and optional functions for additional configurations.
// It returns a LookupResult and an error if the request fails.
func Lookup(ctx context.Context, osmType string, osmID int, opts ...optFunc) ([]LookupResult, error) {
	urlVals := url.Values{}
	urlVals.Set("format", "json")
	urlVals.Set("osm_type", osmType)
	urlVals.Set("osm_id", fmt.Sprintf("%d", osmID))
	res, err := runRequest[[]LookupResult](ctx, apiUrl+"lookup?"+urlVals.Encode(), opts...)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return []LookupResult{}, nil
	}
	return *res, nil
}
