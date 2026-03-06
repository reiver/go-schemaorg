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

	ID opt.Optional[string] `json:"@id"`
	Type json.Const[string] `json:"@type" json.value:"GeoCoordinates"`

	AdditionalType            opt.Optional[string] `json:"additionalType,omitempty"`            // https://schema.org/additionalType
	AlternateName             opt.Optional[string] `json:"alternateName,omitempty"`             // https://schema.org/alternateName
	Description               opt.Optional[string] `json:"description,omitempty"`               // https://schema.org/description
	DisambiguatingDescription opt.Optional[string] `json:"disambiguatingDescription,omitempty"` // https://schema.org/disambiguatingDescription
	Identifier                opt.Optional[string] `json:"identifier,omitempty"`                // https://schema.org/identifier
	Name                      opt.Optional[string] `json:"name,omitempty"`                      // https://schema.org/name
	URL                       opt.Optional[string] `json:"url,omitempty"`                       // https://schema.org/url

	Elevation                 opt.Optional[string] `json:"elevation,omitempty"`                 // https://schema.org/elevation
	Latitude                  opt.Optional[string] `json:"latitude,omitempty"`                  // https://schema.org/latitude
	Longitude                 opt.Optional[string] `json:"longitude,omitempty"`                 // https://schema.org/longitude
}
