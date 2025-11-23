package gdtfxml_test

import XMLTypes "github.com/Patch2PDF/GDTF-Parser/internal/types/gdtfxml"

var FixtureTypeXML = `
	<FixtureType Name="Basic moving head" ShortName="BMH" LongName="Moving head with basic functionality" Description="This is a very simple moving head that only has basic features" Manufacturer="GDTF example" FixtureTypeID="00000000-9D13-0000-C92D-C1C76F583F46" Thumbnail="picture_fixture.png">
	</FixtureType>
`

var FixtureTypeStruct = XMLTypes.FixtureType{
	Name:          "Basic moving head",
	ShortName:     "BMH",
	LongName:      "Moving head with basic functionality",
	Description:   "This is a very simple moving head that only has basic features",
	Manufacturer:  "GDTF example",
	FixtureTypeID: "00000000-9D13-0000-C92D-C1C76F583F46",
	Thumbnail:     "picture_fixture.png",
}
