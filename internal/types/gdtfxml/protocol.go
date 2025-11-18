package XMLTypes

import Types "github.com/Patch2PDF/GDTF-Parser/pkg/types"

type Protocol struct {
	RDM              []RDM              `xml:"FTRDM"`
	ArtNet           []ArtNet           `xml:"Art-Net"`
	SACN             []SACN             `xml:"sACN"`
	PosiStageNet     []PosiStageNet     `xml:"PosiStageNet"`
	OpenSoundControl []OpenSoundControl `xml:"OpenSoundControl"`
	CITP             []CITP             `xml:"CITP"`
}

func (protocol Protocol) Parse() Types.Protocol {
	return Types.Protocol{
		RDM:    ParseList(&protocol.RDM),
		ArtNet: ParseList(&protocol.ArtNet),
		SACN:   ParseList(&protocol.SACN),
		// PosiStageNet: [],
		// OpenSoundControl: [],
		// CITP: [],
	}
}

type RDM struct {
	ManufacturerID   Hex                  `xml:",attr"`
	DeviceModelID    Hex                  `xml:",attr"`
	SoftwareVersions []RDMSoftwareVersion `xml:"SoftwareVersionID"`
}

func (protcol RDM) Parse() Types.RDM {
	return Types.RDM{
		ManufacturerID:   Types.Hex(protcol.ManufacturerID),
		DeviceModelID:    Types.Hex(protcol.DeviceModelID),
		SoftwareVersions: ParseList(&protcol.SoftwareVersions),
	}
}

type RDMSoftwareVersion struct {
	Value            Hex                 `xml:",attr"`
	DMXPersonalities []RDMDMXPersonality `xml:"DMXPersonality"`
}

func (protcol RDMSoftwareVersion) Parse() Types.RDMSoftwareVersion {
	return Types.RDMSoftwareVersion{
		Value:            Types.Hex(protcol.Value),
		DMXPersonalities: ParseList(&protcol.DMXPersonalities),
	}
}

type RDMDMXPersonality struct {
	Value   Hex              `xml:",attr"`
	DMXMode XMLNodeReference `xml:",attr"` // reference to DMX Mode
}

func (protcol RDMDMXPersonality) Parse() Types.RDMDMXPersonality {
	return Types.RDMDMXPersonality{
		Value: Types.Hex(protcol.Value),
		DMXMode: Types.NodeReference[Types.DMXMode]{
			String: protcol.DMXMode,
		},
	}
}

type Map struct {
	Key   uint `xml:",attr"` // Art-Net value
	Value uint `xml:",attr"` // DMX value
}

func (object Map) Parse() Types.Map {
	return Types.Map(object)
}

// key: artnet value, value: dmx value
type ArtNet struct {
	Maps []Map `xml:"Map"`
}

func (protcol ArtNet) Parse() Types.ArtNet {
	return Types.ArtNet{
		Maps: ParseList(&protcol.Maps),
	}
}

type SACN struct {
	Maps []Map `xml:"Map"`
}

func (protcol SACN) Parse() Types.SACN {
	return Types.SACN{
		Maps: ParseList(&protcol.Maps),
	}
}

type PosiStageNet struct {
	// not defined in standard yet
}

type OpenSoundControl struct {
	// not defined in standard yet
}

type CITP struct {
	// not defined in standard yet
}
