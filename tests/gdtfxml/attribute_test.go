package gdtfxml_test

import (
	"encoding/xml"
	"reflect"
	"testing"

	XMLTypes "github.com/Patch2PDF/GDTF-Parser/internal/types/gdtfxml"
)

func TestParseActivationGroups(t *testing.T) {
	type activationGroupTest struct {
		ActivationGroups []XMLTypes.ActivationGroup `xml:"ActivationGroup"`
	}

	xmlData := `
	<ActivationGroups>
		<ActivationGroup Name="ColorRGB" />
	</ActivationGroups>
	`

	var result activationGroupTest
	err := xml.Unmarshal([]byte(xmlData), &result)
	if err != nil {
		t.Errorf("Activation Group: XML Unmarshal threw error: %s", err)
	}

	want := []XMLTypes.ActivationGroup{
		{
			Name: "ColorRGB",
		},
	}

	if !reflect.DeepEqual(result.ActivationGroups, want) {
		t.Errorf("Activation Group: XML Unmarshaling Output does not match")
	}
}

func TestParseFeatureGroup(t *testing.T) {
	type featureGroupTest struct {
		FeatureGroups []XMLTypes.FeatureGroup `xml:"FeatureGroup"`
	}

	xmlData := `
	<FeatureGroups>
    <FeatureGroup Name="Color">
        <Feature Name="RGB" />
        <Feature Name="Color" />
    </FeatureGroup>
    <FeatureGroup Name="Beam">
        <Feature Name="Beam" />
    </FeatureGroup>
    <FeatureGroup Name="Dimmer">
        <Feature Name="Dimmer" />
    </FeatureGroup>
    <FeatureGroup Name="Gobo">
        <Feature Name="Gobo" />
    </FeatureGroup>
    <FeatureGroup Name="Focus">
        <Feature Name="Focus" />
    </FeatureGroup>
    <FeatureGroup Name="Position">
        <Feature Name="PanTilt" />
    </FeatureGroup>
    <FeatureGroup Name="Control">
        <Feature Name="Control" />
    </FeatureGroup>
    <FeatureGroup Name="Shapers">
        <Feature Name="Shapers" />
    </FeatureGroup>
	</FeatureGroups>
	`

	var result featureGroupTest
	err := xml.Unmarshal([]byte(xmlData), &result)
	if err != nil {
		t.Errorf("Feature Group: XML Unmarshal threw error: %s", err)
	}

	want := []XMLTypes.FeatureGroup{
		{
			Name: "Color",
			Features: []XMLTypes.Feature{
				{Name: "RGB"},
				{Name: "Color"},
			},
		},
		{
			Name: "Beam",
			Features: []XMLTypes.Feature{
				{Name: "Beam"},
			},
		},
		{
			Name: "Dimmer",
			Features: []XMLTypes.Feature{
				{Name: "Dimmer"},
			},
		},
		{
			Name: "Gobo",
			Features: []XMLTypes.Feature{
				{Name: "Gobo"},
			},
		},
		{
			Name: "Focus",
			Features: []XMLTypes.Feature{
				{Name: "Focus"},
			},
		},
		{
			Name: "Position",
			Features: []XMLTypes.Feature{
				{Name: "PanTilt"},
			},
		},
		{
			Name: "Control",
			Features: []XMLTypes.Feature{
				{Name: "Control"},
			},
		},
		{
			Name: "Shapers",
			Features: []XMLTypes.Feature{
				{Name: "Shapers"},
			},
		},
	}

	if !reflect.DeepEqual(result.FeatureGroups, want) {
		t.Errorf("Feature Group: XML Unmarshaling Output does not match")
	}
}

func TestParseAttributes(t *testing.T) {
	type attributeTest struct {
		Attributes []XMLTypes.Attribute `xml:"Attribute"`
	}

	xmlData := `
	<Attributes>
    <Attribute Name="Shutter1" Pretty="Sh1" Feature="Beam.Beam" />
    <Attribute Name="Shutter1Strobe" Pretty="Strobe" MainAttribute="Shutter1" Feature="Beam.Beam" PhysicalUnit="Frequency" />
    <Attribute Name="Shutter1StrobeRandom" Pretty="Random" MainAttribute="Shutter1" Feature="Beam.Beam" PhysicalUnit="Frequency" />
    <Attribute Name="Dimmer" Pretty="Dim" Feature="Dimmer.Dimmer" PhysicalUnit="LuminousIntensity" />
    <Attribute Name="ColorAdd_R" Pretty="R" ActivationGroup="ColorRGB" Feature="Color.RGB" PhysicalUnit="ColorComponent" Color="0.64,0.33,21.3" />
    <Attribute Name="ColorAdd_G" Pretty="G" ActivationGroup="ColorRGB" Feature="Color.RGB" PhysicalUnit="ColorComponent" Color="0.3,0.6,71.5" />
    <Attribute Name="ColorAdd_B" Pretty="B" ActivationGroup="ColorRGB" Feature="Color.RGB" PhysicalUnit="ColorComponent" Color="0.15,0.06,7.2" />
    <Attribute Name="Color1" Pretty="C1" ActivationGroup="ColorRGB" Feature="Color.Color" />
    <Attribute Name="Color1WheelSpin" Pretty="Wheel Spin" MainAttribute="Color1" ActivationGroup="ColorRGB" Feature="Color.Color" PhysicalUnit="AngularSpeed" />
    <Attribute Name="Gobo1" Pretty="G1" ActivationGroup="Gobo" Feature="Gobo.Gobo" />
    <Attribute Name="Gobo1WheelSpin" Pretty="Wheel Spin" MainAttribute="Gobo1" ActivationGroup="Gobo" Feature="Gobo.Gobo" PhysicalUnit="AngularSpeed" />
    <Attribute Name="Gobo1Pos" Pretty="G1 &lt;&gt;" Feature="Gobo.Gobo" PhysicalUnit="Angle" />
    <Attribute Name="Gobo1PosRotate" Pretty="Rotate" MainAttribute="Gobo1Pos" Feature="Gobo.Gobo" PhysicalUnit="AngularSpeed" />
    <Attribute Name="Gobo2" Pretty="G2" ActivationGroup="Gobo" Feature="Gobo.Gobo" />
    <Attribute Name="Gobo2WheelSpin" Pretty="Wheel Spin" MainAttribute="Gobo2" ActivationGroup="Gobo" Feature="Gobo.Gobo" PhysicalUnit="AngularSpeed" />
    <Attribute Name="Prism1" Pretty="Prism1" ActivationGroup="Prism" Feature="Beam.Beam" />
    <Attribute Name="Iris" Pretty="Iris" Feature="Beam.Beam" />
    <Attribute Name="IrisStrobe" Pretty="Strobe" MainAttribute="Iris" Feature="Beam.Beam" PhysicalUnit="Frequency" />
    <Attribute Name="Zoom" Pretty="Zoom" Feature="Focus.Focus" PhysicalUnit="Angle" />
    <Attribute Name="Focus1" Pretty="Focus1" Feature="Focus.Focus" />
    <Attribute Name="Pan" Pretty="P" ActivationGroup="PanTilt" Feature="Position.PanTilt" PhysicalUnit="Angle" />
    <Attribute Name="Tilt" Pretty="T" ActivationGroup="PanTilt" Feature="Position.PanTilt" PhysicalUnit="Angle" />
    <Attribute Name="Control1" Pretty="Ctrl1" Feature="Control.Control" />
    <Attribute Name="Effects1" Pretty="FX1" Feature="Beam.Beam" />
    <Attribute Name="Effects1Rate" Pretty="FX1 Rate" Feature="Beam.Beam" PhysicalUnit="Speed" />
    <Attribute Name="Effects2" Pretty="FX2" Feature="Beam.Beam" />
    <Attribute Name="Effects2Rate" Pretty="FX2 Rate" Feature="Beam.Beam" PhysicalUnit="Speed" />
    <Attribute Name="EffectsSync" Pretty="FX Sync" Feature="Beam.Beam" />
	</Attributes>
	`

	var result attributeTest
	err := xml.Unmarshal([]byte(xmlData), &result)
	if err != nil {
		t.Errorf("Attributes: XML Unmarshal threw error: %s", err)
	}

	want := []XMLTypes.Attribute{
		{
			Name:    "Shutter1",
			Pretty:  "Sh1",
			Feature: "Beam.Beam",
		},
		{
			Name:          "Shutter1Strobe",
			Pretty:        "Strobe",
			MainAttribute: "Shutter1",
			Feature:       "Beam.Beam",
			PhysicalUnit:  "Frequency",
		},
		{
			Name:          "Shutter1StrobeRandom",
			Pretty:        "Random",
			MainAttribute: "Shutter1",
			Feature:       "Beam.Beam",
			PhysicalUnit:  "Frequency",
		},
		{
			Name:         "Dimmer",
			Pretty:       "Dim",
			Feature:      "Dimmer.Dimmer",
			PhysicalUnit: "LuminousIntensity",
		},
		{
			Name:            "ColorAdd_R",
			Pretty:          "R",
			ActivationGroup: "ColorRGB",
			Feature:         "Color.RGB",
			PhysicalUnit:    "ColorComponent",
			Color:           XMLTypes.ColorCIE{X: 0.64, Y: 0.33, Y2: 21.3},
		},
		{
			Name:            "ColorAdd_G",
			Pretty:          "G",
			ActivationGroup: "ColorRGB",
			Feature:         "Color.RGB",
			PhysicalUnit:    "ColorComponent",
			Color:           XMLTypes.ColorCIE{X: 0.3, Y: 0.6, Y2: 71.5},
		},
		{
			Name:            "ColorAdd_B",
			Pretty:          "B",
			ActivationGroup: "ColorRGB",
			Feature:         "Color.RGB",
			PhysicalUnit:    "ColorComponent",
			Color:           XMLTypes.ColorCIE{X: 0.15, Y: 0.06, Y2: 7.2},
		},
		{
			Name:            "Color1",
			Pretty:          "C1",
			ActivationGroup: "ColorRGB",
			Feature:         "Color.Color",
		},
		{
			Name:            "Color1WheelSpin",
			Pretty:          "Wheel Spin",
			MainAttribute:   "Color1",
			ActivationGroup: "ColorRGB",
			Feature:         "Color.Color",
			PhysicalUnit:    "AngularSpeed",
		},
		{
			Name:            "Gobo1",
			Pretty:          "G1",
			ActivationGroup: "Gobo",
			Feature:         "Gobo.Gobo",
		},
		{
			Name:            "Gobo1WheelSpin",
			Pretty:          "Wheel Spin",
			MainAttribute:   "Gobo1",
			ActivationGroup: "Gobo",
			Feature:         "Gobo.Gobo",
			PhysicalUnit:    "AngularSpeed",
		},
		{
			Name:         "Gobo1Pos",
			Pretty:       "G1 <>",
			Feature:      "Gobo.Gobo",
			PhysicalUnit: "Angle",
		},
		{
			Name:          "Gobo1PosRotate",
			Pretty:        "Rotate",
			MainAttribute: "Gobo1Pos",
			Feature:       "Gobo.Gobo",
			PhysicalUnit:  "AngularSpeed",
		},
		{
			Name:            "Gobo2",
			Pretty:          "G2",
			ActivationGroup: "Gobo",
			Feature:         "Gobo.Gobo",
		},
		{
			Name:            "Gobo2WheelSpin",
			Pretty:          "Wheel Spin",
			MainAttribute:   "Gobo2",
			ActivationGroup: "Gobo",
			Feature:         "Gobo.Gobo",
			PhysicalUnit:    "AngularSpeed",
		},
		{
			Name:            "Prism1",
			Pretty:          "Prism1",
			ActivationGroup: "Prism",
			Feature:         "Beam.Beam",
		},
		{
			Name:    "Iris",
			Pretty:  "Iris",
			Feature: "Beam.Beam",
		},
		{
			Name:          "IrisStrobe",
			Pretty:        "Strobe",
			MainAttribute: "Iris",
			Feature:       "Beam.Beam",
			PhysicalUnit:  "Frequency",
		},
		{
			Name:         "Zoom",
			Pretty:       "Zoom",
			Feature:      "Focus.Focus",
			PhysicalUnit: "Angle",
		},
		{
			Name:    "Focus1",
			Pretty:  "Focus1",
			Feature: "Focus.Focus",
		},
		{
			Name:            "Pan",
			Pretty:          "P",
			ActivationGroup: "PanTilt",
			Feature:         "Position.PanTilt",
			PhysicalUnit:    "Angle",
		},
		{
			Name:            "Tilt",
			Pretty:          "T",
			ActivationGroup: "PanTilt",
			Feature:         "Position.PanTilt",
			PhysicalUnit:    "Angle",
		},
		{
			Name:    "Control1",
			Pretty:  "Ctrl1",
			Feature: "Control.Control",
		},
		{
			Name:    "Effects1",
			Pretty:  "FX1",
			Feature: "Beam.Beam",
		},
		{
			Name:         "Effects1Rate",
			Pretty:       "FX1 Rate",
			Feature:      "Beam.Beam",
			PhysicalUnit: "Speed",
		},
		{
			Name:    "Effects2",
			Pretty:  "FX2",
			Feature: "Beam.Beam",
		},
		{
			Name:         "Effects2Rate",
			Pretty:       "FX2 Rate",
			Feature:      "Beam.Beam",
			PhysicalUnit: "Speed",
		},
		{
			Name:    "EffectsSync",
			Pretty:  "FX Sync",
			Feature: "Beam.Beam",
		},
	}

	if !reflect.DeepEqual(result.Attributes, want) {
		t.Errorf("Attributes: XML Unmarshaling Output does not match")
	}
}
