package gdtfxml_test

import (
	"testing"

	XMLTypes "github.com/Patch2PDF/GDTF-Parser/internal/types/gdtfxml"
)

var geometryXML = `
	<Geometries>
    <Geometry Name="Base2" Model="Base2">
        <Axis Name="Yoke2" Model="Yoke2" Position="{1.000000,0.000000,0.000000,0.000000}{0.000000,1.000000,0.000000,0.000000}{0.000000,0.000000,1.000000,-0.225000}{0.000000,0.000000,0.000000,1.000000}">
            <Axis Name="Head" Model="Head" Position="{1.000000,0.000000,0.000000,0.000000}{0.000000,1.000000,0.000000,0.000000}{0.000000,0.000000,1.000000,-0.100000}{0.000000,0.000000,0.000000,1.000000}">
                <Beam Name="Beam" Model="Beam" Position="{1.000000,0.000000,0.000000,0.000000}{0.000000,1.000000,0.000000,0.000000}{0.000000,0.000000,1.000000,-0.150000}{0.000000,0.000000,0.000000,1.000000}" LampType="Discharge" PowerConsumption="350" LuminousFlux="12000" ColorTemperature="6500" BeamAngle="20" FieldAngle="22" BeamRadius="0.025" BeamType="Spot" ColorRenderingIndex="75" />
            </Axis>
        </Axis>
    </Geometry>
    <Geometry Name="Base1" Model="Base1">
        <Axis Name="Yoke1" Model="Yoke1" Position="{1.000000,0.000000,0.000000,0.000000}{0.000000,1.000000,0.000000,0.000000}{0.000000,0.000000,1.000000,-0.200000}{0.000000,0.000000,0.000000,1.000000}">
            <GeometryReference Name="Head1" Position="{1.000000,0.000000,0.000000,-0.550000}{0.000000,1.000000,0.000000,0.000000}{0.000000,0.000000,1.000000,-0.125000}{0.000000,0.000000,0.000000,1.000000}" Geometry="Base2">
                <Break DMXOffset="1" />
            </GeometryReference>
            <GeometryReference Name="Head2" Position="{1.000000,0.000000,0.000000,-0.183000}{0.000000,1.000000,0.000000,0.000000}{0.000000,0.000000,1.000000,-0.125000}{0.000000,0.000000,0.000000,1.000000}" Geometry="Base2">
                <Break DMXOffset="8" />
						</GeometryReference>
            <GeometryReference Name="Head3" Position="{1.000000,0.000000,0.000000,0.183000}{0.000000,1.000000,0.000000,0.000000}{0.000000,0.000000,1.000000,-0.125000}{0.000000,0.000000,0.000000,1.000000}" Geometry="Base2">
                <Break DMXOffset="15" />
            </GeometryReference>
            <GeometryReference Name="Head4" Position="{1.000000,0.000000,0.000000,0.550000}{0.000000,1.000000,0.000000,0.000000}{0.000000,0.000000,1.000000,-0.125000}{0.000000,0.000000,0.000000,1.000000}" Geometry="Base2">
                <Break DMXOffset="22" />
            </GeometryReference>
        </Axis>
    </Geometry>
	</Geometries>
`

var geometryStruct = XMLTypes.Geometries{
	GeometryList: []XMLTypes.Geometry{
		{
			Name:  "Base2",
			Model: "Base2",
			Geometries: XMLTypes.Geometries{
				AxisList: []XMLTypes.Axis{
					{
						Name:  "Yoke2",
						Model: "Yoke2",
						Position: XMLTypes.Matrix{
							{1.000000, 0.000000, 0.000000, 0.000000},
							{0.000000, 1.000000, 0.000000, 0.000000},
							{0.000000, 0.000000, 1.000000, -0.225000},
							{0.000000, 0.000000, 0.000000, 1.000000},
						},
						Geometries: XMLTypes.Geometries{
							AxisList: []XMLTypes.Axis{
								{
									Name:  "Head",
									Model: "Head",
									Position: XMLTypes.Matrix{
										{1.000000, 0.000000, 0.000000, 0.000000},
										{0.000000, 1.000000, 0.000000, 0.000000},
										{0.000000, 0.000000, 1.000000, -0.100000},
										{0.000000, 0.000000, 0.000000, 1.000000},
									},
									Geometries: XMLTypes.Geometries{
										BeamList: []XMLTypes.Beam{
											{
												Name:  "Beam",
												Model: "Beam",
												Position: XMLTypes.Matrix{
													{1.000000, 0.000000, 0.000000, 0.000000},
													{0.000000, 1.000000, 0.000000, 0.000000},
													{0.000000, 0.000000, 1.000000, -0.150000},
													{0.000000, 0.000000, 0.000000, 1.000000},
												},
												LampType:         "Discharge",
												PowerConsumption: 350,
												LuminousFlux:     12000,
												ColorTemperature: 6500,
												BeamAngle:        20,
												FieldAngle:       22,
												BeamRadius:       0.025,
												BeamType:         "Spot",
												CRI:              75,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			Name:  "Base1",
			Model: "Base1",
			Geometries: XMLTypes.Geometries{
				AxisList: []XMLTypes.Axis{
					{
						Name:  "Yoke1",
						Model: "Yoke1",
						Position: XMLTypes.Matrix{
							{1.000000, 0.000000, 0.000000, 0.000000},
							{0.000000, 1.000000, 0.000000, 0.000000},
							{0.000000, 0.000000, 1.000000, -0.200000},
							{0.000000, 0.000000, 0.000000, 1.000000},
						},
						Geometries: XMLTypes.Geometries{
							GeometryReferenceList: []XMLTypes.GeometryReference{
								{
									Name:        "Head1",
									GeometryRef: "Base2",
									Position: XMLTypes.Matrix{
										{1.000000, 0.000000, 0.000000, -0.550000},
										{0.000000, 1.000000, 0.000000, 0.000000},
										{0.000000, 0.000000, 1.000000, -0.125000},
										{0.000000, 0.000000, 0.000000, 1.000000},
									},
									Breaks: []XMLTypes.Break{
										{
											DMXOffset: XMLTypes.DMXAddress{Universe: 0, Address: 1},
										},
									},
								},
								{
									Name:        "Head2",
									GeometryRef: "Base2",
									Position: XMLTypes.Matrix{
										{1.000000, 0.000000, 0.000000, -0.183000},
										{0.000000, 1.000000, 0.000000, 0.000000},
										{0.000000, 0.000000, 1.000000, -0.125000},
										{0.000000, 0.000000, 0.000000, 1.000000},
									},
									Breaks: []XMLTypes.Break{
										{
											DMXOffset: XMLTypes.DMXAddress{Universe: 0, Address: 8},
										},
									},
								},
								{
									Name:        "Head3",
									GeometryRef: "Base2",
									Position: XMLTypes.Matrix{
										{1.000000, 0.000000, 0.000000, 0.183000},
										{0.000000, 1.000000, 0.000000, 0.000000},
										{0.000000, 0.000000, 1.000000, -0.125000},
										{0.000000, 0.000000, 0.000000, 1.000000},
									},
									Breaks: []XMLTypes.Break{
										{
											DMXOffset: XMLTypes.DMXAddress{Universe: 0, Address: 15},
										},
									},
								},
								{
									Name:        "Head4",
									GeometryRef: "Base2",
									Position: XMLTypes.Matrix{
										{1.000000, 0.000000, 0.000000, 0.550000},
										{0.000000, 1.000000, 0.000000, 0.000000},
										{0.000000, 0.000000, 1.000000, -0.125000},
										{0.000000, 0.000000, 0.000000, 1.000000},
									},
									Breaks: []XMLTypes.Break{
										{
											DMXOffset: XMLTypes.DMXAddress{Universe: 0, Address: 22},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}

func TestGeometry(t *testing.T) {
	parsingTest(t, geometryXML, "Geometry", geometryStruct)
}
