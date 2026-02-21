package Types

import "github.com/Patch2PDF/GDTF-Mesh-Reader/v2/pkg/MeshTypes"

type MeshModel struct {
	Mesh         MeshTypes.Mesh
	GeometryType GeometryType
	GeometryPtr  GeometryModel
}

func (obj MeshModel) Copy() MeshModel {
	return MeshModel{
		Mesh:         obj.Mesh.Copy(),
		GeometryType: obj.GeometryType,
		GeometryPtr:  obj.GeometryPtr,
	}
}
