package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)


func main(){
	if len(os.Args) < 2 {
		fmt.Println("The path of template should be input.")
		os.Exit(0)
	}
	tmlPath := os.Args[1]

	if !strings.HasSuffix(tmlPath, ".go.yml") {
		fmt.Println(`template path should has suffix ".go.yml".`)
		os.Exit(0)
	}

	data, err := os.ReadFile(tmlPath)
	if err != nil {
		fmt.Println(`Failed in reading yaml file`, tmlPath, ". Error:", err.Error())
		os.Exit(0)
	}

	// var templateData templateData
	// err = yaml.Unmarshal(data, &templateData)
	// if err != nil {
	// 	fmt.Println(`Failed in translating yaml file to data.`)
	// 	os.Exit(0)
	// }

	var templateData map[string]interface{}
	err = yaml.Unmarshal(data, &templateData)
	if err != nil {
		fmt.Println(`Failed in translating yaml file to data. Error:`, err.Error())
		os.Exit(0)
	}

	tml, err := template.New("").Parse(templateData["template"].(string))
	if err != nil {
		fmt.Println("Failed in creating template. Error:", err.Error())
		os.Exit(0)
	}

	// to prevent template from using itself as a parameter
	delete(templateData, "template")

	var bb bytes.Buffer
	err = tml.Execute(&bb, templateData)
	if err != nil {
		fmt.Println("Failed in executing template. Error:", err.Error())
		os.Exit(0)
	}

	goFile, _ := strings.CutSuffix(tmlPath, ".yml")
	goFile, _ = strings.CutPrefix(goFile, "gen.")
	err = os.WriteFile("./" + goFile, bb.Bytes(), 0644)
	if err != nil {
		fmt.Println("Failed in writing file", goFile, "Error:" , err.Error())
	}
}