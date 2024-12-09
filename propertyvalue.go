package schemaorg

import (
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"
)

// https://schema.org/PropertyValue
type PropertyValue struct {
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

	// The upper value of some characteristic or property.
	MaxValue                  opt.Optional[string] `json:"maxValue,omitempty,bare"`             // https://schema.org/maxValue

	// The lower value of some characteristic or property.
	MinValue                  opt.Optional[string] `json:"minValue,omitempty,bare"`             // https://schema.org/minValue

	// The name of the thing.
	Name                      opt.Optional[string] `json:"name,omitempty"`                      // https://schema.org/name

	// A commonly used identifier for the characteristic represented by the property, e.g. a manufacturer or a standard code for a property.
	// propertyID can be (1) a prefixed string, mainly meant to be used with standards for product properties; (2) a site-specific, non-prefixed string (e.g. the primary key of the property or the vendor-specific ID of the property), or (3) a URL indicating the type of the property, either pointing to an external vocabulary, or a Web resource that describes the property (e.g. a glossary entry).
	PropertyID                opt.Optional[string] `json:"propertyID,omitempty"`                // https://schema.org/propertyID

	Type                        json.Const[string] `json:"type" json.value:"PropertyValue"`

	// The unit of measurement given using the UN/CEFACT Common Code (3 characters) or a URL.
	// Other codes than the UN/CEFACT Common Code may be used with a prefix followed by a colon.
	UnitCode                  opt.Optional[string] `json:"unitCode,omitempty"`                  // https://schema.org/unitCode

	// A string or text indicating the unit of measurement.
	// Useful if you cannot provide a standard unit code for unitCode.
	UnitText                  opt.Optional[string] `json:"unitText,omitempty"`                  // https://schema.org/unitText

	// URL of the thing.
	URL                       opt.Optional[string] `json:"url,omitempty"`                       // https://schema.org/url

	Value                     opt.Optional[string] `json:"value,omitempty"`                     // https://schema.org/value
}
