package XMLTypes

type XMLAttributeDefinitions struct {
	ActivationGroups []XMLActivationGroup `xml:"ActivationGroups>ActivationGroup"`
	FeatureGroups    []XMLFeatureGroup    `xml:"FeatureGroups>FeatureGroup"`
	Attributes       []XMLAttribute       `xml:"Attributes>Attribute"`
}

type XMLActivationGroup struct {
	Name string `xml:",attr"`
}

type XMLFeatureGroup struct {
	Name     string       `xml:",attr"`
	Pretty   string       `xml:",attr"`
	Features []XMLFeature `xml:"Feature"`
}

type XMLFeature struct {
	Name string `xml:",attr"`
}

type XMLAttribute struct {
	Name             string               `xml:",attr"`
	Pretty           string               `xml:",attr"`
	ActivationGroup  XMLNodeReference     `xml:",attr"` // XMLActivationGroup
	Feature          XMLNodeReference     `xml:",attr"` // XMLFeature
	MainAttribute    XMLNodeReference     `xml:",attr"` // XMLAttribute
	PhysicalUnit     string               `xml:",attr"`
	Color            ColorCIE             `xml:",attr"`
	SubPhysicalUnits []XMLSubPhysicalUnit `xml:"SubPhysicalUnits"`
}

type XMLSubPhysicalUnit struct {
	Type         string  `xml:",attr"`
	PhysicalUnit string  `xml:",attr"`
	PhysicalFrom float32 `xml:",attr"`
	PhysicalTo   float32 `xml:",attr"`
}
