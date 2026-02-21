package Types

import (
	"reflect"
	"strings"

	"github.com/Patch2PDF/GDTF-Mesh-Reader/v2/pkg/MeshTypes"
)

type GeometryType int

const (
	GeometryTypeGeometry = iota
	GeometryTypeAxis
	GeometryTypeFilterBeam
	GeometryTypeFilterColor
	GeometryTypeFilterGobo
	GeometryTypeFilterShaper
	GeometryTypeBeam
	GeometryTypeMediaServerLayer
	GeometryTypeMediaServerCamera
	GeometryTypeMediaServerMaster
	GeometryTypeDisplay
	GeometryTypeGeometryReference
	GeometryTypeLaser
	GeometryTypeWiringObject
	GeometryTypeInventory
	GeometryTypeStructure
	GeometryTypeSupport
	GeometryTypeMagnet
)

type GeometryNodeReference struct {
	Ptr  GeometryModel
	Type string
}

type GeometryModel interface {
	GetName() string
	GetNamePtr() *string
	SetName(string)

	GetModel() NodeReference[Model]
	GetModelPtr() *NodeReference[Model]
	SetModel(NodeReference[Model])

	GetPosition() Matrix
	GetPositionPtr() *Matrix
	SetPosition(Matrix)

	GenerateMesh(parentTransformation MeshTypes.Matrix, modelList []MeshModel) []MeshModel
}

type GeometryModelWithChildGeometries interface {
	GetName() string
	GetNamePtr() *string
	SetName(string)

	GetModel() NodeReference[Model]
	GetModelPtr() *NodeReference[Model]
	SetModel(NodeReference[Model])

	GetPosition() Matrix
	GetPositionPtr() *Matrix
	SetPosition(Matrix)

	GenerateMesh(parentTransformation MeshTypes.Matrix, modelList []MeshModel) []MeshModel

	GetGeometries() Geometries
	GetGeometriesPtr() *Geometries
	SetGeometries(Geometries)
}

func CreateGeometryWithChildsReferencePointer(obj GeometryModelWithChildGeometries, refPointers *ReferencePointers, parentPrefix string) {
	newParentPrefix := strings.Trim(parentPrefix+"."+obj.GetName(), ".")
	refPointers.Geometries[newParentPrefix] = &GeometryNodeReference{
		Ptr:  obj,
		Type: reflect.TypeOf(obj).String(),
	}
	obj.GetGeometriesPtr().CreateGeometryReferencePointer(refPointers, newParentPrefix)
}

func ResolveGeometryWithChildsReference(obj GeometryModelWithChildGeometries, refPointers *ReferencePointers) {
	obj.GetModelPtr().Ptr = refPointers.Models[obj.GetModel().String]
	obj.GetGeometriesPtr().ResolveReference(refPointers)
}

// type MeshGenerator interface {
// 	GenerateMesh(parentTransformation MeshTypes.Matrix, modelList []MeshModel)
// }

func GenerateMesh(obj GeometryModel, parentTransformation MeshTypes.Matrix, modelList []MeshModel, geometryType GeometryType) ([]MeshModel, MeshTypes.Matrix) {
	var mesh MeshTypes.Mesh
	localTransformation := obj.GetPosition().toMeshMatrix()
	transformation := parentTransformation.Mul(localTransformation)
	if obj.GetModel().Ptr != nil && obj.GetModel().Ptr.Mesh != nil {
		mesh = obj.GetModel().Ptr.Mesh.Copy()
		mesh.RotateAndTranslate(transformation)
	}
	modelList = append(modelList, MeshModel{
		Mesh:         mesh,
		GeometryType: geometryType,
		GeometryPtr:  obj,
	})
	return modelList, transformation
}

func GenerateMeshWithChildren(obj GeometryModelWithChildGeometries, parentTransformation MeshTypes.Matrix, modelList []MeshModel, geometryType GeometryType) []MeshModel {
	modelList, transformation := GenerateMesh(obj, parentTransformation, modelList, geometryType)
	modelList = obj.GetGeometriesPtr().GenerateMesh(transformation, modelList)
	return modelList
}

func GenerateMeshes[T GeometryModel](source *[]T, parentTransformation MeshTypes.Matrix, modelList []MeshModel) []MeshModel {
	if source == nil {
		return modelList
	}
	for i := range *source {
		modelList = (*source)[i].GenerateMesh(parentTransformation, modelList)
	}
	return modelList
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

func (obj *Geometries) GenerateMesh(parentTransformation MeshTypes.Matrix, modelList []MeshModel) []MeshModel {
	modelList = GenerateMeshes(&obj.GeometryList, parentTransformation, modelList)
	modelList = GenerateMeshes(&obj.AxisList, parentTransformation, modelList)
	modelList = GenerateMeshes(&obj.FilterBeamList, parentTransformation, modelList)
	modelList = GenerateMeshes(&obj.FilterColorList, parentTransformation, modelList)
	modelList = GenerateMeshes(&obj.FilterGoboList, parentTransformation, modelList)
	modelList = GenerateMeshes(&obj.FilterShaperList, parentTransformation, modelList)
	modelList = GenerateMeshes(&obj.BeamList, parentTransformation, modelList)
	modelList = GenerateMeshes(&obj.MediaServerLayerList, parentTransformation, modelList)
	modelList = GenerateMeshes(&obj.MediaServerCameraList, parentTransformation, modelList)
	modelList = GenerateMeshes(&obj.MediaServerMasterList, parentTransformation, modelList)
	modelList = GenerateMeshes(&obj.DisplayList, parentTransformation, modelList)
	modelList = GenerateMeshes(&obj.LaserList, parentTransformation, modelList)
	modelList = GenerateMeshes(&obj.GeometryReferenceList, parentTransformation, modelList)
	modelList = GenerateMeshes(&obj.WiringObjectList, parentTransformation, modelList)
	modelList = GenerateMeshes(&obj.InventoryList, parentTransformation, modelList)
	modelList = GenerateMeshes(&obj.StructureList, parentTransformation, modelList)
	modelList = GenerateMeshes(&obj.SupportList, parentTransformation, modelList)
	modelList = GenerateMeshes(&obj.MagnetList, parentTransformation, modelList)
	return modelList
}

type Geometry struct {
	Name     string
	Model    NodeReference[Model]
	Position Matrix
	Geometries
}

func (obj *Geometry) GetGeometries() Geometries {
	return obj.Geometries
}

func (obj *Geometry) GetGeometriesPtr() *Geometries {
	return &obj.Geometries
}

func (obj *Geometry) GetModel() NodeReference[Model] {
	return obj.Model
}

func (obj *Geometry) GetModelPtr() *NodeReference[Model] {
	return &obj.Model
}

func (obj *Geometry) GetName() string {
	return obj.Name
}

func (obj *Geometry) GetNamePtr() *string {
	return &obj.Name
}

func (obj *Geometry) GetPosition() Matrix {
	return obj.Position
}

func (obj *Geometry) GetPositionPtr() *Matrix {
	return &obj.Position
}

func (obj *Geometry) SetGeometries(geometries Geometries) {
	obj.Geometries = geometries
}

func (obj *Geometry) SetModel(model NodeReference[Model]) {
	obj.Model = model
}

func (obj *Geometry) SetName(name string) {
	obj.Name = name
}

func (obj *Geometry) SetPosition(matrix Matrix) {
	obj.Position = matrix
}

func (obj *Geometry) ResolveReference(refPointers *ReferencePointers) {
	ResolveGeometryWithChildsReference(obj, refPointers)
}

func (obj *Geometry) CreateGeometryReferencePointer(refPointers *ReferencePointers, parentPrefix string) {
	CreateGeometryWithChildsReferencePointer(obj, refPointers, parentPrefix)
}

func (obj *Geometry) GenerateMesh(parentTransformation MeshTypes.Matrix, modelList []MeshModel) []MeshModel {
	modelList = GenerateMeshWithChildren(obj, parentTransformation, modelList, GeometryTypeGeometry)
	return modelList
}

type Axis struct {
	Name     string
	Model    NodeReference[Model]
	Position Matrix
	Geometries
}

func (obj *Axis) GetGeometries() Geometries {
	return obj.Geometries
}

func (obj *Axis) GetGeometriesPtr() *Geometries {
	return &obj.Geometries
}

func (obj *Axis) GetModel() NodeReference[Model] {
	return obj.Model
}

func (obj *Axis) GetModelPtr() *NodeReference[Model] {
	return &obj.Model
}

func (obj *Axis) GetName() string {
	return obj.Name
}

func (obj *Axis) GetNamePtr() *string {
	return &obj.Name
}

func (obj *Axis) GetPosition() Matrix {
	return obj.Position
}

func (obj *Axis) GetPositionPtr() *Matrix {
	return &obj.Position
}

func (obj *Axis) SetGeometries(geometries Geometries) {
	obj.Geometries = geometries
}

func (obj *Axis) SetModel(model NodeReference[Model]) {
	obj.Model = model
}

func (obj *Axis) SetName(name string) {
	obj.Name = name
}

func (obj *Axis) SetPosition(matrix Matrix) {
	obj.Position = matrix
}

func (obj *Axis) ResolveReference(refPointers *ReferencePointers) {
	ResolveGeometryWithChildsReference(obj, refPointers)
}

func (obj *Axis) CreateGeometryReferencePointer(refPointers *ReferencePointers, parentPrefix string) {
	CreateGeometryWithChildsReferencePointer(obj, refPointers, parentPrefix)
}

func (obj *Axis) GenerateMesh(parentTransformation MeshTypes.Matrix, modelList []MeshModel) []MeshModel {
	modelList = GenerateMeshWithChildren(obj, parentTransformation, modelList, GeometryTypeAxis)
	return modelList
}

type FilterBeam struct {
	Name     string
	Model    NodeReference[Model]
	Position Matrix
	Geometries
}

func (obj *FilterBeam) GetGeometries() Geometries {
	return obj.Geometries
}

func (obj *FilterBeam) GetGeometriesPtr() *Geometries {
	return &obj.Geometries
}

func (obj *FilterBeam) GetModel() NodeReference[Model] {
	return obj.Model
}

func (obj *FilterBeam) GetModelPtr() *NodeReference[Model] {
	return &obj.Model
}

func (obj *FilterBeam) GetName() string {
	return obj.Name
}

func (obj *FilterBeam) GetNamePtr() *string {
	return &obj.Name
}

func (obj *FilterBeam) GetPosition() Matrix {
	return obj.Position
}

func (obj *FilterBeam) GetPositionPtr() *Matrix {
	return &obj.Position
}

func (obj *FilterBeam) SetGeometries(geometries Geometries) {
	obj.Geometries = geometries
}

func (obj *FilterBeam) SetModel(model NodeReference[Model]) {
	obj.Model = model
}

func (obj *FilterBeam) SetName(name string) {
	obj.Name = name
}

func (obj *FilterBeam) SetPosition(matrix Matrix) {
	obj.Position = matrix
}

func (obj *FilterBeam) ResolveReference(refPointers *ReferencePointers) {
	ResolveGeometryWithChildsReference(obj, refPointers)
}

func (obj *FilterBeam) CreateGeometryReferencePointer(refPointers *ReferencePointers, parentPrefix string) {
	CreateGeometryWithChildsReferencePointer(obj, refPointers, parentPrefix)
}

func (obj *FilterBeam) GenerateMesh(parentTransformation MeshTypes.Matrix, modelList []MeshModel) []MeshModel {
	modelList = GenerateMeshWithChildren(obj, parentTransformation, modelList, GeometryTypeFilterBeam)
	return modelList
}

type FilterColor struct {
	Name     string
	Model    NodeReference[Model]
	Position Matrix
	Geometries
}

func (obj *FilterColor) GetGeometries() Geometries {
	return obj.Geometries
}

func (obj *FilterColor) GetGeometriesPtr() *Geometries {
	return &obj.Geometries
}

func (obj *FilterColor) GetModel() NodeReference[Model] {
	return obj.Model
}

func (obj *FilterColor) GetModelPtr() *NodeReference[Model] {
	return &obj.Model
}

func (obj *FilterColor) GetName() string {
	return obj.Name
}

func (obj *FilterColor) GetNamePtr() *string {
	return &obj.Name
}

func (obj *FilterColor) GetPosition() Matrix {
	return obj.Position
}

func (obj *FilterColor) GetPositionPtr() *Matrix {
	return &obj.Position
}

func (obj *FilterColor) SetGeometries(geometries Geometries) {
	obj.Geometries = geometries
}

func (obj *FilterColor) SetModel(model NodeReference[Model]) {
	obj.Model = model
}

func (obj *FilterColor) SetName(name string) {
	obj.Name = name
}

func (obj *FilterColor) SetPosition(matrix Matrix) {
	obj.Position = matrix
}

func (obj *FilterColor) ResolveReference(refPointers *ReferencePointers) {
	ResolveGeometryWithChildsReference(obj, refPointers)
}

func (obj *FilterColor) CreateGeometryReferencePointer(refPointers *ReferencePointers, parentPrefix string) {
	CreateGeometryWithChildsReferencePointer(obj, refPointers, parentPrefix)
}

func (obj *FilterColor) GenerateMesh(parentTransformation MeshTypes.Matrix, modelList []MeshModel) []MeshModel {
	modelList = GenerateMeshWithChildren(obj, parentTransformation, modelList, GeometryTypeFilterColor)
	return modelList
}

type FilterGobo struct {
	Name     string
	Model    NodeReference[Model]
	Position Matrix
	Geometries
}

func (obj *FilterGobo) GetGeometries() Geometries {
	return obj.Geometries
}

func (obj *FilterGobo) GetGeometriesPtr() *Geometries {
	return &obj.Geometries
}

func (obj *FilterGobo) GetModel() NodeReference[Model] {
	return obj.Model
}

func (obj *FilterGobo) GetModelPtr() *NodeReference[Model] {
	return &obj.Model
}

func (obj *FilterGobo) GetName() string {
	return obj.Name
}

func (obj *FilterGobo) GetNamePtr() *string {
	return &obj.Name
}

func (obj *FilterGobo) GetPosition() Matrix {
	return obj.Position
}

func (obj *FilterGobo) GetPositionPtr() *Matrix {
	return &obj.Position
}

func (obj *FilterGobo) SetGeometries(geometries Geometries) {
	obj.Geometries = geometries
}

func (obj *FilterGobo) SetModel(model NodeReference[Model]) {
	obj.Model = model
}

func (obj *FilterGobo) SetName(name string) {
	obj.Name = name
}

func (obj *FilterGobo) SetPosition(matrix Matrix) {
	obj.Position = matrix
}

func (obj *FilterGobo) ResolveReference(refPointers *ReferencePointers) {
	ResolveGeometryWithChildsReference(obj, refPointers)
}

func (obj *FilterGobo) CreateGeometryReferencePointer(refPointers *ReferencePointers, parentPrefix string) {
	CreateGeometryWithChildsReferencePointer(obj, refPointers, parentPrefix)
}

func (obj *FilterGobo) GenerateMesh(parentTransformation MeshTypes.Matrix, modelList []MeshModel) []MeshModel {
	modelList = GenerateMeshWithChildren(obj, parentTransformation, modelList, GeometryTypeFilterGobo)
	return modelList
}

type FilterShaper struct {
	Name     string
	Model    NodeReference[Model]
	Position Matrix
	Geometries
}

func (obj *FilterShaper) GetGeometries() Geometries {
	return obj.Geometries
}

func (obj *FilterShaper) GetGeometriesPtr() *Geometries {
	return &obj.Geometries
}

func (obj *FilterShaper) GetModel() NodeReference[Model] {
	return obj.Model
}

func (obj *FilterShaper) GetModelPtr() *NodeReference[Model] {
	return &obj.Model
}

func (obj *FilterShaper) GetName() string {
	return obj.Name
}

func (obj *FilterShaper) GetNamePtr() *string {
	return &obj.Name
}

func (obj *FilterShaper) GetPosition() Matrix {
	return obj.Position
}

func (obj *FilterShaper) GetPositionPtr() *Matrix {
	return &obj.Position
}

func (obj *FilterShaper) SetGeometries(geometries Geometries) {
	obj.Geometries = geometries
}

func (obj *FilterShaper) SetModel(model NodeReference[Model]) {
	obj.Model = model
}

func (obj *FilterShaper) SetName(name string) {
	obj.Name = name
}

func (obj *FilterShaper) SetPosition(matrix Matrix) {
	obj.Position = matrix
}

func (obj *FilterShaper) ResolveReference(refPointers *ReferencePointers) {
	ResolveGeometryWithChildsReference(obj, refPointers)
}

func (obj *FilterShaper) CreateGeometryReferencePointer(refPointers *ReferencePointers, parentPrefix string) {
	CreateGeometryWithChildsReferencePointer(obj, refPointers, parentPrefix)
}

func (obj *FilterShaper) GenerateMesh(parentTransformation MeshTypes.Matrix, modelList []MeshModel) []MeshModel {
	modelList = GenerateMeshWithChildren(obj, parentTransformation, modelList, GeometryTypeFilterShaper)
	return modelList
}

type Beam struct {
	Name     string
	Model    NodeReference[Model]
	Position Matrix
	Geometries
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

func (obj *Beam) GetGeometries() Geometries {
	return obj.Geometries
}

func (obj *Beam) GetGeometriesPtr() *Geometries {
	return &obj.Geometries
}

func (obj *Beam) GetModel() NodeReference[Model] {
	return obj.Model
}

func (obj *Beam) GetModelPtr() *NodeReference[Model] {
	return &obj.Model
}

func (obj *Beam) GetName() string {
	return obj.Name
}

func (obj *Beam) GetNamePtr() *string {
	return &obj.Name
}

func (obj *Beam) GetPosition() Matrix {
	return obj.Position
}

func (obj *Beam) GetPositionPtr() *Matrix {
	return &obj.Position
}

func (obj *Beam) SetGeometries(geometries Geometries) {
	obj.Geometries = geometries
}

func (obj *Beam) SetModel(model NodeReference[Model]) {
	obj.Model = model
}

func (obj *Beam) SetName(name string) {
	obj.Name = name
}

func (obj *Beam) SetPosition(matrix Matrix) {
	obj.Position = matrix
}

func (obj *Beam) ResolveReference(refPointers *ReferencePointers) {
	obj.EmitterSpectrum.Ptr = refPointers.Emitters[obj.EmitterSpectrum.String]
	ResolveGeometryWithChildsReference(obj, refPointers)
}

func (obj *Beam) CreateGeometryReferencePointer(refPointers *ReferencePointers, parentPrefix string) {
	CreateGeometryWithChildsReferencePointer(obj, refPointers, parentPrefix)
}

func (obj *Beam) GenerateMesh(parentTransformation MeshTypes.Matrix, modelList []MeshModel) []MeshModel {
	modelList = GenerateMeshWithChildren(obj, parentTransformation, modelList, GeometryTypeBeam)
	return modelList
}

type MediaServerLayer struct {
	Name     string
	Model    NodeReference[Model]
	Position Matrix
	Geometries
}

func (obj *MediaServerLayer) GetGeometries() Geometries {
	return obj.Geometries
}

func (obj *MediaServerLayer) GetGeometriesPtr() *Geometries {
	return &obj.Geometries
}

func (obj *MediaServerLayer) GetModel() NodeReference[Model] {
	return obj.Model
}

func (obj *MediaServerLayer) GetModelPtr() *NodeReference[Model] {
	return &obj.Model
}

func (obj *MediaServerLayer) GetName() string {
	return obj.Name
}

func (obj *MediaServerLayer) GetNamePtr() *string {
	return &obj.Name
}

func (obj *MediaServerLayer) GetPosition() Matrix {
	return obj.Position
}

func (obj *MediaServerLayer) GetPositionPtr() *Matrix {
	return &obj.Position
}

func (obj *MediaServerLayer) SetGeometries(geometries Geometries) {
	obj.Geometries = geometries
}

func (obj *MediaServerLayer) SetModel(model NodeReference[Model]) {
	obj.Model = model
}

func (obj *MediaServerLayer) SetName(name string) {
	obj.Name = name
}

func (obj *MediaServerLayer) SetPosition(matrix Matrix) {
	obj.Position = matrix
}

func (obj *MediaServerLayer) ResolveReference(refPointers *ReferencePointers) {
	ResolveGeometryWithChildsReference(obj, refPointers)
}

func (obj *MediaServerLayer) CreateGeometryReferencePointer(refPointers *ReferencePointers, parentPrefix string) {
	CreateGeometryWithChildsReferencePointer(obj, refPointers, parentPrefix)
}

func (obj *MediaServerLayer) GenerateMesh(parentTransformation MeshTypes.Matrix, modelList []MeshModel) []MeshModel {
	modelList = GenerateMeshWithChildren(obj, parentTransformation, modelList, GeometryTypeMediaServerLayer)
	return modelList
}

type MediaServerCamera struct {
	Name     string
	Model    NodeReference[Model]
	Position Matrix
	Geometries
}

func (obj *MediaServerCamera) GetGeometries() Geometries {
	return obj.Geometries
}

func (obj *MediaServerCamera) GetGeometriesPtr() *Geometries {
	return &obj.Geometries
}

func (obj *MediaServerCamera) GetModel() NodeReference[Model] {
	return obj.Model
}

func (obj *MediaServerCamera) GetModelPtr() *NodeReference[Model] {
	return &obj.Model
}

func (obj *MediaServerCamera) GetName() string {
	return obj.Name
}

func (obj *MediaServerCamera) GetNamePtr() *string {
	return &obj.Name
}

func (obj *MediaServerCamera) GetPosition() Matrix {
	return obj.Position
}

func (obj *MediaServerCamera) GetPositionPtr() *Matrix {
	return &obj.Position
}

func (obj *MediaServerCamera) SetGeometries(geometries Geometries) {
	obj.Geometries = geometries
}

func (obj *MediaServerCamera) SetModel(model NodeReference[Model]) {
	obj.Model = model
}

func (obj *MediaServerCamera) SetName(name string) {
	obj.Name = name
}

func (obj *MediaServerCamera) SetPosition(matrix Matrix) {
	obj.Position = matrix
}

func (obj *MediaServerCamera) ResolveReference(refPointers *ReferencePointers) {
	ResolveGeometryWithChildsReference(obj, refPointers)
}

func (obj *MediaServerCamera) CreateGeometryReferencePointer(refPointers *ReferencePointers, parentPrefix string) {
	CreateGeometryWithChildsReferencePointer(obj, refPointers, parentPrefix)
}

func (obj *MediaServerCamera) GenerateMesh(parentTransformation MeshTypes.Matrix, modelList []MeshModel) []MeshModel {
	modelList = GenerateMeshWithChildren(obj, parentTransformation, modelList, GeometryTypeMediaServerCamera)
	return modelList
}

type MediaServerMaster struct {
	Name     string
	Model    NodeReference[Model]
	Position Matrix
	Geometries
}

func (obj *MediaServerMaster) GetGeometries() Geometries {
	return obj.Geometries
}

func (obj *MediaServerMaster) GetGeometriesPtr() *Geometries {
	return &obj.Geometries
}

func (obj *MediaServerMaster) GetModel() NodeReference[Model] {
	return obj.Model
}

func (obj *MediaServerMaster) GetModelPtr() *NodeReference[Model] {
	return &obj.Model
}

func (obj *MediaServerMaster) GetName() string {
	return obj.Name
}

func (obj *MediaServerMaster) GetNamePtr() *string {
	return &obj.Name
}

func (obj *MediaServerMaster) GetPosition() Matrix {
	return obj.Position
}

func (obj *MediaServerMaster) GetPositionPtr() *Matrix {
	return &obj.Position
}

func (obj *MediaServerMaster) SetGeometries(geometries Geometries) {
	obj.Geometries = geometries
}

func (obj *MediaServerMaster) SetModel(model NodeReference[Model]) {
	obj.Model = model
}

func (obj *MediaServerMaster) SetName(name string) {
	obj.Name = name
}

func (obj *MediaServerMaster) SetPosition(matrix Matrix) {
	obj.Position = matrix
}

func (obj *MediaServerMaster) ResolveReference(refPointers *ReferencePointers) {
	ResolveGeometryWithChildsReference(obj, refPointers)
}

func (obj *MediaServerMaster) CreateGeometryReferencePointer(refPointers *ReferencePointers, parentPrefix string) {
	CreateGeometryWithChildsReferencePointer(obj, refPointers, parentPrefix)
}

func (obj *MediaServerMaster) GenerateMesh(parentTransformation MeshTypes.Matrix, modelList []MeshModel) []MeshModel {
	modelList = GenerateMeshWithChildren(obj, parentTransformation, modelList, GeometryTypeMediaServerMaster)
	return modelList
}

type Display struct {
	Name     string
	Model    NodeReference[Model]
	Position Matrix
	Geometries
	Texture FileReference
}

func (obj *Display) GetGeometries() Geometries {
	return obj.Geometries
}

func (obj *Display) GetGeometriesPtr() *Geometries {
	return &obj.Geometries
}

func (obj *Display) GetModel() NodeReference[Model] {
	return obj.Model
}

func (obj *Display) GetModelPtr() *NodeReference[Model] {
	return &obj.Model
}

func (obj *Display) GetName() string {
	return obj.Name
}

func (obj *Display) GetNamePtr() *string {
	return &obj.Name
}

func (obj *Display) GetPosition() Matrix {
	return obj.Position
}

func (obj *Display) GetPositionPtr() *Matrix {
	return &obj.Position
}

func (obj *Display) SetGeometries(geometries Geometries) {
	obj.Geometries = geometries
}

func (obj *Display) SetModel(model NodeReference[Model]) {
	obj.Model = model
}

func (obj *Display) SetName(name string) {
	obj.Name = name
}

func (obj *Display) SetPosition(matrix Matrix) {
	obj.Position = matrix
}

func (obj *Display) ResolveReference(refPointers *ReferencePointers) {
	ResolveGeometryWithChildsReference(obj, refPointers)
}

func (obj *Display) CreateGeometryReferencePointer(refPointers *ReferencePointers, parentPrefix string) {
	CreateGeometryWithChildsReferencePointer(obj, refPointers, parentPrefix)
}

func (obj *Display) GenerateMesh(parentTransformation MeshTypes.Matrix, modelList []MeshModel) []MeshModel {
	modelList = GenerateMeshWithChildren(obj, parentTransformation, modelList, GeometryTypeDisplay)
	return modelList
}

type GeometryReference struct {
	Name        string
	Model       NodeReference[Model]
	Position    Matrix
	GeometryRef NodeReference[GeometryNodeReference] // only top level geometries allowed to be referenced
	Breaks      []*Break
	// do we need to link a parent for this Geometry Type?
}

func (obj *GeometryReference) GetModel() NodeReference[Model] {
	return obj.Model
}

func (obj *GeometryReference) GetModelPtr() *NodeReference[Model] {
	return &obj.Model
}

func (obj *GeometryReference) GetName() string {
	return obj.Name
}

func (obj *GeometryReference) GetNamePtr() *string {
	return &obj.Name
}

func (obj *GeometryReference) GetPosition() Matrix {
	return obj.Position
}

func (obj *GeometryReference) GetPositionPtr() *Matrix {
	return &obj.Position
}

func (obj *GeometryReference) SetModel(model NodeReference[Model]) {
	obj.Model = model
}

func (obj *GeometryReference) SetName(name string) {
	obj.Name = name
}

func (obj *GeometryReference) SetPosition(matrix Matrix) {
	obj.Position = matrix
}

func (obj *GeometryReference) ResolveReference(refPointers *ReferencePointers) {
	obj.GeometryRef.Ptr = refPointers.Geometries[obj.GeometryRef.String]
	obj.Model.Ptr = refPointers.Models[obj.Model.String]
}

// func (obj *GeometryReference) CreateGeometryReferencePointer(refPointers *ReferencePointers, parentPrefix string) {
// 	CreateGeometryWithChildsReferencePointer(obj, refPointers, parentPrefix)
// }

func (obj *GeometryReference) GenerateMesh(parentTransformation MeshTypes.Matrix, modelList []MeshModel) []MeshModel {
	// var mesh *MeshTypes.Mesh
	localTransformation := obj.Position.toMeshMatrix()
	transformation := parentTransformation.Mul(localTransformation)
	// if own model, replace parent mesh
	modelList, _ = GenerateMesh(obj, transformation, modelList, GeometryTypeGeometryReference)
	return modelList
	// if obj.Model.Ptr != nil {
	// 	// temp := obj.Model.Ptr.Mesh.Copy()
	// 	// mesh = &temp
	// 	// geometries := obj.GeometryRef.Ptr.Ptr
	// } else {
	// 	ptr := obj.GeometryRef.Ptr.Ptr
	// 	mesh := ptr.GenerateMesh(transformation)
	// }
}

type Break struct {
	DMXOffset DMXAddress
	DMXBreak  uint
}

type Laser struct {
	Name     string
	Model    NodeReference[Model]
	Position Matrix
	Geometries
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

func (obj *Laser) GetGeometries() Geometries {
	return obj.Geometries
}

func (obj *Laser) GetGeometriesPtr() *Geometries {
	return &obj.Geometries
}

func (obj *Laser) GetModel() NodeReference[Model] {
	return obj.Model
}

func (obj *Laser) GetModelPtr() *NodeReference[Model] {
	return &obj.Model
}

func (obj *Laser) GetName() string {
	return obj.Name
}

func (obj *Laser) GetNamePtr() *string {
	return &obj.Name
}

func (obj *Laser) GetPosition() Matrix {
	return obj.Position
}

func (obj *Laser) GetPositionPtr() *Matrix {
	return &obj.Position
}

func (obj *Laser) SetGeometries(geometries Geometries) {
	obj.Geometries = geometries
}

func (obj *Laser) SetModel(model NodeReference[Model]) {
	obj.Model = model
}

func (obj *Laser) SetName(name string) {
	obj.Name = name
}

func (obj *Laser) SetPosition(matrix Matrix) {
	obj.Position = matrix
}

func (obj *Laser) ResolveReference(refPointers *ReferencePointers) {
	obj.Emitter.Ptr = refPointers.Emitters[obj.Emitter.String]
	ResolveGeometryWithChildsReference(obj, refPointers)
}

func (obj *Laser) CreateGeometryReferencePointer(refPointers *ReferencePointers, parentPrefix string) {
	CreateGeometryWithChildsReferencePointer(obj, refPointers, parentPrefix)
}

func (obj *Laser) GenerateMesh(parentTransformation MeshTypes.Matrix, modelList []MeshModel) []MeshModel {
	modelList = GenerateMeshWithChildren(obj, parentTransformation, modelList, GeometryTypeLaser)
	return modelList
}

type LaserProtocol struct {
	Name string
}

type WiringObject struct {
	Name     string
	Model    NodeReference[Model]
	Position Matrix
	Geometries
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
	ResolveGeometryWithChildsReference(obj, refPointers)
}

func (obj *WiringObject) CreateGeometryReferencePointer(refPointers *ReferencePointers, parentPrefix string) {
	CreateGeometryWithChildsReferencePointer(obj, refPointers, parentPrefix)
}

func (obj *WiringObject) GetGeometries() Geometries {
	return obj.Geometries
}

func (obj *WiringObject) GetGeometriesPtr() *Geometries {
	return &obj.Geometries
}

func (obj *WiringObject) GetModel() NodeReference[Model] {
	return obj.Model
}

func (obj *WiringObject) GetModelPtr() *NodeReference[Model] {
	return &obj.Model
}

func (obj *WiringObject) GetName() string {
	return obj.Name
}

func (obj *WiringObject) GetNamePtr() *string {
	return &obj.Name
}

func (obj *WiringObject) GetPosition() Matrix {
	return obj.Position
}

func (obj *WiringObject) GetPositionPtr() *Matrix {
	return &obj.Position
}

func (obj *WiringObject) SetGeometries(geometries Geometries) {
	obj.Geometries = geometries
}

func (obj *WiringObject) SetModel(model NodeReference[Model]) {
	obj.Model = model
}

func (obj *WiringObject) SetName(name string) {
	obj.Name = name
}

func (obj *WiringObject) SetPosition(matrix Matrix) {
	obj.Position = matrix
}

func (obj *WiringObject) GenerateMesh(parentTransformation MeshTypes.Matrix, modelList []MeshModel) []MeshModel {
	modelList = GenerateMeshWithChildren(obj, parentTransformation, modelList, GeometryTypeWiringObject)
	return modelList
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
	Name     string
	Model    NodeReference[Model]
	Position Matrix
	Geometries
	Count int
}

func (obj *Inventory) GetGeometries() Geometries {
	return obj.Geometries
}

func (obj *Inventory) GetGeometriesPtr() *Geometries {
	return &obj.Geometries
}

func (obj *Inventory) GetModel() NodeReference[Model] {
	return obj.Model
}

func (obj *Inventory) GetModelPtr() *NodeReference[Model] {
	return &obj.Model
}

func (obj *Inventory) GetName() string {
	return obj.Name
}

func (obj *Inventory) GetNamePtr() *string {
	return &obj.Name
}

func (obj *Inventory) GetPosition() Matrix {
	return obj.Position
}

func (obj *Inventory) GetPositionPtr() *Matrix {
	return &obj.Position
}

func (obj *Inventory) SetGeometries(geometries Geometries) {
	obj.Geometries = geometries
}

func (obj *Inventory) SetModel(model NodeReference[Model]) {
	obj.Model = model
}

func (obj *Inventory) SetName(name string) {
	obj.Name = name
}

func (obj *Inventory) SetPosition(matrix Matrix) {
	obj.Position = matrix
}

func (obj *Inventory) ResolveReference(refPointers *ReferencePointers) {
	ResolveGeometryWithChildsReference(obj, refPointers)
}

func (obj *Inventory) CreateGeometryReferencePointer(refPointers *ReferencePointers, parentPrefix string) {
	CreateGeometryWithChildsReferencePointer(obj, refPointers, parentPrefix)
}

func (obj *Inventory) GenerateMesh(parentTransformation MeshTypes.Matrix, modelList []MeshModel) []MeshModel {
	modelList = GenerateMeshWithChildren(obj, parentTransformation, modelList, GeometryTypeInventory)
	return modelList
}

type Structure struct {
	Name     string
	Model    NodeReference[Model]
	Position Matrix
	Geometries
	LinkedGeometry            string // for now (analyse if this can be a NodeReference instead)
	StructureType             string // enum
	CrossSectionType          string // enum
	CrossSectionHeight        float32
	CrossSectionWallThickness float32
	TrussCrossSection         string
}

func (obj *Structure) GetGeometries() Geometries {
	return obj.Geometries
}

func (obj *Structure) GetGeometriesPtr() *Geometries {
	return &obj.Geometries
}

func (obj *Structure) GetModel() NodeReference[Model] {
	return obj.Model
}

func (obj *Structure) GetModelPtr() *NodeReference[Model] {
	return &obj.Model
}

func (obj *Structure) GetName() string {
	return obj.Name
}

func (obj *Structure) GetNamePtr() *string {
	return &obj.Name
}

func (obj *Structure) GetPosition() Matrix {
	return obj.Position
}

func (obj *Structure) GetPositionPtr() *Matrix {
	return &obj.Position
}

func (obj *Structure) SetGeometries(geometries Geometries) {
	obj.Geometries = geometries
}

func (obj *Structure) SetModel(model NodeReference[Model]) {
	obj.Model = model
}

func (obj *Structure) SetName(name string) {
	obj.Name = name
}

func (obj *Structure) SetPosition(matrix Matrix) {
	obj.Position = matrix
}

func (obj *Structure) ResolveReference(refPointers *ReferencePointers) {
	ResolveGeometryWithChildsReference(obj, refPointers)
}

func (obj *Structure) CreateGeometryReferencePointer(refPointers *ReferencePointers, parentPrefix string) {
	CreateGeometryWithChildsReferencePointer(obj, refPointers, parentPrefix)
}

func (obj *Structure) GenerateMesh(parentTransformation MeshTypes.Matrix, modelList []MeshModel) []MeshModel {
	modelList = GenerateMeshWithChildren(obj, parentTransformation, modelList, GeometryTypeStructure)
	return modelList
}

type Support struct {
	Name     string
	Model    NodeReference[Model]
	Position Matrix
	Geometries
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

func (obj *Support) GetGeometries() Geometries {
	return obj.Geometries
}

func (obj *Support) GetGeometriesPtr() *Geometries {
	return &obj.Geometries
}

func (obj *Support) GetModel() NodeReference[Model] {
	return obj.Model
}

func (obj *Support) GetModelPtr() *NodeReference[Model] {
	return &obj.Model
}

func (obj *Support) GetName() string {
	return obj.Name
}

func (obj *Support) GetNamePtr() *string {
	return &obj.Name
}

func (obj *Support) GetPosition() Matrix {
	return obj.Position
}

func (obj *Support) GetPositionPtr() *Matrix {
	return &obj.Position
}

func (obj *Support) SetGeometries(geometries Geometries) {
	obj.Geometries = geometries
}

func (obj *Support) SetModel(model NodeReference[Model]) {
	obj.Model = model
}

func (obj *Support) SetName(name string) {
	obj.Name = name
}

func (obj *Support) SetPosition(matrix Matrix) {
	obj.Position = matrix
}

func (obj *Support) ResolveReference(refPointers *ReferencePointers) {
	ResolveGeometryWithChildsReference(obj, refPointers)
}

func (obj *Support) CreateGeometryReferencePointer(refPointers *ReferencePointers, parentPrefix string) {
	CreateGeometryWithChildsReferencePointer(obj, refPointers, parentPrefix)
}

func (obj *Support) GenerateMesh(parentTransformation MeshTypes.Matrix, modelList []MeshModel) []MeshModel {
	modelList = GenerateMeshWithChildren(obj, parentTransformation, modelList, GeometryTypeSupport)
	return modelList
}

type Magnet struct {
	Name     string
	Model    NodeReference[Model]
	Position Matrix
	Geometries
}

func (obj *Magnet) GetGeometries() Geometries {
	return obj.Geometries
}

func (obj *Magnet) GetGeometriesPtr() *Geometries {
	return &obj.Geometries
}

func (obj *Magnet) GetModel() NodeReference[Model] {
	return obj.Model
}

func (obj *Magnet) GetModelPtr() *NodeReference[Model] {
	return &obj.Model
}

func (obj *Magnet) GetName() string {
	return obj.Name
}

func (obj *Magnet) GetNamePtr() *string {
	return &obj.Name
}

func (obj *Magnet) GetPosition() Matrix {
	return obj.Position
}

func (obj *Magnet) GetPositionPtr() *Matrix {
	return &obj.Position
}

func (obj *Magnet) SetGeometries(geometries Geometries) {
	obj.Geometries = geometries
}

func (obj *Magnet) SetModel(model NodeReference[Model]) {
	obj.Model = model
}

func (obj *Magnet) SetName(name string) {
	obj.Name = name
}

func (obj *Magnet) SetPosition(matrix Matrix) {
	obj.Position = matrix
}

func (obj *Magnet) ResolveReference(refPointers *ReferencePointers) {
	ResolveGeometryWithChildsReference(obj, refPointers)
}

func (obj *Magnet) CreateGeometryReferencePointer(refPointers *ReferencePointers, parentPrefix string) {
	CreateGeometryWithChildsReferencePointer(obj, refPointers, parentPrefix)
}

func (obj *Magnet) GenerateMesh(parentTransformation MeshTypes.Matrix, modelList []MeshModel) []MeshModel {
	modelList = GenerateMeshWithChildren(obj, parentTransformation, modelList, GeometryTypeMagnet)
	return modelList
}
