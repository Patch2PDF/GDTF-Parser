package XMLTypes

import Types "github.com/Patch2PDF/GDTF-Parser/pkg/types"

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
	Protocols            Protocol              `xml:"Protocols"`
}

func (fixtureType FixtureType) Parse() Types.FixtureType {
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
		Wheels:               ParseList(&fixtureType.Wheels),
		PhysicalDescriptions: ParseList(&fixtureType.PhysicalDescriptions),
		Models:               ParseList(&fixtureType.Models),
		Geometries:           fixtureType.Geometries.Parse(),
		DMXModes:             ParseMap(&fixtureType.DMXModes),
		Revisions:            ParseList(&fixtureType.Revisions),
		FTPresets:            nil, // not defined yet in spec
		Protocols:            fixtureType.Protocols.Parse(),
	}
}
