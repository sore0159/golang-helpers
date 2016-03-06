package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

func main() {
	fmt.Println("")
	log.Println("Hello and welcome to mulegen Code Generation Services")
	structTP := template.Must(template.ParseFiles("tmpl/item.tmpl"))
	managerTP := template.Must(template.ParseFiles("tmpl/overmind.tmpl"))
	intfTP := template.Must(template.ParseFiles("tmpl/intf.tmpl"))
	dirInfo, err := ioutil.ReadDir("in")
	if err != nil {
		log.Println("DIR READ ERROR:", err)
		return
	}
	var allData []StructData
	for _, f := range dirInfo {
		name := f.Name()
		parts := strings.Split(name, ".")
		if parts[len(parts)-1] != "go" {
			continue
		}
		inBytes, err := ioutil.ReadFile("in/" + name)
		if err != nil {
			log.Println(name+" FILE READ ERROR:", err)
			continue
		}
		sD, err := ParseStructData(inBytes)
		if err != nil {
			log.Println(name+" FILE PARSE ERROR:", err)
			continue
		}
		outFile, err := os.Create("out/" + name)
		if err != nil {
			log.Println(name+" FILE CREATE ERROR:", err)
			continue
		}
		err = structTP.ExecuteTemplate(outFile, "main", sD)
		if err != nil {
			log.Println("TEMPLATE ERROR:", err)
			outFile.Close()
			continue
		}
		allData = append(allData, sD)
		outFile.Close()
		GoFmt("out/" + name)
	}

	if len(allData) > 0 {
		outFile, err := os.Create("out/overManager.go")
		if err != nil {
			log.Println("MANAGER FILE CREATE ERROR:", err)
			return
		}
		err = managerTP.ExecuteTemplate(outFile, "main", allData)
		if err != nil {
			log.Println("MANAGER TEMPLATE ERROR:", err)
			outFile.Close()
			return
		}
		outFile.Close()
		GoFmt("out/overManager.go")

		outFile, err = os.Create("out/interfaceCollection.go")
		if err != nil {
			log.Println("INTERFACE FILE CREATE ERROR:", err)
			return
		}
		intfP := strings.Trim(allData[0].IntfPack, ".")
		if intfP == "" {
			intfP = allData[0].PackName
		}
		mainData := MainData{
			IntfPack: intfP,
			Structs:  allData,
		}
		for _, sd := range allData {
			mainData.Imports = append(mainData.Imports, sd.Imports...)
		}
		err = intfTP.ExecuteTemplate(outFile, "main", mainData)
		if err != nil {
			log.Println("INTERFACE TEMPLATE ERROR:", err)
			outFile.Close()
			return
		}
		outFile.Close()
		GoFmt("out/interfaceCollection.go")

	}
	log.Println("Thank you for your patronage\n")
	return
}

func GoFmt(fileName string) {
	cmd := exec.Command("go", "fmt", fileName)
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Println("ERROR GO FMTing output for", fileName+":", err)
	} else {
		log.Println("File", fileName, "generated")
	}
}
