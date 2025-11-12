package XMLTypes

import Types "github.com/Patch2PDF/GDTF-Parser/types"

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

func (physicalDescription PhysicalDescription) Parse() Types.PhysicalDescription {
	return Types.PhysicalDescription{
		Emitters:              ParseList(&physicalDescription.Emitters),
		Filters:               ParseList(&physicalDescription.Filters),
		ColorSpace:            physicalDescription.ColorSpace.Parse(),
		AdditionalColorSpaces: ParseList(&physicalDescription.AdditionalColorSpaces),
		Gamuts:                ParseList(&physicalDescription.Gamuts),
		DMXProfiles:           ParseList(&physicalDescription.DMXProfiles),
		CRIs:                  ParseList(&physicalDescription.CRIs),
		Connectors:            ParseList(&physicalDescription.Connectors),
		Properties:            physicalDescription.Properties.Parse(),
	}
}

type Emitter struct {
	Name               string        `xml:",attr"`
	Color              ColorCIE      `xml:",attr"`
	DominantWaveLength float32       `xml:",attr"`
	DiodePart          string        `xml:",attr"`
	Measurements       []Measurement `xml:"Measurement"`
}

func (object Emitter) Parse() Types.Emitter {
	return Types.Emitter{
		Name:               object.Name,
		Color:              object.Color.Parse(),
		DominantWaveLength: object.DominantWaveLength,
		DiodePart:          object.DiodePart,
		Measurements:       ParseList(&object.Measurements),
	}
}

type Measurement struct {
	Phyiscal          float32             `xml:",attr"`
	LuminousIntensity float32             `xml:",attr"`
	Transmission      float32             `xml:",attr"`
	InterpolationTo   string              `xml:",attr"`
	MeasurementPoints *[]MeasurementPoint `xml:"MeasurementPoint"`
}

func (object Measurement) Parse() Types.Measurement {
	return Types.Measurement{
		Phyiscal:          object.Phyiscal,
		LuminousIntensity: object.LuminousIntensity,
		Transmission:      object.Transmission,
		InterpolationTo:   object.InterpolationTo,
		MeasurementPoints: ParseList(object.MeasurementPoints),
	}
}

type MeasurementPoint struct {
	WaveLength float32 `xml:",attr"`
	Energy     float32 `xml:",attr"`
}

func (object MeasurementPoint) Parse() Types.MeasurementPoint {
	return Types.MeasurementPoint(object)
}

type Filter struct {
	Name         string        `xml:",attr"`
	Color        ColorCIE      `xml:",attr"`
	Measurements []Measurement `xml:"Measurement"`
}

func (object Filter) Parse() Types.Filter {
	return Types.Filter{
		Name:         object.Name,
		Color:        Types.ColorCIE(object.Color),
		Measurements: ParseList(&object.Measurements),
	}
}

type ColorSpace struct {
	Name       string   `xml:",attr"`
	Mode       string   `xml:",attr"` //enum + TODO: default data (see https://www.gdtf.eu/gdtf/file-spec/physical-descriptions/#table-20-color-space-attributes)
	Red        ColorCIE `xml:",attr"`
	Green      ColorCIE `xml:",attr"`
	Blue       ColorCIE `xml:",attr"`
	WhitePoint ColorCIE `xml:",attr"`
}

func (object ColorSpace) Parse() Types.ColorSpace {
	return Types.ColorSpace{
		Name:       object.Name,
		Mode:       object.Mode,
		Red:        Types.ColorCIE(object.Red),
		Green:      Types.ColorCIE(object.Green),
		Blue:       Types.ColorCIE(object.Blue),
		WhitePoint: Types.ColorCIE(object.WhitePoint),
	}
}

type Gamut struct {
	Name   string     `xml:",attr"`
	Points []ColorCIE `xml:",attr"`
}

func (object Gamut) Parse() Types.Gamut {
	return Types.Gamut{
		Name:   object.Name,
		Points: ParseList(&object.Points),
	}
}

type DMXProfile struct {
	Name   string  `xml:",attr"`
	Points []Point `xml:"Point"`
}

func (object DMXProfile) Parse() Types.DMXProfile {
	return Types.DMXProfile{
		Name:   object.Name,
		Points: ParseList(&object.Points),
	}
}

type Point struct {
	DMXPercentage float32 `xml:",attr"`
	CFC0          float32 `xml:",attr"`
	CFC1          float32 `xml:",attr"`
	CFC2          float32 `xml:",attr"`
	CFC3          float32 `xml:",attr"`
}

func (object Point) Parse() Types.Point {
	return Types.Point(object)
}

type CRIGroup struct {
	ColorTemperature float32 `xml:",attr"`
	CRIs             *[]CRI  `xml:"CRI"`
}

func (object CRIGroup) Parse() Types.CRIGroup {
	return Types.CRIGroup{
		ColorTemperature: object.ColorTemperature,
		CRIs:             ParseList(object.CRIs),
	}
}

type CRI struct {
	CES string `xml:",attr"`
	CRI uint   `xml:",attr"`
}

func (object CRI) Parse() Types.CRI {
	return Types.CRI(object)
}

type Connector struct {
	Name     string  `xml:",attr"`
	Type     string  `xml:",attr"`
	DMXBreak uint    `xml:",attr"` // obsolete (see https://www.gdtf.eu/gdtf/file-spec/physical-descriptions/#table-27-connector-attributes)
	Gender   int     `xml:",attr"` // obsolete
	Length   float32 `xml:",attr"` // obsolete
}

func (object Connector) Parse() Types.Connector {
	return Types.Connector(object)
}

type Properties struct {
	OperatingTemperature *OperatingTemperature `xml:"OperatingTemperature"`
	Weight               *Weight               `xml:"Weight"`
	PowerConsumption     any                   `xml:"PowerConsumption"`
	LegHeight            *LegHeight            `xml:"LegHeight"`
}

func (object Properties) Parse() Types.Properties {
	return Types.Properties{
		OperatingTemperature: (*Types.OperatingTemperature)(object.OperatingTemperature),
		Weight:               (*Types.Weight)(object.Weight),
		PowerConsumption:     object.PowerConsumption,
		LegHeight:            (*Types.LegHeight)(object.LegHeight),
	}
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
