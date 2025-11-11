package XMLTypes

import Types "github.com/Patch2PDF/GDTF-Parser/types"

type Wheel struct {
	Name             string `xml:",attr"`
	WheelSlots       []WheelSlot
	PrismFacets      []PrismFacet
	AnimationSystems []AnimationSystem
}

func (attr Wheel) Parse() Types.Wheel {
	wheelSlots := ParseList(&attr.WheelSlots)
	prismFacets := ParseList(&attr.PrismFacets)
	animationSystems := ParseList(&attr.AnimationSystems)
	return Types.Wheel{
		Name:             attr.Name,
		WheelSlots:       wheelSlots,
		PrismFacets:      prismFacets,
		AnimationSystems: animationSystems,
	}
}

type WheelSlot struct {
	Name          string            `xml:",attr"`
	Color         *ColorCIE         `xml:",attr,omitempty"`
	Filter        *XMLNodeReference `xml:",attr,omitempty"` // ref to Physical/Filter
	MediaFileName *string           `xml:",attr,omitempty"`
}

func (attr WheelSlot) Parse() Types.WheelSlot {
	return Types.WheelSlot{
		Name:  attr.Name,
		Color: Types.ColorCIE(*attr.Color),
		Filter: Types.NodeReference[Types.Filter]{
			String: *attr.Filter,
		},
		MediaFileName: attr.MediaFileName,
	}
}

type PrismFacet struct {
	Color    ColorCIE `xml:",attr"`
	Rotation Rotation `xml:",attr"`
}

func (attr PrismFacet) Parse() Types.PrismFacet {
	return Types.PrismFacet{
		Color:    Types.ColorCIE(attr.Color),
		Rotation: Types.Rotation(attr.Rotation),
	}
}

type AnimationSystem struct {
	P1      []float32 `xml:",attr"`
	P2      []float32 `xml:",attr"`
	P3      []float32 `xml:",attr"`
	Radiues float32   `xml:",attr"`
}

func (attr AnimationSystem) Parse() Types.AnimationSystem {
	return Types.AnimationSystem(attr)
}
