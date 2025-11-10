package XMLTypes

type XMLFixtureType struct {
	FixtureTypeID    string    `xml:",attr"`
	Name             string    `xml:",attr"`
	ShortName        string    `xml:",attr"`
	LongName         string    `xml:",attr"`
	Manufacturer     string    `xml:",attr"`
	Description      string    `xml:",attr"`
	Thumbnail        *string   `xml:",attr"`
	ThumbnailOffsetX int       `xml:",attr"`
	ThumbnailOffsetY int       `xml:",attr"`
	CanHaveChildren  YesNoBool `xml:",attr"`
	RefFT            *string   `xml:",attr"`

	AttributeDefinitions XMLAttributeDefinitions `xml:"AttributeDefinitions"`
	Wheels               []XMLWheel              `xml:"Wheels>Wheel"`
	PhysicalDescriptions []PhysicalDescription   `xml:"PhysicalDescriptions"`
	Models               []XMLModel              `xml:"Models>Model"`
	Geometries           Geometries              `xml:"Geometries"`
	DMXModes             []XMLDMXMode            `xml:"DMXModes>DMXMode"`
	Revisions            []XMLRevision           `xml:"Revisions>Revision"`
	FTPresets            []string                `xml:"FTPresets"`
	Protocols            []string                `xml:"Protocols"`
}
