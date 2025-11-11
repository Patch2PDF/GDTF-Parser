package XMLTypes

import Types "github.com/Patch2PDF/GDTF-Parser/types"

type FixtureType struct {
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

	AttributeDefinitions AttributeDefinitions  `xml:"AttributeDefinitions"`
	Wheels               []Wheel               `xml:"Wheels>Wheel"`
	PhysicalDescriptions []PhysicalDescription `xml:"PhysicalDescriptions"`
	Models               []Model               `xml:"Models>Model"`
	Geometries           Geometries            `xml:"Geometries"`
	DMXModes             []DMXMode             `xml:"DMXModes>DMXMode"`
	Revisions            []Revision            `xml:"Revisions>Revision"`
	FTPresets            []string              `xml:"FTPresets"`
	Protocols            []Protocol            `xml:"Protocols"`
}

func (fixtureType FixtureType) Parse() Types.FixtureType {
	wheels := ParseList(&fixtureType.Wheels)
	physicalDescriptions := ParseList(&fixtureType.PhysicalDescriptions)
	models := ParseList(&fixtureType.Models)
	dmxModes := ParseList(&fixtureType.DMXModes)
	revisions := ParseList(&fixtureType.Revisions)
	return Types.FixtureType{
		FixtureTypeID:        fixtureType.FixtureTypeID,
		Name:                 fixtureType.Name,
		ShortName:            fixtureType.ShortName,
		LongName:             fixtureType.LongName,
		Manufacturer:         fixtureType.Manufacturer,
		Description:          fixtureType.Description,
		Thumbnail:            fixtureType.Thumbnail,
		ThumbnailOffsetX:     fixtureType.ThumbnailOffsetX,
		ThumbnailOffsetY:     fixtureType.ThumbnailOffsetY,
		CanHaveChildren:      bool(fixtureType.CanHaveChildren),
		RefFT:                fixtureType.RefFT,
		AttributeDefinitions: fixtureType.AttributeDefinitions.Parse(),
		Wheels:               wheels,
		PhysicalDescriptions: physicalDescriptions,
		Models:               models,
		Geometries:           fixtureType.Geometries.Parse(),
		DMXModes:             dmxModes,
		Revisions:            revisions,
		FTPresets:            nil, // not defined yet in spec
		Protocols:            nil, // not defined yet in spec
	}
}
