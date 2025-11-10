package Types

type DMXMode struct {
	Name        string
	Description string
	Geometry    *Geometry
	DMXChannels []DMXChannel
	Relations   []Relation
	FTMacros    []FTMacro
}

type DMXChannel struct {
	DMXBreak        int
	Offset          []int
	InitialFunction string //TODO:
	Highlight       DMXValue
	Geometry        *Geometry
	LogicalChannels []LogicalChannel
}

type LogicalChannel struct {
	Attribute          *Attribute
	Snap               string // enum
	Master             string // enum
	MibFade            float32
	DMXChangeTimeLimit float32
	ChannelFunctions   []ChannelFunction
}

type ChannelFunction struct {
	Name              string
	Attribute         *Attribute
	OriginalAttribute string
	DMXFrom           DMXValue
	Default           DMXValue
	PhysicalFrom      float32
	PhysicalTo        float32
	RealFade          float32
	RealAcceleration  float32
	Wheel             *Wheel
	Emitter           string // TODO: link to emitter
	Filter            string // TODO: link to filter
	ColorSpace        string //TODO:
	Gamut             string //TODO:
	ModeMaster        string //TODO:
	ModeFrom          DMXValue
	ModeTo            DMXValue
	DMXProfile        string  //TODO:
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
	WheelSlot    *WheelSlot
}

type SubChannelSet struct {
	Name            string
	DMXFrom         DMXValue
	PhysicalFrom    float32
	PhysicalTo      float32
	SubPhysicalUnit *SubPhysicalUnit
	DMXProfile      string //TODO:
}

type Relation struct {
	Name     string
	Master   *DMXChannel
	Follower *DMXChannel
	Type     string //enum with "Multiply" or "Override"
}

// TODO:
type FTMacro struct {
}
