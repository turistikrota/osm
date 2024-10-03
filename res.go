package osm

import "time"

// ReverseResult represents the result of a reverse geocoding request.
type ReverseResult struct {
	PlaceID     int     `json:"place_id"`     // PlaceID is the unique identifier for the place.
	Licence     string  `json:"licence"`      // Licence is the licence information for the data.
	OsmType     string  `json:"osm_type"`     // OsmType is the type of the OpenStreetMap object (e.g., node, way, relation).
	OsmID       int     `json:"osm_id"`       // OsmID is the unique identifier for the OpenStreetMap object.
	Lat         string  `json:"lat"`          // Lat is the latitude of the place.
	Lon         string  `json:"lon"`          // Lon is the longitude of the place.
	Class       string  `json:"class"`        // Class is the classification of the place (e.g., place, highway).
	Type        string  `json:"type"`         // Type is the type of the place (e.g., city, street).
	PlaceRank   int     `json:"place_rank"`   // PlaceRank is the rank of the place.
	Importance  float64 `json:"importance"`   // Importance is the importance score of the place.
	Addresstype string  `json:"addresstype"`  // Addresstype is the type of address.
	Name        string  `json:"name"`         // Name is the name of the place.
	DisplayName string  `json:"display_name"` // DisplayName is the human-readable name of the place.
	Address     struct {
		Amenity       string `json:"amenity"`        // Amenity is the amenity at the place.
		HouseNumber   string `json:"house_number"`   // HouseNumber is the house number of the place.
		Road          string `json:"road"`           // Road is the road of the place.
		Quarter       string `json:"quarter"`        // Quarter is the quarter of the place.
		Neighbourhood string `json:"neighbourhood"`  // Neighbourhood is the neighbourhood of the place.
		Suburb        string `json:"suburb"`         // Suburb is the suburb of the place.
		County        string `json:"county"`         // County is the county of the place.
		City          string `json:"city"`           // City is the city of the place.
		State         string `json:"state"`          // State is the state of the place.
		ISO31662Lvl4  string `json:"ISO3166-2-lvl4"` // ISO31662Lvl4 is the ISO 3166-2 level 4 code of the place.
		Postcode      string `json:"postcode"`       // Postcode is the postal code of the place.
		Country       string `json:"country"`        // Country is the country of the place.
		CountryCode   string `json:"country_code"`   // CountryCode is the country code of the place.
	} `json:"address"` // Address contains detailed address information.
	Boundingbox []string `json:"boundingbox"` // Boundingbox is the bounding box of the place.
}

// DetailsResult represents the detailed information about a specific OpenStreetMap object.
type DetailsResult struct {
	PlaceID       int    `json:"place_id"`        // PlaceID is the unique identifier for the place.
	ParentPlaceID int    `json:"parent_place_id"` // ParentPlaceID is the unique identifier for the parent place.
	OsmType       string `json:"osm_type"`        // OsmType is the type of the OpenStreetMap object (e.g., node, way, relation).
	OsmID         int    `json:"osm_id"`          // OsmID is the unique identifier for the OpenStreetMap object.
	Category      string `json:"category"`        // Category is the category of the place.
	Type          string `json:"type"`            // Type is the type of the place.
	AdminLevel    any    `json:"admin_level"`     // AdminLevel is the administrative level of the place.
	Localname     string `json:"localname"`       // Localname is the local name of the place.
	Names         struct {
		Name   string `json:"name"`    // Name is the name of the place.
		NameBe string `json:"name:be"` // NameBe is the name of the place in Belarusian.
		NameDe string `json:"name:de"` // NameDe is the name of the place in German.
		NameEs string `json:"name:es"` // NameEs is the name of the place in Spanish.
		NameHe string `json:"name:he"` // NameHe is the name of the place in Hebrew.
		NameKo string `json:"name:ko"` // NameKo is the name of the place in Korean.
		NameLa string `json:"name:la"` // NameLa is the name of the place in Latin.
		NameRu string `json:"name:ru"` // NameRu is the name of the place in Russian.
		NameUk string `json:"name:uk"` // NameUk is the name of the place in Ukrainian.
		NameZh string `json:"name:zh"` // NameZh is the name of the place in Chinese.
	} `json:"names,omitempty"` // Names contains the names of the place in different languages.
	Addresstags struct {
		Postcode string `json:"postcode"` // Postcode is the postal code of the place.
	} `json:"addresstags,omitempty"` // Addresstags contains address-related tags.
	Housenumber          interface{} `json:"housenumber"` // Housenumber is the house number of the place.
	CalculatedPostcode   string      `json:"calculated_postcode"`
	CountryCode          string      `json:"country_code"`
	IndexedDate          time.Time   `json:"indexed_date"`
	Importance           float64     `json:"importance"`
	CalculatedImportance float64     `json:"calculated_importance"`
	Extratags            struct {
		Wikidata  string `json:"wikidata"`
		Wikipedia string `json:"wikipedia"`
	} `json:"extratags,omitempty"`
	CalculatedWikipedia string `json:"calculated_wikipedia"`
	RankAddress         int    `json:"rank_address"`
	RankSearch          int    `json:"rank_search"`
	Isarea              bool   `json:"isarea"`
	Centroid            struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"centroid"`
	Geometry struct {
		Type        string    `json:"type"`
		Coordinates []float64 `json:"coordinates"`
	} `json:"geometry"`
}

type SearchResult struct {
	Address struct {
		ISO31662Lvl4  string `json:"ISO3166-2-lvl4"`
		Borough       string `json:"borough"`
		City          string `json:"city"`
		Country       string `json:"country"`
		CountryCode   string `json:"country_code"`
		Historic      string `json:"historic"`
		HouseNumber   string `json:"house_number"`
		Neighbourhood string `json:"neighbourhood"`
		Postcode      string `json:"postcode"`
		Road          string `json:"road"`
		Suburb        string `json:"suburb"`
	} `json:"address"`
	Boundingbox []string `json:"boundingbox"`
	Class       string   `json:"class"`
	DisplayName string   `json:"display_name"`
	Importance  float64  `json:"importance"`
	Lat         string   `json:"lat"`
	Licence     string   `json:"licence"`
	Lon         string   `json:"lon"`
	OsmID       int      `json:"osm_id"`
	OsmType     string   `json:"osm_type"`
	PlaceID     int      `json:"place_id"`
	Svg         string   `json:"svg"`
	Type        string   `json:"type"`
}

type LookupResult struct {
	PlaceID     int      `json:"place_id"`
	Licence     string   `json:"licence"`
	OsmType     string   `json:"osm_type"`
	OsmID       int      `json:"osm_id"`
	Boundingbox []string `json:"boundingbox"`
	Lat         string   `json:"lat"`
	Lon         string   `json:"lon"`
	DisplayName string   `json:"display_name"`
	Class       string   `json:"class"`
	Type        string   `json:"type"`
	Importance  float64  `json:"importance"`
	Address     struct {
		Tourism     string `json:"tourism"`
		Road        string `json:"road"`
		Suburb      string `json:"suburb"`
		City        string `json:"city"`
		State       string `json:"state"`
		Postcode    string `json:"postcode"`
		Country     string `json:"country"`
		CountryCode string `json:"country_code"`
	} `json:"address"`
	Extratags struct {
		Image              string `json:"image"`
		Heritage           string `json:"heritage"`
		Wikidata           string `json:"wikidata"`
		Architect          string `json:"architect"`
		Wikipedia          string `json:"wikipedia"`
		Wheelchair         string `json:"wheelchair"`
		Description        string `json:"description"`
		HeritageWebsite    string `json:"heritage:website"`
		HeritageOperator   string `json:"heritage:operator"`
		ArchitectWikidata  string `json:"architect:wikidata"`
		YearOfConstruction string `json:"year_of_construction"`
	} `json:"extratags"`
}

type ErrorResult struct {
	Details struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func (e *ErrorResult) Error() string {
	return e.Details.Message
}
