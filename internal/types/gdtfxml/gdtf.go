package XMLTypes

import Types "github.com/Patch2PDF/GDTF-Parser/pkg/types"

type GDTF struct {
	DataVersion string `xml:",attr"`
	FixtureType FixtureType
}

func (gdtf GDTF) Parse() Types.GDTF {
	return Types.GDTF{
		DataVersion: gdtf.DataVersion,
		FixtureType: gdtf.FixtureType.Parse(),
	}
}
