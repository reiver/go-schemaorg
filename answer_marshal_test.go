package schemaorg_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-schemaorg"
)

func TestAnswer_marshal(t *testing.T) {

	const context string = `"@context":{"schema":"https://schema.org/","@id":"schema:@id","@type":"schema:@type","about":"schema:about","abstract":"schema:abstract","accessMode":"schema:accessMode","accessModeSufficient":"schema:accessModeSufficient","accessibilityAPI":"schema:accessibilityAPI","accessibilityControl":"schema:accessibilityControl","accessibilityFeature":"schema:accessibilityFeature","accessibilityHazard":"schema:accessibilityHazard","accessibilitySummary":"schema:accessibilitySummary","accountablePerson":"schema:accountablePerson","acquireLicensePage":"schema:acquireLicensePage","additionalType":"schema:additionalType","aggregateRating":"schema:aggregateRating","alternateName":"schema:alternateName","alternativeHeadline":"schema:alternativeHeadline","answerExplanation":"schema:answerExplanation","archivedAt":"schema:archivedAt","assesses":"schema:assesses","associatedMedia":"schema:associatedMedia","audience":"schema:audience","audio":"schema:audio","author":"schema:author","award":"schema:award","character":"schema:character","citation":"schema:citation","comment":"schema:comment","commentCount":"schema:commentCount","conditionsOfAccess":"schema:conditionsOfAccess","contentLocation":"schema:contentLocation","contentRating":"schema:contentRating","contentReferenceTime":"schema:contentReferenceTime","contributor":"schema:contributor","copyrightHolder":"schema:copyrightHolder","copyrightNotice":"schema:copyrightNotice","copyrightYear":"schema:copyrightYear","correction":"schema:correction","countryOfOrigin":"schema:countryOfOrigin","creativeWorkStatus":"schema:creativeWorkStatus","creator":"schema:creator","creditText":"schema:creditText","dateCreated":"schema:dateCreated","dateModified":"schema:dateModified","datePublished":"schema:datePublished","description":"schema:description","digitalSourceType":"schema:digitalSourceType","disambiguatingDescription":"schema:disambiguatingDescription","discussionUrl":"schema:discussionUrl","displayLocation":"schema:displayLocation","downvoteCount":"schema:downvoteCount","editEIDR":"schema:editEIDR","editor":"schema:editor","educationalAlignment":"schema:educationalAlignment","educationalLevel":"schema:educationalLevel","educationalUse":"schema:educationalUse","encoding":"schema:encoding","encodingFormat":"schema:encodingFormat","exampleOfWork":"schema:exampleOfWork","expires":"schema:expires","funder":"schema:funder","funding":"schema:funding","genre":"schema:genre","hasPart":"schema:hasPart","headline":"schema:headline","identifier":"schema:identifier","image":"schema:image","inLanguage":"schema:inLanguage","interactionStatistic":"schema:interactionStatistic","interactivityType":"schema:interactivityType","interpretedAsClaim":"schema:interpretedAsClaim","isAccessibleForFree":"schema:isAccessibleForFree","isBasedOn":"schema:isBasedOn","isFamilyFriendly":"schema:isFamilyFriendly","isPartOf":"schema:isPartOf","keywords":"schema:keywords","learningResourceType":"schema:learningResourceType","license":"schema:license","locationCreated":"schema:locationCreated","mainEntity":"schema:mainEntity","mainEntityOfPage":"schema:mainEntityOfPage","maintainer":"schema:maintainer","material":"schema:material","materialExtent":"schema:materialExtent","mentions":"schema:mentions","name":"schema:name","offers":"schema:offers","owner":"schema:owner","parentItem":"schema:parentItem","pattern":"schema:pattern","position":"schema:position","potentialAction":"schema:potentialAction","producer":"schema:producer","provider":"schema:provider","publication":"schema:publication","publisher":"schema:publisher","publisherImprint":"schema:publisherImprint","publishingPrinciples":"schema:publishingPrinciples","recordedAt":"schema:recordedAt","releasedEvent":"schema:releasedEvent","review":"schema:review","sameAs":"schema:sameAs","schemaVersion":"schema:schemaVersion","sdDatePublished":"schema:sdDatePublished","sdLicense":"schema:sdLicense","sdPublisher":"schema:sdPublisher","sharedContent":"schema:sharedContent","size":"schema:size","sourceOrganization":"schema:sourceOrganization","spatial":"schema:spatial","spatialCoverage":"schema:spatialCoverage","sponsor":"schema:sponsor","subjectOf":"schema:subjectOf","teaches":"schema:teaches","temporal":"schema:temporal","temporalCoverage":"schema:temporalCoverage","text":"schema:text","thumbnail":"schema:thumbnail","thumbnailUrl":"schema:thumbnailUrl","timeRequired":"schema:timeRequired","translationOfWork":"schema:translationOfWork","translator":"schema:translator","typicalAgeRange":"schema:typicalAgeRange","upvoteCount":"schema:upvoteCount","url":"schema:url","usageInfo":"schema:usageInfo","version":"schema:version","video":"schema:video","wordCount":"schema:wordCount","workExample":"schema:workExample","workTranslation":"schema:workTranslation"}`

	type EmbeddedAnswer struct {
		Answer schemaorg.Answer `json:"answer,omitempty"`
	}

	tests := []struct{
		Value any
		Expected []byte
	}{
		{
			Value: schemaorg.Answer{},
			Expected: []byte(`{`+context+`,"@type":"Answer"}`),
		},
		{
			Value: EmbeddedAnswer{},
			Expected: []byte(`{`+context+`}`),
		},



		{
			Value: EmbeddedAnswer{
				Answer: schemaorg.Answer{

					Name:              opt.Something("Test Answer"),
					Description:       opt.Something("A test answer."),
					Author:            opt.Something("John Doe"),
					UpvoteCount:       opt.Something("5"),
					AnswerExplanation: opt.Something("Step by step explanation."),
				},
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"answer":{`+
						`"@type":"Answer"`+
						`,`+
						`"description":"A test answer."`+
						`,`+
						`"name":"Test Answer"`+
						`,`+
						`"author":"John Doe"`+
						`,`+
						`"upvoteCount":5`+
						`,`+
						`"answerExplanation":"Step by step explanation."`+
					`}`+
				`}`),
		},
	}

	for testNumber, test := range tests {

		actual, err := jsonld.Marshal(test.Value)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: (%T) %s", err, err)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}

		expected := test.Expected

		if !bytes.Equal(expected, actual) {
			t.Errorf("For test #%d, the actual marshaled value is not what was expected.", testNumber)
			t.Logf("EXPECTED: (%d)\n%s", len(expected), expected)
			t.Logf("ACTUAL:   (%d)\n%s", len(actual), actual)
			t.Logf("VALUE: (%T) %#v", test.Value, test.Value)
			continue
		}
	}
}
