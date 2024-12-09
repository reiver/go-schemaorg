package schemaorg

import (
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"
)

// https://schema.org/Role
type Role struct {
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

	// The end date and time of the item (in ISO 8601 date format).
	EndDate                   opt.Optional[string] `json:"endDate,omitempty"`                   // https://schema.org/endDate

	// The identifier property represents any kind of identifier for any kind, such as ISBNs, GTIN codes, UUIDs etc.
	Identifier                opt.Optional[string] `json:"identifier,omitempty"`                // https://schema.org/identifier

	// The name of the thing.
	Name                      opt.Optional[string] `json:"name,omitempty"`                      // https://schema.org/name

	// A role played, performed or filled by a person or organization.
	// For example, the team of creators for a comic book might fill the roles named 'inker', 'penciller', and 'letterer'; or an athlete in a SportsTeam might play in the position named 'Quarterback'.
	RoleName                  opt.Optional[string] `json:"roleName,omitempty"`                  // https://schema.org/roleName

	// The start date and time of the item (in ISO 8601 date format).
	StartDate                 opt.Optional[string] `json:"startDate,omitempty"`                 // https://schema.org/startDate

	Type                        json.Const[string] `json:"type" json.value:"Role"`

	// URL of the thing.
	URL                       opt.Optional[string] `json:"url,omitempty"`                       // https://schema.org/url
}
