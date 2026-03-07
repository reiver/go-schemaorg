package schemaorg_test

import (
	"testing"

	"bytes"

	"github.com/reiver/go-jsonld"
	"github.com/reiver/go-opt"

	"github.com/reiver/go-schemaorg"
)

func TestAnyAction_marshal(t *testing.T) {

	const context string = `"@context":{"schema":"https://schema.org/","@id":"schema:@id","@type":"schema:@type","actionProcess":"schema:actionProcess","actionStatus":"schema:actionStatus","additionalType":"schema:additionalType","agent":"schema:agent","alternateName":"schema:alternateName","description":"schema:description","disambiguatingDescription":"schema:disambiguatingDescription","endTime":"schema:endTime","error":"schema:error","identifier":"schema:identifier","image":"schema:image","instrument":"schema:instrument","location":"schema:location","mainEntityOfPage":"schema:mainEntityOfPage","name":"schema:name","object":"schema:object","owner":"schema:owner","participant":"schema:participant","potentialAction":"schema:potentialAction","provider":"schema:provider","result":"schema:result","sameAs":"schema:sameAs","startTime":"schema:startTime","subjectOf":"schema:subjectOf","target":"schema:target","url":"schema:url"}`

	type EmbeddedAnyAction struct {
		AnyAction schemaorg.AnyAction `json:"anyAction,omitempty"`
	}

	tests := []struct{
		Value any
		Expected []byte
	}{
		{
			Value: schemaorg.AnyAction{},
			Expected: []byte(`{`+context+`}`),
		},
		{
			Value: EmbeddedAnyAction{},
			Expected: []byte(`{`+context+`}`),
		},



		{
			Value: EmbeddedAnyAction{
				AnyAction: schemaorg.AnyAction{

					Name:        opt.Something("Join Event"),
					Description: opt.Something("Register for the online seminar."),
					URL:         opt.Something("https://example.com/actions/join"),

					ActionStatus: opt.Something("PotentialActionStatus"),
					Agent:        opt.Something("Jane Doe"),
					StartTime:    opt.Something("2024-09-01T10:00:00Z"),
					EndTime:      opt.Something("2024-09-01T12:00:00Z"),
					Location:     opt.Something("https://meet.example.com/room"),
					Target:       opt.Something("https://example.com/join"),
				},
			},
			Expected: []byte(
				`{`+
					context+
					`,`+
					`"anyAction":{`+
						`"description":"Register for the online seminar."`+
						`,`+
						`"name":"Join Event"`+
						`,`+
						`"url":"https://example.com/actions/join"`+
						`,`+
						`"actionStatus":"PotentialActionStatus"`+
						`,`+
						`"agent":"Jane Doe"`+
						`,`+
						`"endTime":"2024-09-01T12:00:00Z"`+
						`,`+
						`"location":"https://meet.example.com/room"`+
						`,`+
						`"startTime":"2024-09-01T10:00:00Z"`+
						`,`+
						`"target":"https://example.com/join"`+
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
