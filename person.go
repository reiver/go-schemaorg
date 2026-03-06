package schemaorg

import (
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
	jsonldstrings "github.com/reiver/go-jsonld/strings"
	"github.com/reiver/go-opt"
)

// https://schema.org/Person
type Person struct {
	NameSpace jsonld.NameSpace `jsonld:"https://schema.org/"`
	Prefix    jsonld.Prefix    `jsonld:"schema"`

	ID opt.Optional[string] `json:"@id,omitempty"`
	Type json.Const[string] `json:"@type" json.value:"Person"`

	// ---- Thing properties ----

	AdditionalType            opt.Optional[string] `json:"additionalType,omitempty"`            // https://schema.org/additionalType
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

	// ---- Person properties ----

	AdditionalName            opt.Optional[string] `json:"additionalName,omitempty"`            // https://schema.org/additionalName
	Address                   opt.Optional[string] `json:"address,omitempty"`                   // https://schema.org/address
	Affiliation               opt.Optional[string] `json:"affiliation,omitempty"`               // https://schema.org/affiliation
	AgentInteractionStatistic opt.Optional[string] `json:"agentInteractionStatistic,omitempty"` // https://schema.org/agentInteractionStatistic
	AlumniOf                  opt.Optional[string] `json:"alumniOf,omitempty"`                  // https://schema.org/alumniOf
	Award                     opt.Optional[string] `json:"award,omitempty"`                     // https://schema.org/award
	BirthDate                 opt.Optional[string] `json:"birthDate,omitempty"`                 // https://schema.org/birthDate
	BirthPlace                opt.Optional[string] `json:"birthPlace,omitempty"`                // https://schema.org/birthPlace
	Brand                     opt.Optional[string] `json:"brand,omitempty"`                     // https://schema.org/brand
	CallSign                  opt.Optional[string] `json:"callSign,omitempty"`                  // https://schema.org/callSign
	Children                  opt.Optional[string] `json:"children,omitempty"`                  // https://schema.org/children
	Colleague                 opt.Optional[string] `json:"colleague,omitempty"`                 // https://schema.org/colleague
	ContactPoint              opt.Optional[string] `json:"contactPoint,omitempty"`              // https://schema.org/contactPoint
	DeathDate                 opt.Optional[string] `json:"deathDate,omitempty"`                 // https://schema.org/deathDate
	DeathPlace                opt.Optional[string] `json:"deathPlace,omitempty"`                // https://schema.org/deathPlace
	Duns                      opt.Optional[string] `json:"duns,omitempty"`                      // https://schema.org/duns
	Email                     opt.Optional[string] `json:"email,omitempty"`                     // https://schema.org/email
	FamilyName                opt.Optional[string] `json:"familyName,omitempty"`                // https://schema.org/familyName
	FaxNumber                 opt.Optional[string] `json:"faxNumber,omitempty"`                 // https://schema.org/faxNumber
	Follows                   opt.Optional[string] `json:"follows,omitempty"`                   // https://schema.org/follows
	Funder                    opt.Optional[string] `json:"funder,omitempty"`                    // https://schema.org/funder
	Funding                   opt.Optional[string] `json:"funding,omitempty"`                   // https://schema.org/funding
	Gender                    opt.Optional[string] `json:"gender,omitempty"`                    // https://schema.org/gender
	GivenName                 opt.Optional[string] `json:"givenName,omitempty"`                 // https://schema.org/givenName
	GlobalLocationNumber      opt.Optional[string] `json:"globalLocationNumber,omitempty"`      // https://schema.org/globalLocationNumber
	HasCertification          opt.Optional[string] `json:"hasCertification,omitempty"`          // https://schema.org/hasCertification
	HasCredential             opt.Optional[string] `json:"hasCredential,omitempty"`             // https://schema.org/hasCredential
	HasOccupation             opt.Optional[string] `json:"hasOccupation,omitempty"`             // https://schema.org/hasOccupation
	HasOfferCatalog           opt.Optional[string] `json:"hasOfferCatalog,omitempty"`           // https://schema.org/hasOfferCatalog
	HasPOS                    opt.Optional[string] `json:"hasPOS,omitempty"`                    // https://schema.org/hasPOS
	Height                    opt.Optional[string] `json:"height,omitempty"`                    // https://schema.org/height
	HomeLocation              opt.Optional[string] `json:"homeLocation,omitempty"`              // https://schema.org/homeLocation
	HonorificPrefix           opt.Optional[string] `json:"honorificPrefix,omitempty"`           // https://schema.org/honorificPrefix
	HonorificSuffix           opt.Optional[string] `json:"honorificSuffix,omitempty"`           // https://schema.org/honorificSuffix
	InteractionStatistic      opt.Optional[string] `json:"interactionStatistic,omitempty"`      // https://schema.org/interactionStatistic
	IsicV4                    opt.Optional[string] `json:"isicV4,omitempty"`                    // https://schema.org/isicV4
	JobTitle                  opt.Optional[string] `json:"jobTitle,omitempty"`                  // https://schema.org/jobTitle
	Knows                     opt.Optional[string] `json:"knows,omitempty"`                     // https://schema.org/knows
	KnowsAbout                opt.Optional[string] `json:"knowsAbout,omitempty"`                // https://schema.org/knowsAbout
	KnowsLanguage             opt.Optional[string] `json:"knowsLanguage,omitempty"`             // https://schema.org/knowsLanguage
	LifeEvent                 opt.Optional[string] `json:"lifeEvent,omitempty"`                 // https://schema.org/lifeEvent
	MakesOffer                opt.Optional[string] `json:"makesOffer,omitempty"`                // https://schema.org/makesOffer
	MemberOf                  opt.Optional[string] `json:"memberOf,omitempty"`                  // https://schema.org/memberOf
	Naics                     opt.Optional[string] `json:"naics,omitempty"`                     // https://schema.org/naics
	Nationality               opt.Optional[string] `json:"nationality,omitempty"`               // https://schema.org/nationality
	NetWorth                  opt.Optional[string] `json:"netWorth,omitempty"`                  // https://schema.org/netWorth
	Owns                      opt.Optional[string] `json:"owns,omitempty"`                      // https://schema.org/owns
	Parent                    opt.Optional[string] `json:"parent,omitempty"`                    // https://schema.org/parent
	PerformerIn               opt.Optional[string] `json:"performerIn,omitempty"`               // https://schema.org/performerIn
	Pronouns                  opt.Optional[string] `json:"pronouns,omitempty"`                  // https://schema.org/pronouns
	PublishingPrinciples      opt.Optional[string] `json:"publishingPrinciples,omitempty"`      // https://schema.org/publishingPrinciples
	RelatedTo                 opt.Optional[string] `json:"relatedTo,omitempty"`                 // https://schema.org/relatedTo
	Seeks                     opt.Optional[string] `json:"seeks,omitempty"`                     // https://schema.org/seeks
	Sibling                   opt.Optional[string] `json:"sibling,omitempty"`                   // https://schema.org/sibling
	Skills                    opt.Optional[string] `json:"skills,omitempty"`                    // https://schema.org/skills
	Sponsor                   opt.Optional[string] `json:"sponsor,omitempty"`                   // https://schema.org/sponsor
	Spouse                    opt.Optional[string] `json:"spouse,omitempty"`                    // https://schema.org/spouse
	Telephone                 opt.Optional[string] `json:"telephone,omitempty"`                 // https://schema.org/telephone
	VatID                     opt.Optional[string] `json:"vatID,omitempty"`                     // https://schema.org/vatID
	Weight                    opt.Optional[string] `json:"weight,omitempty"`                    // https://schema.org/weight
	WorkLocation              opt.Optional[string] `json:"workLocation,omitempty"`              // https://schema.org/workLocation
	WorksFor                  opt.Optional[string] `json:"worksFor,omitempty"`                  // https://schema.org/worksFor
}

var _ ProtoThing = Person{}

func (receiver Person) ProtoThing() AnyThing {
	const thingType string = TypePerson

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
