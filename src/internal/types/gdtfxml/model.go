package XMLTypes

import Types "github.com/Patch2PDF/GDTF-Parser/types"

type Model struct {
	Name            string  `xml:",attr"`
	Length          float32 `xml:",attr"`
	Width           float32 `xml:",attr"`
	Height          float32 `xml:",attr"`
	PrimitiveType   string  `xml:",attr"` // enum
	File            *string `xml:",attr"`
	SVGOffsetX      float32 `xml:",attr"`
	SVGOffsetY      float32 `xml:",attr"`
	SVGSideOffsetX  float32 `xml:",attr"`
	SVGSideOffsetY  float32 `xml:",attr"`
	SVGFrontOffsetX float32 `xml:",attr"`
	SVGFrontOffsetY float32 `xml:",attr"`
}

func (model Model) Parse() Types.Model {
	return Types.Model(model) // for now, potentially need to adjust when parsing 3d model data
}
