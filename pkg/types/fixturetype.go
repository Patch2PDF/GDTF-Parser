package Types

import "image"

type FixtureType struct {
	FixtureTypeID    string
	Name             string
	ShortName        string
	LongName         string
	Manufacturer     string
	Description      string
	Thumbnail        NodeReference[image.Image]
	ThumbnailOffsetX int
	ThumbnailOffsetY int
	CanHaveChildren  bool
	RefFT            *string

	AttributeDefinitions AttributeDefinitions
	Wheels               []*Wheel
	PhysicalDescriptions []*PhysicalDescription
	Models               []*Model
	Geometries           Geometries
	DMXModes             map[string]*DMXMode
	Revisions            []*Revision
	FTPresets            *[]string
	Protocols            Protocol
}

func (obj *FixtureType) CreateReferencePointer(refPointers *ReferencePointers) {
	obj.AttributeDefinitions.CreateReferencePointer(refPointers)
	CreateReferencePointers(refPointers, &obj.Wheels)
	CreateReferencePointers(refPointers, &obj.PhysicalDescriptions)
	CreateReferencePointers(refPointers, &obj.Models)
	obj.Geometries.CreateGeometryReferencePointer(refPointers, "")
	CreateReferencePointersMap(refPointers, &obj.DMXModes)
	// CreateReferencePointers(refPointers, &obj.Revisions)
	// obj.Protocols.CreateReferencePointer(refPointers)
}

func (obj *FixtureType) ResolveReference(refPointers *ReferencePointers) {
	obj.AttributeDefinitions.ResolveReference(refPointers)
	ResolveReferences(refPointers, &obj.Wheels)
	obj.Geometries.ResolveReference(refPointers)
	ResolveReferencesMap(refPointers, &obj.DMXModes)
}
