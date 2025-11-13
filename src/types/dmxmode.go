package Types

type DMXMode struct {
	Name        string
	Description string
	Geometry    NodeReference[GeometryNodeReference]
	DMXChannels []*DMXChannel
	Relations   []*Relation
	FTMacros    []*FTMacro
}

func (obj *DMXMode) CreateReferencePointer() {
	refPointers.DMXModes[obj.Name] = obj
	for _, element := range obj.DMXChannels {
		refPointers.DMXChannels[obj.Name+"."+element.GetName()] = element
		for _, logicalChannel := range element.LogicalChannels {
			for _, channelFunction := range logicalChannel.ChannelFunctions {
				channelFunctionName := (obj.Name +
					"." +
					element.GetName() +
					"." +
					refPointers.Attributes[logicalChannel.Attribute.String].Name + //TODO: this might not work, as it could be not yet filled, need to definetely ensure we do not hit a race condition
					"." +
					channelFunction.Name)
				refPointers.ChannelFunctions[channelFunctionName] = channelFunction
			}
		}
	}
}

func (obj *DMXMode) ResolveReference() {
	obj.Geometry.Ptr = &GeometryNodeReference{
		// Ptr: ,
		// Type
		// TODO:
	}
	ResolveReferences(&obj.DMXChannels)
	ResolveReferences(&obj.FTMacros)
}

type DMXChannel struct {
	DMXBreak        int
	Offset          []int
	InitialFunction string //TODO:
	Highlight       DMXValue
	Geometry        NodeReference[GeometryNodeReference]
	LogicalChannels []*LogicalChannel
}

func (element DMXChannel) GetName() string {
	// according to https://www.gdtf.eu/gdtf/file-spec/dmx-mode-collect/#dmx-channel-collect
	return element.Geometry.String + "_" + element.LogicalChannels[0].Attribute.String
}

func (obj *DMXChannel) ResolveReference() {
	//TODO: Geometry Reference
	ResolveReferences(&obj.LogicalChannels)
}

type LogicalChannel struct {
	Attribute          NodeReference[Attribute]
	Snap               string // enum
	Master             string // enum
	MibFade            float32
	DMXChangeTimeLimit float32
	ChannelFunctions   []*ChannelFunction
}

func (obj *LogicalChannel) ResolveReference() {
	obj.Attribute.Ptr = refPointers.Attributes[obj.Attribute.String]
	ResolveReferences(&obj.ChannelFunctions)
}

type ChannelFunction struct {
	Name              string
	Attribute         NodeReference[Attribute]
	OriginalAttribute string
	DMXFrom           DMXValue
	Default           DMXValue
	PhysicalFrom      float32
	PhysicalTo        float32
	RealFade          float32
	RealAcceleration  float32
	Wheel             NodeReference[Wheel]
	Emitter           NodeReference[Emitter]
	Filter            NodeReference[Filter]
	ColorSpace        NodeReference[ColorSpace]
	Gamut             NodeReference[Gamut]
	ModeMaster        string //TODO:
	ModeFrom          DMXValue
	ModeTo            DMXValue
	DMXProfile        NodeReference[DMXProfile]
	Min               float32 // Fallback to PhysicalFrom if 0.0
	Max               float32 // Fallback to PhysicalTo if 0.0
	CustomName        string  // Default: Node Name of the Channel function Example: Head_Dimmer.Dimmer.Dimmer
	ChannelSets       []*ChannelSet
	SubChannelSets    []*SubChannelSet
}

func (obj *ChannelFunction) ResolveReference() {
	obj.Attribute.Ptr = refPointers.Attributes[obj.Attribute.String]
	obj.Wheel.Ptr = refPointers.Wheels[obj.Wheel.String]
	obj.Emitter.Ptr = refPointers.Emitters[obj.Emitter.String]
	obj.Filter.Ptr = refPointers.Filters[obj.Filter.String]
	//TODO: ColorSpace
	obj.Gamut.Ptr = refPointers.Gamuts[obj.Gamut.String]
	obj.DMXProfile.Ptr = refPointers.DMXProfiles[obj.DMXProfile.String]

	ResolveReferences(&obj.ChannelSets)
	ResolveReferences(&obj.SubChannelSets)
}

type ChannelSet struct {
	Name         string
	DMXFrom      DMXValue
	PhysicalFrom float32
	PhysicalTo   float32
	WheelSlot    NodeReference[WheelSlot]
}

func (obj *ChannelSet) ResolveReference() {
	obj.WheelSlot.Ptr = refPointers.WheelSlots[obj.WheelSlot.String]
}

type SubChannelSet struct {
	Name            string
	PhysicalFrom    float32
	PhysicalTo      float32
	SubPhysicalUnit NodeReference[SubPhysicalUnit]
	DMXProfile      NodeReference[DMXProfile]
}

func (obj *SubChannelSet) ResolveReference() {
	obj.SubPhysicalUnit.Ptr = refPointers.SubPhysicalUnits[obj.SubPhysicalUnit.String]
	obj.DMXProfile.Ptr = refPointers.DMXProfiles[obj.DMXProfile.String]
}

type Relation struct {
	Name     string
	Master   *DMXChannel
	Follower *DMXChannel
	Type     string //enum with "Multiply" or "Override"
}

type FTMacro struct {
	Name            string
	ChannelFunction NodeReference[ChannelFunction]
	MacroDMXs       []*MacroDMX
}

func (obj *FTMacro) ResolveReference() {
	obj.ChannelFunction.Ptr = refPointers.ChannelFunctions[obj.ChannelFunction.String]
	ResolveReferences(&obj.MacroDMXs)
}

type MacroDMX struct {
	Steps []*MacroDMXStep
}

func (obj *MacroDMX) ResolveReference() {
	ResolveReferences(&obj.Steps)
}

type MacroDMXStep struct {
	Duration  float32
	DMXValues []*MacroDMXValue
}

func (obj *MacroDMXStep) ResolveReference() {
	ResolveReferences(&obj.DMXValues)
}

type MacroDMXValue struct {
	Value      DMXValue
	DMXChannel NodeReference[DMXChannel]
}

func (obj *MacroDMXValue) ResolveReference() {
	obj.DMXChannel.Ptr = refPointers.DMXChannels[obj.DMXChannel.String]
}
