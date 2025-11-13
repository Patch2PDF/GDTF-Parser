package XMLTypes

import (
	"strconv"

	Types "github.com/Patch2PDF/GDTF-Parser/types"
)

type DMXMode struct {
	Name        string           `xml:",attr"`
	Description string           `xml:",attr"`
	Geometry    XMLNodeReference `xml:",attr"`
	DMXChannels []DMXChannel     `xml:"DMXChannels>DMXChannel"`
	Relations   []Relation       `xml:"Relations>Relation"`
	FTMacros    []FTMacro        `xml:"FTMacros>FTMacro"`
}

func (dmxMode DMXMode) Parse() Types.DMXMode {
	return Types.DMXMode{
		Name:        dmxMode.Name,
		Description: dmxMode.Description,
		Geometry: Types.NodeReference[Types.GeometryNodeReference]{
			String: dmxMode.Geometry,
		},
		DMXChannels: ParseList(&dmxMode.DMXChannels),
		Relations:   ParseList(&dmxMode.Relations),
		FTMacros:    ParseList(&dmxMode.FTMacros),
	}
}

type DMXChannel struct {
	DMXBreak        int              `xml:",attr"`
	Offset          IntList          `xml:",attr"`
	InitialFunction XMLNodeReference `xml:",attr"`
	Highlight       XMLDMXValue      `xml:",attr"`
	Geometry        XMLNodeReference `xml:",attr"`
	LogicalChannels []LogicalChannel `xml:"LogicalChannel"`
}

func (dmx DMXChannel) Parse() Types.DMXChannel {
	return Types.DMXChannel{
		DMXBreak:        dmx.DMXBreak,
		Offset:          dmx.Offset,
		InitialFunction: dmx.InitialFunction,
		Highlight:       dmx.Highlight,
		Geometry: Types.NodeReference[Types.GeometryNodeReference]{
			String: dmx.Geometry,
		},
		LogicalChannels: ParseList(&dmx.LogicalChannels),
	}
}

type LogicalChannel struct {
	Attribute          XMLNodeReference  `xml:",attr"`
	Snap               string            `xml:",attr"` // enum
	Master             string            `xml:",attr"` // enum
	MibFade            float32           `xml:",attr"`
	DMXChangeTimeLimit float32           `xml:",attr"`
	ChannelFunctions   []ChannelFunction `xml:"ChannelFunction"`
}

func (dmx LogicalChannel) Parse() Types.LogicalChannel {
	return Types.LogicalChannel{
		Attribute: Types.NodeReference[Types.Attribute]{
			String: dmx.Attribute,
		},
		Snap:               dmx.Snap,
		Master:             dmx.Master,
		MibFade:            dmx.MibFade,
		DMXChangeTimeLimit: dmx.DMXChangeTimeLimit,
		ChannelFunctions:   ParseList(&dmx.ChannelFunctions),
	}
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

func (dmx ChannelFunction) Parse() Types.ChannelFunction {
	return Types.ChannelFunction{
		Name: dmx.Name,
		Attribute: Types.NodeReference[Types.Attribute]{
			String: dmx.Attribute,
		},
		OriginalAttribute: dmx.OriginalAttribute,
		DMXFrom:           dmx.DMXFrom,
		Default:           dmx.Default,
		PhysicalFrom:      dmx.PhysicalFrom,
		PhysicalTo:        dmx.PhysicalTo,
		RealFade:          dmx.RealFade,
		RealAcceleration:  dmx.RealAcceleration,
		Wheel: Types.NodeReference[Types.Wheel]{
			String: dmx.Wheel,
		},
		Emitter: Types.NodeReference[Types.Emitter]{
			String: dmx.Emitter,
		},
		Filter: Types.NodeReference[Types.Filter]{
			String: dmx.Filter,
		},
		ColorSpace: Types.NodeReference[Types.ColorSpace]{
			String: dmx.ColorSpace,
		},
		Gamut: Types.NodeReference[Types.Gamut]{
			String: dmx.Gamut,
		},
		ModeMaster: dmx.ModeMaster,
		ModeFrom:   dmx.ModeFrom,
		ModeTo:     dmx.ModeTo,
		DMXProfile: Types.NodeReference[Types.DMXProfile]{
			String: dmx.DMXProfile,
		},
		Min:            dmx.Min,
		Max:            dmx.Max,
		CustomName:     dmx.CustomName,
		ChannelSets:    ParseList(&dmx.ChannelSets),
		SubChannelSets: ParseList(&dmx.SubChannelSets),
	}
}

type ChannelSet struct {
	Name           string      `xml:",attr"`
	DMXFrom        XMLDMXValue `xml:",attr"`
	PhysicalFrom   float32     `xml:",attr"`
	PhysicalTo     float32     `xml:",attr"`
	WheelSlotIndex int         `xml:",attr"`
}

func (dmx ChannelSet) Parse() Types.ChannelSet {
	return Types.ChannelSet{
		Name:         dmx.Name,
		DMXFrom:      dmx.DMXFrom,
		PhysicalFrom: dmx.PhysicalFrom,
		PhysicalTo:   dmx.PhysicalTo,
		WheelSlot: Types.NodeReference[Types.WheelSlot]{
			String: strconv.FormatInt(int64(dmx.WheelSlotIndex), 10),
		},
	}
}

type SubChannelSet struct {
	Name            string           `xml:",attr"`
	PhysicalFrom    float32          `xml:",attr"`
	PhysicalTo      float32          `xml:",attr"`
	SubPhysicalUnit XMLNodeReference `xml:",attr"`
	DMXProfile      XMLNodeReference `xml:",attr"`
}

func (dmx SubChannelSet) Parse() Types.SubChannelSet {
	return Types.SubChannelSet{
		Name:         dmx.Name,
		PhysicalFrom: dmx.PhysicalFrom,
		PhysicalTo:   dmx.PhysicalTo,
		SubPhysicalUnit: Types.NodeReference[Types.SubPhysicalUnit]{
			String: dmx.SubPhysicalUnit,
		},
		DMXProfile: Types.NodeReference[Types.DMXProfile]{
			String: dmx.DMXProfile,
		},
	}
}

type Relation struct {
	Name     string           `xml:",attr"`
	Master   XMLNodeReference `xml:",attr"`
	Follower XMLNodeReference `xml:",attr"`
	Type     string           `xml:",attr"` //enum with "Multiply" or "Override"
}

func (dmx Relation) Parse() Types.Relation {
	return Types.Relation{}
}

type FTMacro struct {
	Name            string            `xml:",attr"`
	ChannelFunction *XMLNodeReference `xml:",attr"`
	MacroDMXs       []MacroDMX        `xml:"MacroDMX"`
}

func (dmx FTMacro) Parse() Types.FTMacro {
	return Types.FTMacro{
		Name: dmx.Name,
		ChannelFunction: Types.NodeReference[Types.ChannelFunction]{
			String: *dmx.ChannelFunction,
		},
		MacroDMXs: ParseList(&dmx.MacroDMXs),
	}
}

type MacroDMX struct {
	Steps []MacroDMXStep `xml:"MacroDMXStep"`
}

func (dmx MacroDMX) Parse() Types.MacroDMX {
	return Types.MacroDMX{
		Steps: ParseList(&dmx.Steps),
	}
}

type MacroDMXStep struct {
	Duration  float32         `xml:",attr"`
	DMXValues []MacroDMXValue `xml:"DMXValue"`
}

func (dmx MacroDMXStep) Parse() Types.MacroDMXStep {
	return Types.MacroDMXStep{
		Duration:  dmx.Duration,
		DMXValues: ParseList(&dmx.DMXValues),
	}
}

type MacroDMXValue struct {
	Value      XMLDMXValue      `xml:",attr"`
	DMXChannel XMLNodeReference `xml:",attr"`
}

func (dmx MacroDMXValue) Parse() Types.MacroDMXValue {
	return Types.MacroDMXValue{
		Value: dmx.Value,
		DMXChannel: Types.NodeReference[Types.DMXChannel]{
			String: dmx.DMXChannel,
		},
	}
}
