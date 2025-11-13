package Types

type GDTF struct {
	DataVersion string
	FixtureType FixtureType
}

func (obj *GDTF) CreateReferencePointer() {
	obj.FixtureType.CreateReferencePointer()
}

func (obj *GDTF) ResolveReference() {
	obj.FixtureType.ResolveReference()
}
