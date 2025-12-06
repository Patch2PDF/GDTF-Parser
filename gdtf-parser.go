package GDTFParser

import (
	"archive/zip"
	"bytes"
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

// Parse a GDTF file
//
// # Args:
//
// - zipfile: GDTF File as zip.Reader
//
// - readMeshes: wether to read model meshes (true) or leave nil (false)
//
// - readThumbnail: wether to read model thumbnail (true) or leave nil (false)
func ParseGDTFZipReader(zipfile *zip.Reader, readMeshes bool, readThumbnail bool) (*Types.GDTF, error) {
	// create map of files
	var fileMap map[string]*zip.File = make(map[string]*zip.File)
	for _, file := range zipfile.File {
		fileMap[file.Name] = file
	}

	xmlFile, err := fileMap["description.xml"].Open()
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}
	xmlFile.Close()

	var gdtfContent XMLTypes.GDTF
	err = xml.Unmarshal(data, &gdtfContent)
	if err != nil {
		return nil, err
	}

	// parse XML structs into desired destination structs (kept seperate for breaking changes in GDTF version)
	parsedGDTF := gdtfContent.Parse()

	refPointers := Types.CreateRefPointersMap()

	// create pointers for referencing
	parsedGDTF.CreateReferencePointer(refPointers)

	// resolve pointer references
	parsedGDTF.ResolveReference(refPointers)

	if readMeshes {
		for _, model := range parsedGDTF.FixtureType.Models {
			conf := GDTFMeshReader.ModelReaderConf{
				File:          nil,
				Filename:      nil,
				PrimitiveType: model.PrimitiveType,
			}

			// model file has precedence
			if model.File != nil && *model.File != "" {
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
				&MeshTypes.Vector{
					X: float64(model.Length),
					Y: float64(model.Width),
					Z: float64(model.Height),
				},
			)
			if err != nil {
				// return nil, err
				log.Printf("invalid model: %s: %s", model.Name, err)
				mesh = &MeshTypes.Mesh{}
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

// Parse a GDTF file
//
// # Args:
//
// - filename: path to GDTF File (can be relative or absolute)
//
// - readMeshes: wether to read model meshes (true) or leave nil (false)
//
// - readThumbnail: wether to read model thumbnail (true) or leave nil (false)
func ParseGDTFByFilename(filename string, readMeshes bool, readThumbnail bool) (*Types.GDTF, error) {
	if filepath.Ext(filename) != ".gdtf" {
		return nil, fmt.Errorf("%s is not a GDTF file", filename)
	}
	// unzip gdtf file
	gdtf, err := zip.OpenReader(filename)
	if err != nil {
		return nil, err
	}
	defer gdtf.Close()

	return ParseGDTFZipReader(&gdtf.Reader, readMeshes, readThumbnail)
}

// Parse a GDTF file
//
// # Args:
//
// - file: GDTF File as io.Reader
//
// - readMeshes: wether to read model meshes (true) or leave nil (false)
//
// - readThumbnail: wether to read model thumbnail (true) or leave nil (false)
func ParseGDTFByFile(file io.Reader, readMeshes bool, readThumbnail bool) (*Types.GDTF, error) {
	var in io.Reader
	var size int64

	if _, ok := in.(io.ReaderAt); !ok {
		buffer, err := io.ReadAll(file)

		if err != nil {
			return nil, err
		}

		in = bytes.NewReader(buffer)
		size = int64(len(buffer))
	}

	// unzip gdtf file
	gdtf, err := zip.NewReader(in.(io.ReaderAt), size)
	if err != nil {
		return nil, err
	}

	return ParseGDTFZipReader(gdtf, readMeshes, readThumbnail)
}

// get path to model file by model name
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
