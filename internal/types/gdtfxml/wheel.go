package XMLTypes

import Types "github.com/Patch2PDF/GDTF-Parser/pkg/types"

type Wheel struct {
	Name       string      `xml:",attr"`
	WheelSlots []WheelSlot `xml:"Slot"`
}

func (attr Wheel) Parse() Types.Wheel {
	return Types.Wheel{
		Name:       attr.Name,
		WheelSlots: ParseList(&attr.WheelSlots),
	}
}

type WheelSlot struct {
	Name             string            `xml:",attr"`
	Color            *ColorCIE         `xml:",attr,omitempty"`
	Filter           *XMLNodeReference `xml:",attr,omitempty"` // ref to Physical/Filter
	MediaFileName    *string           `xml:",attr,omitempty"`
	PrismFacets      []PrismFacet      `xml:"Facet"`
	AnimationSystems []AnimationSystem `xml:"AnimationSystem"`
}

func (attr WheelSlot) Parse() Types.WheelSlot {
	filter := ""
	if attr.Filter != nil {
		filter = *attr.Filter
	}
	return Types.WheelSlot{
		Name:  attr.Name,
		Color: Types.ColorCIE(*attr.Color),
		Filter: Types.NodeReference[Types.Filter]{
			String: filter,
		},
		MediaFileName:    attr.MediaFileName,
		PrismFacets:      ParseList(&attr.PrismFacets),
		AnimationSystems: ParseList(&attr.AnimationSystems),
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
