package gdtfxml_test

import (
	"testing"

	XMLTypes "github.com/Patch2PDF/GDTF-Parser/internal/types/gdtfxml"
)

func TestParseWheel(t *testing.T) {
	type wheelTest struct {
		Wheels []XMLTypes.Wheel `xml:"Wheel"`
	}

	xmlData := `
	<Wheels>
    <Wheel Name="ColorWheel">
        <Slot Name="Open" />
        <Slot Name="Red" Color="0.64,0.33,21.26" />
        <Slot Name="Green" Color="0.3,0.6,71.52" />
        <Slot Name="Blue" Color="0.15,0.06,7.22" />
        <Slot Name="Cyan" Color="0.2247,0.3287,78.74" />
        <Slot Name="Magenta" Color="0.3209,0.1542,28.48" />
        <Slot Name="Yellow" Color="0.4193,0.5053,92.78" />
    </Wheel>
    <Wheel Name="GoboWheel">
        <Slot Name="Open" />
        <Slot Name="Gobo 1" MediaFileName="Gobo1" />
        <Slot Name="Gobo 2" MediaFileName="Gobo2" />
        <Slot Name="Gobo 3" MediaFileName="Gobo3" />
        <Slot Name="Gobo 4" MediaFileName="Gobo4" />
        <Slot Name="Gobo 5" MediaFileName="Gobo5" />
    </Wheel>
		<Wheel Name="PrismWheel">
			<Slot Name="Open" />
			<Slot Name="Prism">
				<Facet Rotation="{0.6, 0.0, 0.0}{ 0.0, 0.6, 0.0}{ -0.12,  0.15, 1.0}"/>
				<Facet Rotation="{0.6, 0.0, 0.0}{ 0.0, 0.6, 0.0}{  0.17,  0.0,  1.0}"/>
				<Facet Rotation="{0.6, 0.0, 0.0}{ 0.0, 0.6, 0.0}{ -0.12, -0.15, 1.0}"/>
			</Slot>
		</Wheel>
	</Wheels>
	`

	gobo1 := "Gobo1"
	gobo2 := "Gobo2"
	gobo3 := "Gobo3"
	gobo4 := "Gobo4"
	gobo5 := "Gobo5"

	want := []XMLTypes.Wheel{
		{
			Name: "ColorWheel",
			WheelSlots: []XMLTypes.WheelSlot{
				{
					Name: "Open",
				},
				{
					Name: "Red",
					Color: &XMLTypes.ColorCIE{
						X:  0.64,
						Y:  0.33,
						Y2: 21.26,
					},
				},
				{
					Name: "Green",
					Color: &XMLTypes.ColorCIE{
						X:  0.3,
						Y:  0.6,
						Y2: 71.52,
					},
				},
				{
					Name: "Blue",
					Color: &XMLTypes.ColorCIE{
						X:  0.15,
						Y:  0.06,
						Y2: 7.22,
					},
				},
				{
					Name: "Cyan",
					Color: &XMLTypes.ColorCIE{
						X:  0.2247,
						Y:  0.3287,
						Y2: 78.74,
					},
				},
				{
					Name: "Magenta",
					Color: &XMLTypes.ColorCIE{
						X:  0.3209,
						Y:  0.1542,
						Y2: 28.48,
					},
				},
				{
					Name: "Yellow",
					Color: &XMLTypes.ColorCIE{
						X:  0.4193,
						Y:  0.5053,
						Y2: 92.78,
					},
				},
			},
		},
		{
			Name: "GoboWheel",
			WheelSlots: []XMLTypes.WheelSlot{
				{
					Name: "Open",
				},
				{
					Name:          "Gobo 1",
					MediaFileName: &gobo1,
				},
				{
					Name:          "Gobo 2",
					MediaFileName: &gobo2,
				},
				{
					Name:          "Gobo 3",
					MediaFileName: &gobo3,
				},
				{
					Name:          "Gobo 4",
					MediaFileName: &gobo4,
				},
				{
					Name:          "Gobo 5",
					MediaFileName: &gobo5,
				},
			},
		},
		{
			Name: "PrismWheel",
			WheelSlots: []XMLTypes.WheelSlot{
				{
					Name: "Open",
				},
				{
					Name: "Prism",
					PrismFacets: []XMLTypes.PrismFacet{
						{
							Rotation: XMLTypes.Rotation{
								[3]float32{0.6, 0.0, 0.0},
								[3]float32{0.0, 0.6, 0.0},
								[3]float32{-0.12, 0.15, 1.0},
							},
						},
						{
							Rotation: XMLTypes.Rotation{
								[3]float32{0.6, 0.0, 0.0},
								[3]float32{0.0, 0.6, 0.0},
								[3]float32{0.17, 0.0, 1.0},
							},
						},
						{
							Rotation: XMLTypes.Rotation{
								[3]float32{0.6, 0.0, 0.0},
								[3]float32{0.0, 0.6, 0.0},
								[3]float32{-0.12, -0.15, 1.0},
							},
						},
					},
				},
			},
		},
	}

	parsingTest(t, xmlData, "Wheel", wheelTest{Wheels: want})
}
