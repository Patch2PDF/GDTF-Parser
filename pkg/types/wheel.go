package Types

type Wheel struct {
	Name       string
	WheelSlots []*WheelSlot
}

func (obj *Wheel) CreateReferencePointer() {
	refPointers.Wheels[obj.Name] = obj
	for _, element := range obj.WheelSlots {
		refPointers.WheelSlots[obj.Name+"."+element.Name] = element
	}
}

func (obj *Wheel) ResolveReference() {
	ResolveReferences(&obj.WheelSlots)
}

type WheelSlot struct {
	Name             string
	Color            ColorCIE
	Filter           NodeReference[Filter] // ref to Physical/Filter
	MediaFileName    *string
	PrismFacets      []*PrismFacet
	AnimationSystems []*AnimationSystem
}

func (obj *WheelSlot) ResolveReference() {
	obj.Filter.Ptr = refPointers.Filters[obj.Filter.String]
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
