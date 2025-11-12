package Types

type GeometryNodeReference struct {
	Geometry          Geometry
	Axis              Axis
	FilterBeam        FilterBeam
	FilterColor       FilterColor
	FilterGobo        FilterGobo
	FilterShaper      FilterShaper
	Beam              Beam
	MediaServerLayer  MediaServerLayer
	MediaServerCamera MediaServerCamera
	MediaServerMaster MediaServerMaster
	Display           Display
	Laser             Laser
	GeometryReference GeometryReference
	WiringObject      WiringObject
	Inventory         Inventory
	Structure         Structure
	Support           Support
	Magnet            Magnet
}

type Geometries struct {
	GeometryList          []Geometry
	AxisList              []Axis
	FilterBeamList        []FilterBeam
	FilterColorList       []FilterColor
	FilterGoboList        []FilterGobo
	FilterShaperList      []FilterShaper
	BeamList              []Beam
	MediaServerLayerList  []MediaServerLayer
	MediaServerCameraList []MediaServerCamera
	MediaServerMasterList []MediaServerMaster
	DisplayList           []Display
	LaserList             []Laser
	GeometryReferenceList []GeometryReference
	WiringObjectList      []WiringObject
	InventoryList         []Inventory
	StructureList         []Structure
	SupportList           []Support
	MagnetList            []Magnet
}

type Geometry struct {
	Name     string
	Model    string
	Position Matrix
	Geometries
}

type Axis struct {
	Name     string
	Model    string
	Position Matrix
	Geometries
}

type FilterBeam struct {
	Name     string
	Model    string
	Position Matrix
	Geometries
}

type FilterColor struct {
	Name     string
	Model    string
	Position Matrix
	Geometries
}

type FilterGobo struct {
	Name     string
	Model    string
	Position Matrix
	Geometries
}

type FilterShaper struct {
	Name     string
	Model    string
	Position Matrix
	Geometries
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

type MediaServerLayer struct {
	Name     string
	Model    string
	Position Matrix
	Geometries
}

type MediaServerCamera struct {
	Name     string
	Model    string
	Position Matrix
	Geometries
}

type MediaServerMaster struct {
	Name     string
	Model    string
	Position Matrix
	Geometries
}

type Display struct {
	Name     string
	Model    string
	Position Matrix
	Texture  FileReference
	Geometries
}

type GeometryReference struct {
	Name        string
	Model       string
	Position    Matrix
	GeometryRef string // only top level geometries allowed to be referenced
	Breaks      []Break
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
	Protocols         []LaserProtocol
	Geometries
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
	PinPatches        []PinPatch
	Geometries
}

type PinPatch struct {
	ToWiringObject NodeReference[WiringObject]
	FromPin        int
	ToPin          int
}

type Inventory struct {
	Name     string
	Model    string
	Position Matrix
	Count    int
	Geometries
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

type Magnet struct {
	Name     string
	Model    string
	Position Matrix
	Geometries
}
