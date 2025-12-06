package Types

type AttributeDefinitions struct {
	ActivationGroups []*ActivationGroup
	FeatureGroups    []*FeatureGroup
	Attributes       []*Attribute
}

func (obj *AttributeDefinitions) CreateReferencePointer(refPointers *ReferencePointers) {
	CreateReferencePointers(refPointers, &obj.ActivationGroups)
	CreateReferencePointers(refPointers, &obj.FeatureGroups)
	CreateReferencePointers(refPointers, &obj.Attributes)
}

func (obj *AttributeDefinitions) ResolveReference(refPointers *ReferencePointers) {
	ResolveReferences(refPointers, &obj.Attributes)
}

type ActivationGroup struct {
	Name       string
	Attributes []*Attribute
}

func (obj *ActivationGroup) CreateReferencePointer(refPointers *ReferencePointers) {
	refPointers.ActivationGroups[obj.Name] = obj
}

type FeatureGroup struct {
	Name     string
	Pretty   string
	Features []*Feature
}

func (obj *FeatureGroup) CreateReferencePointer(refPointers *ReferencePointers) {
	for _, element := range obj.Features {
		refPointers.Features[obj.Name+"."+element.Name] = element
	}
}

type Feature struct {
	Name       string
	Attributes []*Attribute
}

type Attribute struct {
	Name             string
	Pretty           string
	ActivationGroup  NodeReference[ActivationGroup]
	Feature          NodeReference[Feature]
	MainAttribute    NodeReference[Attribute]
	PhysicalUnit     string
	Color            ColorCIE
	SubPhysicalUnits *[]*SubPhysicalUnit
}

func (obj *Attribute) CreateReferencePointer(refPointers *ReferencePointers) {
	refPointers.Attributes[obj.Name] = obj
	// TODO: find out reference name for SubPhysicalUnits
	// for _, element := range *obj.SubPhysicalUnits {
	// 	refPointers.SubPhysicalUnits[obj.Name + "." + element]
	// }
}

func (obj *Attribute) ResolveReference(refPointers *ReferencePointers) {
	if obj.ActivationGroup.String != "" {
		obj.ActivationGroup.Ptr = refPointers.ActivationGroups[obj.ActivationGroup.String]
		refPointers.ActivationGroups[obj.ActivationGroup.String].Attributes =
			append(refPointers.ActivationGroups[obj.ActivationGroup.String].Attributes, obj)
	}
	if obj.Feature.String != "" {
		obj.Feature.Ptr = refPointers.Features[obj.Feature.String]
		refPointers.Features[obj.Feature.String].Attributes =
			append(refPointers.Features[obj.Feature.String].Attributes, obj)
	}
	if obj.MainAttribute.String != "" {
		obj.MainAttribute.Ptr = refPointers.Attributes[obj.MainAttribute.String]
	}
}

type SubPhysicalUnit struct {
	Type         string
	PhysicalUnit string
	PhysicalFrom float32
	PhysicalTo   float32
}
