package schemaorg_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
	jsonldstrings "github.com/reiver/go-jsonld/strings"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-schemaorg"
)

func TestGeoCoordinates_marshal(t *testing.T) {

	const context string = `"@context":{"schema":"https://schema.org/","@id":"schema:@id","@type":"schema:@type","additionalType":"schema:additionalType","alternateName":"schema:alternateName","description":"schema:description","disambiguatingDescription":"schema:disambiguatingDescription","elevation":"schema:elevation","identifier":"schema:identifier","image":"schema:image","latitude":"schema:latitude","longitude":"schema:longitude","mainEntityOfPage":"schema:mainEntityOfPage","name":"schema:name","owner":"schema:owner","potentialAction":"schema:potentialAction","sameAs":"schema:sameAs","subjectOf":"schema:subjectOf","url":"schema:url"}`

	type EmbeddedGeoCoordinates struct {
		GeoCoordinates schemaorg.GeoCoordinates `json:"thing,omitempty"`
	}

	tests := []struct{
		Value any
		Expected []byte
	}{
		{
			Value: schemaorg.GeoCoordinates{},
			Expected: []byte(`{`+context+`,"@type":"GeoCoordinates"}`),
		},
		{
			Value: EmbeddedGeoCoordinates{},
			Expected: []byte(`{`+context+`}`),
		},



		{
			Value: EmbeddedGeoCoordinates{
				GeoCoordinates: schemaorg.GeoCoordinates{

					AdditionalType:            jsonldstrings.Something("That-GeoCoordinates"),
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
						`"@type":"GeoCoordinates"`+
						`,`+
						`"additionalType":"That-GeoCoordinates"`+
						`,`+
						`"alternateName":"id-123abc"`+
						`,`+
						`"description":"This is something that is some type of thing."`+
						`,`+
						`"disambiguatingDescription":"It's mine."`+
						`,`+
						`"identifier":"055d2ddb-fd55-48c0-892b-28f1adcc0465"`+
						`,`+
						`"name":"Something"`+
						`,`+
						`"url":"http://example.com/somewhere.php"`+
						`,`+
						`"elevation":"20.17"`+
						`,`+
						`"latitude":"17.18"`+
						`,`+
						`"longitude":"10.03"`+
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
