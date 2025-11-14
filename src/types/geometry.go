package Types

import (
	"reflect"
	"strings"
)

type GeometryNodeReference struct {
	Ptr  any
	Type string
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
	Name     string
	Model    string
	Position Matrix
	Geometries
}

func (obj *Geometry) CreateGeometryReferencePointer(parentPrefix string) {
	newParentPrefix := strings.Trim(parentPrefix+"."+obj.Name, ".")
	refPointers.Geometries[newParentPrefix] = &GeometryNodeReference{
		Ptr:  obj,
		Type: reflect.TypeOf(obj).String(),
	}
	obj.Geometries.CreateGeometryReferencePointer(newParentPrefix)
}

type Axis struct {
	Name     string
	Model    string
	Position Matrix
	Geometries
}

func (obj *Axis) CreateGeometryReferencePointer(parentPrefix string) {
	newParentPrefix := strings.Trim(parentPrefix+"."+obj.Name, ".")
	refPointers.Geometries[newParentPrefix] = &GeometryNodeReference{
		Ptr:  obj,
		Type: reflect.TypeOf(obj).String(),
	}
	obj.Geometries.CreateGeometryReferencePointer(newParentPrefix)
}

type FilterBeam struct {
	Name     string
	Model    string
	Position Matrix
	Geometries
}

func (obj *FilterBeam) CreateGeometryReferencePointer(parentPrefix string) {
	newParentPrefix := strings.Trim(parentPrefix+"."+obj.Name, ".")
	refPointers.Geometries[newParentPrefix] = &GeometryNodeReference{
		Ptr:  obj,
		Type: reflect.TypeOf(obj).String(),
	}
	obj.Geometries.CreateGeometryReferencePointer(newParentPrefix)
}

type FilterColor struct {
	Name     string
	Model    string
	Position Matrix
	Geometries
}

func (obj *FilterColor) CreateGeometryReferencePointer(parentPrefix string) {
	newParentPrefix := strings.Trim(parentPrefix+"."+obj.Name, ".")
	refPointers.Geometries[newParentPrefix] = &GeometryNodeReference{
		Ptr:  obj,
		Type: reflect.TypeOf(obj).String(),
	}
	obj.Geometries.CreateGeometryReferencePointer(newParentPrefix)
}

type FilterGobo struct {
	Name     string
	Model    string
	Position Matrix
	Geometries
}

func (obj *FilterGobo) CreateGeometryReferencePointer(parentPrefix string) {
	newParentPrefix := strings.Trim(parentPrefix+"."+obj.Name, ".")
	refPointers.Geometries[newParentPrefix] = &GeometryNodeReference{
		Ptr:  obj,
		Type: reflect.TypeOf(obj).String(),
	}
	obj.Geometries.CreateGeometryReferencePointer(newParentPrefix)
}

type FilterShaper struct {
	Name     string
	Model    string
	Position Matrix
	Geometries
}

func (obj *FilterShaper) CreateGeometryReferencePointer(parentPrefix string) {
	newParentPrefix := strings.Trim(parentPrefix+"."+obj.Name, ".")
	refPointers.Geometries[newParentPrefix] = &GeometryNodeReference{
		Ptr:  obj,
		Type: reflect.TypeOf(obj).String(),
	}
	obj.Geometries.CreateGeometryReferencePointer(newParentPrefix)
}

type Beam struct {
	Name             string
	Model            string
	Position         Matrix
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
	Geometries
}

func (obj *Beam) CreateGeometryReferencePointer(parentPrefix string) {
	newParentPrefix := strings.Trim(parentPrefix+"."+obj.Name, ".")
	refPointers.Geometries[newParentPrefix] = &GeometryNodeReference{
		Ptr:  obj,
		Type: reflect.TypeOf(obj).String(),
	}
	obj.Geometries.CreateGeometryReferencePointer(newParentPrefix)
}

func (obj *Beam) ResolveReference() {
	obj.EmitterSpectrum.Ptr = refPointers.Emitters[obj.EmitterSpectrum.String]
}

type MediaServerLayer struct {
	Name     string
	Model    string
	Position Matrix
	Geometries
}

func (obj *MediaServerLayer) CreateGeometryReferencePointer(parentPrefix string) {
	newParentPrefix := strings.Trim(parentPrefix+"."+obj.Name, ".")
	refPointers.Geometries[newParentPrefix] = &GeometryNodeReference{
		Ptr:  obj,
		Type: reflect.TypeOf(obj).String(),
	}
	obj.Geometries.CreateGeometryReferencePointer(newParentPrefix)
}

type MediaServerCamera struct {
	Name     string
	Model    string
	Position Matrix
	Geometries
}

func (obj *MediaServerCamera) CreateGeometryReferencePointer(parentPrefix string) {
	newParentPrefix := strings.Trim(parentPrefix+"."+obj.Name, ".")
	refPointers.Geometries[newParentPrefix] = &GeometryNodeReference{
		Ptr:  obj,
		Type: reflect.TypeOf(obj).String(),
	}
	obj.Geometries.CreateGeometryReferencePointer(newParentPrefix)
}

type MediaServerMaster struct {
	Name     string
	Model    string
	Position Matrix
	Geometries
}

func (obj *MediaServerMaster) CreateGeometryReferencePointer(parentPrefix string) {
	newParentPrefix := strings.Trim(parentPrefix+"."+obj.Name, ".")
	refPointers.Geometries[newParentPrefix] = &GeometryNodeReference{
		Ptr:  obj,
		Type: reflect.TypeOf(obj).String(),
	}
	obj.Geometries.CreateGeometryReferencePointer(newParentPrefix)
}

type Display struct {
	Name     string
	Model    string
	Position Matrix
	Texture  FileReference
	Geometries
}

func (obj *Display) CreateGeometryReferencePointer(parentPrefix string) {
	newParentPrefix := strings.Trim(parentPrefix+"."+obj.Name, ".")
	refPointers.Geometries[newParentPrefix] = &GeometryNodeReference{
		Ptr:  obj,
		Type: reflect.TypeOf(obj).String(),
	}
	obj.Geometries.CreateGeometryReferencePointer(newParentPrefix)
}

type GeometryReference struct {
	Name        string
	Model       string
	Position    Matrix
	GeometryRef NodeReference[GeometryNodeReference] // only top level geometries allowed to be referenced
	Breaks      []*Break
	// do we need to link a parent for this Geometry Type?
}

func (obj *GeometryReference) ResolveReference() {
	obj.GeometryRef.Ptr = refPointers.Geometries[obj.GeometryRef.String]
}

type Break struct {
	DMXOffset DMXAddress
	DMXBreak  uint
}

type Laser struct {
	Name              string
	Model             string
	Position          Matrix
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
	Geometries
}

func (obj *Laser) CreateGeometryReferencePointer(parentPrefix string) {
	newParentPrefix := strings.Trim(parentPrefix+"."+obj.Name, ".")
	refPointers.Geometries[newParentPrefix] = &GeometryNodeReference{
		Ptr:  obj,
		Type: reflect.TypeOf(obj).String(),
	}
	obj.Geometries.CreateGeometryReferencePointer(newParentPrefix)
}

func (obj *Laser) ResolveReference() {
	obj.Emitter.Ptr = refPointers.Emitters[obj.Emitter.String]
}

type LaserProtocol struct {
	Name string
}

type WiringObject struct {
	Name              string
	Model             string
	ConnectorType     string
	Position          Matrix
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
	Geometries
}

func (obj *WiringObject) CreateGeometryReferencePointer(parentPrefix string) {
	newParentPrefix := strings.Trim(parentPrefix+"."+obj.Name, ".")
	refPointers.Geometries[newParentPrefix] = &GeometryNodeReference{
		Ptr:  obj,
		Type: reflect.TypeOf(obj).String(),
	}
	obj.Geometries.CreateGeometryReferencePointer(newParentPrefix)
}

func (obj *WiringObject) ResolveReference() {
	ResolveReferences(&obj.PinPatches)
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
	Name     string
	Model    string
	Position Matrix
	Count    int
	Geometries
}

func (obj *Inventory) CreateGeometryReferencePointer(parentPrefix string) {
	newParentPrefix := strings.Trim(parentPrefix+"."+obj.Name, ".")
	refPointers.Geometries[newParentPrefix] = &GeometryNodeReference{
		Ptr:  obj,
		Type: reflect.TypeOf(obj).String(),
	}
	obj.Geometries.CreateGeometryReferencePointer(newParentPrefix)
}

type Structure struct {
	Name                      string
	Model                     string
	Position                  Matrix
	LinkedGeometry            string // for now (analyse if this can be a NodeReference instead)
	StructureType             string // enum
	CrossSectionType          string // enum
	CrossSectionHeight        float32
	CrossSectionWallThickness float32
	TrussCrossSection         string
	Geometries
}

func (obj *Structure) CreateGeometryReferencePointer(parentPrefix string) {
	newParentPrefix := strings.Trim(parentPrefix+"."+obj.Name, ".")
	refPointers.Geometries[newParentPrefix] = &GeometryNodeReference{
		Ptr:  obj,
		Type: reflect.TypeOf(obj).String(),
	}
	obj.Geometries.CreateGeometryReferencePointer(newParentPrefix)
}

type Support struct {
	Name             string
	Model            string
	Position         Matrix
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
	Geometries
}

func (obj *Support) CreateGeometryReferencePointer(parentPrefix string) {
	newParentPrefix := strings.Trim(parentPrefix+"."+obj.Name, ".")
	refPointers.Geometries[newParentPrefix] = &GeometryNodeReference{
		Ptr:  obj,
		Type: reflect.TypeOf(obj).String(),
	}
	obj.Geometries.CreateGeometryReferencePointer(newParentPrefix)
}

type Magnet struct {
	Name     string
	Model    string
	Position Matrix
	Geometries
}

func (obj *Magnet) CreateGeometryReferencePointer(parentPrefix string) {
	newParentPrefix := strings.Trim(parentPrefix+"."+obj.Name, ".")
	refPointers.Geometries[newParentPrefix] = &GeometryNodeReference{
		Ptr:  obj,
		Type: reflect.TypeOf(obj).String(),
	}
	obj.Geometries.CreateGeometryReferencePointer(newParentPrefix)
}
