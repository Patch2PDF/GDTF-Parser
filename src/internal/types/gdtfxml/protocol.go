package XMLTypes

import Types "github.com/Patch2PDF/GDTF-Parser/types"

type Protocol struct {
	RDM              []RDM              `xml:"FTRDM"`
	ArtNet           []ArtNet           `xml:"Art-Net"`
	SACN             []SACN             `xml:"sACN"`
	PosiStageNet     []PosiStageNet     `xml:"PosiStageNet"`
	OpenSoundControl []OpenSoundControl `xml:"OpenSoundControl"`
	CITP             []CITP             `xml:"CITP"`
}

// TODO:
func (protocol Protocol) Parse() Types.Protocol {
	return Types.Protocol{}
}

type RDM struct {
	ManufacturerID   Hex                  `xml:",attr"`
	DeviceModelID    Hex                  `xml:",attr"`
	SoftwareVersions []RDMSoftwareVersion `xml:"SoftwareVersionID"`
}

type RDMSoftwareVersion struct {
	Value            Hex                 `xml:",attr"`
	DMXPersonalities []RDMDMXPersonality `xml:"DMXPersonality"`
}

type RDMDMXPersonality struct {
	Value   Hex              `xml:",attr"`
	DMXMode XMLNodeReference `xml:",attr"` // reference to DMX Mode
}

type Map struct {
	Key   uint `xml:",attr"` // Art-Net value
	Value uint `xml:",attr"` // DMX value
}

// key: artnet value, value: dmx value
type ArtNet struct {
	Maps []Map `xml:"Map"`
}

type SACN struct {
	Maps []Map `xml:"Map"`
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
