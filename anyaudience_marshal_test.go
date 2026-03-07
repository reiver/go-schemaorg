package schemaorg_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-schemaorg"
)

func TestAnyAudience_marshal(t *testing.T) {

	const context string = `"@context":{"schema":"https://schema.org/","@id":"schema:@id","@type":"schema:@type","additionalType":"schema:additionalType","alternateName":"schema:alternateName","audienceType":"schema:audienceType","description":"schema:description","disambiguatingDescription":"schema:disambiguatingDescription","geographicArea":"schema:geographicArea","identifier":"schema:identifier","image":"schema:image","mainEntityOfPage":"schema:mainEntityOfPage","name":"schema:name","owner":"schema:owner","potentialAction":"schema:potentialAction","sameAs":"schema:sameAs","subjectOf":"schema:subjectOf","url":"schema:url"}`

	type EmbeddedAnyAudience struct {
		AnyAudience schemaorg.AnyAudience `json:"anyAudience,omitempty"`
	}

	tests := []struct{
		Value any
		Expected []byte
	}{
		{
			Value: schemaorg.AnyAudience{},
			Expected: []byte(`{`+context+`}`),
		},
		{
			Value: EmbeddedAnyAudience{},
			Expected: []byte(`{`+context+`}`),
		},



		{
			Value: EmbeddedAnyAudience{
				AnyAudience: schemaorg.AnyAudience{

					Name:        opt.Something("Music Enthusiasts"),
					Description: opt.Something("People who enjoy live music events."),
					URL:         opt.Something("https://example.com/audiences/music"),

					AudienceType:   opt.Something("musicians"),
					GeographicArea: opt.Something("North America"),
				},
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"anyAudience":{`+
						`"description":"People who enjoy live music events."`+
						`,`+
						`"name":"Music Enthusiasts"`+
						`,`+
						`"url":"https://example.com/audiences/music"`+
						`,`+
						`"audienceType":"musicians"`+
						`,`+
						`"geographicArea":"North America"`+
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
