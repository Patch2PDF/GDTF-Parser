package XMLTypes

type XMLWheel struct {
	Name             string `xml:",attr"`
	WheelSlots       []XMLWheelSlot
	PrismFacets      []XMLPrismFacet
	AnimationSystems []XMLAnimationSystem
}

type XMLWheelSlot struct {
	Name          string            `xml:",attr"`
	Color         *ColorCIE         `xml:",attr,omitempty"`
	Filter        *XMLNodeReference `xml:",attr,omitempty"` // ref to Physical/Filter
	MediaFileName *string           `xml:",attr,omitempty"`
}

type XMLPrismFacet struct {
	Color    ColorCIE `xml:",attr"`
	Rotation Rotation `xml:",attr"`
}

type XMLAnimationSystem struct {
	P1      []float32 `xml:",attr"`
	P2      []float32 `xml:",attr"`
	P3      []float32 `xml:",attr"`
	Radiues float32   `xml:",attr"`
}
