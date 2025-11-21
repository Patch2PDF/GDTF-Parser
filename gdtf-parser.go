package main

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"image"
	_ "image/png"
	"io"
	"log"
	"path/filepath"

	GDTFMeshReader "github.com/Patch2PDF/GDTF-Mesh-Reader"
	"github.com/Patch2PDF/GDTF-Mesh-Reader/pkg/MeshTypes"
	XMLTypes "github.com/Patch2PDF/GDTF-Parser/internal/types/gdtfxml"
	Types "github.com/Patch2PDF/GDTF-Parser/pkg/types"
)

func ParseGDTF(filename string, readMeshes bool, readThumbnail bool) (*Types.GDTF, error) {
	if filepath.Ext(filename) != ".gdtf" {
		return nil, fmt.Errorf("%s is not a GDTF file", filename)
	}
	// unzip gdtf file
	gdtf, err := zip.OpenReader(filename)

	if err != nil {
		return nil, err
	}
	defer gdtf.Close()

	// create map of files
	var fileMap map[string]*zip.File = make(map[string]*zip.File)
	for _, file := range gdtf.File {
		fileMap[file.Name] = file
	}

	xmlFile, err := fileMap["description.xml"].Open()
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()
	data, err := io.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}

	var gdtfContent XMLTypes.GDTF
	err = xml.Unmarshal(data, &gdtfContent)
	if err != nil {
		return nil, err
	}

	// parse XML structs into desired destination structs (kept seperate for breaking changes in GDTF version)
	parsedGDTF := gdtfContent.Parse()

	// create pointers for referencing
	parsedGDTF.CreateReferencePointer()

	// resolve pointer references
	parsedGDTF.ResolveReference()

	if readMeshes {
		for _, model := range parsedGDTF.FixtureType.Models {
			conf := GDTFMeshReader.ModelReaderConf{
				File:          nil,
				Filename:      nil,
				PrimitiveType: model.PrimitiveType,
			}

			if model.File != nil && *model.File != "" && model.PrimitiveType == "Undefined" {
				filename := getGDTFModelFileName(fileMap, *model.File)
				if filename == nil {
					return nil, fmt.Errorf("could not find model file '%s' in GDTF File", *model.File)
				}
				file, err := fileMap[*filename].Open()
				if err != nil {
					return nil, err
				}
				conf.File = file
				conf.Filename = filename
			}

			mesh, err := GDTFMeshReader.GetModel(
				conf,
				MeshTypes.Vector{
					X: float64(model.Length),
					Y: float64(model.Width),
					Z: float64(model.Height),
				},
			)
			if err != nil {
				return nil, err
			}
			model.Mesh = mesh
		}
	}

	if readThumbnail {
		pngPath := parsedGDTF.FixtureType.Thumbnail.String + ".png"
		if fileMap[pngPath] != nil {
			file, err := fileMap[pngPath].Open()
			if err != nil {
				return nil, err
			}
			thumbnail, _, err := image.Decode(file)
			if err != nil {
				return nil, err
			}
			parsedGDTF.FixtureType.Thumbnail.Ptr = &thumbnail
		}
	}

	return &parsedGDTF, nil
}

func getGDTFModelFileName(fileMap map[string]*zip.File, modelName string) *string {
	// prefer gltf, otherwise look for 3ds
	modelPaths := []string{
		"models/gltf/" + modelName + ".gltf",
		"models/gltf/" + modelName + ".glb",
		"models/3ds/" + modelName + ".3ds",
	}
	for _, path := range modelPaths {
		if fileMap[path] != nil {
			return &path
		}
	}
	return nil
}

func main() {
	// Load Primitive Meshes
	err := GDTFMeshReader.LoadPrimitives()
	if err != nil {
		log.Fatal(err)
	}

	gdtf, err := ParseGDTF("test.gdtf", true, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", *gdtf)

	_, err = gdtf.BuildMesh("32Ch")
	if err != nil {
		log.Fatal(err)
	}
}
