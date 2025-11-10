package XMLTypes

type DMXMode struct {
	Name        string           `xml:",attr"`
	Description string           `xml:",attr"`
	Geometry    XMLNodeReference `xml:",attr"`
	DMXChannels []DMXChannel     `xml:"DMXChannels>DMXChannel"`
	Relations   []Relation       `xml:"Relations>Relation"`
	FTMacros    []FTMacro        `xml:"FTMacros>FTMacro"`
}

type DMXChannel struct {
	DMXBreak        int              `xml:",attr"`
	Offset          []int            `xml:",attr"`
	InitialFunction XMLNodeReference `xml:",attr"`
	Highlight       XMLDMXValue      `xml:",attr"`
	Geometry        XMLNodeReference `xml:",attr"`
	LogicalChannels []LogicalChannel `xml:"LogicalChannel"`
}

type LogicalChannel struct {
	Attribute          XMLNodeReference  `xml:",attr"`
	Snap               string            `xml:",attr"` // enum
	Master             string            `xml:",attr"` // enum
	MibFade            float32           `xml:",attr"`
	DMXChangeTimeLimit float32           `xml:",attr"`
	ChannelFunctions   []ChannelFunction `xml:"ChannelFunction"`
}

type ChannelFunction struct {
	Name              string           `xml:",attr"`
	Attribute         XMLNodeReference `xml:",attr"`
	OriginalAttribute string           `xml:",attr"`
	DMXFrom           XMLDMXValue      `xml:",attr"`
	Default           XMLDMXValue      `xml:",attr"`
	PhysicalFrom      float32          `xml:",attr"`
	PhysicalTo        float32          `xml:",attr"`
	RealFade          float32          `xml:",attr"`
	RealAcceleration  float32          `xml:",attr"`
	Wheel             XMLNodeReference `xml:",attr"`
	Emitter           XMLNodeReference `xml:",attr"`
	Filter            XMLNodeReference `xml:",attr"`
	ColorSpace        XMLNodeReference `xml:",attr"`
	Gamut             XMLNodeReference `xml:",attr"`
	ModeMaster        XMLNodeReference `xml:",attr"`
	ModeFrom          XMLDMXValue      `xml:",attr"`
	ModeTo            XMLDMXValue      `xml:",attr"`
	DMXProfile        XMLNodeReference `xml:",attr"`
	Min               float32          `xml:",attr"` // Fallback to PhysicalFrom if 0.0
	Max               float32          `xml:",attr"` // Fallback to PhysicalTo if 0.0
	CustomName        string           `xml:",attr"` // Default: Node Name of the Channel function Example: Head_Dimmer.Dimmer.Dimmer
	ChannelSets       []ChannelSet     `xml:"ChannelSet"`
	SubChannelSets    []SubChannelSet  `xml:"SubChannelSet"`
}

type ChannelSet struct {
	Name           string      `xml:",attr"`
	DMXFrom        XMLDMXValue `xml:",attr"`
	PhysicalFrom   float32     `xml:",attr"`
	PhysicalTo     float32     `xml:",attr"`
	WheelSlotIndex int         `xml:",attr"`
}

type SubChannelSet struct {
	Name            string           `xml:",attr"`
	PhysicalFrom    float32          `xml:",attr"`
	PhysicalTo      float32          `xml:",attr"`
	SubPhysicalUnit XMLNodeReference `xml:",attr"`
	DMXProfile      XMLNodeReference `xml:",attr"`
}

type Relation struct {
	Name     string           `xml:",attr"`
	Master   XMLNodeReference `xml:",attr"`
	Follower XMLNodeReference `xml:",attr"`
	Type     string           `xml:",attr"` //enum with "Multiply" or "Override"
}

type FTMacro struct {
	Name            string            `xml:",attr"`
	ChannelFunction *XMLNodeReference `xml:",attr"`
	MacroDMXs       []MacroDMX        `xml:"MacroDMX"`
}

type MacroDMX struct {
	Steps []MacroDMXStep `xml:"MacroDMXStep"`
}

type MacroDMXStep struct {
	Duration  float32         `xml:",attr"`
	DMXValues []MacroDMXValue `xml:"DMXValue"`
}

type MacroDMXValue struct {
	Value      XMLDMXValue      `xml:",attr"`
	DMXChannel XMLNodeReference `xml:",attr"`
}
