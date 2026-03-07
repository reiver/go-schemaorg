package schemaorg

import (
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
	jsonldstrings "github.com/reiver/go-jsonld/strings"
	"github.com/reiver/go-opt"
)

// https://schema.org/VirtualLocation
type VirtualLocation struct {
	NameSpace jsonld.NameSpace `jsonld:"https://schema.org/"`
	Prefix    jsonld.Prefix    `jsonld:"schema"`

	ID opt.Optional[string] `json:"@id,omitempty"`
	Type json.Const[string] `json:"@type" json.value:"VirtualLocation"`

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
}

var _ ProtoThing = VirtualLocation{}

func (receiver VirtualLocation) ProtoThing() AnyThing {
	const thingType string = TypeVirtualLocation

	return AnyThing{
		ID:                        receiver.ID,
		Type:                      jsonldstrings.Something(thingType),
		AdditionalType:            receiver.AdditionalType,
		AlternateName:             receiver.AlternateName,
		Description:               receiver.Description,
		DisambiguatingDescription: receiver.DisambiguatingDescription,
		Identifier:                receiver.Identifier,
		Image:                     receiver.Image,
		MainEntityOfPage:          receiver.MainEntityOfPage,
		Name:                      receiver.Name,
		Owner:                     receiver.Owner,
		PotentialAction:           receiver.PotentialAction,
		SameAs:                    receiver.SameAs,
		SubjectOf:                 receiver.SubjectOf,
		URL:                       receiver.URL,
	}
}

var _ ProtoIntangible = VirtualLocation{}

func (receiver VirtualLocation) ProtoIntangible() AnyIntangible {
	const thingType string = TypeVirtualLocation

	return AnyIntangible{
		ID:                        receiver.ID,
		Type:                      jsonldstrings.Something(thingType),
		AdditionalType:            receiver.AdditionalType,
		AlternateName:             receiver.AlternateName,
		Description:               receiver.Description,
		DisambiguatingDescription: receiver.DisambiguatingDescription,
		Identifier:                receiver.Identifier,
		Image:                     receiver.Image,
		MainEntityOfPage:          receiver.MainEntityOfPage,
		Name:                      receiver.Name,
		Owner:                     receiver.Owner,
		PotentialAction:           receiver.PotentialAction,
		SameAs:                    receiver.SameAs,
		SubjectOf:                 receiver.SubjectOf,
		URL:                       receiver.URL,
	}
}
