package Types

type PhysicalDescription struct {
	Emitters              []Emitter
	Filters               []Filter
	ColorSpace            ColorSpace
	AdditionalColorSpaces []ColorSpace
	Gamuts                []Gamut
	DMXProfiles           []DMXProfile
	CRIs                  []CRIGroup
	Connectors            []Connector
	Properties            Properties
}

type Emitter struct {
	Name               string
	Color              ColorCIE
	DominantWaveLength float32
	DiodePart          string
	Measurements       []Measurement
}

type Measurement struct {
	Phyiscal          float32
	LuminousIntensity float32
	Transmission      float32
	InterpolationTo   string
	MeasurementPoints *[]MeasurementPoint
}

type MeasurementPoint struct {
	WaveLength float32
	Energy     float32
}

type Filter struct {
	Name         string
	Color        ColorCIE
	Measurements []Measurement
}

type ColorSpace struct {
	Name       string
	Mode       string //enum + TODO: default data (see https://www.gdtf.eu/gdtf/file-spec/physical-descriptions/#table-20-color-space-attributes)
	Red        ColorCIE
	Green      ColorCIE
	Blue       ColorCIE
	WhitePoint ColorCIE
}

type Gamut struct {
	Name   string
	Points []ColorCIE
}

type DMXProfile struct {
	Name   string
	Points []Point
}

type Point struct {
	DMXPercentage float32
	CFC0          float32
	CFC1          float32
	CFC2          float32
	CFC3          float32
}

type CRIGroup struct {
	ColorTemperature float32
	CRIs             *[]CRI
}

type CRI struct {
	CES string
	CRI uint
}

type Connector struct {
	Name     string
	Type     string
	DMXBreak uint    // obsolete (see https://www.gdtf.eu/gdtf/file-spec/physical-descriptions/#table-27-connector-attributes)
	Gender   int     // obsolete
	Length   float32 // obsolete
}

type Properties struct {
	OperatingTemperature *OperatingTemperature
	Weight               *Weight
	PowerConsumption     any
	LegHeight            *LegHeight
}

type OperatingTemperature struct {
	Low  float32
	High float32
}

type Weight struct {
	Value float32
}

type LegHeight struct {
	Value float32
}
