package XMLTypes

type GDTF struct {
	DataVersion string `xml:",attr"`
	FixtureType FixtureType
}
