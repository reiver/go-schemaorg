package schemaorg_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-schemaorg"
)

func TestGeoCoordinates_marshal(t *testing.T) {

	const context string = `"@context":{"schema":"https://schema.org/","additionalType":"schema:additionalType","alternateName":"schema:alternateName","description":"schema:description","disambiguatingDescription":"schema:disambiguatingDescription","elevation":"schema:elevation","identifier":"schema:identifier","latitude":"schema:latitude","longitude":"schema:longitude","name":"schema:name","type":"schema:type","url":"schema:url"}`

	type EmbeddedGeoCoordinates struct {
		GeoCoordinates schemaorg.GeoCoordinates `json:"thing,omitempty"`
	}

	tests := []struct{
		Value any
		Expected []byte
	}{
		{
			Value: schemaorg.GeoCoordinates{},
			Expected: []byte(`{`+context+`,"type":"GeoCoordinates"}`),
		},
		{
			Value: EmbeddedGeoCoordinates{},
			Expected: []byte(`{`+context+`}`),
		},



		{
			Value: EmbeddedGeoCoordinates{
				GeoCoordinates: schemaorg.GeoCoordinates{

					AdditionalType:            opt.Something("That-GeoCoordinates"),
					AlternateName:             opt.Something("id-123abc"),
					Description:               opt.Something("This is something that is some type of thing."),
					DisambiguatingDescription: opt.Something("It's mine."),
					Elevation:                 opt.Something("20.17"),
					Identifier:                opt.Something("055d2ddb-fd55-48c0-892b-28f1adcc0465"),
					Latitude:                  opt.Something("17.18"),
					Longitude:                 opt.Something("10.03"),
					Name:                      opt.Something("Something"),
					URL:                       opt.Something("http://example.com/somewhere.php"),
				},
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"thing":{`+
						`"additionalType":"That-GeoCoordinates"`+
						`,`+
						`"alternateName":"id-123abc"`+
						`,`+
						`"description":"This is something that is some type of thing."`+
						`,`+
						`"disambiguatingDescription":"It's mine."`+
						`,`+
						`"elevation":"20.17"`+
						`,`+
						`"identifier":"055d2ddb-fd55-48c0-892b-28f1adcc0465"`+
						`,`+
						`"latitude":"17.18"`+
						`,`+
						`"longitude":"10.03"`+
						`,`+
						`"name":"Something"`+
						`,`+
						`"type":"GeoCoordinates"`+
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
