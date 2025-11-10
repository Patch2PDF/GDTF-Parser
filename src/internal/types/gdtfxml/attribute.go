package XMLTypes

type AttributeDefinitions struct {
	ActivationGroups []ActivationGroup `xml:"ActivationGroups>ActivationGroup"`
	FeatureGroups    []FeatureGroup    `xml:"FeatureGroups>FeatureGroup"`
	Attributes       []Attribute       `xml:"Attributes>Attribute"`
}

type ActivationGroup struct {
	Name string `xml:",attr"`
}

type FeatureGroup struct {
	Name     string    `xml:",attr"`
	Pretty   string    `xml:",attr"`
	Features []Feature `xml:"Feature"`
}

type Feature struct {
	Name string `xml:",attr"`
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

type SubPhysicalUnit struct {
	Type         string  `xml:",attr"`
	PhysicalUnit string  `xml:",attr"`
	PhysicalFrom float32 `xml:",attr"`
	PhysicalTo   float32 `xml:",attr"`
}
