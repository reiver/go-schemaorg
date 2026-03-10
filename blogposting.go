package schemaorg

import (
	"github.com/reiver/go-json"
	"github.com/reiver/go-jsonld"
	jsonldstrings "github.com/reiver/go-jsonld/strings"
	"github.com/reiver/go-opt"
)

// https://schema.org/BlogPosting
type BlogPosting struct {
	NameSpace jsonld.NameSpace `jsonld:"https://schema.org/"`
	Prefix    jsonld.Prefix    `jsonld:"schema"`

	ID opt.Optional[string] `json:"@id,omitempty"`
	Type json.Const[string] `json:"@type" json.value:"BlogPosting"`

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

	// ---- CreativeWork properties ----

	About                     ProtoThing           `json:"about,omitempty"`                     // https://schema.org/about
	Abstract                  opt.Optional[string] `json:"abstract,omitempty"`                  // https://schema.org/abstract
	AccessMode                opt.Optional[string] `json:"accessMode,omitempty"`                // https://schema.org/accessMode
	AccessModeSufficient      opt.Optional[string] `json:"accessModeSufficient,omitempty"`      // https://schema.org/accessModeSufficient
	AccessibilityAPI          opt.Optional[string] `json:"accessibilityAPI,omitempty"`          // https://schema.org/accessibilityAPI
	AccessibilityControl      opt.Optional[string] `json:"accessibilityControl,omitempty"`      // https://schema.org/accessibilityControl
	AccessibilityFeature      opt.Optional[string] `json:"accessibilityFeature,omitempty"`      // https://schema.org/accessibilityFeature
	AccessibilityHazard       opt.Optional[string] `json:"accessibilityHazard,omitempty"`       // https://schema.org/accessibilityHazard
	AccessibilitySummary      opt.Optional[string] `json:"accessibilitySummary,omitempty"`      // https://schema.org/accessibilitySummary
	AccountablePerson         opt.Optional[string] `json:"accountablePerson,omitempty"`         // https://schema.org/accountablePerson
	AcquireLicensePage        opt.Optional[string] `json:"acquireLicensePage,omitempty"`        // https://schema.org/acquireLicensePage
	AggregateRating           opt.Optional[string] `json:"aggregateRating,omitempty"`           // https://schema.org/aggregateRating
	AlternativeHeadline       opt.Optional[string] `json:"alternativeHeadline,omitempty"`       // https://schema.org/alternativeHeadline
	ArchivedAt                opt.Optional[string] `json:"archivedAt,omitempty"`                // https://schema.org/archivedAt
	Assesses                  opt.Optional[string] `json:"assesses,omitempty"`                  // https://schema.org/assesses
	AssociatedMedia           opt.Optional[string] `json:"associatedMedia,omitempty"`           // https://schema.org/associatedMedia
	Audience                  opt.Optional[string] `json:"audience,omitempty"`                  // https://schema.org/audience
	Audio                     opt.Optional[string] `json:"audio,omitempty"`                     // https://schema.org/audio
	Author                    opt.Optional[string] `json:"author,omitempty"`                    // https://schema.org/author
	Award                     opt.Optional[string] `json:"award,omitempty"`                     // https://schema.org/award
	Character                 opt.Optional[string] `json:"character,omitempty"`                 // https://schema.org/character
	Citation                  opt.Optional[string] `json:"citation,omitempty"`                  // https://schema.org/citation
	CommentField              opt.Optional[string] `json:"comment,omitempty"`                   // https://schema.org/comment
	CommentCount              opt.Optional[string] `json:"commentCount,omitempty,bare"`         // https://schema.org/commentCount
	ConditionsOfAccess        opt.Optional[string] `json:"conditionsOfAccess,omitempty"`        // https://schema.org/conditionsOfAccess
	ContentLocation           opt.Optional[string] `json:"contentLocation,omitempty"`           // https://schema.org/contentLocation
	ContentRating             opt.Optional[string] `json:"contentRating,omitempty"`             // https://schema.org/contentRating
	ContentReferenceTime      opt.Optional[string] `json:"contentReferenceTime,omitempty"`      // https://schema.org/contentReferenceTime
	Contributor               opt.Optional[string] `json:"contributor,omitempty"`               // https://schema.org/contributor
	CopyrightHolder           opt.Optional[string] `json:"copyrightHolder,omitempty"`           // https://schema.org/copyrightHolder
	CopyrightNotice           opt.Optional[string] `json:"copyrightNotice,omitempty"`           // https://schema.org/copyrightNotice
	CopyrightYear             opt.Optional[string] `json:"copyrightYear,omitempty,bare"`        // https://schema.org/copyrightYear
	Correction                opt.Optional[string] `json:"correction,omitempty"`                // https://schema.org/correction
	CountryOfOrigin           opt.Optional[string] `json:"countryOfOrigin,omitempty"`           // https://schema.org/countryOfOrigin
	CreativeWorkStatus        opt.Optional[string] `json:"creativeWorkStatus,omitempty"`        // https://schema.org/creativeWorkStatus
	Creator                   opt.Optional[string] `json:"creator,omitempty"`                   // https://schema.org/creator
	CreditText                opt.Optional[string] `json:"creditText,omitempty"`                // https://schema.org/creditText
	DateCreated               opt.Optional[string] `json:"dateCreated,omitempty"`               // https://schema.org/dateCreated
	DateModified              opt.Optional[string] `json:"dateModified,omitempty"`              // https://schema.org/dateModified
	DatePublished             opt.Optional[string] `json:"datePublished,omitempty"`             // https://schema.org/datePublished
	DigitalSourceType         opt.Optional[string] `json:"digitalSourceType,omitempty"`         // https://schema.org/digitalSourceType
	DiscussionURL             opt.Optional[string] `json:"discussionUrl,omitempty"`             // https://schema.org/discussionUrl
	DisplayLocation           opt.Optional[string] `json:"displayLocation,omitempty"`           // https://schema.org/displayLocation
	EditEIDR                  opt.Optional[string] `json:"editEIDR,omitempty"`                  // https://schema.org/editEIDR
	Editor                    opt.Optional[string] `json:"editor,omitempty"`                    // https://schema.org/editor
	EducationalAlignment      opt.Optional[string] `json:"educationalAlignment,omitempty"`      // https://schema.org/educationalAlignment
	EducationalLevel          opt.Optional[string] `json:"educationalLevel,omitempty"`          // https://schema.org/educationalLevel
	EducationalUse            opt.Optional[string] `json:"educationalUse,omitempty"`            // https://schema.org/educationalUse
	Encoding                  opt.Optional[string] `json:"encoding,omitempty"`                  // https://schema.org/encoding
	EncodingFormat            opt.Optional[string] `json:"encodingFormat,omitempty"`            // https://schema.org/encodingFormat
	ExampleOfWork             ProtoCreativeWork   `json:"exampleOfWork,omitempty"`             // https://schema.org/exampleOfWork
	Expires                   opt.Optional[string] `json:"expires,omitempty"`                   // https://schema.org/expires
	Funder                    opt.Optional[string] `json:"funder,omitempty"`                    // https://schema.org/funder
	Funding                   opt.Optional[string] `json:"funding,omitempty"`                   // https://schema.org/funding
	Genre                     opt.Optional[string] `json:"genre,omitempty"`                     // https://schema.org/genre
	HasPart                   ProtoCreativeWork   `json:"hasPart,omitempty"`                   // https://schema.org/hasPart
	Headline                  opt.Optional[string] `json:"headline,omitempty"`                  // https://schema.org/headline
	InLanguage                opt.Optional[string] `json:"inLanguage,omitempty"`                // https://schema.org/inLanguage
	InteractionStatistic      opt.Optional[string] `json:"interactionStatistic,omitempty"`      // https://schema.org/interactionStatistic
	InteractivityType         opt.Optional[string] `json:"interactivityType,omitempty"`         // https://schema.org/interactivityType
	InterpretedAsClaim        opt.Optional[string] `json:"interpretedAsClaim,omitempty"`        // https://schema.org/interpretedAsClaim
	IsAccessibleForFree       opt.Optional[string] `json:"isAccessibleForFree,omitempty,bare"`  // https://schema.org/isAccessibleForFree
	IsBasedOn                 opt.Optional[string] `json:"isBasedOn,omitempty"`                 // https://schema.org/isBasedOn
	IsFamilyFriendly          opt.Optional[string] `json:"isFamilyFriendly,omitempty,bare"`     // https://schema.org/isFamilyFriendly
	IsPartOf                  opt.Optional[string] `json:"isPartOf,omitempty"`                  // https://schema.org/isPartOf
	Keywords                  opt.Optional[string] `json:"keywords,omitempty"`                  // https://schema.org/keywords
	LearningResourceType      opt.Optional[string] `json:"learningResourceType,omitempty"`      // https://schema.org/learningResourceType
	License                   opt.Optional[string] `json:"license,omitempty"`                   // https://schema.org/license
	LocationCreated           opt.Optional[string] `json:"locationCreated,omitempty"`           // https://schema.org/locationCreated
	MainEntity                []ProtoThing         `json:"mainEntity,omitempty"`                // https://schema.org/mainEntity
	Maintainer                opt.Optional[string] `json:"maintainer,omitempty"`                // https://schema.org/maintainer
	Material                  opt.Optional[string] `json:"material,omitempty"`                  // https://schema.org/material
	MaterialExtent            opt.Optional[string] `json:"materialExtent,omitempty"`            // https://schema.org/materialExtent
	Mentions                  ProtoThing           `json:"mentions,omitempty"`                  // https://schema.org/mentions
	Offers                    opt.Optional[string] `json:"offers,omitempty"`                    // https://schema.org/offers
	Pattern                   opt.Optional[string] `json:"pattern,omitempty"`                   // https://schema.org/pattern
	Position                  opt.Optional[string] `json:"position,omitempty,bare"`             // https://schema.org/position
	Producer                  opt.Optional[string] `json:"producer,omitempty"`                  // https://schema.org/producer
	Provider                  opt.Optional[string] `json:"provider,omitempty"`                  // https://schema.org/provider
	Publication               opt.Optional[string] `json:"publication,omitempty"`               // https://schema.org/publication
	Publisher                 opt.Optional[string] `json:"publisher,omitempty"`                  // https://schema.org/publisher
	PublisherImprint          opt.Optional[string] `json:"publisherImprint,omitempty"`          // https://schema.org/publisherImprint
	PublishingPrinciples      opt.Optional[string] `json:"publishingPrinciples,omitempty"`      // https://schema.org/publishingPrinciples
	RecordedAt                opt.Optional[string] `json:"recordedAt,omitempty"`                // https://schema.org/recordedAt
	ReleasedEvent             opt.Optional[string] `json:"releasedEvent,omitempty"`             // https://schema.org/releasedEvent
	Review                    opt.Optional[string] `json:"review,omitempty"`                    // https://schema.org/review
	SchemaVersion             opt.Optional[string] `json:"schemaVersion,omitempty"`             // https://schema.org/schemaVersion
	SDDatePublished           opt.Optional[string] `json:"sdDatePublished,omitempty"`           // https://schema.org/sdDatePublished
	SDLicense                 opt.Optional[string] `json:"sdLicense,omitempty"`                 // https://schema.org/sdLicense
	SDPublisher               opt.Optional[string] `json:"sdPublisher,omitempty"`               // https://schema.org/sdPublisher
	Size                      opt.Optional[string] `json:"size,omitempty"`                      // https://schema.org/size
	SourceOrganization        opt.Optional[string] `json:"sourceOrganization,omitempty"`        // https://schema.org/sourceOrganization
	Spatial                   opt.Optional[string] `json:"spatial,omitempty"`                   // https://schema.org/spatial
	SpatialCoverage           opt.Optional[string] `json:"spatialCoverage,omitempty"`           // https://schema.org/spatialCoverage
	Sponsor                   opt.Optional[string] `json:"sponsor,omitempty"`                   // https://schema.org/sponsor
	Teaches                   opt.Optional[string] `json:"teaches,omitempty"`                   // https://schema.org/teaches
	Temporal                  opt.Optional[string] `json:"temporal,omitempty"`                  // https://schema.org/temporal
	TemporalCoverage          opt.Optional[string] `json:"temporalCoverage,omitempty"`          // https://schema.org/temporalCoverage
	Text                      opt.Optional[string] `json:"text,omitempty"`                      // https://schema.org/text
	Thumbnail                 opt.Optional[string] `json:"thumbnail,omitempty"`                 // https://schema.org/thumbnail
	ThumbnailURL              opt.Optional[string] `json:"thumbnailUrl,omitempty"`              // https://schema.org/thumbnailUrl
	TimeRequired              opt.Optional[string] `json:"timeRequired,omitempty"`              // https://schema.org/timeRequired
	TranslationOfWork         ProtoCreativeWork   `json:"translationOfWork,omitempty"`         // https://schema.org/translationOfWork
	Translator                opt.Optional[string] `json:"translator,omitempty"`                // https://schema.org/translator
	TypicalAgeRange           opt.Optional[string] `json:"typicalAgeRange,omitempty"`           // https://schema.org/typicalAgeRange
	UsageInfo                 opt.Optional[string] `json:"usageInfo,omitempty"`                 // https://schema.org/usageInfo
	Version                   opt.Optional[string] `json:"version,omitempty"`                   // https://schema.org/version
	Video                     opt.Optional[string] `json:"video,omitempty"`                     // https://schema.org/video
	WordCount                 opt.Optional[string] `json:"wordCount,omitempty,bare"`            // https://schema.org/wordCount
	WorkExample               ProtoCreativeWork   `json:"workExample,omitempty"`               // https://schema.org/workExample
	WorkTranslation           ProtoCreativeWork   `json:"workTranslation,omitempty"`           // https://schema.org/workTranslation

	// ---- Article properties ----

	ArticleBody               opt.Optional[string] `json:"articleBody,omitempty"`               // https://schema.org/articleBody
	ArticleSection            opt.Optional[string] `json:"articleSection,omitempty"`            // https://schema.org/articleSection
	Backstory                 opt.Optional[string] `json:"backstory,omitempty"`                 // https://schema.org/backstory
	PageEnd                   opt.Optional[string] `json:"pageEnd,omitempty,bare"`              // https://schema.org/pageEnd
	PageStart                 opt.Optional[string] `json:"pageStart,omitempty,bare"`            // https://schema.org/pageStart
	Pagination                opt.Optional[string] `json:"pagination,omitempty"`                // https://schema.org/pagination
	Speakable                 opt.Optional[string] `json:"speakable,omitempty"`                 // https://schema.org/speakable

	// ---- SocialMediaPosting properties ----

	SharedContent             opt.Optional[string] `json:"sharedContent,omitempty"`             // https://schema.org/sharedContent
}

var _ ProtoThing = BlogPosting{}

func (receiver BlogPosting) ProtoThing() AnyThing {
	const thingType string = TypeBlogPosting

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

var _ ProtoCreativeWork = BlogPosting{}

func (receiver BlogPosting) ProtoCreativeWork() AnyCreativeWork {
	const thingType string = TypeBlogPosting

	return AnyCreativeWork{
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

		About:                     receiver.About,
		Abstract:                  receiver.Abstract,
		AccessMode:                receiver.AccessMode,
		AccessModeSufficient:      receiver.AccessModeSufficient,
		AccessibilityAPI:          receiver.AccessibilityAPI,
		AccessibilityControl:      receiver.AccessibilityControl,
		AccessibilityFeature:      receiver.AccessibilityFeature,
		AccessibilityHazard:       receiver.AccessibilityHazard,
		AccessibilitySummary:      receiver.AccessibilitySummary,
		AccountablePerson:         receiver.AccountablePerson,
		AcquireLicensePage:        receiver.AcquireLicensePage,
		AggregateRating:           receiver.AggregateRating,
		AlternativeHeadline:       receiver.AlternativeHeadline,
		ArchivedAt:                receiver.ArchivedAt,
		Assesses:                  receiver.Assesses,
		AssociatedMedia:           receiver.AssociatedMedia,
		Audience:                  receiver.Audience,
		Audio:                     receiver.Audio,
		Author:                    receiver.Author,
		Award:                     receiver.Award,
		Character:                 receiver.Character,
		Citation:                  receiver.Citation,
		CommentField:              receiver.CommentField,
		CommentCount:              receiver.CommentCount,
		ConditionsOfAccess:        receiver.ConditionsOfAccess,
		ContentLocation:           receiver.ContentLocation,
		ContentRating:             receiver.ContentRating,
		ContentReferenceTime:      receiver.ContentReferenceTime,
		Contributor:               receiver.Contributor,
		CopyrightHolder:           receiver.CopyrightHolder,
		CopyrightNotice:           receiver.CopyrightNotice,
		CopyrightYear:             receiver.CopyrightYear,
		Correction:                receiver.Correction,
		CountryOfOrigin:           receiver.CountryOfOrigin,
		CreativeWorkStatus:        receiver.CreativeWorkStatus,
		Creator:                   receiver.Creator,
		CreditText:                receiver.CreditText,
		DateCreated:               receiver.DateCreated,
		DateModified:              receiver.DateModified,
		DatePublished:             receiver.DatePublished,
		DigitalSourceType:         receiver.DigitalSourceType,
		DiscussionURL:             receiver.DiscussionURL,
		DisplayLocation:           receiver.DisplayLocation,
		EditEIDR:                  receiver.EditEIDR,
		Editor:                    receiver.Editor,
		EducationalAlignment:      receiver.EducationalAlignment,
		EducationalLevel:          receiver.EducationalLevel,
		EducationalUse:            receiver.EducationalUse,
		Encoding:                  receiver.Encoding,
		EncodingFormat:            receiver.EncodingFormat,
		ExampleOfWork:             receiver.ExampleOfWork,
		Expires:                   receiver.Expires,
		Funder:                    receiver.Funder,
		Funding:                   receiver.Funding,
		Genre:                     receiver.Genre,
		HasPart:                   receiver.HasPart,
		Headline:                  receiver.Headline,
		InLanguage:                receiver.InLanguage,
		InteractionStatistic:      receiver.InteractionStatistic,
		InteractivityType:         receiver.InteractivityType,
		InterpretedAsClaim:        receiver.InterpretedAsClaim,
		IsAccessibleForFree:       receiver.IsAccessibleForFree,
		IsBasedOn:                 receiver.IsBasedOn,
		IsFamilyFriendly:          receiver.IsFamilyFriendly,
		IsPartOf:                  receiver.IsPartOf,
		Keywords:                  receiver.Keywords,
		LearningResourceType:      receiver.LearningResourceType,
		License:                   receiver.License,
		LocationCreated:           receiver.LocationCreated,
		MainEntity:                receiver.MainEntity,
		Maintainer:                receiver.Maintainer,
		Material:                  receiver.Material,
		MaterialExtent:            receiver.MaterialExtent,
		Mentions:                  receiver.Mentions,
		Offers:                    receiver.Offers,
		Pattern:                   receiver.Pattern,
		Position:                  receiver.Position,
		Producer:                  receiver.Producer,
		Provider:                  receiver.Provider,
		Publication:               receiver.Publication,
		Publisher:                 receiver.Publisher,
		PublisherImprint:          receiver.PublisherImprint,
		PublishingPrinciples:      receiver.PublishingPrinciples,
		RecordedAt:                receiver.RecordedAt,
		ReleasedEvent:             receiver.ReleasedEvent,
		Review:                    receiver.Review,
		SchemaVersion:             receiver.SchemaVersion,
		SDDatePublished:           receiver.SDDatePublished,
		SDLicense:                 receiver.SDLicense,
		SDPublisher:               receiver.SDPublisher,
		Size:                      receiver.Size,
		SourceOrganization:        receiver.SourceOrganization,
		Spatial:                   receiver.Spatial,
		SpatialCoverage:           receiver.SpatialCoverage,
		Sponsor:                   receiver.Sponsor,
		Teaches:                   receiver.Teaches,
		Temporal:                  receiver.Temporal,
		TemporalCoverage:          receiver.TemporalCoverage,
		Text:                      receiver.Text,
		Thumbnail:                 receiver.Thumbnail,
		ThumbnailURL:              receiver.ThumbnailURL,
		TimeRequired:              receiver.TimeRequired,
		TranslationOfWork:         receiver.TranslationOfWork,
		Translator:                receiver.Translator,
		TypicalAgeRange:           receiver.TypicalAgeRange,
		UsageInfo:                 receiver.UsageInfo,
		Version:                   receiver.Version,
		Video:                     receiver.Video,
		WordCount:                 receiver.WordCount,
		WorkExample:               receiver.WorkExample,
		WorkTranslation:           receiver.WorkTranslation,
	}
}

var _ ProtoArticle = BlogPosting{}

func (receiver BlogPosting) ProtoArticle() AnyArticle {
	const thingType string = TypeBlogPosting

	return AnyArticle{
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

		About:                     receiver.About,
		Abstract:                  receiver.Abstract,
		AccessMode:                receiver.AccessMode,
		AccessModeSufficient:      receiver.AccessModeSufficient,
		AccessibilityAPI:          receiver.AccessibilityAPI,
		AccessibilityControl:      receiver.AccessibilityControl,
		AccessibilityFeature:      receiver.AccessibilityFeature,
		AccessibilityHazard:       receiver.AccessibilityHazard,
		AccessibilitySummary:      receiver.AccessibilitySummary,
		AccountablePerson:         receiver.AccountablePerson,
		AcquireLicensePage:        receiver.AcquireLicensePage,
		AggregateRating:           receiver.AggregateRating,
		AlternativeHeadline:       receiver.AlternativeHeadline,
		ArchivedAt:                receiver.ArchivedAt,
		Assesses:                  receiver.Assesses,
		AssociatedMedia:           receiver.AssociatedMedia,
		Audience:                  receiver.Audience,
		Audio:                     receiver.Audio,
		Author:                    receiver.Author,
		Award:                     receiver.Award,
		Character:                 receiver.Character,
		Citation:                  receiver.Citation,
		CommentField:              receiver.CommentField,
		CommentCount:              receiver.CommentCount,
		ConditionsOfAccess:        receiver.ConditionsOfAccess,
		ContentLocation:           receiver.ContentLocation,
		ContentRating:             receiver.ContentRating,
		ContentReferenceTime:      receiver.ContentReferenceTime,
		Contributor:               receiver.Contributor,
		CopyrightHolder:           receiver.CopyrightHolder,
		CopyrightNotice:           receiver.CopyrightNotice,
		CopyrightYear:             receiver.CopyrightYear,
		Correction:                receiver.Correction,
		CountryOfOrigin:           receiver.CountryOfOrigin,
		CreativeWorkStatus:        receiver.CreativeWorkStatus,
		Creator:                   receiver.Creator,
		CreditText:                receiver.CreditText,
		DateCreated:               receiver.DateCreated,
		DateModified:              receiver.DateModified,
		DatePublished:             receiver.DatePublished,
		DigitalSourceType:         receiver.DigitalSourceType,
		DiscussionURL:             receiver.DiscussionURL,
		DisplayLocation:           receiver.DisplayLocation,
		EditEIDR:                  receiver.EditEIDR,
		Editor:                    receiver.Editor,
		EducationalAlignment:      receiver.EducationalAlignment,
		EducationalLevel:          receiver.EducationalLevel,
		EducationalUse:            receiver.EducationalUse,
		Encoding:                  receiver.Encoding,
		EncodingFormat:            receiver.EncodingFormat,
		ExampleOfWork:             receiver.ExampleOfWork,
		Expires:                   receiver.Expires,
		Funder:                    receiver.Funder,
		Funding:                   receiver.Funding,
		Genre:                     receiver.Genre,
		HasPart:                   receiver.HasPart,
		Headline:                  receiver.Headline,
		InLanguage:                receiver.InLanguage,
		InteractionStatistic:      receiver.InteractionStatistic,
		InteractivityType:         receiver.InteractivityType,
		InterpretedAsClaim:        receiver.InterpretedAsClaim,
		IsAccessibleForFree:       receiver.IsAccessibleForFree,
		IsBasedOn:                 receiver.IsBasedOn,
		IsFamilyFriendly:          receiver.IsFamilyFriendly,
		IsPartOf:                  receiver.IsPartOf,
		Keywords:                  receiver.Keywords,
		LearningResourceType:      receiver.LearningResourceType,
		License:                   receiver.License,
		LocationCreated:           receiver.LocationCreated,
		MainEntity:                receiver.MainEntity,
		Maintainer:                receiver.Maintainer,
		Material:                  receiver.Material,
		MaterialExtent:            receiver.MaterialExtent,
		Mentions:                  receiver.Mentions,
		Offers:                    receiver.Offers,
		Pattern:                   receiver.Pattern,
		Position:                  receiver.Position,
		Producer:                  receiver.Producer,
		Provider:                  receiver.Provider,
		Publication:               receiver.Publication,
		Publisher:                 receiver.Publisher,
		PublisherImprint:          receiver.PublisherImprint,
		PublishingPrinciples:      receiver.PublishingPrinciples,
		RecordedAt:                receiver.RecordedAt,
		ReleasedEvent:             receiver.ReleasedEvent,
		Review:                    receiver.Review,
		SchemaVersion:             receiver.SchemaVersion,
		SDDatePublished:           receiver.SDDatePublished,
		SDLicense:                 receiver.SDLicense,
		SDPublisher:               receiver.SDPublisher,
		Size:                      receiver.Size,
		SourceOrganization:        receiver.SourceOrganization,
		Spatial:                   receiver.Spatial,
		SpatialCoverage:           receiver.SpatialCoverage,
		Sponsor:                   receiver.Sponsor,
		Teaches:                   receiver.Teaches,
		Temporal:                  receiver.Temporal,
		TemporalCoverage:          receiver.TemporalCoverage,
		Text:                      receiver.Text,
		Thumbnail:                 receiver.Thumbnail,
		ThumbnailURL:              receiver.ThumbnailURL,
		TimeRequired:              receiver.TimeRequired,
		TranslationOfWork:         receiver.TranslationOfWork,
		Translator:                receiver.Translator,
		TypicalAgeRange:           receiver.TypicalAgeRange,
		UsageInfo:                 receiver.UsageInfo,
		Version:                   receiver.Version,
		Video:                     receiver.Video,
		WordCount:                 receiver.WordCount,
		WorkExample:               receiver.WorkExample,
		WorkTranslation:           receiver.WorkTranslation,

		ArticleBody:               receiver.ArticleBody,
		ArticleSection:            receiver.ArticleSection,
		Backstory:                 receiver.Backstory,
		PageEnd:                   receiver.PageEnd,
		PageStart:                 receiver.PageStart,
		Pagination:                receiver.Pagination,
		Speakable:                 receiver.Speakable,
	}
}
