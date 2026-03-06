package schemaorg

import (
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"
)

//https://schema.org/MonetaryAmount
type MonetaryAmount struct {
	NameSpace jsonld.NameSpace `jsonld:"https://schema.org/"`
	Prefix    jsonld.Prefix    `jsonld:"schema"`

	ID opt.Optional[string] `json:"@id,omitempty"`
	Type json.Const[string] `json:"@type" json.value:"MonetaryAmount"`

	AdditionalType            opt.Optional[string] `json:"additionalType,omitempty"`            // https://schema.org/additionalType
	AlternateName             opt.Optional[string] `json:"alternateName,omitempty"`             // https://schema.org/alternateName
	Currency                  opt.Optional[string] `json:"currency,omitempty"`                  // https://schema.org/currency
	Description               opt.Optional[string] `json:"description,omitempty"`               // https://schema.org/description
	DisambiguatingDescription opt.Optional[string] `json:"disambiguatingDescription,omitempty"` // https://schema.org/disambiguatingDescription
	Identifier                opt.Optional[string] `json:"identifier,omitempty"`                // https://schema.org/identifier
	Image                     opt.Optional[string] `json:"image,omitempty"`                     // https://schema.org/image
	MainEntityOfPage       []ProtoCreativeWork `json:"mainEntityOfPage,omitempty"`          // https://schema.org/mainEntityOfPage
	MaxValue                  opt.Optional[string] `json:"maxValue,omitempty,bare"`             // https://schema.org/maxValue
	MinValue                  opt.Optional[string] `json:"minValue,omitempty,bare"`             // https://schema.org/minValue
	Name                      opt.Optional[string] `json:"name,omitempty"`                      // https://schema.org/name
	Owner                     opt.Optional[string] `json:"owner,omitempty"`                     // https://schema.org/owner
	PotentialAction           opt.Optional[string] `json:"potentialAction,omitempty"`           // https://schema.org/potentialAction
	SameAs                    opt.Optional[string] `json:"sameAs,omitempty"`                    // https://schema.org/sameAs
	SubjectOf                 opt.Optional[string] `json:"subjectOf,omitempty"`                 // https://schema.org/subjectOf
	URL                       opt.Optional[string] `json:"url,omitempty"`                       // https://schema.org/url
	Value                     opt.Optional[string] `json:"value,omitempty"`                     // https://schema.org/value
}
