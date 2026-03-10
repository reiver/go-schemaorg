package schemaorg

import (
	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"
)

// AnyAction represents a general schema.org Action, that could be used to represent any schema.org Action type including sub-types.
type AnyAction struct {
	NameSpace jsonld.NameSpace `jsonld:"https://schema.org/"`
	Prefix    jsonld.Prefix    `jsonld:"schema"`

	ID   opt.Optional[string] `json:"@id,omitempty"`
	Type jsonld.Strings       `json:"@type,omitempty"`

	// ---- Thing properties ----

	AdditionalType            jsonld.Strings       `json:"additionalType,omitempty"`            // https://schema.org/additionalType
	AlternateName             opt.Optional[string] `json:"alternateName,omitempty"`             // https://schema.org/alternateName
	Description               opt.Optional[string] `json:"description,omitempty"`               // https://schema.org/description
	DisambiguatingDescription opt.Optional[string] `json:"disambiguatingDescription,omitempty"` // https://schema.org/disambiguatingDescription
	Identifier                opt.Optional[string] `json:"identifier,omitempty"`                // https://schema.org/identifier
	Image                     opt.Optional[string] `json:"image,omitempty"`                     // https://schema.org/image
	MainEntityOfPage          []ProtoCreativeWork  `json:"mainEntityOfPage,omitempty"`          // https://schema.org/mainEntityOfPage
	Name                      opt.Optional[string] `json:"name,omitempty"`                      // https://schema.org/name
	Owner                     opt.Optional[string] `json:"owner,omitempty"`                     // https://schema.org/owner
	PotentialAction           opt.Optional[string] `json:"potentialAction,omitempty"`           // https://schema.org/potentialAction
	SameAs                    opt.Optional[string] `json:"sameAs,omitempty"`                    // https://schema.org/sameAs
	SubjectOf                 opt.Optional[string] `json:"subjectOf,omitempty"`                 // https://schema.org/subjectOf
	URL                       opt.Optional[string] `json:"url,omitempty"`                       // https://schema.org/url

	// ---- Action properties ----

	ActionProcess             opt.Optional[string] `json:"actionProcess,omitempty"`             // https://schema.org/actionProcess
	ActionStatus              opt.Optional[string] `json:"actionStatus,omitempty"`              // https://schema.org/actionStatus
	Agent                     opt.Optional[string] `json:"agent,omitempty"`                     // https://schema.org/agent
	EndTime                   opt.Optional[string] `json:"endTime,omitempty"`                   // https://schema.org/endTime
	Error                     ProtoThing           `json:"error,omitempty"`                     // https://schema.org/error
	Instrument                ProtoThing           `json:"instrument,omitempty"`                // https://schema.org/instrument
	Location                  opt.Optional[string] `json:"location,omitempty"`                  // https://schema.org/location
	Object                    ProtoThing           `json:"object,omitempty"`                    // https://schema.org/object
	Participant               opt.Optional[string] `json:"participant,omitempty"`               // https://schema.org/participant
	Provider                  opt.Optional[string] `json:"provider,omitempty"`                  // https://schema.org/provider
	Result                    ProtoThing           `json:"result,omitempty"`                    // https://schema.org/result
	StartTime                 opt.Optional[string] `json:"startTime,omitempty"`                 // https://schema.org/startTime
	Target                    opt.Optional[string] `json:"target,omitempty"`                    // https://schema.org/target
}
