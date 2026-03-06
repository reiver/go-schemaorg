package schemaorg_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-schemaorg"
)

func TestPerson_marshal(t *testing.T) {

	const context string = `"@context":{"schema":"https://schema.org/","@id":"schema:@id","@type":"schema:@type","additionalName":"schema:additionalName","additionalType":"schema:additionalType","address":"schema:address","affiliation":"schema:affiliation","agentInteractionStatistic":"schema:agentInteractionStatistic","alternateName":"schema:alternateName","alumniOf":"schema:alumniOf","award":"schema:award","birthDate":"schema:birthDate","birthPlace":"schema:birthPlace","brand":"schema:brand","callSign":"schema:callSign","children":"schema:children","colleague":"schema:colleague","contactPoint":"schema:contactPoint","deathDate":"schema:deathDate","deathPlace":"schema:deathPlace","description":"schema:description","disambiguatingDescription":"schema:disambiguatingDescription","duns":"schema:duns","email":"schema:email","familyName":"schema:familyName","faxNumber":"schema:faxNumber","follows":"schema:follows","funder":"schema:funder","funding":"schema:funding","gender":"schema:gender","givenName":"schema:givenName","globalLocationNumber":"schema:globalLocationNumber","hasCertification":"schema:hasCertification","hasCredential":"schema:hasCredential","hasOccupation":"schema:hasOccupation","hasOfferCatalog":"schema:hasOfferCatalog","hasPOS":"schema:hasPOS","height":"schema:height","homeLocation":"schema:homeLocation","honorificPrefix":"schema:honorificPrefix","honorificSuffix":"schema:honorificSuffix","identifier":"schema:identifier","image":"schema:image","interactionStatistic":"schema:interactionStatistic","isicV4":"schema:isicV4","jobTitle":"schema:jobTitle","knows":"schema:knows","knowsAbout":"schema:knowsAbout","knowsLanguage":"schema:knowsLanguage","lifeEvent":"schema:lifeEvent","mainEntityOfPage":"schema:mainEntityOfPage","makesOffer":"schema:makesOffer","memberOf":"schema:memberOf","naics":"schema:naics","name":"schema:name","nationality":"schema:nationality","netWorth":"schema:netWorth","owner":"schema:owner","owns":"schema:owns","parent":"schema:parent","performerIn":"schema:performerIn","potentialAction":"schema:potentialAction","pronouns":"schema:pronouns","publishingPrinciples":"schema:publishingPrinciples","relatedTo":"schema:relatedTo","sameAs":"schema:sameAs","seeks":"schema:seeks","sibling":"schema:sibling","skills":"schema:skills","sponsor":"schema:sponsor","spouse":"schema:spouse","subjectOf":"schema:subjectOf","telephone":"schema:telephone","url":"schema:url","vatID":"schema:vatID","weight":"schema:weight","workLocation":"schema:workLocation","worksFor":"schema:worksFor"}`

	type EmbeddedPerson struct {
		Person schemaorg.Person `json:"person,omitempty"`
	}

	tests := []struct{
		Value any
		Expected []byte
	}{
		{
			Value: schemaorg.Person{},
			Expected: []byte(`{`+context+`,"@type":"Person"}`),
		},
		{
			Value: EmbeddedPerson{},
			Expected: []byte(`{`+context+`}`),
		},



		{
			Value: EmbeddedPerson{
				Person: schemaorg.Person{

					Name: opt.Something("Jane Doe"),
					URL:  opt.Something("https://example.com/janedoe"),

					Email:      opt.Something("jane@example.com"),
					FamilyName: opt.Something("Doe"),
					GivenName:  opt.Something("Jane"),
					JobTitle:   opt.Something("Software Engineer"),
				},
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"person":{`+
						`"@type":"Person"`+
						`,`+
						`"name":"Jane Doe"`+
						`,`+
						`"url":"https://example.com/janedoe"`+
						`,`+
						`"email":"jane@example.com"`+
						`,`+
						`"familyName":"Doe"`+
						`,`+
						`"givenName":"Jane"`+
						`,`+
						`"jobTitle":"Software Engineer"`+
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
