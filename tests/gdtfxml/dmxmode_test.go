package gdtfxml_test

import (
	"testing"

	XMLTypes "github.com/Patch2PDF/GDTF-Parser/internal/types/gdtfxml"
)

var dmxModeXML = `
	<DMXModes>
    <DMXMode Name="RGB Mode" Geometry="Body">
        <DMXChannels>
            <DMXChannel Offset="1" Highlight="255/1" Geometry="Body">
                <LogicalChannel Attribute="ColorAdd_R">
                    <ChannelFunction Attribute="ColorAdd_R" DMXFrom="0/1" PhysicalFrom="0" PhysicalTo="1" />
                </LogicalChannel>
            </DMXChannel>
            <DMXChannel Offset="2" Highlight="255/1" Geometry="Body">
                <LogicalChannel Attribute="ColorAdd_G">
                    <ChannelFunction Attribute="ColorAdd_G" DMXFrom="0/1" PhysicalFrom="0" PhysicalTo="1" />
                </LogicalChannel>
            </DMXChannel>
            <DMXChannel Offset="3" Highlight="255/1" Geometry="Body">
                <LogicalChannel Attribute="ColorAdd_B">
                    <ChannelFunction Attribute="ColorAdd_B" DMXFrom="0/1" PhysicalFrom="0" PhysicalTo="1" />
                </LogicalChannel>
            </DMXChannel>
        </DMXChannels>
        <Relations />
    </DMXMode>
    <DMXMode Name="Intensity + RGB Mode" Geometry="Body">
        <DMXChannels>
            <DMXChannel Offset="1" Highlight="255/1" Geometry="Body">
                <LogicalChannel Attribute="Dimmer" Master="Grand">
                    <ChannelFunction Attribute="Dimmer" DMXFrom="0/1" PhysicalFrom="0" PhysicalTo="1" />
                </LogicalChannel>
            </DMXChannel>
            <DMXChannel Offset="2" Default="255/1" Highlight="255/1" Geometry="Body">
                <LogicalChannel Attribute="ColorAdd_R">
                    <ChannelFunction Attribute="ColorAdd_R" DMXFrom="0/1" PhysicalFrom="0" PhysicalTo="1" />
                </LogicalChannel>
            </DMXChannel>
            <DMXChannel Offset="3" Default="255/1" Highlight="255/1" Geometry="Body">
                <LogicalChannel Attribute="ColorAdd_G">
                    <ChannelFunction Attribute="ColorAdd_G" DMXFrom="0/1" PhysicalFrom="0" PhysicalTo="1" />
                </LogicalChannel>
            </DMXChannel>
            <DMXChannel Offset="4" Default="255/1" Highlight="255/1" Geometry="Body">
                <LogicalChannel Attribute="ColorAdd_B">
                    <ChannelFunction Attribute="ColorAdd_B" DMXFrom="0/1" PhysicalFrom="0" PhysicalTo="1" />
                </LogicalChannel>
            </DMXChannel>
        </DMXChannels>
        <Relations />
    </DMXMode>
    <DMXMode Name="Intensity + ColorMacro Mode" Geometry="Body">
        <DMXChannels>
            <DMXChannel Offset="1" Highlight="255/1" Geometry="Body">
                <LogicalChannel Attribute="Dimmer">
                    <ChannelFunction Attribute="Dimmer" DMXFrom="0/1" PhysicalFrom="0" PhysicalTo="1" />
                </LogicalChannel>
            </DMXChannel>
            <DMXChannel Offset="2" Highlight="0/1" Geometry="Body">
                <LogicalChannel Attribute="ColorMacro1">
                    <ChannelFunction Attribute="ColorMacro1" DMXFrom="0/1" PhysicalFrom="0" PhysicalTo="1" />
                </LogicalChannel>
            </DMXChannel>
        </DMXChannels>
        <Relations />
    </DMXMode>
	</DMXModes>
`

var dmxModeStruct = []XMLTypes.DMXMode{
	{
		Name:        "RGB Mode",
		Description: "",
		Geometry:    "Body",
		DMXChannels: []XMLTypes.DMXChannel{
			{
				DMXBreak:        0,
				Offset:          XMLTypes.IntList{1},
				InitialFunction: "",
				Highlight:       "255/1",
				Geometry:        "Body",
				LogicalChannels: []XMLTypes.LogicalChannel{
					{
						Attribute:          "ColorAdd_R",
						Snap:               "",
						Master:             "",
						MibFade:            0,
						DMXChangeTimeLimit: 0,
						ChannelFunctions: []XMLTypes.ChannelFunction{
							{
								Name:              "",
								Attribute:         "ColorAdd_R",
								OriginalAttribute: "",
								DMXFrom:           "0/1",
								Default:           "",
								PhysicalFrom:      0,
								PhysicalTo:        1,
								RealFade:          0,
								RealAcceleration:  0,
								Wheel:             "",
								Emitter:           "",
								Filter:            "",
								ColorSpace:        "",
								Gamut:             "",
								ModeMaster:        "",
								ModeFrom:          "",
								ModeTo:            "",
								DMXProfile:        "",
								Min:               0,
								Max:               0,
								CustomName:        "",
								ChannelSets:       nil,
								SubChannelSets:    nil,
							},
						},
					},
				},
			},
			{
				DMXBreak:        0,
				Offset:          XMLTypes.IntList{2},
				InitialFunction: "",
				Highlight:       "255/1",
				Geometry:        "Body",
				LogicalChannels: []XMLTypes.LogicalChannel{
					{
						Attribute:          "ColorAdd_G",
						Snap:               "",
						Master:             "",
						MibFade:            0,
						DMXChangeTimeLimit: 0,
						ChannelFunctions: []XMLTypes.ChannelFunction{
							{
								Attribute:    "ColorAdd_G",
								DMXFrom:      "0/1",
								PhysicalFrom: 0,
								PhysicalTo:   1,
							},
						},
					},
				},
			},
			{
				DMXBreak:        0,
				Offset:          XMLTypes.IntList{3},
				InitialFunction: "",
				Highlight:       "255/1",
				Geometry:        "Body",
				LogicalChannels: []XMLTypes.LogicalChannel{
					{
						Attribute: "ColorAdd_B",
						ChannelFunctions: []XMLTypes.ChannelFunction{
							{
								Attribute:    "ColorAdd_B",
								DMXFrom:      "0/1",
								PhysicalFrom: 0,
								PhysicalTo:   1,
							},
						},
					},
				},
			},
		},
		Relations: nil,
		FTMacros:  nil,
	},
	{
		Name:        "Intensity + RGB Mode",
		Description: "",
		Geometry:    "Body",
		DMXChannels: []XMLTypes.DMXChannel{
			{
				Offset:    XMLTypes.IntList{1},
				Highlight: "255/1",
				Geometry:  "Body",
				LogicalChannels: []XMLTypes.LogicalChannel{
					{
						Attribute: "Dimmer",
						Master:    "Grand",
						ChannelFunctions: []XMLTypes.ChannelFunction{
							{
								Attribute:    "Dimmer",
								DMXFrom:      "0/1",
								PhysicalFrom: 0,
								PhysicalTo:   1,
							},
						},
					},
				},
			},
			{
				Offset:    XMLTypes.IntList{2},
				Highlight: "255/1",
				Geometry:  "Body",
				LogicalChannels: []XMLTypes.LogicalChannel{
					{
						Attribute: "ColorAdd_R",
						ChannelFunctions: []XMLTypes.ChannelFunction{
							{
								Attribute:    "ColorAdd_R",
								DMXFrom:      "0/1",
								PhysicalFrom: 0,
								PhysicalTo:   1,
							},
						},
					},
				},
			},
			{
				Offset:    XMLTypes.IntList{3},
				Highlight: "255/1",
				Geometry:  "Body",
				LogicalChannels: []XMLTypes.LogicalChannel{
					{
						Attribute: "ColorAdd_G",
						ChannelFunctions: []XMLTypes.ChannelFunction{
							{
								Attribute:    "ColorAdd_G",
								DMXFrom:      "0/1",
								PhysicalFrom: 0,
								PhysicalTo:   1,
							},
						},
					},
				},
			},
			{
				Offset:    XMLTypes.IntList{4},
				Highlight: "255/1",
				Geometry:  "Body",
				LogicalChannels: []XMLTypes.LogicalChannel{
					{
						Attribute: "ColorAdd_B",
						ChannelFunctions: []XMLTypes.ChannelFunction{
							{
								Attribute:    "ColorAdd_B",
								DMXFrom:      "0/1",
								PhysicalFrom: 0,
								PhysicalTo:   1,
							},
						},
					},
				},
			},
		},
	},
	{
		Name:     "Intensity + ColorMacro Mode",
		Geometry: "Body",
		DMXChannels: []XMLTypes.DMXChannel{
			{
				Offset:    XMLTypes.IntList{1},
				Highlight: "255/1",
				Geometry:  "Body",
				LogicalChannels: []XMLTypes.LogicalChannel{
					{
						Attribute: "Dimmer",
						ChannelFunctions: []XMLTypes.ChannelFunction{
							{
								Attribute:    "Dimmer",
								DMXFrom:      "0/1",
								PhysicalFrom: 0,
								PhysicalTo:   1,
							},
						},
					},
				},
			},
			{
				Offset:    XMLTypes.IntList{2},
				Highlight: "0/1",
				Geometry:  "Body",
				LogicalChannels: []XMLTypes.LogicalChannel{
					{
						Attribute: "ColorMacro1",
						ChannelFunctions: []XMLTypes.ChannelFunction{
							{
								Attribute:    "ColorMacro1",
								DMXFrom:      "0/1",
								PhysicalFrom: 0,
								PhysicalTo:   1,
							},
						},
					},
				},
			},
		},
	},
}

type dmxModeTest struct {
	DMXModes []XMLTypes.DMXMode `xml:"DMXMode"`
}

func TestDMXMode(t *testing.T) {
	parsingTest(t, dmxModeXML, "DMX Mode", dmxModeTest{DMXModes: dmxModeStruct})
}

var dmxModeDependenciesXML = `
<DMXModes>
    <DMXMode Name="Default" Geometry="Head">
        <DMXChannels>
            <DMXChannel Offset="1"  Geometry="Head">
                <LogicalChannel Attribute="Gobo1">
                    <ChannelFunction Attribute="Gobo1" DMXFrom="0/1" Wheel="GoboWheel1">
                        <ChannelSet Name="Open" DMXFrom="0/1" PhysicalFrom="0" PhysicalTo="0" WheelSlotIndex="1" />
                        <ChannelSet Name="Gobo 1 index" DMXFrom="11/1" PhysicalFrom="0" PhysicalTo="0" WheelSlotIndex="2" />
                        <ChannelSet Name="Gobo 2 index" DMXFrom="21/1" PhysicalFrom="0" PhysicalTo="0" WheelSlotIndex="3" />
                        <ChannelSet Name="Gobo 3 index" DMXFrom="31/1" PhysicalFrom="0" PhysicalTo="0" WheelSlotIndex="4" />
                        <ChannelSet Name="Gobo 4 index" DMXFrom="41/1" PhysicalFrom="0" PhysicalTo="0" WheelSlotIndex="5" />
                        <ChannelSet Name="Gobo 5 index" DMXFrom="51/1" PhysicalFrom="0" PhysicalTo="0" WheelSlotIndex="6" />
                        <ChannelSet Name="Gobo 1 rot" DMXFrom="61/1" PhysicalFrom="0" PhysicalTo="0" WheelSlotIndex="2" />
                        <ChannelSet Name="Gobo 2 rot" DMXFrom="71/1" PhysicalFrom="0" PhysicalTo="0" WheelSlotIndex="3" />
                        <ChannelSet Name="Gobo 3 rot" DMXFrom="81/1" PhysicalFrom="0" PhysicalTo="0" WheelSlotIndex="4" />
                        <ChannelSet Name="Gobo 4 rot" DMXFrom="91/1" PhysicalFrom="0" PhysicalTo="0" WheelSlotIndex="5" />
                        <ChannelSet Name="Gobo 5 rot" DMXFrom="101/1" PhysicalFrom="0" PhysicalTo="0" WheelSlotIndex="6" />
                    </ChannelFunction>
                    <ChannelFunction Attribute="Gobo1WheelSpin" DMXFrom="111/1" >
                        <ChannelSet Name="CCW" DMXFrom="111/1" PhysicalFrom="-60" PhysicalTo="0" />
                        <ChannelSet Name="Stop" DMXFrom="183/1" PhysicalFrom="0" PhysicalTo="0" />
                        <ChannelSet Name="CW" DMXFrom="184/1" PhysicalFrom="0" PhysicalTo="60" />
                    </ChannelFunction>
                </LogicalChannel>
            </DMXChannel>
            <DMXChannel Offset="2" Geometry="Head">
                <LogicalChannel Attribute="Gobo1Pos">
                    <ChannelFunction Attribute="Gobo1Pos" DMXFrom="0/1" PhysicalFrom="0" PhysicalTo="540" ModeMaster="Head_Gobo1" ModeFrom="0/1" ModeTo="60/1" />
                    <ChannelFunction Attribute="Gobo1PosRotate" DMXFrom="0/1" PhysicalFrom="-150" PhysicalTo="150" ModeMaster="Head_Gobo1" ModeFrom="61/1" ModeTo="110/1">
                        <ChannelSet Name="CW" DMXFrom="0/1" PhysicalFrom="-150" PhysicalTo="0" />
                        <ChannelSet Name="Stop" DMXFrom="128/1" PhysicalFrom="0" PhysicalTo="0" />
                        <ChannelSet Name="CCW" DMXFrom="129/1" PhysicalFrom="0" PhysicalTo="150"/>
                    </ChannelFunction>
                </LogicalChannel>
            </DMXChannel>
        </DMXChannels>
        <Relations />
    </DMXMode>
</DMXModes>
`

var dmxModeDependenciesStruct = []XMLTypes.DMXMode{
	{
		Name:     "Default",
		Geometry: "Head",
		DMXChannels: []XMLTypes.DMXChannel{
			{
				Offset:   XMLTypes.IntList{1},
				Geometry: "Head",
				LogicalChannels: []XMLTypes.LogicalChannel{
					{
						Attribute: "Gobo1",
						ChannelFunctions: []XMLTypes.ChannelFunction{
							{
								Attribute: "Gobo1",
								DMXFrom:   "0/1",
								Wheel:     "GoboWheel1",
								ChannelSets: []XMLTypes.ChannelSet{
									{Name: "Open", DMXFrom: "0/1", PhysicalFrom: 0, PhysicalTo: 0, WheelSlotIndex: 1},
									{Name: "Gobo 1 index", DMXFrom: "11/1", PhysicalFrom: (0), PhysicalTo: (0), WheelSlotIndex: (2)},
									{Name: "Gobo 2 index", DMXFrom: "21/1", PhysicalFrom: (0), PhysicalTo: (0), WheelSlotIndex: (3)},
									{Name: "Gobo 3 index", DMXFrom: "31/1", PhysicalFrom: (0), PhysicalTo: (0), WheelSlotIndex: (4)},
									{Name: "Gobo 4 index", DMXFrom: "41/1", PhysicalFrom: (0), PhysicalTo: (0), WheelSlotIndex: (5)},
									{Name: "Gobo 5 index", DMXFrom: "51/1", PhysicalFrom: (0), PhysicalTo: (0), WheelSlotIndex: (6)},
									{Name: "Gobo 1 rot", DMXFrom: "61/1", PhysicalFrom: (0), PhysicalTo: (0), WheelSlotIndex: (2)},
									{Name: "Gobo 2 rot", DMXFrom: "71/1", PhysicalFrom: (0), PhysicalTo: (0), WheelSlotIndex: (3)},
									{Name: "Gobo 3 rot", DMXFrom: "81/1", PhysicalFrom: (0), PhysicalTo: (0), WheelSlotIndex: (4)},
									{Name: "Gobo 4 rot", DMXFrom: "91/1", PhysicalFrom: (0), PhysicalTo: (0), WheelSlotIndex: (5)},
									{Name: "Gobo 5 rot", DMXFrom: "101/1", PhysicalFrom: (0), PhysicalTo: (0), WheelSlotIndex: (6)},
								},
							},
							{
								Attribute: "Gobo1WheelSpin",
								DMXFrom:   "111/1",
								ChannelSets: []XMLTypes.ChannelSet{
									{Name: "CCW", DMXFrom: "111/1", PhysicalFrom: (-60), PhysicalTo: (0)},
									{Name: "Stop", DMXFrom: "183/1", PhysicalFrom: (0), PhysicalTo: (0)},
									{Name: "CW", DMXFrom: "184/1", PhysicalFrom: (0), PhysicalTo: (60)},
								},
							},
						},
					},
				},
			},
			{
				Offset:   XMLTypes.IntList{2},
				Geometry: "Head",
				LogicalChannels: []XMLTypes.LogicalChannel{
					{
						Attribute: "Gobo1Pos",
						ChannelFunctions: []XMLTypes.ChannelFunction{
							{
								Attribute:    "Gobo1Pos",
								DMXFrom:      "0/1",
								PhysicalFrom: (0),
								PhysicalTo:   (540),
								ModeMaster:   "Head_Gobo1",
								ModeFrom:     "0/1",
								ModeTo:       "60/1",
							},
							{
								Attribute:    "Gobo1PosRotate",
								DMXFrom:      "0/1",
								PhysicalFrom: (-150),
								PhysicalTo:   (150),
								ModeMaster:   "Head_Gobo1",
								ModeFrom:     "61/1",
								ModeTo:       "110/1",
								ChannelSets: []XMLTypes.ChannelSet{
									{Name: "CW", DMXFrom: "0/1", PhysicalFrom: (-150), PhysicalTo: (0)},
									{Name: "Stop", DMXFrom: "128/1", PhysicalFrom: (0), PhysicalTo: (0)},
									{Name: "CCW", DMXFrom: "129/1", PhysicalFrom: (0), PhysicalTo: (150)},
								},
							},
						},
					},
				},
			},
		},
		Relations: nil,
	},
}

func TestDMXModeDependencies(t *testing.T) {
	parsingTest(t, dmxModeDependenciesXML, "DMX Mode", dmxModeTest{DMXModes: dmxModeDependenciesStruct})
}

var dmxRelationsXML = `
<DMXModes>
	<DMXMode Name="Mode 1" Geometry="Body">
			<DMXChannels>
					<DMXChannel Offset="1" Default="255/1" Highlight="255/1" Geometry="Pixel">
							<LogicalChannel Attribute="ColorAdd_R" >
									<ChannelFunction Attribute="ColorAdd_R" DMXFrom="0/1" PhysicalFrom="0" PhysicalTo="1" />
							</LogicalChannel>
					</DMXChannel>
					<DMXChannel Offset="2" Default="255/1" Highlight="255/1" Geometry="Pixel">
							<LogicalChannel Attribute="ColorAdd_G" >
									<ChannelFunction Attribute="ColorAdd_G" DMXFrom="0/1" PhysicalFrom="0" PhysicalTo="1" />
							</LogicalChannel>
					</DMXChannel>
					<DMXChannel Offset="3" Default="255/1" Highlight="255/1" Geometry="Pixel">
							<LogicalChannel Attribute="ColorAdd_B" >
									<ChannelFunction Attribute="ColorAdd_B" DMXFrom="0/255" PhysicalFrom="0" PhysicalTo="1" />
							</LogicalChannel>
					</DMXChannel>
					<DMXChannel Highlight="255/1" Geometry="Pixel">
							<LogicalChannel Attribute="Dimmer" >
									<ChannelFunction Attribute="Dimmer" DMXFrom="0/1" PhysicalFrom="0" PhysicalTo="1" />
							</LogicalChannel>
					</DMXChannel>
			</DMXChannels>
			<Relations>
					<Relation Name="Virtual Dimmer R" Master="Pixel_Dimmer" Follower="Pixel_ColorAdd_R.ColorAdd_R.ColorAdd_R 1" Type="Multiply" />
					<Relation Name="Virtual Dimmer G" Master="Pixel_Dimmer" Follower="Pixel_ColorAdd_G.ColorAdd_G.ColorAdd_G 1" Type="Multiply" />
					<Relation Name="Virtual Dimmer B" Master="Pixel_Dimmer" Follower="Pixel_ColorAdd_B.ColorAdd_B.ColorAdd_B 1" Type="Multiply" />
			</Relations>
	</DMXMode>
</DMXModes>
`

var dmxRelationsStruct = []XMLTypes.DMXMode{
	{
		Name:     "Mode 1",
		Geometry: "Body",
		DMXChannels: []XMLTypes.DMXChannel{
			{
				Offset:    XMLTypes.IntList{1},
				Highlight: "255/1",
				Geometry:  "Pixel",
				LogicalChannels: []XMLTypes.LogicalChannel{
					{
						Attribute: "ColorAdd_R",
						ChannelFunctions: []XMLTypes.ChannelFunction{
							{
								Attribute:    "ColorAdd_R",
								DMXFrom:      "0/1",
								PhysicalFrom: 0,
								PhysicalTo:   1,
							},
						},
					},
				},
			},
			{
				Offset:    XMLTypes.IntList{2},
				Highlight: "255/1",
				Geometry:  "Pixel",
				LogicalChannels: []XMLTypes.LogicalChannel{
					{
						Attribute: "ColorAdd_G",
						ChannelFunctions: []XMLTypes.ChannelFunction{
							{
								Attribute:    "ColorAdd_G",
								DMXFrom:      "0/1",
								PhysicalFrom: 0,
								PhysicalTo:   1,
							},
						},
					},
				},
			},
			{
				Offset:    XMLTypes.IntList{3},
				Highlight: "255/1",
				Geometry:  "Pixel",
				LogicalChannels: []XMLTypes.LogicalChannel{
					{
						Attribute: "ColorAdd_B",
						ChannelFunctions: []XMLTypes.ChannelFunction{
							{
								Attribute:    "ColorAdd_B",
								DMXFrom:      "0/255",
								PhysicalFrom: 0,
								PhysicalTo:   1,
							},
						},
					},
				},
			},
			{
				Offset:    nil,
				Highlight: "255/1",
				Geometry:  "Pixel",
				LogicalChannels: []XMLTypes.LogicalChannel{
					{
						Attribute: "Dimmer",
						ChannelFunctions: []XMLTypes.ChannelFunction{
							{
								Attribute:    "Dimmer",
								DMXFrom:      "0/1",
								PhysicalFrom: 0,
								PhysicalTo:   1,
							},
						},
					},
				},
			},
		},
		Relations: []XMLTypes.Relation{
			{
				Name:     "Virtual Dimmer R",
				Master:   "Pixel_Dimmer",
				Follower: "Pixel_ColorAdd_R.ColorAdd_R.ColorAdd_R 1",
				Type:     "Multiply",
			},
			{
				Name:     "Virtual Dimmer G",
				Master:   "Pixel_Dimmer",
				Follower: "Pixel_ColorAdd_G.ColorAdd_G.ColorAdd_G 1",
				Type:     "Multiply",
			},
			{
				Name:     "Virtual Dimmer B",
				Master:   "Pixel_Dimmer",
				Follower: "Pixel_ColorAdd_B.ColorAdd_B.ColorAdd_B 1",
				Type:     "Multiply",
			},
		},
	},
}

func TestDMXRelationsDependencies(t *testing.T) {
	parsingTest(t, dmxRelationsXML, "DMX Relations", dmxModeTest{DMXModes: dmxRelationsStruct})
}

var dmxBreaksXML = `
<DMXMode Name="Mode 1" Geometry="Body">
	<DMXChannels>
			<DMXChannel Offset="1" Highlight="255/1" Geometry="Body">
					<LogicalChannel Attribute="Dimmer" Master="Grand">
							<ChannelFunction Attribute="Dimmer" DMXFrom="0/255" PhysicalFrom="0" PhysicalTo="1" RealFade="0" />
					</LogicalChannel>
			</DMXChannel>
			<DMXChannel DMXBreak="2" Offset="1" Highlight="0/1" Geometry="Body">
					<LogicalChannel Attribute="Color1" >
							<ChannelFunction Attribute="Color1" DMXFrom="0/1" Wheel="ColorScroller">
									<ChannelSet Name="Gel1" DMXFrom="0/1" PhysicalFrom="0" PhysicalTo="0.5" WheelSlotIndex="1" />
									<ChannelSet Name="Gel2" DMXFrom="51/1" PhysicalFrom="-0.5" PhysicalTo="0.5" WheelSlotIndex="2" />
									<ChannelSet Name="Gel3" DMXFrom="102/1" PhysicalFrom="-0.5" PhysicalTo="0.5" WheelSlotIndex="3" />
									<ChannelSet Name="Gel4" DMXFrom="153/1" PhysicalFrom="-0.5" PhysicalTo="0.5" WheelSlotIndex="4" />
									<ChannelSet Name="Gel5" DMXFrom="204/1" PhysicalFrom="-0.5" PhysicalTo="0" WheelSlotIndex="5" />
							</ChannelFunction>
					</LogicalChannel>
			</DMXChannel>
	</DMXChannels>
	<Relations />
</DMXMode>
`

var dmxBreaksStruct = []XMLTypes.DMXMode{
	{
		Name:        "Mode 1",
		Description: "",
		Geometry:    "Body",
		DMXChannels: []XMLTypes.DMXChannel{
			{
				DMXBreak:        0,
				Offset:          XMLTypes.IntList{1},
				InitialFunction: "",
				Highlight:       "255/1",
				Geometry:        "Body",
				LogicalChannels: []XMLTypes.LogicalChannel{
					{
						Attribute:          "Dimmer",
						Snap:               "",
						Master:             "Grand",
						MibFade:            0,
						DMXChangeTimeLimit: 0,
						ChannelFunctions: []XMLTypes.ChannelFunction{
							{
								Name:              "",
								Attribute:         "Dimmer",
								OriginalAttribute: "",
								DMXFrom:           "0/255",
								Default:           "",
								PhysicalFrom:      0,
								PhysicalTo:        1,
								RealFade:          0,
								RealAcceleration:  0,
								Wheel:             "",
								Emitter:           "",
								Filter:            "",
								ColorSpace:        "",
								Gamut:             "",
								ModeMaster:        "",
								ModeFrom:          "",
								ModeTo:            "",
								DMXProfile:        "",
								Min:               0,
								Max:               0,
								CustomName:        "",
								ChannelSets:       nil,
								SubChannelSets:    nil,
							},
						},
					},
				},
			},
			{
				DMXBreak:        2,
				Offset:          XMLTypes.IntList{1},
				InitialFunction: "",
				Highlight:       "0/1",
				Geometry:        "Body",
				LogicalChannels: []XMLTypes.LogicalChannel{
					{
						Attribute:          "Color1",
						Snap:               "",
						Master:             "",
						MibFade:            0,
						DMXChangeTimeLimit: 0,
						ChannelFunctions: []XMLTypes.ChannelFunction{
							{
								Name:              "",
								Attribute:         "Color1",
								OriginalAttribute: "",
								DMXFrom:           "0/1",
								Default:           "",
								PhysicalFrom:      0,
								PhysicalTo:        0, // top-level ChannelFunction physical range is set via ChannelSets below
								RealFade:          0,
								RealAcceleration:  0,
								Wheel:             "ColorScroller",
								Emitter:           "",
								Filter:            "",
								ColorSpace:        "",
								Gamut:             "",
								ModeMaster:        "",
								ModeFrom:          "",
								ModeTo:            "",
								DMXProfile:        "",
								Min:               0,
								Max:               0,
								CustomName:        "",
								ChannelSets: []XMLTypes.ChannelSet{
									{
										Name:           "Gel1",
										DMXFrom:        "0/1",
										PhysicalFrom:   0,
										PhysicalTo:     0.5,
										WheelSlotIndex: 1,
									},
									{
										Name:           "Gel2",
										DMXFrom:        "51/1",
										PhysicalFrom:   -0.5,
										PhysicalTo:     0.5,
										WheelSlotIndex: 2,
									},
									{
										Name:           "Gel3",
										DMXFrom:        "102/1",
										PhysicalFrom:   -0.5,
										PhysicalTo:     0.5,
										WheelSlotIndex: 3,
									},
									{
										Name:           "Gel4",
										DMXFrom:        "153/1",
										PhysicalFrom:   -0.5,
										PhysicalTo:     0.5,
										WheelSlotIndex: 4,
									},
									{
										Name:           "Gel5",
										DMXFrom:        "204/1",
										PhysicalFrom:   -0.5,
										PhysicalTo:     0,
										WheelSlotIndex: 5,
									},
								},
								SubChannelSets: nil,
							},
						},
					},
				},
			},
		},
		Relations: nil,
		FTMacros:  nil,
	},
}

func TestDMXBreaks(t *testing.T) {
	parsingTest(t, dmxBreaksXML, "DMX Breaks", dmxBreaksStruct)
}

var dmxChannelCollectXML = `
<DMXChannels>
    <DMXChannel Offset="1" Highlight="255/1" Geometry="Head">
        <LogicalChannel Attribute="Dimmer" >
            <ChannelFunction Attribute="Dimmer" DMXFrom="0/1" PhysicalFrom="0" PhysicalTo="1" >
                <ChannelSet Name="Closed" DMXFrom="0/1" />
                <ChannelSet DMXFrom="1/1" />
                <ChannelSet Name="10%" DMXFrom="25/1" />
                <ChannelSet DMXFrom="26/1" />
                <ChannelSet Name="20%" DMXFrom="51/1" />
                <ChannelSet DMXFrom="52/1" />
                <ChannelSet Name="30%" DMXFrom="76/1" />
                <ChannelSet DMXFrom="77/1" />
                <ChannelSet Name="40%" DMXFrom="102/1" />
                <ChannelSet DMXFrom="103/1" />
                <ChannelSet Name="50%" DMXFrom="128/1" />
                <ChannelSet DMXFrom="129/1" />
                <ChannelSet Name="60%" DMXFrom="153/1" />
                <ChannelSet DMXFrom="154/1" />
                <ChannelSet Name="70%" DMXFrom="179/1" />
                <ChannelSet DMXFrom="180/1" />
                <ChannelSet Name="80%" DMXFrom="204/1" />
                <ChannelSet DMXFrom="205/1" />
                <ChannelSet Name="90%" DMXFrom="230/1" />
                <ChannelSet DMXFrom="231/1" />
                <ChannelSet Name="Open" DMXFrom="255/1" />
            </ChannelFunction>
        </LogicalChannel>
    </DMXChannel>
    <DMXChannel Offset="2,3" Default="32768/2" Geometry="Yoke">
        <LogicalChannel Attribute="Pan">
            <ChannelFunction Attribute="Pan" DMXFrom="0/1" PhysicalFrom="-270" PhysicalTo=" 270" >
                <ChannelSet DMXFrom="0/1" />
                <ChannelSet Name="Home" DMXFrom="32768/2" />
                <ChannelSet DMXFrom="32769/2" />
            </ChannelFunction>
        </LogicalChannel>
    </DMXChannel>
    <DMXChannel Offset="4,5" Default="32768/2" Geometry="Head">
        <LogicalChannel Attribute="Tilt">
            <ChannelFunction Attribute="Tilt" DMXFrom="0/1" PhysicalFrom="-130" PhysicalTo=" 130" >
                <ChannelSet DMXFrom="0/1" />
                <ChannelSet Name="Home" DMXFrom="32768/2" />
                <ChannelSet DMXFrom="32769/2" />
            </ChannelFunction>
        </LogicalChannel>
    </DMXChannel>
    <DMXChannel Offset="6" Highlight="0/1" Geometry="Head">
        <LogicalChannel Attribute="Gobo1">
            <ChannelFunction Attribute="Gobo1" DMXFrom="0/1" Wheel="GoboWheel">
                <ChannelSet Name="Open" DMXFrom="0/1" PhysicalFrom="0" PhysicalTo="0" WheelSlotIndex="1" />
                <ChannelSet Name="Gobo 1" DMXFrom="42/1" PhysicalFrom="0" PhysicalTo="0" WheelSlotIndex="2" />
                <ChannelSet Name="Gobo 2" DMXFrom="85/1" PhysicalFrom="0" PhysicalTo="0" WheelSlotIndex="3" />
                <ChannelSet Name="Gobo 3" DMXFrom="127/1" PhysicalFrom="0" PhysicalTo="0" WheelSlotIndex="4" />
                <ChannelSet Name="Gobo 4" DMXFrom="169/1" PhysicalFrom="0" PhysicalTo="0" WheelSlotIndex="5" />
                <ChannelSet Name="Gobo 5" DMXFrom="212/1" PhysicalFrom="0" PhysicalTo="0" WheelSlotIndex="6" />
            </ChannelFunction>
        </LogicalChannel>
    </DMXChannel>
    <DMXChannel Offset="7" Highlight="0/1" Geometry="Head">
        <LogicalChannel Attribute="Color1">
            <ChannelFunction Attribute="Color1" DMXFrom="0/1" Wheel="ColorWheel">
                <ChannelSet Name="Open" DMXFrom="0/1" PhysicalFrom="0" PhysicalTo="0.5" WheelSlotIndex="1" />
                <ChannelSet Name="Red" DMXFrom="36/1" PhysicalFrom="-0.5" PhysicalTo="0.5" WheelSlotIndex="2" />
                <ChannelSet Name="Green" DMXFrom="73/1" PhysicalFrom="-0.5" PhysicalTo="0.5" WheelSlotIndex="3" />
                <ChannelSet Name="Blue" DMXFrom="109/1" PhysicalFrom="-0.5" PhysicalTo="0.5" WheelSlotIndex="4" />
                <ChannelSet Name="Cyan" DMXFrom="146/1" PhysicalFrom="-0.5" PhysicalTo="0.5" WheelSlotIndex="5" />
                <ChannelSet Name="Magenta" DMXFrom="182/1" PhysicalFrom="-0.5" PhysicalTo="0.5" WheelSlotIndex="6" />
                <ChannelSet Name="Yellow" DMXFrom="219/1" PhysicalFrom="-0.5" PhysicalTo="0" WheelSlotIndex="7" />
            </ChannelFunction>
        </LogicalChannel>
    </DMXChannel>
</DMXChannels>
`

var dmxChannelCollectStruct = []XMLTypes.DMXChannel{
	{
		DMXBreak:        0,
		Offset:          XMLTypes.IntList{1},
		InitialFunction: "",
		Highlight:       "255/1",
		Geometry:        "Head",
		LogicalChannels: []XMLTypes.LogicalChannel{
			{
				Attribute:          "Dimmer",
				Snap:               "",
				Master:             "",
				MibFade:            0,
				DMXChangeTimeLimit: 0,
				ChannelFunctions: []XMLTypes.ChannelFunction{
					{
						Name:              "",
						Attribute:         "Dimmer",
						OriginalAttribute: "",
						DMXFrom:           "0/1",
						Default:           "",
						PhysicalFrom:      0,
						PhysicalTo:        1,
						RealFade:          0,
						RealAcceleration:  0,
						Wheel:             "",
						Emitter:           "",
						Filter:            "",
						ColorSpace:        "",
						Gamut:             "",
						ModeMaster:        "",
						ModeFrom:          "",
						ModeTo:            "",
						DMXProfile:        "",
						Min:               0,
						Max:               0,
						CustomName:        "",
						ChannelSets: []XMLTypes.ChannelSet{
							{Name: "Closed", DMXFrom: "0/1"},
							{Name: "", DMXFrom: "1/1"},
							{Name: "10%", DMXFrom: "25/1"},
							{Name: "", DMXFrom: "26/1"},
							{Name: "20%", DMXFrom: "51/1"},
							{Name: "", DMXFrom: "52/1"},
							{Name: "30%", DMXFrom: "76/1"},
							{Name: "", DMXFrom: "77/1"},
							{Name: "40%", DMXFrom: "102/1"},
							{Name: "", DMXFrom: "103/1"},
							{Name: "50%", DMXFrom: "128/1"},
							{Name: "", DMXFrom: "129/1"},
							{Name: "60%", DMXFrom: "153/1"},
							{Name: "", DMXFrom: "154/1"},
							{Name: "70%", DMXFrom: "179/1"},
							{Name: "", DMXFrom: "180/1"},
							{Name: "80%", DMXFrom: "204/1"},
							{Name: "", DMXFrom: "205/1"},
							{Name: "90%", DMXFrom: "230/1"},
							{Name: "", DMXFrom: "231/1"},
							{Name: "Open", DMXFrom: "255/1"},
						},
						SubChannelSets: nil,
					},
				},
			},
		},
	},
	{
		DMXBreak:        0,
		Offset:          XMLTypes.IntList{2, 3},
		InitialFunction: "",
		Highlight:       "",
		Geometry:        "Yoke",
		LogicalChannels: []XMLTypes.LogicalChannel{
			{
				Attribute:          "Pan",
				Snap:               "",
				Master:             "",
				MibFade:            0,
				DMXChangeTimeLimit: 0,
				ChannelFunctions: []XMLTypes.ChannelFunction{
					{
						Name:              "",
						Attribute:         "Pan",
						OriginalAttribute: "",
						DMXFrom:           "0/1",
						// Default:           "32768/2",
						// note: trim spaces in " 270" -> 270
						PhysicalFrom:     -270,
						PhysicalTo:       270,
						RealFade:         0,
						RealAcceleration: 0,
						Wheel:            "",
						Emitter:          "",
						Filter:           "",
						ColorSpace:       "",
						Gamut:            "",
						ModeMaster:       "",
						ModeFrom:         "",
						ModeTo:           "",
						DMXProfile:       "",
						Min:              0,
						Max:              0,
						CustomName:       "",
						ChannelSets: []XMLTypes.ChannelSet{
							{Name: "", DMXFrom: "0/1"},
							{Name: "Home", DMXFrom: "32768/2"},
							{Name: "", DMXFrom: "32769/2"},
						},
						SubChannelSets: nil,
					},
				},
			},
		},
	},
	{
		DMXBreak:        0,
		Offset:          XMLTypes.IntList{4, 5},
		InitialFunction: "",
		Highlight:       "",
		Geometry:        "Head",
		LogicalChannels: []XMLTypes.LogicalChannel{
			{
				Attribute:          "Tilt",
				Snap:               "",
				Master:             "",
				MibFade:            0,
				DMXChangeTimeLimit: 0,
				ChannelFunctions: []XMLTypes.ChannelFunction{
					{
						Name:              "",
						Attribute:         "Tilt",
						OriginalAttribute: "",
						DMXFrom:           "0/1",
						// Default:           "32768/2",
						PhysicalFrom:     -130,
						PhysicalTo:       130,
						RealFade:         0,
						RealAcceleration: 0,
						Wheel:            "",
						Emitter:          "",
						Filter:           "",
						ColorSpace:       "",
						Gamut:            "",
						ModeMaster:       "",
						ModeFrom:         "",
						ModeTo:           "",
						DMXProfile:       "",
						Min:              0,
						Max:              0,
						CustomName:       "",
						ChannelSets: []XMLTypes.ChannelSet{
							{Name: "", DMXFrom: "0/1"},
							{Name: "Home", DMXFrom: "32768/2"},
							{Name: "", DMXFrom: "32769/2"},
						},
						SubChannelSets: nil,
					},
				},
			},
		},
	},
	{
		DMXBreak:        0,
		Offset:          XMLTypes.IntList{6},
		InitialFunction: "",
		Highlight:       "0/1",
		Geometry:        "Head",
		LogicalChannels: []XMLTypes.LogicalChannel{
			{
				Attribute:          "Gobo1",
				Snap:               "",
				Master:             "",
				MibFade:            0,
				DMXChangeTimeLimit: 0,
				ChannelFunctions: []XMLTypes.ChannelFunction{
					{
						Name:              "",
						Attribute:         "Gobo1",
						OriginalAttribute: "",
						DMXFrom:           "0/1",
						Default:           "",
						PhysicalFrom:      0,
						PhysicalTo:        0,
						RealFade:          0,
						RealAcceleration:  0,
						Wheel:             "GoboWheel",
						Emitter:           "",
						Filter:            "",
						ColorSpace:        "",
						Gamut:             "",
						ModeMaster:        "",
						ModeFrom:          "",
						ModeTo:            "",
						DMXProfile:        "",
						Min:               0,
						Max:               0,
						CustomName:        "",
						ChannelSets: []XMLTypes.ChannelSet{
							{Name: "Open", DMXFrom: "0/1", PhysicalFrom: 0, PhysicalTo: 0, WheelSlotIndex: 1},
							{Name: "Gobo 1", DMXFrom: "42/1", PhysicalFrom: 0, PhysicalTo: 0, WheelSlotIndex: 2},
							{Name: "Gobo 2", DMXFrom: "85/1", PhysicalFrom: 0, PhysicalTo: 0, WheelSlotIndex: 3},
							{Name: "Gobo 3", DMXFrom: "127/1", PhysicalFrom: 0, PhysicalTo: 0, WheelSlotIndex: 4},
							{Name: "Gobo 4", DMXFrom: "169/1", PhysicalFrom: 0, PhysicalTo: 0, WheelSlotIndex: 5},
							{Name: "Gobo 5", DMXFrom: "212/1", PhysicalFrom: 0, PhysicalTo: 0, WheelSlotIndex: 6},
						},
						SubChannelSets: nil,
					},
				},
			},
		},
	},
	{
		DMXBreak:        0,
		Offset:          XMLTypes.IntList{7},
		InitialFunction: "",
		Highlight:       "0/1",
		Geometry:        "Head",
		LogicalChannels: []XMLTypes.LogicalChannel{
			{
				Attribute:          "Color1",
				Snap:               "",
				Master:             "",
				MibFade:            0,
				DMXChangeTimeLimit: 0,
				ChannelFunctions: []XMLTypes.ChannelFunction{
					{
						Name:              "",
						Attribute:         "Color1",
						OriginalAttribute: "",
						DMXFrom:           "0/1",
						Default:           "",
						PhysicalFrom:      0,
						PhysicalTo:        0,
						RealFade:          0,
						RealAcceleration:  0,
						Wheel:             "ColorWheel",
						Emitter:           "",
						Filter:            "",
						ColorSpace:        "",
						Gamut:             "",
						ModeMaster:        "",
						ModeFrom:          "",
						ModeTo:            "",
						DMXProfile:        "",
						Min:               0,
						Max:               0,
						CustomName:        "",
						ChannelSets: []XMLTypes.ChannelSet{
							{Name: "Open", DMXFrom: "0/1", PhysicalFrom: 0, PhysicalTo: 0.5, WheelSlotIndex: 1},
							{Name: "Red", DMXFrom: "36/1", PhysicalFrom: -0.5, PhysicalTo: 0.5, WheelSlotIndex: 2},
							{Name: "Green", DMXFrom: "73/1", PhysicalFrom: -0.5, PhysicalTo: 0.5, WheelSlotIndex: 3},
							{Name: "Blue", DMXFrom: "109/1", PhysicalFrom: -0.5, PhysicalTo: 0.5, WheelSlotIndex: 4},
							{Name: "Cyan", DMXFrom: "146/1", PhysicalFrom: -0.5, PhysicalTo: 0.5, WheelSlotIndex: 5},
							{Name: "Magenta", DMXFrom: "182/1", PhysicalFrom: -0.5, PhysicalTo: 0.5, WheelSlotIndex: 6},
							{Name: "Yellow", DMXFrom: "219/1", PhysicalFrom: -0.5, PhysicalTo: 0, WheelSlotIndex: 7},
						},
						SubChannelSets: nil,
					},
				},
			},
		},
	},
}

type dmxChannelTest struct {
	DMXChannels []XMLTypes.DMXChannel `xml:"DMXChannel"`
}

func TestDMXChannelCollect(t *testing.T) {
	parsingTest(t, dmxChannelCollectXML, "DMX Channel Collect", dmxChannelTest{DMXChannels: dmxChannelCollectStruct})
}
