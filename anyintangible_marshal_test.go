package schemaorg_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-schemaorg"
)

func TestAnyIntangible_marshal(t *testing.T) {

	const context string = `"@context":{"schema":"https://schema.org/","@id":"schema:@id","@type":"schema:@type","additionalType":"schema:additionalType","alternateName":"schema:alternateName","description":"schema:description","disambiguatingDescription":"schema:disambiguatingDescription","identifier":"schema:identifier","image":"schema:image","mainEntityOfPage":"schema:mainEntityOfPage","name":"schema:name","owner":"schema:owner","potentialAction":"schema:potentialAction","sameAs":"schema:sameAs","subjectOf":"schema:subjectOf","url":"schema:url"}`

	type EmbeddedAnyIntangible struct {
		AnyIntangible schemaorg.AnyIntangible `json:"anyIntangible,omitempty"`
	}

	tests := []struct{
		Value any
		Expected []byte
	}{
		{
			Value: schemaorg.AnyIntangible{},
			Expected: []byte(`{`+context+`}`),
		},
		{
			Value: EmbeddedAnyIntangible{},
			Expected: []byte(`{`+context+`}`),
		},



		{
			Value: EmbeddedAnyIntangible{
				AnyIntangible: schemaorg.AnyIntangible{

					Name:        opt.Something("An Intangible"),
					Description: opt.Something("Some intangible thing."),
					URL:         opt.Something("https://example.com/intangible/1"),
				},
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"anyIntangible":{`+
						`"description":"Some intangible thing."`+
						`,`+
						`"name":"An Intangible"`+
						`,`+
						`"url":"https://example.com/intangible/1"`+
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
