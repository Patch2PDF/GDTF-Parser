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
	Wheels               []*Wheel
	PhysicalDescriptions []*PhysicalDescription
	Models               []*Model
	Geometries           Geometries
	DMXModes             []*DMXMode
	Revisions            []*Revision
	FTPresets            *[]string
	Protocols            Protocol
}

func (obj *FixtureType) CreateReferencePointer() {
	obj.AttributeDefinitions.CreateReferencePointer()
	CreateReferencePointers(&obj.Wheels)
	CreateReferencePointers(&obj.PhysicalDescriptions)
	// CreateReferencePointers(&obj.Models)
	obj.Geometries.CreateReferencePointer()
	CreateReferencePointers(&obj.DMXModes)
	// CreateReferencePointers(&obj.Revisions)
	// obj.Protocols.CreateReferencePointer()
}

func (obj *FixtureType) ResolveReference() {
	obj.AttributeDefinitions.ResolveReference()
	ResolveReferences(&obj.Wheels)
	obj.Geometries.ResolveReference()
	ResolveReferences(&obj.DMXModes)
}
