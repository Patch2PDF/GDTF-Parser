package STL

import (
	"encoding/binary"
	"io"

	"github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"
)

// helper for `WriteBinary`
func computeNormal(t *MeshTypes.Triangle) MeshTypes.Vector {
	// edge vectors
	a := t.V1.Position.Sub(t.V0.Position)

	b := t.V2.Position.Sub(t.V0.Position)

	n := a.Cross(b)

	return n.Normalize()
}

// helper for testing extracted vertices, will be removed
func WriteBinary(w io.Writer, mesh *MeshTypes.Mesh) error {
	// 80-byte header
	header := make([]byte, 80)
	if _, err := w.Write(header); err != nil {
		return err
	}

	triCount := uint32(len(mesh.Triangles))
	if err := binary.Write(w, binary.LittleEndian, triCount); err != nil {
		return err
	}

	for _, t := range mesh.Triangles {

		// Use vertex normal if present, else compute one
		var n MeshTypes.Vector
		if t.V0.Normal != nil {
			n = *t.V0.Normal
		} else {
			n = computeNormal(t)
		}

		writeVec := func(v MeshTypes.Vector) error {
			if err := binary.Write(w, binary.LittleEndian, float32(v.X)); err != nil {
				return err
			}
			if err := binary.Write(w, binary.LittleEndian, float32(v.Y)); err != nil {
				return err
			}
			return binary.Write(w, binary.LittleEndian, float32(v.Z))
		}

		// normal + 3 vertices + attribute byte count
		if err := writeVec(n); err != nil {
			return err
		}
		if err := writeVec(t.V0.Position); err != nil {
			return err
		}
		if err := writeVec(t.V1.Position); err != nil {
			return err
		}
		if err := writeVec(t.V2.Position); err != nil {
			return err
		}

		binary.Write(w, binary.LittleEndian, uint16(0)) // attribute
	}

	return nil
}
