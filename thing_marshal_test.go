package schemaorg_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-schemaorg"
)

func TestThing_marshal(t *testing.T) {

	const context string = `"@context":{"schema":"https://schema.org/","additionalType":"schema:additionalType","alternateName":"schema:alternateName","description":"schema:description","disambiguatingDescription":"schema:disambiguatingDescription","identifier":"schema:identifier","name":"schema:name","type":"schema:type","url":"schema:url"}`

	type EmbeddedThing struct {
		Thing schemaorg.Thing `json:"thing,omitempty"`
	}

	tests := []struct{
		Value any
		Expected []byte
	}{
		{
			Value: schemaorg.Thing{},
			Expected: []byte(`{`+context+`}`),
		},
		{
			Value: EmbeddedThing{},
			Expected: []byte(`{`+context+`}`),
		},



		{
			Value: EmbeddedThing{
				Thing: schemaorg.Thing{

					AdditionalType:            opt.Something("That-Thing"),
					Description:               opt.Something("This is something that is some type of thing."),
					DisambiguatingDescription: opt.Something("It's mine."),
					Identifier:                opt.Something("055d2ddb-fd55-48c0-892b-28f1adcc0465"),
					Name:                      opt.Something("Something"),
					Type:                      opt.Something("Some-Type"),
					URL:                       opt.Something("http://example.com/somewhere.php"),
				},
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"thing":{`+
						`"additionalType":"That-Thing"`+
						`,`+
						`"description":"This is something that is some type of thing."`+
						`,`+
						`"disambiguatingDescription":"It's mine."`+
						`,`+
						`"identifier":"055d2ddb-fd55-48c0-892b-28f1adcc0465"`+
						`,`+
						`"name":"Something"`+
						`,`+
						`"type":"Some-Type"`+
						`,`+
						`"url":"http://example.com/somewhere.php"`+
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
