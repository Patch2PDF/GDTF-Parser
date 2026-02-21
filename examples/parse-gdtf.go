package main

import (
	"fmt"
	"log"
	"os"

	GDTFMeshReader "github.com/Patch2PDF/GDTF-Mesh-Reader/v2"
	"github.com/Patch2PDF/GDTF-Mesh-Reader/v2/pkg/MeshTypes"
	GDTFParser "github.com/Patch2PDF/GDTF-Parser"
	STL "github.com/Patch2PDF/GDTF-Parser/examples/stl"
)

func main() {
	// Load Primitive Meshes
	err := GDTFMeshReader.LoadPrimitives()
	if err != nil {
		log.Fatal(err)
	}

	// open and read gdtf file
	gdtf, err := GDTFParser.ParseGDTFByFilename("test.gdtf", true, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", *gdtf)

	// build gdtf model (by dmx mode)
	models, err := gdtf.BuildMesh("32Ch")
	if err != nil {
		log.Fatal(err)
	}
	// write mesh as STL
	mesh := &MeshTypes.Mesh{}
	for _, entry := range models {
		mesh.Add(&entry.Mesh)
	}
	f, _ := os.Create("Test.stl")
	STL.WriteBinary(f, mesh)
	f.Close()
}
