package Types

type DMXMode struct {
	Name        string
	Description string
	Geometry    NodeReference[GeometryNodeReference]
	DMXChannels []DMXChannel
	Relations   []Relation
	FTMacros    []FTMacro
}

type DMXChannel struct {
	DMXBreak        int
	Offset          []int
	InitialFunction string //TODO:
	Highlight       DMXValue
	Geometry        NodeReference[GeometryNodeReference]
	LogicalChannels []LogicalChannel
}

type LogicalChannel struct {
	Attribute          NodeReference[Attribute]
	Snap               string // enum
	Master             string // enum
	MibFade            float32
	DMXChangeTimeLimit float32
	ChannelFunctions   []ChannelFunction
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
	ChannelSets       []ChannelSet
	SubChannelSets    []SubChannelSet
}

type ChannelSet struct {
	Name         string
	DMXFrom      DMXValue
	PhysicalFrom float32
	PhysicalTo   float32
	WheelSlot    NodeReference[WheelSlot]
}

type SubChannelSet struct {
	Name            string
	PhysicalFrom    float32
	PhysicalTo      float32
	SubPhysicalUnit NodeReference[SubPhysicalUnit]
	DMXProfile      NodeReference[DMXProfile]
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
	MacroDMXs       []MacroDMX
}

type MacroDMX struct {
	Steps []MacroDMXStep
}

type MacroDMXStep struct {
	Duration  float32
	DMXValues []MacroDMXValue
}

type MacroDMXValue struct {
	Value      DMXValue
	DMXChannel NodeReference[DMXChannel]
}
