package gdtfxml_test

import (
	"testing"

	XMLTypes "github.com/Patch2PDF/GDTF-Parser/internal/types/gdtfxml"
)

func TestModels(t *testing.T) {
	type modelTest struct {
		Models []XMLTypes.Model `xml:"Model"`
	}

	xmlData := `
	<Models>
    <Model Name="Base1" Length="0.3" Width="0.3" Height="0.2" PrimitiveType="Cube" File="Base" />
    <Model Name="Yoke1" Length="1.3" Width="0.2" Height="0.2" PrimitiveType="Cube" File="Body" />
    <Model Name="Base2" Length="0.1" Width="0.1" Height="0.05" PrimitiveType="Cube" File="Yoke1" />
    <Model Name="Yoke2" Length="0.2" Width="0.1" Height="0.2" PrimitiveType="Cube" File="Yoke2" />
    <Model Name="Head" Length="0.15" Width="0.15" Height="0.25" PrimitiveType="Cylinder" File="Head" />
    <Model Name="Beam" Length="0.05" Width="0.05" Height="0.02" PrimitiveType="Cylinder" File="Beam" />
	</Models>
	`

	base := "Base"
	body := "Body"
	yoke1 := "Yoke1"
	yoke2 := "Yoke2"
	head := "Head"
	beam := "Beam"

	want := []XMLTypes.Model{
		{
			Name:          "Base1",
			Length:        0.3,
			Width:         0.3,
			Height:        0.2,
			PrimitiveType: "Cube",
			File:          &base,
		},
		{
			Name:          "Yoke1",
			Length:        1.3,
			Width:         0.2,
			Height:        0.2,
			PrimitiveType: "Cube",
			File:          &body,
		},
		{
			Name:          "Base2",
			Length:        0.1,
			Width:         0.1,
			Height:        0.05,
			PrimitiveType: "Cube",
			File:          &yoke1,
		},
		{
			Name:          "Yoke2",
			Length:        0.2,
			Width:         0.1,
			Height:        0.2,
			PrimitiveType: "Cube",
			File:          &yoke2,
		},
		{
			Name:          "Head",
			Length:        0.15,
			Width:         0.15,
			Height:        0.25,
			PrimitiveType: "Cylinder",
			File:          &head,
		},
		{
			Name:          "Beam",
			Length:        0.05,
			Width:         0.05,
			Height:        0.02,
			PrimitiveType: "Cylinder",
			File:          &beam,
		},
	}

	parsingTest(t, xmlData, "Model", modelTest{Models: want})
}
