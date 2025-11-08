package gdtfparser

import "time"

// TODO: DMXValue Type
// see gdtf https://www.gdtf.eu/gdtf/file-spec/file-format-definition/#attrtype-dmxvalue
type DMXValue = string

type FixtureType struct {
	FixtureTypeID    string
	Name             string
	ShortName        string
	LongName         string
	Manufacturer     string
	Description      string
	Thumbnail        *string
	ThumbnailOffsetX int
	ThumbnailOffsetY int
	CanHaveChildren  bool
	RefFT            *string

	AttributeDefinitions AttributeDefinitions
	Wheels               *[]Wheel
	PhysicalDescriptions *[]string
	Models               *[]Model
	Geometries           []Geometry
	DMXModes             []DMXMode
	Revisions            *[]string
	FTPresets            *[]string
	Protocols            *[]string
}

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

type Wheel struct {
	Name string
}

// TODO:
type WheelSlot struct {
	Name          string
	Color         string
	Filter        string // TODO:
	MediaFileName string
}

type Model struct {
	Name            string
	Length          float32
	Width           float32
	Height          float32
	PrimitiveType   string // TODO: enum
	File            *string
	SVGOffsetX      float32
	SVGOffsetY      float32
	SVGSideOffsetX  float32
	SVGSideOffsetY  float32
	SVGFrontOffsetX float32
	SVGFrontOffsetY float32
}

type Geometry struct {
}

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
	Snap               string // TODO: enum
	Master             string // TODO: enum
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
	Type     string //TODO: enum with "Multiply" or "Override"
}

// TODO:
type FTMacro struct {
}

type Revision struct {
	Text       string
	Date       time.Time
	UserID     uint
	ModifiedBy string
}

//TODO: geometry and protocol
