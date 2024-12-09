package schemaorg_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-schemaorg"
)

func TestPropertyValue_marshal(t *testing.T) {

	const context string = `"@context":{"schema":"https://schema.org/","additionalType":"schema:additionalType","alternateName":"schema:alternateName","description":"schema:description","disambiguatingDescription":"schema:disambiguatingDescription","identifier":"schema:identifier","maxValue":"schema:maxValue","minValue":"schema:minValue","name":"schema:name","propertyID":"schema:propertyID","type":"schema:type","unitCode":"schema:unitCode","unitText":"schema:unitText","url":"schema:url","value":"schema:value"}`

	type EmbeddedPropertyValue struct {
		PropertyValue schemaorg.PropertyValue `json:"thing,omitempty"`
	}

	tests := []struct{
		Value any
		Expected []byte
	}{
		{
			Value: schemaorg.PropertyValue{},
			Expected: []byte(`{`+context+`,"type":"PropertyValue"}`),
		},
		{
			Value: EmbeddedPropertyValue{},
			Expected: []byte(`{`+context+`}`),
		},



		{
			Value: EmbeddedPropertyValue{
				PropertyValue: schemaorg.PropertyValue{

					AdditionalType:            opt.Something("That-PropertyValue"),
					AlternateName:             opt.Something("where-i-live"),
					Description:               opt.Something("This is something that is some type of thing."),
					DisambiguatingDescription: opt.Something("It's mine."),
					Identifier:                opt.Something("055d2ddb-fd55-48c0-892b-28f1adcc0465"),
					MaxValue:                  opt.Something("101.232"),
					MinValue:                  opt.Something("-123.45"),
					Name:                      opt.Something("Location"),
					PropertyID:                opt.Something("123-abc"),
					UnitCode:                  opt.Something("asdf"),
					UnitText:                  opt.Something("ghjk"),
					URL:                       opt.Something("http://example.com/somewhere.php"),
					Value:                     opt.Something("Isfahan"),
				},
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"thing":{`+
						`"additionalType":"That-PropertyValue"`+
						`,`+
						`"alternateName":"where-i-live"`+
						`,`+
						`"description":"This is something that is some type of thing."`+
						`,`+
						`"disambiguatingDescription":"It's mine."`+
						`,`+
						`"identifier":"055d2ddb-fd55-48c0-892b-28f1adcc0465"`+
						`,`+
						`"maxValue":101.232`+
						`,`+
						`"minValue":-123.45`+
						`,`+
						`"name":"Location"`+
						`,`+
						`"propertyID":"123-abc"`+
						`,`+
						`"type":"PropertyValue"`+
						`,`+
						`"unitCode":"asdf"`+
						`,`+
						`"unitText":"ghjk"`+
						`,`+
						`"url":"http://example.com/somewhere.php"`+
						`,`+
						`"value":"Isfahan"`+
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
