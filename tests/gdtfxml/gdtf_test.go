package gdtfxml_test

import (
	"testing"

	XMLTypes "github.com/Patch2PDF/GDTF-Parser/internal/types/gdtfxml"
)

func TestGDTF(t *testing.T) {
	xmlData := `
	<GDTF DataVersion="1.1">
		` + FixtureTypeXML + `
	</GDTF>
	`

	want := XMLTypes.GDTF{
		DataVersion: "1.1",
		FixtureType: FixtureTypeStruct,
	}

	parsingTest(t, xmlData, "GDTF", want)
}
