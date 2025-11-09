package main

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"log"

	XMLTypes "github.com/Patch2PDF/GDTF-Parser/internal/types/gdtfxml"
)

func main() {
	gdtf, err := zip.OpenReader("test.gdtf")

	if err != nil {
		log.Fatal(err)
	}
	defer gdtf.Close()

	var fileNameMap map[string]int = make(map[string]int)
	for index, file := range gdtf.File {
		fileNameMap[file.Name] = index
	}

	xmlFile, err := gdtf.File[fileNameMap["description.xml"]].Open()
	if err != nil {
		log.Fatal(err)
	}
	defer xmlFile.Close()
	data, err := io.ReadAll(xmlFile)
	if err != nil {
		log.Fatal(err)
	}

	var gdtfContent XMLTypes.XMLGDTF
	err = xml.Unmarshal(data, &gdtfContent)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(gdtfContent.FixtureType)
	fmt.Printf("%+v\n", gdtfContent.FixtureType.Geometries)
}
