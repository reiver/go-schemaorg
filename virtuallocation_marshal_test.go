package schemaorg_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-schemaorg"
)

func TestVirtualLocation_marshal(t *testing.T) {

	const context string = `"@context":{"schema":"https://schema.org/","@id":"schema:@id","@type":"schema:@type","additionalType":"schema:additionalType","alternateName":"schema:alternateName","description":"schema:description","disambiguatingDescription":"schema:disambiguatingDescription","identifier":"schema:identifier","image":"schema:image","mainEntityOfPage":"schema:mainEntityOfPage","name":"schema:name","owner":"schema:owner","potentialAction":"schema:potentialAction","sameAs":"schema:sameAs","subjectOf":"schema:subjectOf","url":"schema:url"}`

	type EmbeddedVirtualLocation struct {
		VirtualLocation schemaorg.VirtualLocation `json:"virtualLocation,omitempty"`
	}

	tests := []struct{
		Value any
		Expected []byte
	}{
		{
			Value: schemaorg.VirtualLocation{},
			Expected: []byte(`{`+context+`,"@type":"VirtualLocation"}`),
		},
		{
			Value: EmbeddedVirtualLocation{},
			Expected: []byte(`{`+context+`}`),
		},



		{
			Value: EmbeddedVirtualLocation{
				VirtualLocation: schemaorg.VirtualLocation{

					Name:        opt.Something("Online Seminar Room"),
					Description: opt.Something("Virtual meeting space for the weekly seminar."),
					URL:         opt.Something("https://meet.example.com/seminar-room"),
				},
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"virtualLocation":{`+
						`"@type":"VirtualLocation"`+
						`,`+
						`"description":"Virtual meeting space for the weekly seminar."`+
						`,`+
						`"name":"Online Seminar Room"`+
						`,`+
						`"url":"https://meet.example.com/seminar-room"`+
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
