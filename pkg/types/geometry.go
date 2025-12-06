package Types

import (
	"reflect"
	"strings"

	"github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"
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

func (obj *GeometryBase) CreateGeometryReferencePointer(refPointers *ReferencePointers, parentPrefix string) {
	newParentPrefix := strings.Trim(parentPrefix+"."+obj.Name, ".")
	refPointers.Geometries[newParentPrefix] = &GeometryNodeReference{
		Ptr:  obj,
		Type: reflect.TypeOf(obj).String(),
	}
	obj.Geometries.CreateGeometryReferencePointer(refPointers, newParentPrefix)
}

func (obj *GeometryBase) ResolveReference(refPointers *ReferencePointers) {
	obj.Model.Ptr = refPointers.Models[obj.Model.String]
	obj.Geometries.ResolveReference(refPointers)
}

type MeshGenerator interface {
	GenerateMesh(parentTransformation MeshTypes.Matrix) *MeshTypes.Mesh
}

func (obj *GeometryBase) GenerateMesh(parentTransformation MeshTypes.Matrix) *MeshTypes.Mesh {
	var mesh1 MeshTypes.Mesh
	localTransformation := obj.Position.toMeshMatrix()
	transformation := parentTransformation.Mul(localTransformation)
	if obj.Model.Ptr != nil && obj.Model.Ptr.Mesh != nil {
		mesh1 = obj.Model.Ptr.Mesh.Copy()
		mesh1.RotateAndTranslate(transformation)
	}
	mesh2 := obj.Geometries.GenerateMesh(transformation)
	return mesh1.Add(mesh2)
}

func GenerateMeshes[T MeshGenerator](source *[]T, parentTransformation MeshTypes.Matrix) *MeshTypes.Mesh {
	mesh := &MeshTypes.Mesh{}
	if source == nil {
		return mesh
	}
	for i := range *source {
		mesh.Add((*source)[i].GenerateMesh(parentTransformation))
	}
	return mesh
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

func (obj *Geometries) CreateGeometryReferencePointer(refPointers *ReferencePointers, parentPrefix string) {
	CreateGeometryReferencePointers(refPointers, &obj.GeometryList, parentPrefix)
	CreateGeometryReferencePointers(refPointers, &obj.AxisList, parentPrefix)
	CreateGeometryReferencePointers(refPointers, &obj.FilterBeamList, parentPrefix)
	CreateGeometryReferencePointers(refPointers, &obj.FilterColorList, parentPrefix)
	CreateGeometryReferencePointers(refPointers, &obj.FilterGoboList, parentPrefix)
	CreateGeometryReferencePointers(refPointers, &obj.FilterShaperList, parentPrefix)
	CreateGeometryReferencePointers(refPointers, &obj.BeamList, parentPrefix)
	CreateGeometryReferencePointers(refPointers, &obj.MediaServerLayerList, parentPrefix)
	CreateGeometryReferencePointers(refPointers, &obj.MediaServerCameraList, parentPrefix)
	CreateGeometryReferencePointers(refPointers, &obj.MediaServerMasterList, parentPrefix)
	CreateGeometryReferencePointers(refPointers, &obj.DisplayList, parentPrefix)
	CreateGeometryReferencePointers(refPointers, &obj.LaserList, parentPrefix)
	// CreateReferencePointers(refPointers, &obj.GeometryReferenceList, parentPrefix)
	CreateGeometryReferencePointers(refPointers, &obj.WiringObjectList, parentPrefix)
	CreateGeometryReferencePointers(refPointers, &obj.InventoryList, parentPrefix)
	CreateGeometryReferencePointers(refPointers, &obj.StructureList, parentPrefix)
	CreateGeometryReferencePointers(refPointers, &obj.SupportList, parentPrefix)
	CreateGeometryReferencePointers(refPointers, &obj.MagnetList, parentPrefix)
}

func (obj *Geometries) ResolveReference(refPointers *ReferencePointers) {
	ResolveReferences(refPointers, &obj.GeometryList)
	ResolveReferences(refPointers, &obj.AxisList)
	ResolveReferences(refPointers, &obj.FilterBeamList)
	ResolveReferences(refPointers, &obj.FilterColorList)
	ResolveReferences(refPointers, &obj.FilterGoboList)
	ResolveReferences(refPointers, &obj.FilterShaperList)
	ResolveReferences(refPointers, &obj.BeamList)
	ResolveReferences(refPointers, &obj.MediaServerLayerList)
	ResolveReferences(refPointers, &obj.MediaServerCameraList)
	ResolveReferences(refPointers, &obj.MediaServerMasterList)
	ResolveReferences(refPointers, &obj.DisplayList)
	ResolveReferences(refPointers, &obj.LaserList)
	ResolveReferences(refPointers, &obj.GeometryReferenceList)
	ResolveReferences(refPointers, &obj.WiringObjectList)
	ResolveReferences(refPointers, &obj.InventoryList)
	ResolveReferences(refPointers, &obj.StructureList)
	ResolveReferences(refPointers, &obj.SupportList)
	ResolveReferences(refPointers, &obj.MagnetList)
}

func (obj *Geometries) GenerateMesh(parentTransformation MeshTypes.Matrix) *MeshTypes.Mesh {
	mesh := GenerateMeshes(&obj.GeometryList, parentTransformation)
	mesh.Add(GenerateMeshes(&obj.AxisList, parentTransformation))
	mesh.Add(GenerateMeshes(&obj.FilterBeamList, parentTransformation))
	mesh.Add(GenerateMeshes(&obj.FilterColorList, parentTransformation))
	mesh.Add(GenerateMeshes(&obj.FilterGoboList, parentTransformation))
	mesh.Add(GenerateMeshes(&obj.FilterShaperList, parentTransformation))
	mesh.Add(GenerateMeshes(&obj.BeamList, parentTransformation))
	mesh.Add(GenerateMeshes(&obj.MediaServerLayerList, parentTransformation))
	mesh.Add(GenerateMeshes(&obj.MediaServerCameraList, parentTransformation))
	mesh.Add(GenerateMeshes(&obj.MediaServerMasterList, parentTransformation))
	mesh.Add(GenerateMeshes(&obj.DisplayList, parentTransformation))
	mesh.Add(GenerateMeshes(&obj.LaserList, parentTransformation))
	mesh.Add(GenerateMeshes(&obj.GeometryReferenceList, parentTransformation))
	mesh.Add(GenerateMeshes(&obj.WiringObjectList, parentTransformation))
	mesh.Add(GenerateMeshes(&obj.InventoryList, parentTransformation))
	mesh.Add(GenerateMeshes(&obj.StructureList, parentTransformation))
	mesh.Add(GenerateMeshes(&obj.SupportList, parentTransformation))
	mesh.Add(GenerateMeshes(&obj.MagnetList, parentTransformation))
	return mesh
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

func (obj *Beam) ResolveReference(refPointers *ReferencePointers) {
	obj.EmitterSpectrum.Ptr = refPointers.Emitters[obj.EmitterSpectrum.String]
	obj.GeometryBase.ResolveReference(refPointers)
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

func (obj *GeometryReference) ResolveReference(refPointers *ReferencePointers) {
	obj.GeometryRef.Ptr = refPointers.Geometries[obj.GeometryRef.String]
	obj.Model.Ptr = refPointers.Models[obj.Model.String]
}

func (obj *GeometryReference) GenerateMesh(parentTransformation MeshTypes.Matrix) *MeshTypes.Mesh {
	var mesh *MeshTypes.Mesh
	localTransformation := obj.Position.toMeshMatrix()
	transformation := parentTransformation.Mul(localTransformation)
	// if own model, replace parent mesh
	if obj.Model.Ptr != nil {
		temp := obj.Model.Ptr.Mesh.Copy()
		mesh = &temp
		geometries := obj.GeometryRef.Ptr.Ptr.(Geometries)
		mesh.Add(geometries.GenerateMesh(transformation))
	} else {
		ptr := obj.GeometryRef.Ptr.Ptr.(MeshGenerator)
		mesh = ptr.GenerateMesh(transformation)
	}
	return mesh
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

func (obj *Laser) ResolveReference(refPointers *ReferencePointers) {
	obj.Emitter.Ptr = refPointers.Emitters[obj.Emitter.String]
	obj.GeometryBase.ResolveReference(refPointers)
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

func (obj *WiringObject) ResolveReference(refPointers *ReferencePointers) {
	ResolveReferences(refPointers, &obj.PinPatches)
	obj.GeometryBase.ResolveReference(refPointers)
}

type PinPatch struct {
	ToWiringObject NodeReference[WiringObject]
	FromPin        int
	ToPin          int
}

func (obj *PinPatch) ResolveReference(refPointers *ReferencePointers) {
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
