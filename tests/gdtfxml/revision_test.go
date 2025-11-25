package gdtfxml_test

import (
	"encoding/xml"
	"testing"
	"time"

	XMLTypes "github.com/Patch2PDF/GDTF-Parser/internal/types/gdtfxml"

	"github.com/google/go-cmp/cmp"
)

func TestRevision(t *testing.T) {
	type revisionTest struct {
		Revisions []XMLTypes.Revision `xml:"Revision"`
	}

	xmlData := `
	<Revisions>
    <Revision Text="Created basic structure of fixture type." Date="2018-02-23T09:56:53" UserID="128" />
    <Revision Text="Added wheels" Date="2018-03-01T10:17:09" UserID="128" />
    <Revision Text="Added geometries" Date="2018-03-07T10:56:14" UserID="128" />
    <Revision Text="Added relations" Date="2018-03-26T08:45:56" UserID="128" />
	</Revisions>
	`

	var result revisionTest
	err := xml.Unmarshal([]byte(xmlData), &result)
	if err != nil {
		t.Errorf("Revision: XML Unmarshal threw error: %s", err)
	}

	want := []XMLTypes.Revision{
		{
			Text:   "Created basic structure of fixture type.",
			Date:   XMLTypes.XMLTime{Time: time.Date(2018, 02, 23, 9, 56, 53, 0, time.UTC)},
			UserID: 128,
		},
		{
			Text:   "Added wheels",
			Date:   XMLTypes.XMLTime{Time: time.Date(2018, 03, 01, 10, 17, 9, 0, time.UTC)},
			UserID: 128,
		},
		{
			Text:   "Added geometries",
			Date:   XMLTypes.XMLTime{Time: time.Date(2018, 03, 07, 10, 56, 14, 0, time.UTC)},
			UserID: 128,
		},
		{
			Text:   "Added relations",
			Date:   XMLTypes.XMLTime{Time: time.Date(2018, 03, 26, 8, 45, 56, 0, time.UTC)},
			UserID: 128,
		},
	}

	if !cmp.Equal(result.Revisions, want) {
		t.Errorf("Revision: XML Unmarshaling Output does not match")
		t.Errorf("%s", cmp.Diff(result.Revisions, want))
	}
}
