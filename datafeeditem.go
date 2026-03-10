package schemaorg

import (
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
	jsonldstrings "github.com/reiver/go-jsonld/strings"
	"github.com/reiver/go-opt"
)

// https://schema.org/DataFeedItem
type DataFeedItem struct {
	NameSpace jsonld.NameSpace `jsonld:"https://schema.org/"`
	Prefix    jsonld.Prefix    `jsonld:"schema"`

	ID opt.Optional[string] `json:"@id,omitempty"`
	Type json.Const[string] `json:"@type" json.value:"DataFeedItem"`

	// ---- Thing properties ----

	AdditionalType            jsonld.Strings       `json:"additionalType,omitempty"`            // https://schema.org/additionalType
	AlternateName             opt.Optional[string] `json:"alternateName,omitempty"`             // https://schema.org/alternateName
	Description               opt.Optional[string] `json:"description,omitempty"`               // https://schema.org/description
	DisambiguatingDescription opt.Optional[string] `json:"disambiguatingDescription,omitempty"` // https://schema.org/disambiguatingDescription
	Identifier                opt.Optional[string] `json:"identifier,omitempty"`                // https://schema.org/identifier
	Image                     opt.Optional[string] `json:"image,omitempty"`                     // https://schema.org/image
	MainEntityOfPage       []ProtoCreativeWork `json:"mainEntityOfPage,omitempty"`          // https://schema.org/mainEntityOfPage
	Name                      opt.Optional[string] `json:"name,omitempty"`                      // https://schema.org/name
	Owner                     opt.Optional[string] `json:"owner,omitempty"`                     // https://schema.org/owner
	PotentialAction           opt.Optional[string] `json:"potentialAction,omitempty"`           // https://schema.org/potentialAction
	SameAs                    opt.Optional[string] `json:"sameAs,omitempty"`                    // https://schema.org/sameAs
	SubjectOf                 opt.Optional[string] `json:"subjectOf,omitempty"`                 // https://schema.org/subjectOf
	URL                       opt.Optional[string] `json:"url,omitempty"`                       // https://schema.org/url

	// ---- DataFeedItem properties ----

	DateCreated               opt.Optional[string] `json:"dateCreated,omitempty"`               // https://schema.org/dateCreated
	DateDeleted               opt.Optional[string] `json:"dateDeleted,omitempty"`               // https://schema.org/dateDeleted
	DateModified              opt.Optional[string] `json:"dateModified,omitempty"`              // https://schema.org/dateModified
	Item                      ProtoThing           `json:"item,omitempty"`                      // https://schema.org/item
}

var _ ProtoThing = DataFeedItem{}

func (receiver DataFeedItem) ProtoThing() AnyThing {
	const thingType string = TypeDataFeedItem

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

var _ ProtoIntangible = DataFeedItem{}

func (receiver DataFeedItem) ProtoIntangible() AnyIntangible {
	const thingType string = TypeDataFeedItem

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
