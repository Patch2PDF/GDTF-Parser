package XMLTypes

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

type Geometry struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Geometries
}

type Axis struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Geometries
}

type FilterBeam struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Geometries
}

type FilterColor struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Geometries
}

type FilterGobo struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Geometries
}

type FilterShaper struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Geometries
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

type MediaServerLayer struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Geometries
}

type MediaServerCamera struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Geometries
}

type MediaServerMaster struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Geometries
}

type Display struct {
	Name     string        `xml:",attr"`
	Model    string        `xml:",attr"`
	Position Matrix        `xml:",attr"`
	Texture  FileReference `xml:",attr"`
	Geometries
}

type GeometryReference struct {
	Name        string  `xml:",attr"`
	Model       string  `xml:",attr"`
	Position    Matrix  `xml:",attr"`
	GeometryRef string  `xml:"Geometry,attr"` // only top level geometries allowed to be referenced
	Breaks      []Break `xml:"Break"`
}

type Break struct {
	DMXOffset DMXAddress `xml:",attr"`
	DMXBreak  uint       `xml:",attr"`
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

type LaserProtocol struct {
	Name string `xml:",attr"`
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

type PinPatch struct {
	ToWiringObject XMLNodeReference `xml:",attr"`
	FromPin        int              `xml:",attr"`
	ToPin          int              `xml:",attr"`
}

type Inventory struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Count    int    `xml:",attr"`
	Geometries
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

type Magnet struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Geometries
}
