package Types

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
	Wheels               []Wheel
	PhysicalDescriptions []PhysicalDescription
	Models               []Model
	Geometries           Geometries
	DMXModes             []DMXMode
	Revisions            []Revision
	FTPresets            *[]string
	Protocols            *[]string
}
