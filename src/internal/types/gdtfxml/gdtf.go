package XMLTypes

type XMLGDTF struct {
	DataVersion string `xml:",attr"`
	FixtureType XMLFixtureType
}
