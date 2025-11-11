package Types

type Wheel struct {
	Name             string
	WheelSlots       []WheelSlot
	PrismFacets      []PrismFacet
	AnimationSystems []AnimationSystem
}

type WheelSlot struct {
	Name          string
	Color         ColorCIE
	Filter        NodeReference[Filter] // ref to Physical/Filter
	MediaFileName *string
}

type PrismFacet struct {
	Color    ColorCIE
	Rotation Rotation
}

type AnimationSystem struct {
	P1      []float32
	P2      []float32
	P3      []float32
	Radiues float32
}
