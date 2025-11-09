package Type

type FixtureType struct {
	FixtureTypeID    string
	Name             string
	ShortName        string
	LongName         string
	Manufacturer     string
	Description      string
	Thumbnail        *string
	ThumbnailOffsetX int
	ThumbnailOffsetY int
	CanHaveChildren  bool
	RefFT            *string

	AttributeDefinitions AttributeDefinitions
	Wheels               *[]Wheel
	PhysicalDescriptions *[]string
	Models               *[]Model
	Geometries           []Geometry
	DMXModes             []DMXMode
	Revisions            *[]string
	FTPresets            *[]string
	Protocols            *[]string
}
