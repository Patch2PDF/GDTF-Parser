package XMLTypes

import Types "github.com/Patch2PDF/GDTF-Parser/types"

type AttributeDefinitions struct {
	ActivationGroups []ActivationGroup `xml:"ActivationGroups>ActivationGroup"`
	FeatureGroups    []FeatureGroup    `xml:"FeatureGroups>FeatureGroup"`
	Attributes       []Attribute       `xml:"Attributes>Attribute"`
}

func (attr AttributeDefinitions) Parse() Types.AttributeDefinitions {
	activationGroups := ParseMap(&attr.ActivationGroups)
	featureGroups := ParseMap(&attr.FeatureGroups)
	attributes := ParseMap(&attr.Attributes)
	return Types.AttributeDefinitions{
		ActivationGroups: &activationGroups,
		FeatureGroups:    featureGroups,
		Attributes:       attributes,
	}
}

type ActivationGroup struct {
	Name string `xml:",attr"`
}

func (attr ActivationGroup) Parse() Types.ActivationGroup {
	return Types.ActivationGroup{
		Name: attr.Name,
	}
}

func (attr ActivationGroup) ParseKey() string {
	return attr.Name
}

type FeatureGroup struct {
	Name     string    `xml:",attr"`
	Pretty   string    `xml:",attr"`
	Features []Feature `xml:"Feature"`
}

func (attr FeatureGroup) Parse() Types.FeatureGroup {
	features := ParseMap(&attr.Features)
	return Types.FeatureGroup{
		Name:     attr.Name,
		Pretty:   attr.Pretty,
		Features: features,
	}
}

func (attr FeatureGroup) ParseKey() string {
	return attr.Name
}

type Feature struct {
	Name string `xml:",attr"`
}

func (attr Feature) Parse() Types.Feature {
	return Types.Feature{
		Name: attr.Name,
	}
}

func (attr Feature) ParseKey() string {
	return attr.Name
}

type Attribute struct {
	Name             string            `xml:",attr"`
	Pretty           string            `xml:",attr"`
	ActivationGroup  XMLNodeReference  `xml:",attr"` // ActivationGroup
	Feature          XMLNodeReference  `xml:",attr"` // Feature
	MainAttribute    XMLNodeReference  `xml:",attr"` // Attribute
	PhysicalUnit     string            `xml:",attr"`
	Color            ColorCIE          `xml:",attr"`
	SubPhysicalUnits []SubPhysicalUnit `xml:"SubPhysicalUnits"`
}

func (attr Attribute) Parse() Types.Attribute {
	subPhysicalUnits := ParseList(&attr.SubPhysicalUnits)
	return Types.Attribute{
		Name:   attr.Name,
		Pretty: attr.Pretty,
		ActivationGroup: Types.NodeReference[Types.ActivationGroup]{
			String: attr.ActivationGroup,
		},
		Feature: Types.NodeReference[Types.Feature]{
			String: attr.Feature,
		},
		MainAttribute: Types.NodeReference[Types.Attribute]{
			String: attr.MainAttribute,
		},
		PhysicalUnit:     attr.PhysicalUnit,
		Color:            Types.ColorCIE(attr.Color),
		SubPhysicalUnits: &subPhysicalUnits,
	}
}

func (attr Attribute) ParseKey() string {
	return attr.Name
}

type SubPhysicalUnit struct {
	Type         string  `xml:",attr"`
	PhysicalUnit string  `xml:",attr"`
	PhysicalFrom float32 `xml:",attr"`
	PhysicalTo   float32 `xml:",attr"`
}

func (attr SubPhysicalUnit) Parse() Types.SubPhysicalUnit {
	return Types.SubPhysicalUnit(attr)
}
