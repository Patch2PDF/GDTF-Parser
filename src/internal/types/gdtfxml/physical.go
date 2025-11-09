package XMLTypes

type PhysicalDescription struct {
	Emitters              []Emitter    `xml:"Emitters>Emitter"`
	Filters               []Filter     `xml:"Filters>Filter"`
	ColorSpace            ColorSpace   `xml:"ColorSpace"`
	AdditionalColorSpaces []ColorSpace `xml:"AdditionalColorSpaces>ColorSpace"`
	Gamuts                []Gamut      `xml:"Gamuts>Gamut"`
	DMXProfiles           []DMXProfile `xml:"DMXProfiles>DMXProfile"`
	CRIs                  []CRIGroup   `xml:"CRIs>CRIGroup"`
	Connectors            []Connector  `xml:"Connectors>Connector"`
	Properties            Properties   `xml:"Properties"`
}

type Emitter struct {
	Name               string        `xml:",attr"`
	Color              ColorCIE      `xml:",attr"`
	DominantWaveLength float32       `xml:",attr"`
	DiodePart          string        `xml:",attr"`
	Measurements       []Measurement `xml:"Measurement"`
}

type Measurement struct {
	Phyiscal          float32             `xml:",attr"`
	LuminousIntensity float32             `xml:",attr"`
	Transmission      float32             `xml:",attr"`
	InterpolationTo   string              `xml:",attr"`
	MeasurementPoints *[]MeasurementPoint `xml:"MeasurementPoint"`
}

type MeasurementPoint struct {
	WaveLength float32 `xml:",attr"`
	Energy     float32 `xml:",attr"`
}

type Filter struct {
	Name         string        `xml:",attr"`
	Color        ColorCIE      `xml:",attr"`
	Measurements []Measurement `xml:"Measurement"`
}

type ColorSpace struct {
	Name       string   `xml:",attr"`
	Mode       string   `xml:",attr"` //TODO: enum + default data (see https://www.gdtf.eu/gdtf/file-spec/physical-descriptions/#table-20-color-space-attributes)
	Red        ColorCIE `xml:",attr"`
	Green      ColorCIE `xml:",attr"`
	Blue       ColorCIE `xml:",attr"`
	WhitePoint ColorCIE `xml:",attr"`
}

type Gamut struct {
	Name   string     `xml:",attr"`
	Points []ColorCIE `xml:",attr"`
}

type DMXProfile struct {
	Name   string  `xml:",attr"`
	Points []Point `xml:"Point"`
}

type Point struct {
	DMXPercentage float32 `xml:",attr"`
	CFC0          float32 `xml:",attr"`
	CFC1          float32 `xml:",attr"`
	CFC2          float32 `xml:",attr"`
	CFC3          float32 `xml:",attr"`
}

type CRIGroup struct {
	ColorTemperature float32 `xml:",attr"`
	CRIs             *[]CRI  `xml:"CRI"`
}

type CRI struct {
	CES string `xml:",attr"`
	CRI uint   `xml:",attr"`
}

type Connector struct {
	Name     string  `xml:",attr"`
	Type     string  `xml:",attr"`
	DMXBreak uint    `xml:",attr"` // obsolete (see https://www.gdtf.eu/gdtf/file-spec/physical-descriptions/#table-27-connector-attributes)
	Gender   int     `xml:",attr"` // obsolete
	Length   float32 `xml:",attr"` // obsolete
}

type Properties struct {
	OperatingTemperature *OperatingTemperature `xml:"OperatingTemperature"`
	Weight               *Weight               `xml:"Weight"`
	PowerConsumption     any                   `xml:"PowerConsumption"`
	LegHeight            *LegHeight            `xml:"LegHeight"`
}

type OperatingTemperature struct {
	Low  float32 `xml:",attr"`
	High float32 `xml:",attr"`
}

type Weight struct {
	Value float32 `xml:",attr"`
}

type LegHeight struct {
	Value float32 `xml:",attr"`
}
