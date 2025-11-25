package gdtfxml_test

import (
	"testing"

	XMLTypes "github.com/Patch2PDF/GDTF-Parser/internal/types/gdtfxml"
)

var physicalXmlData = `
	<PhysicalDescriptions>
    <ColorSpace/>
    <Filters/>

    <Emitters>
        <Emitter Color="0.6951,0.3044,100" Name="measured R">
            <Measurement LuminousIntensity="534" Physical="100"> 
                <MeasurementPoint Energy="0.048200" WaveLength="634"/>
                <MeasurementPoint Energy="0.069000" WaveLength="637"/>
                <MeasurementPoint Energy="0.073700" WaveLength="640"/>
                <MeasurementPoint Energy="0.047800" WaveLength="643"/>
                <MeasurementPoint Energy="0.021700" WaveLength="646"/>
            </Measurement>
        </Emitter>
        <Emitter Color="0.3002,0.5998,71.55" Name="measured G">
            <Measurement LuminousIntensity="974" Physical="100" />
        </Emitter>
        <Emitter Color="0.1503,0.0602,7.246" Name="measured B">
            <Measurement LuminousIntensity="236" Physical="100" />
        </Emitter>
        <Emitter Color="0.473,0.4623,58.526" Name="measured A">
            <Measurement LuminousIntensity="1076" Physical="100" />
        </Emitter>
        <Emitter Color="0.3104,0.3242,96.1" Name="measured W">
            <Measurement LuminousIntensity="1344" Physical="100" />
        </Emitter>
    </Emitters>

    <DMXProfiles/>
    <CRIs/>
	</PhysicalDescriptions>
`

var physicalStruct = XMLTypes.PhysicalDescription{
	Emitters: []XMLTypes.Emitter{
		{
			Color: XMLTypes.ColorCIE{
				X:  0.6951,
				Y:  0.3044,
				Y2: 100,
			},
			Name: "measured R",
			Measurements: []XMLTypes.Measurement{
				{
					LuminousIntensity: 534,
					Physical:          100,
					MeasurementPoints: &[]XMLTypes.MeasurementPoint{
						{
							Energy:     0.048200,
							WaveLength: 634,
						},
						{
							Energy:     0.069000,
							WaveLength: 637,
						},
						{
							Energy:     0.073700,
							WaveLength: 640,
						},
						{
							Energy:     0.047800,
							WaveLength: 643,
						},
						{
							Energy:     0.021700,
							WaveLength: 646,
						},
					},
				},
			},
		},
		{
			Name: "measured G",
			Color: XMLTypes.ColorCIE{
				X: 0.3002, Y: 0.5998, Y2: 71.55,
			},
			Measurements: []XMLTypes.Measurement{
				{
					LuminousIntensity: 974,
					Physical:          100,
				},
			},
		},
		{
			Name: "measured B",
			Color: XMLTypes.ColorCIE{
				X: 0.1503, Y: 0.0602, Y2: 7.246,
			},
			Measurements: []XMLTypes.Measurement{
				{
					LuminousIntensity: 236,
					Physical:          100,
				},
			},
		},
		{
			Name: "measured A",
			Color: XMLTypes.ColorCIE{
				X: 0.473, Y: 0.4623, Y2: 58.526,
			},
			Measurements: []XMLTypes.Measurement{
				{
					LuminousIntensity: 1076,
					Physical:          100,
				},
			},
		},
		{
			Name: "measured W",
			Color: XMLTypes.ColorCIE{
				X: 0.3104, Y: 0.3242, Y2: 96.1,
			},
			Measurements: []XMLTypes.Measurement{
				{
					LuminousIntensity: 1344,
					Physical:          100,
				},
			},
		},
	},
}

func TestPhysicalXML(t *testing.T) {
	parsingTest(t, physicalXmlData, "Physical", physicalStruct)
}
