package schemaorg

import (
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"
)

// https://schema.org/GeoCoordinates
type GeoCoordinates struct {
	NameSpace jsonld.NameSpace `jsonld:"https://schema.org/"`
	Prefix    jsonld.Prefix    `jsonld:"schema"`

	// An additional type for the thing.
	AdditionalType            opt.Optional[string] `json:"additionalType,omitempty"`            // https://schema.org/additionalType

	// An alias for the item.
	AlternateName             opt.Optional[string] `json:"alternateName,omitempty"`             // https://schema.org/alternateName

	// A description of the thing.
	Description               opt.Optional[string] `json:"description,omitempty"`               // https://schema.org/description

	// A sub property of description.
	// A short description of the thing used to disambiguate from other, similar things.
	// Information from other properties (in particular, name) may be necessary for the description to be useful for disambiguation.
	DisambiguatingDescription opt.Optional[string] `json:"disambiguatingDescription,omitempty"` // https://schema.org/disambiguatingDescription

	// The elevation of a location (WGS 84).
	// Values may be of the form 'NUMBER UNIT_OF_MEASUREMENT' (e.g., '1,000 m', '3,200 ft') while numbers alone should be assumed to be a value in meters.
	Elevation                 opt.Optional[string] `json:"elevation,omitempty"`                 // https://schema.org/elevation

	// The identifier property represents any kind of identifier for any kind, such as ISBNs, GTIN codes, UUIDs etc.
	Identifier                opt.Optional[string] `json:"identifier,omitempty"`                // https://schema.org/identifier

	// The latitude of a location. For example 37.42242 (WGS 84).
	Latitude                  opt.Optional[string] `json:"latitude,omitempty"`                  // https://schema.org/latitude

	// The longitude of a location. For example -122.08585 (WGS 84).
	Longitude                 opt.Optional[string] `json:"longitude,omitempty"`                 // https://schema.org/longitude

	// The name of the thing.
	Name                      opt.Optional[string] `json:"name,omitempty"`                      // https://schema.org/name

	Type                        json.Const[string] `json:"type" json.value:"GeoCoordinates"`

	// URL of the thing.
	URL                       opt.Optional[string] `json:"url,omitempty"`                       // https://schema.org/url
}
