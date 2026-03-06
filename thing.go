package schemaorg

import (
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"
)

// https://schema.org/Thing
type Thing struct {
	NameSpace jsonld.NameSpace `jsonld:"https://schema.org/"`
	Prefix    jsonld.Prefix    `jsonld:"schema"`

	ID opt.Optional[string] `json:"@id"`
	Type opt.Optional[string] `json:"type,omitempty"`

	AdditionalType            opt.Optional[string] `json:"additionalType,omitempty"`            // https://schema.org/additionalType
	AlternateName             opt.Optional[string] `json:"alternateName,omitempty"`             // https://schema.org/alternateName
	Description               opt.Optional[string] `json:"description,omitempty"`               // https://schema.org/description
	DisambiguatingDescription opt.Optional[string] `json:"disambiguatingDescription,omitempty"` // https://schema.org/disambiguatingDescription
	Identifier                opt.Optional[string] `json:"identifier,omitempty"`                // https://schema.org/identifier
	Name                      opt.Optional[string] `json:"name,omitempty"`                      // https://schema.org/name
	URL                       opt.Optional[string] `json:"url,omitempty"`                       // https://schema.org/url
}
