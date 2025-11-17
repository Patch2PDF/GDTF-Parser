package Types

import (
	"reflect"
	"strings"
)

type GeometryNodeReference struct {
	Ptr  any
	Type string
}

type GeometryBase struct {
	Name     string
	Model    NodeReference[Model]
	Position Matrix
	Geometries
}

func (obj *GeometryBase) CreateGeometryReferencePointer(parentPrefix string) {
	newParentPrefix := strings.Trim(parentPrefix+"."+obj.Name, ".")
	refPointers.Geometries[newParentPrefix] = &GeometryNodeReference{
		Ptr:  obj,
		Type: reflect.TypeOf(obj).String(),
	}
	obj.Geometries.CreateGeometryReferencePointer(newParentPrefix)
}

func (obj *GeometryBase) ResolveReference() {
	obj.Model.Ptr = refPointers.Models[obj.Model.String]
	obj.Geometries.ResolveReference()
}

type Geometries struct {
	GeometryList          []*Geometry
	AxisList              []*Axis
	FilterBeamList        []*FilterBeam
	FilterColorList       []*FilterColor
	FilterGoboList        []*FilterGobo
	FilterShaperList      []*FilterShaper
	BeamList              []*Beam
	MediaServerLayerList  []*MediaServerLayer
	MediaServerCameraList []*MediaServerCamera
	MediaServerMasterList []*MediaServerMaster
	DisplayList           []*Display
	LaserList             []*Laser
	GeometryReferenceList []*GeometryReference
	WiringObjectList      []*WiringObject
	InventoryList         []*Inventory
	StructureList         []*Structure
	SupportList           []*Support
	MagnetList            []*Magnet

	Parent *GeometryNodeReference
}

func (obj *Geometries) CreateGeometryReferencePointer(parentPrefix string) {
	CreateGeometryReferencePointers(&obj.GeometryList, parentPrefix)
	CreateGeometryReferencePointers(&obj.AxisList, parentPrefix)
	CreateGeometryReferencePointers(&obj.FilterBeamList, parentPrefix)
	CreateGeometryReferencePointers(&obj.FilterColorList, parentPrefix)
	CreateGeometryReferencePointers(&obj.FilterGoboList, parentPrefix)
	CreateGeometryReferencePointers(&obj.FilterShaperList, parentPrefix)
	CreateGeometryReferencePointers(&obj.BeamList, parentPrefix)
	CreateGeometryReferencePointers(&obj.MediaServerLayerList, parentPrefix)
	CreateGeometryReferencePointers(&obj.MediaServerCameraList, parentPrefix)
	CreateGeometryReferencePointers(&obj.MediaServerMasterList, parentPrefix)
	CreateGeometryReferencePointers(&obj.DisplayList, parentPrefix)
	CreateGeometryReferencePointers(&obj.LaserList, parentPrefix)
	// CreateReferencePointers(&obj.GeometryReferenceList, parentPrefix)
	CreateGeometryReferencePointers(&obj.WiringObjectList, parentPrefix)
	CreateGeometryReferencePointers(&obj.InventoryList, parentPrefix)
	CreateGeometryReferencePointers(&obj.StructureList, parentPrefix)
	CreateGeometryReferencePointers(&obj.SupportList, parentPrefix)
	CreateGeometryReferencePointers(&obj.MagnetList, parentPrefix)
}

func (obj *Geometries) ResolveReference() {
	ResolveReferences(&obj.GeometryList)
	ResolveReferences(&obj.AxisList)
	ResolveReferences(&obj.FilterBeamList)
	ResolveReferences(&obj.FilterColorList)
	ResolveReferences(&obj.FilterGoboList)
	ResolveReferences(&obj.FilterShaperList)
	ResolveReferences(&obj.BeamList)
	ResolveReferences(&obj.MediaServerLayerList)
	ResolveReferences(&obj.MediaServerCameraList)
	ResolveReferences(&obj.MediaServerMasterList)
	ResolveReferences(&obj.DisplayList)
	ResolveReferences(&obj.LaserList)
	ResolveReferences(&obj.GeometryReferenceList)
	ResolveReferences(&obj.WiringObjectList)
	ResolveReferences(&obj.InventoryList)
	ResolveReferences(&obj.StructureList)
	ResolveReferences(&obj.SupportList)
	ResolveReferences(&obj.MagnetList)
}

type Geometry struct {
	GeometryBase
}

type Axis struct {
	GeometryBase
}

type FilterBeam struct {
	GeometryBase
}

type FilterColor struct {
	GeometryBase
}

type FilterGobo struct {
	GeometryBase
}

type FilterShaper struct {
	GeometryBase
}

type Beam struct {
	GeometryBase
	LampType         string
	PowerConsumption float32
	LuminousFlux     float32
	ColorTemperature float32
	BeamAngle        float32
	FieldAngle       float32
	ThrowRatio       float32
	RectangleRatio   float32
	BeamRadius       float32
	BeamType         string
	CRI              uint
	EmitterSpectrum  NodeReference[Emitter]
}

func (obj *Beam) ResolveReference() {
	obj.EmitterSpectrum.Ptr = refPointers.Emitters[obj.EmitterSpectrum.String]
	obj.GeometryBase.ResolveReference()
}

type MediaServerLayer struct {
	GeometryBase
}

type MediaServerCamera struct {
	GeometryBase
}

type MediaServerMaster struct {
	GeometryBase
}

type Display struct {
	GeometryBase
	Texture FileReference
}

type GeometryReference struct {
	Name        string
	Model       NodeReference[Model]
	Position    Matrix
	GeometryRef NodeReference[GeometryNodeReference] // only top level geometries allowed to be referenced
	Breaks      []*Break
	// do we need to link a parent for this Geometry Type?
}

func (obj *GeometryReference) ResolveReference() {
	obj.GeometryRef.Ptr = refPointers.Geometries[obj.GeometryRef.String]
	obj.Model.Ptr = refPointers.Models[obj.Model.String]
}

type Break struct {
	DMXOffset DMXAddress
	DMXBreak  uint
}

type Laser struct {
	GeometryBase
	ColorType         string  // enum
	Color             float32 // Required if ColorType is “SingleWaveLength”; Unit:nm (nanometers)
	OutputStrength    float32
	Emitter           NodeReference[Emitter]
	BeamDiameter      float32
	BeamDivergenceMin float32
	BeamDivergenceMax float32
	ScanAnglePan      float32
	ScanAngleTilt     float32
	ScanSpeed         float32
	Protocols         []*LaserProtocol
}

func (obj *Laser) ResolveReference() {
	obj.Emitter.Ptr = refPointers.Emitters[obj.Emitter.String]
	obj.GeometryBase.ResolveReference()
}

type LaserProtocol struct {
	Name string
}

type WiringObject struct {
	GeometryBase
	ConnectorType     string
	ComponentType     string //enum
	SignalType        string
	PinCount          int
	ElectricalPayLoad float32
	VoltageRangeMax   float32
	VoltageRangeMin   float32
	FrequencyRangeMax float32
	FrequencyRangeMin float32
	MaxPayLoad        float32
	Voltage           float32
	SignalLayer       int
	CosPhi            float32
	FuseCurrent       float32 // in ampere
	FuseRating        string  //enum
	Orientation       string  //enum
	WireGroup         string
	PinPatches        []*PinPatch
}

func (obj *WiringObject) ResolveReference() {
	ResolveReferences(&obj.PinPatches)
	obj.GeometryBase.ResolveReference()
}

type PinPatch struct {
	ToWiringObject NodeReference[WiringObject]
	FromPin        int
	ToPin          int
}

func (obj *PinPatch) ResolveReference() {
	obj.ToWiringObject.Ptr = refPointers.WiringObjects[obj.ToWiringObject.String]
}

type Inventory struct {
	GeometryBase
	Count int
}

type Structure struct {
	GeometryBase
	LinkedGeometry            string // for now (analyse if this can be a NodeReference instead)
	StructureType             string // enum
	CrossSectionType          string // enum
	CrossSectionHeight        float32
	CrossSectionWallThickness float32
	TrussCrossSection         string
}

type Support struct {
	GeometryBase
	SupportType      string //enum
	RopeCrossSection string
	RopeOffset       Vector3
	CapacityX        float32
	CapacityY        float32
	CapacityZ        float32
	CapacityXX       float32
	CapacityYY       float32
	CapacityZZ       float32
	ResistanceX      float32
	ResistanceY      float32
	ResistanceZ      float32
	ResistanceXX     float32
	ResistanceYY     float32
	ResistanceZZ     float32
}

type Magnet struct {
	GeometryBase
}
