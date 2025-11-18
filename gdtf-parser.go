package main

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"path/filepath"

	XMLTypes "github.com/Patch2PDF/GDTF-Parser/internal/types/gdtfxml"
	Types "github.com/Patch2PDF/GDTF-Parser/pkg/types"
)

func ParseGDTF(filename string) (*Types.GDTF, error) {
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

	return &parsedGDTF, nil
}

func main() {
	gdtf, err := ParseGDTF("test.gdtf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", *gdtf)
}

// TODO: func for 3d model parsing / conversion or rather a flag in upper `ParseGDTF`?
