package gencode

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"

	"bytes"

	"gopkg.in/yaml.v2"
)

func GenerateFromTemplate(inputPath, outputPath string, mode Mode) error {
	if !strings.HasSuffix(inputPath, ".go.yml") {
		fmt.Println(`template path should has suffix ".go.yml".`)
		os.Exit(0)
	}

	if outputPath == "" {
		outputPath, _ = strings.CutSuffix(inputPath, ".yml")
	}

	data := GetTemplateData(inputPath)
	code := GenerateCode(data)

	err := WriteToFile(code, outputPath, mode)
	return err
}

func GetTemplateData(inputPath string) map[string]interface{} {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Println(`Failed in reading yaml file`, inputPath, ". Error:", err.Error())
		os.Exit(0)
	}

	var templateData map[string]interface{}
	err = yaml.Unmarshal(data, &templateData)
	if err != nil {
		fmt.Println(`Failed in translating yaml file to data. Error:`, err.Error())
		os.Exit(0)
	}

	return templateData
}

func GenerateCode(data map[string]interface{}) []byte {
	tml, err := template.New("").Parse(data["template"].(string))
	if err != nil {
		fmt.Println("Failed in creating template. Error:", err.Error())
		os.Exit(0)
	}

	// to prevent template from using itself as a parameter
	delete(data, "template")

	var bb bytes.Buffer
	err = tml.Execute(&bb, data)
	if err != nil {
		fmt.Println("Failed in executing template. Error:", err.Error())
		os.Exit(0)
	}

	return bb.Bytes()
}

func WriteToFile(code []byte, outputPath string, mode Mode) error {
	if strings.LastIndex(outputPath, "/") == -1 {
		outputPath = "./" + outputPath
	}
	fmt.Println("mode:", mode)
	switch mode {
	case ModeAppend:
		file, err := os.OpenFile(outputPath, os.O_APPEND, os.ModeAppend)
		if err != nil {
			return err
		}
		_, err = file.Write(code)
		if err != nil {
			return err
		}
	case ModeOverwrite:
		err := os.WriteFile(outputPath, code, 0644)
		if err != nil {
			return err
		}
	case ModeGenIfNotExists:
		fmt.Printf("path %s", outputPath)
		if _,err := os.Stat(outputPath); !errors.Is(err, os.ErrNotExist) {
			if err == nil {
				fmt.Printf("gencode %s stopped. path has existed\n", outputPath)
			}
			return err;
		}

		file, err := os.OpenFile(outputPath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			return err
		}
		_, err = file.Write(code)
		if err != nil {
			return err
		}
	}
	return nil
}
