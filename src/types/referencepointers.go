package Types

type ReferencePointers struct {
	ActivationGroups map[string]*ActivationGroup
	Features         map[string]*Feature
	Attributes       map[string]*Attribute
	Wheels           map[string]*Wheel
	WheelSlots       map[string]*WheelSlot
	Emitters         map[string]*Emitter
	Filters          map[string]*Filter
	ColorSpaces      map[string]*ColorSpace // TODO: find example since spec sheet is not specific enough
	Gamuts           map[string]*Gamut
	DMXProfiles      map[string]*DMXProfile
	DMXModes         map[string]*DMXMode
	DMXChannels      map[string]*DMXChannel
	SubPhysicalUnits map[string]*SubPhysicalUnit //TODO: find example since spec sheet is not specific enough
	WiringObjects    map[string]*WiringObject    //TODO: find example since spec sheet is not specific enough
	ChannelFunctions map[string]*ChannelFunction
	Geometries       map[string]*GeometryNodeReference
	Models           map[string]*Model
}

var refPointers ReferencePointers = ReferencePointers{
	ActivationGroups: make(map[string]*ActivationGroup),
	Features:         make(map[string]*Feature),
	Attributes:       make(map[string]*Attribute),
	Wheels:           make(map[string]*Wheel),
	WheelSlots:       make(map[string]*WheelSlot),
	Emitters:         make(map[string]*Emitter),
	Filters:          make(map[string]*Filter),
	ColorSpaces:      make(map[string]*ColorSpace),
	Gamuts:           make(map[string]*Gamut),
	DMXProfiles:      make(map[string]*DMXProfile),
	DMXModes:         make(map[string]*DMXMode),
	DMXChannels:      make(map[string]*DMXChannel),
	SubPhysicalUnits: make(map[string]*SubPhysicalUnit),
	WiringObjects:    make(map[string]*WiringObject),
	ChannelFunctions: make(map[string]*ChannelFunction),
	Geometries:       make(map[string]*GeometryNodeReference),
	Models:           make(map[string]*Model),
}

type ReferenceCreation interface {
	CreateReferencePointer()
}

func CreateReferencePointers[T ReferenceCreation](source *[]T) {
	for _, element := range *source {
		element.CreateReferencePointer()
	}
}

type GeometryReferenceCreation interface {
	CreateGeometryReferencePointer(parentPrefix string)
}

func CreateGeometryReferencePointers[T GeometryReferenceCreation](source *[]T, parentPrefix string) {
	for _, element := range *source {
		element.CreateGeometryReferencePointer(parentPrefix)
	}
}

type ReferenceResolver interface {
	ResolveReference()
}

func ResolveReferences[T ReferenceResolver](source *[]T) {
	if source == nil {
		return
	}
	for i := range *source {
		(*source)[i].ResolveReference()
	}
}
