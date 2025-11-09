package XMLTypes

// TODO:
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
	// AxisList []Axis     `xml:"Axis"`
	// Children Geometries `xml:",any"`
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
	DMXOffset string `xml:",attr"`
	DMXBreak  uint   `xml:",attr"`
}

type Laser struct {
}

type WiringObject struct {
}

type Inventory struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Count    int    `xml:",attr"`
	Geometries
}

type Structure struct {
}

type Support struct {
}

type Magnet struct {
	Name     string `xml:",attr"`
	Model    string `xml:",attr"`
	Position Matrix `xml:",attr"`
	Geometries
}
