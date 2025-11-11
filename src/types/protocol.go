package Types

type Protocol struct {
	RDM              []RDM
	ArtNet           []ArtNet
	SACN             []SACN
	PosiStageNet     []PosiStageNet
	OpenSoundControl []OpenSoundControl
	CITP             []CITP
}

type RDM struct {
	ManufacturerID   Hex
	DeviceModelID    Hex
	SoftwareVersions []RDMSoftwareVersion
}

type RDMSoftwareVersion struct {
	Value            Hex
	DMXPersonalities []RDMDMXPersonality
}

type RDMDMXPersonality struct {
	Value   Hex
	DMXMode NodeReference[DMXMode] // reference to DMX Mode
}

type Map struct {
	Key   uint // Art-Net value
	Value uint // DMX value
}

// key: artnet value, value: dmx value
type ArtNet struct {
	Maps []Map
}

type SACN struct {
	Maps []Map
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
