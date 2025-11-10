package Types

type AttributeDefinitions struct {
	ActivationGroups *map[string]ActivationGroup
	FeatureGroups    map[string]FeatureGroup
	Attributes       map[string]Attribute
}

type ActivationGroup struct {
	Name       string
	Attributes []*Attribute
}

type FeatureGroup struct {
	Name     string
	Pretty   string
	Features map[string]Feature
}

type Feature struct {
	Name       string
	Attributes []*Attribute
}

type Attribute struct {
	Name             string
	Pretty           string
	ActivationGroup  *ActivationGroup
	Feature          *Feature
	MainAttribute    *Attribute
	PhysicalUnit     string
	Color            string
	SubPhysicalUnits *[]SubPhysicalUnit
}

type SubPhysicalUnit struct {
	Type         string
	PhysicalUnit string
	PhysicalFrom float32
	PhysicalTo   float32
}
