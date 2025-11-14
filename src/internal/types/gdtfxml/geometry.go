package XMLTypes

import Types "github.com/Patch2PDF/GDTF-Parser/types"

type Geometries struct {
	GeometryList          []Geometry          `xml:"Geometry"`
	AxisList              []Axis              `xml:"Axis"`
	FilterBeamList        []FilterBeam        `xml:"FilterBeam"`
	FilterColorList       []FilterColor       `xml:"FilterColor"`
	FilterGoboList        []FilterGobo        `xml:"FilterGobo"`
	FilterShaperList      []FilterShaper      `xml:"FilterShaper"`
	BeamList              []Beam              `xml:"Beam"`
	MediaServerLayerList  []MediaServerLayer  `xml:"MediaServerLayer"`
	MediaServerCameraList []MediaServerCamera `xml:"MediaServerCamera"`
	MediaServerMasterList []MediaServerMaster `xml:"MediaServerMaster"`
	DisplayList           []Display           `xml:"Display"`
	LaserList             []Laser             `xml:"Laser"`
	GeometryReferenceList []GeometryReference `xml:"GeometryReference"`
	WiringObjectList      []WiringObject      `xml:"WiringObject"`
	InventoryList         []Inventory         `xml:"Inventory"`
	StructureList         []Structure         `xml:"Structure"`
	SupportList           []Support           `xml:"Support"`
	MagnetList            []Magnet            `xml:"Magnet"`
}

func (geometries Geometries) Parse() Types.Geometries {
	return Types.Geometries{
		GeometryList:          ParseList(&geometries.GeometryList),
		AxisList:              ParseList(&geometries.AxisList),
		FilterBeamList:        ParseList(&geometries.FilterBeamList),
		FilterColorList:       ParseList(&geometries.FilterColorList),
		FilterGoboList:        ParseList(&geometries.FilterGoboList),
		FilterShaperList:      ParseList(&geometries.FilterShaperList),
		BeamList:              ParseList(&geometries.BeamList),
		MediaServerLayerList:  ParseList(&geometries.MediaServerLayerList),
		MediaServerCameraList: ParseList(&geometries.MediaServerCameraList),
		MediaServerMasterList: ParseList(&geometries.MediaServerMasterList),
		DisplayList:           ParseList(&geometries.DisplayList),
		LaserList:             ParseList(&geometries.LaserList),
		GeometryReferenceList: ParseList(&geometries.GeometryReferenceList),
		WiringObjectList:      ParseList(&geometries.WiringObjectList),
		InventoryList:         ParseList(&geometries.InventoryList),
		StructureList:         ParseList(&geometries.StructureList),
		SupportList:           ParseList(&geometries.SupportList),
		MagnetList:            ParseList(&geometries.MagnetList),
	}
}

type Geometry struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Geometries
}

func (geometry Geometry) Parse() Types.Geometry {
	return Types.Geometry{
		Name:       geometry.Name,
		Model:      geometry.Model,
		Position:   Types.Matrix(geometry.Position),
		Geometries: geometry.Geometries.Parse(),
	}
}

type Axis struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Geometries
}

func (geometry Axis) Parse() Types.Axis {
	return Types.Axis{
		Name:       geometry.Name,
		Model:      geometry.Model,
		Position:   Types.Matrix(geometry.Position),
		Geometries: geometry.Geometries.Parse(),
	}
}

type FilterBeam struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Geometries
}

func (geometry FilterBeam) Parse() Types.FilterBeam {
	return Types.FilterBeam{
		Name:       geometry.Name,
		Model:      geometry.Model,
		Position:   Types.Matrix(geometry.Position),
		Geometries: geometry.Geometries.Parse(),
	}
}

type FilterColor struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Geometries
}

func (geometry FilterColor) Parse() Types.FilterColor {
	return Types.FilterColor{
		Name:       geometry.Name,
		Model:      geometry.Model,
		Position:   Types.Matrix(geometry.Position),
		Geometries: geometry.Geometries.Parse(),
	}
}

type FilterGobo struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Geometries
}

func (geometry FilterGobo) Parse() Types.FilterGobo {
	return Types.FilterGobo{
		Name:       geometry.Name,
		Model:      geometry.Model,
		Position:   Types.Matrix(geometry.Position),
		Geometries: geometry.Geometries.Parse(),
	}
}

type FilterShaper struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Geometries
}

func (geometry FilterShaper) Parse() Types.FilterShaper {
	return Types.FilterShaper{
		Name:       geometry.Name,
		Model:      geometry.Model,
		Position:   Types.Matrix(geometry.Position),
		Geometries: geometry.Geometries.Parse(),
	}
}

type Beam struct {
	Name             string           `xml:",attr"`
	Model            string           `xml:",attr"`
	Position         Matrix           `xml:",attr"`
	LampType         string           `xml:",attr"`
	PowerConsumption float32          `xml:",attr"`
	LuminousFlux     float32          `xml:",attr"`
	ColorTemperature float32          `xml:",attr"`
	BeamAngle        float32          `xml:",attr"`
	FieldAngle       float32          `xml:",attr"`
	ThrowRatio       float32          `xml:",attr"`
	RectangleRatio   float32          `xml:",attr"`
	BeamRadius       float32          `xml:",attr"`
	BeamType         string           `xml:",attr"`
	CRI              uint             `xml:"ColorRenderingIndex,attr"`
	EmitterSpectrum  XMLNodeReference `xml:",attr"`
	Geometries
}

func (geometry Beam) Parse() Types.Beam {
	return Types.Beam{
		Name:             geometry.Name,
		Model:            geometry.Model,
		Position:         Types.Matrix(geometry.Position),
		LampType:         geometry.LampType,
		PowerConsumption: geometry.PowerConsumption,
		LuminousFlux:     geometry.LuminousFlux,
		ColorTemperature: geometry.ColorTemperature,
		BeamAngle:        geometry.BeamAngle,
		FieldAngle:       geometry.FieldAngle,
		ThrowRatio:       geometry.ThrowRatio,
		RectangleRatio:   geometry.RectangleRatio,
		BeamRadius:       geometry.BeamRadius,
		BeamType:         geometry.BeamType,
		CRI:              geometry.CRI,
		EmitterSpectrum: Types.NodeReference[Types.Emitter]{
			String: geometry.EmitterSpectrum,
		},
		Geometries: geometry.Geometries.Parse(),
	}
}

type MediaServerLayer struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Geometries
}

func (geometry MediaServerLayer) Parse() Types.MediaServerLayer {
	return Types.MediaServerLayer{
		Name:       geometry.Name,
		Model:      geometry.Model,
		Position:   Types.Matrix(geometry.Position),
		Geometries: geometry.Geometries.Parse(),
	}
}

type MediaServerCamera struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Geometries
}

func (geometry MediaServerCamera) Parse() Types.MediaServerCamera {
	return Types.MediaServerCamera{
		Name:       geometry.Name,
		Model:      geometry.Model,
		Position:   Types.Matrix(geometry.Position),
		Geometries: geometry.Geometries.Parse(),
	}
}

type MediaServerMaster struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Geometries
}

func (geometry MediaServerMaster) Parse() Types.MediaServerMaster {
	return Types.MediaServerMaster{
		Name:       geometry.Name,
		Model:      geometry.Model,
		Position:   Types.Matrix(geometry.Position),
		Geometries: geometry.Geometries.Parse(),
	}
}

type Display struct {
	Name     string        `xml:",attr"`
	Model    string        `xml:",attr"`
	Position Matrix        `xml:",attr"`
	Texture  FileReference `xml:",attr"`
	Geometries
}

func (geometry Display) Parse() Types.Display {
	return Types.Display{
		Name:       geometry.Name,
		Model:      geometry.Model,
		Position:   Types.Matrix(geometry.Position),
		Geometries: geometry.Geometries.Parse(),
	}
}

type GeometryReference struct {
	Name        string  `xml:",attr"`
	Model       string  `xml:",attr"`
	Position    Matrix  `xml:",attr"`
	GeometryRef string  `xml:"Geometry,attr"` // only top level geometries allowed to be referenced
	Breaks      []Break `xml:"Break"`
}

func (geometry GeometryReference) Parse() Types.GeometryReference {
	return Types.GeometryReference{
		Name:     geometry.Name,
		Model:    geometry.Model,
		Position: Types.Matrix(geometry.Position),
		GeometryRef: Types.NodeReference[Types.GeometryNodeReference]{
			String: geometry.GeometryRef,
		},
		Breaks: ParseList(&geometry.Breaks),
	}
}

type Break struct {
	DMXOffset DMXAddress `xml:",attr"`
	DMXBreak  uint       `xml:",attr"`
}

func (geometry Break) Parse() Types.Break {
	return Types.Break{
		DMXOffset: Types.DMXAddress(geometry.DMXOffset),
		DMXBreak:  geometry.DMXBreak,
	}
}

type Laser struct {
	Name              string           `xml:",attr"`
	Model             string           `xml:",attr"`
	Position          Matrix           `xml:",attr"`
	ColorType         string           `xml:",attr"` // enum
	Color             float32          `xml:",attr"` // Required if ColorType is “SingleWaveLength”; Unit:nm (nanometers)
	OutputStrength    float32          `xml:",attr"`
	Emitter           XMLNodeReference `xml:",attr"`
	BeamDiameter      float32          `xml:",attr"`
	BeamDivergenceMin float32          `xml:",attr"`
	BeamDivergenceMax float32          `xml:",attr"`
	ScanAnglePan      float32          `xml:",attr"`
	ScanAngleTilt     float32          `xml:",attr"`
	ScanSpeed         float32          `xml:",attr"`
	Protocols         []LaserProtocol  `xml:"Protocol"`
	Geometries
}

func (geometry Laser) Parse() Types.Laser {
	return Types.Laser{
		Name:           geometry.Name,
		Model:          geometry.Model,
		Position:       Types.Matrix(geometry.Position),
		ColorType:      geometry.ColorType,
		Color:          geometry.Color,
		OutputStrength: geometry.OutputStrength,
		Emitter: Types.NodeReference[Types.Emitter]{
			String: geometry.Emitter,
		},
		BeamDiameter:      geometry.BeamDiameter,
		BeamDivergenceMin: geometry.BeamDivergenceMin,
		BeamDivergenceMax: geometry.BeamDivergenceMax,
		ScanAnglePan:      geometry.ScanAnglePan,
		ScanAngleTilt:     geometry.ScanAngleTilt,
		ScanSpeed:         geometry.ScanSpeed,
		Protocols:         ParseList(&geometry.Protocols),
		Geometries:        geometry.Geometries.Parse(),
	}
}

type LaserProtocol struct {
	Name string `xml:",attr"`
}

func (protocol LaserProtocol) Parse() Types.LaserProtocol {
	return Types.LaserProtocol(protocol)
}

type WiringObject struct {
	Name              string     `xml:",attr"`
	Model             string     `xml:",attr"`
	ConnectorType     string     `xml:",attr"`
	Position          Matrix     `xml:",attr"`
	ComponentType     string     `xml:",attr"` //enum
	SignalType        string     `xml:",attr"`
	PinCount          int        `xml:",attr"`
	ElectricalPayLoad float32    `xml:",attr"`
	VoltageRangeMax   float32    `xml:",attr"`
	VoltageRangeMin   float32    `xml:",attr"`
	FrequencyRangeMax float32    `xml:",attr"`
	FrequencyRangeMin float32    `xml:",attr"`
	MaxPayLoad        float32    `xml:",attr"`
	Voltage           float32    `xml:",attr"`
	SignalLayer       int        `xml:",attr"`
	CosPhi            float32    `xml:",attr"`
	FuseCurrent       float32    `xml:",attr"` // in ampere
	FuseRating        string     `xml:",attr"` //enum
	Orientation       string     `xml:",attr"` //enum
	WireGroup         string     `xml:",attr"`
	PinPatches        []PinPatch `xml:"PinPatch"`
	Geometries
}

func (geometry WiringObject) Parse() Types.WiringObject {
	return Types.WiringObject{
		Name:              geometry.Name,
		Model:             geometry.Model,
		ConnectorType:     geometry.ConnectorType,
		Position:          Types.Matrix(geometry.Position),
		ComponentType:     geometry.ComponentType,
		SignalType:        geometry.SignalType,
		PinCount:          geometry.PinCount,
		ElectricalPayLoad: geometry.ElectricalPayLoad,
		VoltageRangeMax:   geometry.VoltageRangeMax,
		VoltageRangeMin:   geometry.VoltageRangeMin,
		FrequencyRangeMax: geometry.FrequencyRangeMax,
		FrequencyRangeMin: geometry.FrequencyRangeMin,
		MaxPayLoad:        geometry.MaxPayLoad,
		Voltage:           geometry.Voltage,
		SignalLayer:       geometry.SignalLayer,
		CosPhi:            geometry.CosPhi,
		FuseCurrent:       geometry.FuseCurrent,
		FuseRating:        geometry.FuseRating,
		Orientation:       geometry.Orientation,
		WireGroup:         geometry.WireGroup,
		PinPatches:        ParseList(&geometry.PinPatches),
		Geometries:        geometry.Geometries.Parse(),
	}
}

type PinPatch struct {
	ToWiringObject XMLNodeReference `xml:",attr"`
	FromPin        int              `xml:",attr"`
	ToPin          int              `xml:",attr"`
}

func (pinpatch PinPatch) Parse() Types.PinPatch {
	return Types.PinPatch{
		ToWiringObject: Types.NodeReference[Types.WiringObject]{
			String: pinpatch.ToWiringObject,
		},
		FromPin: pinpatch.FromPin,
		ToPin:   pinpatch.ToPin,
	}
}

type Inventory struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Count    int    `xml:",attr"`
	Geometries
}

func (geometry Inventory) Parse() Types.Inventory {
	return Types.Inventory{
		Name:       geometry.Name,
		Model:      geometry.Model,
		Position:   Types.Matrix(geometry.Position),
		Count:      geometry.Count,
		Geometries: geometry.Geometries.Parse(),
	}
}

type Structure struct {
	Name                      string           `xml:",attr"`
	Model                     string           `xml:",attr"`
	Position                  Matrix           `xml:",attr"`
	LinkedGeometry            XMLNodeReference `xml:",attr"`
	StructureType             string           `xml:",attr"` // enum
	CrossSectionType          string           `xml:",attr"` // enum
	CrossSectionHeight        float32          `xml:",attr"`
	CrossSectionWallThickness float32          `xml:",attr"`
	TrussCrossSection         string           `xml:",attr"`
	Geometries
}

func (geometry Structure) Parse() Types.Structure {
	return Types.Structure{
		Name:                      geometry.Name,
		Model:                     geometry.Model,
		Position:                  Types.Matrix(geometry.Position),
		LinkedGeometry:            geometry.LinkedGeometry, // TODO: add pointer? -> find a sample
		StructureType:             geometry.StructureType,
		CrossSectionType:          geometry.CrossSectionType,
		CrossSectionHeight:        geometry.CrossSectionHeight,
		CrossSectionWallThickness: geometry.CrossSectionWallThickness,
		TrussCrossSection:         geometry.TrussCrossSection,
		Geometries:                geometry.Geometries.Parse(),
	}
}

type Support struct {
	Name             string  `xml:",attr"`
	Model            string  `xml:",attr"`
	Position         Matrix  `xml:",attr"`
	SupportType      string  `xml:",attr"` //enum
	RopeCrossSection string  `xml:",attr"`
	RopeOffset       Vector3 `xml:",attr"`
	CapacityX        float32 `xml:",attr"`
	CapacityY        float32 `xml:",attr"`
	CapacityZ        float32 `xml:",attr"`
	CapacityXX       float32 `xml:",attr"`
	CapacityYY       float32 `xml:",attr"`
	CapacityZZ       float32 `xml:",attr"`
	ResistanceX      float32 `xml:",attr"`
	ResistanceY      float32 `xml:",attr"`
	ResistanceZ      float32 `xml:",attr"`
	ResistanceXX     float32 `xml:",attr"`
	ResistanceYY     float32 `xml:",attr"`
	ResistanceZZ     float32 `xml:",attr"`
	Geometries
}

func (geometry Support) Parse() Types.Support {
	return Types.Support{
		Name:             geometry.Name,
		Model:            geometry.Model,
		Position:         Types.Matrix(geometry.Position),
		SupportType:      geometry.SupportType,
		RopeCrossSection: geometry.RopeCrossSection,
		RopeOffset:       Types.Vector3(geometry.RopeOffset),
		CapacityX:        geometry.CapacityX,
		CapacityY:        geometry.CapacityY,
		CapacityZ:        geometry.CapacityZ,
		CapacityXX:       geometry.CapacityXX,
		CapacityYY:       geometry.CapacityYY,
		CapacityZZ:       geometry.CapacityZZ,
		ResistanceX:      geometry.ResistanceX,
		ResistanceY:      geometry.ResistanceY,
		ResistanceZ:      geometry.ResistanceZ,
		ResistanceXX:     geometry.ResistanceXX,
		ResistanceYY:     geometry.ResistanceYY,
		ResistanceZZ:     geometry.ResistanceZZ,
		Geometries:       geometry.Geometries.Parse(),
	}
}

type Magnet struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Geometries
}

func (geometry Magnet) Parse() Types.Magnet {
	return Types.Magnet{
		Name:       geometry.Name,
		Model:      geometry.Model,
		Position:   Types.Matrix(geometry.Position),
		Geometries: geometry.Geometries.Parse(),
	}
}
