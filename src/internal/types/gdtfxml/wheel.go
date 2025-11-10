package XMLTypes

type Wheel struct {
	Name             string `xml:",attr"`
	WheelSlots       []WheelSlot
	PrismFacets      []PrismFacet
	AnimationSystems []AnimationSystem
}

type WheelSlot struct {
	Name          string            `xml:",attr"`
	Color         *ColorCIE         `xml:",attr,omitempty"`
	Filter        *XMLNodeReference `xml:",attr,omitempty"` // ref to Physical/Filter
	MediaFileName *string           `xml:",attr,omitempty"`
}

type PrismFacet struct {
	Color    ColorCIE `xml:",attr"`
	Rotation Rotation `xml:",attr"`
}

type AnimationSystem struct {
	P1      []float32 `xml:",attr"`
	P2      []float32 `xml:",attr"`
	P3      []float32 `xml:",attr"`
	Radiues float32   `xml:",attr"`
}
