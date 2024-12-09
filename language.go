package schemaorg

import (
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"
)

// https://schema.org/Language
type Language struct {
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

	// The identifier property represents any kind of identifier for any kind, such as ISBNs, GTIN codes, UUIDs etc.
	Identifier                opt.Optional[string] `json:"identifier,omitempty"`                // https://schema.org/identifier

	// The name of the thing.
	Name                      opt.Optional[string] `json:"name,omitempty"`                      // https://schema.org/name

	Type                        json.Const[string] `json:"type" json.value:"Language"`

	// URL of the thing.
	URL                       opt.Optional[string] `json:"url,omitempty"`                       // https://schema.org/url
}
