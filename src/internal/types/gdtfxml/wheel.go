package XMLTypes

type XMLWheel struct {
	Name             string `xml:",attr"`
	WheelSlots       []XMLWheelSlot
	PrismFacets      []XMLPrismFacet
	AnimationSystems []XMLAnimationSystem
}

// TODO:
type XMLWheelSlot struct {
	Name          string `xml:",attr"`
	Color         string `xml:",attr"`
	Filter        string `xml:",attr"` // TODO:
	MediaFileName string `xml:",attr"`
}

type XMLPrismFacet struct {
	Color    string `xml:",attr"`
	Rotation string `xml:",attr"`
}

type XMLAnimationSystem struct {
	P1      []float32 `xml:",attr"`
	P2      []float32 `xml:",attr"`
	P3      []float32 `xml:",attr"`
	Radiues float32   `xml:",attr"`
}
