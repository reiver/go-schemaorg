package schemaorg_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-schemaorg"
)

func TestDataFeedItem_marshal(t *testing.T) {

	const context string = `"@context":{"schema":"https://schema.org/","@id":"schema:@id","@type":"schema:@type","additionalType":"schema:additionalType","alternateName":"schema:alternateName","dateCreated":"schema:dateCreated","dateDeleted":"schema:dateDeleted","dateModified":"schema:dateModified","description":"schema:description","disambiguatingDescription":"schema:disambiguatingDescription","identifier":"schema:identifier","image":"schema:image","item":"schema:item","mainEntityOfPage":"schema:mainEntityOfPage","name":"schema:name","owner":"schema:owner","potentialAction":"schema:potentialAction","sameAs":"schema:sameAs","subjectOf":"schema:subjectOf","url":"schema:url"}`

	type EmbeddedDataFeedItem struct {
		DataFeedItem schemaorg.DataFeedItem `json:"dataFeedItem,omitempty"`
	}

	tests := []struct{
		Value any
		Expected []byte
	}{
		{
			Value: schemaorg.DataFeedItem{},
			Expected: []byte(`{`+context+`,"@type":"DataFeedItem"}`),
		},
		{
			Value: EmbeddedDataFeedItem{},
			Expected: []byte(`{`+context+`}`),
		},



		{
			Value: EmbeddedDataFeedItem{
				DataFeedItem: schemaorg.DataFeedItem{

					Name:        opt.Something("Product Update"),
					Description: opt.Something("An updated product listing."),
				},
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"dataFeedItem":{`+
						`"@type":"DataFeedItem"`+
						`,`+
						`"description":"An updated product listing."`+
						`,`+
						`"name":"Product Update"`+
					`}`+
				`}`),
		},



		{
			Value: EmbeddedDataFeedItem{
				DataFeedItem: schemaorg.DataFeedItem{

					Name:         opt.Something("New Item Entry"),
					DateCreated:  opt.Something("2025-01-15"),
					DateModified: opt.Something("2025-02-20"),
					Item:         opt.Something("https://example.com/items/42"),
				},
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"dataFeedItem":{`+
						`"@type":"DataFeedItem"`+
						`,`+
						`"name":"New Item Entry"`+
						`,`+
						`"dateCreated":"2025-01-15"`+
						`,`+
						`"dateModified":"2025-02-20"`+
						`,`+
						`"item":"https://example.com/items/42"`+
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
