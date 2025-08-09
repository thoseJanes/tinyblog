package main

import (
	"errors"
	"strings"
	"github.com/spf13/cobra"
	"github.com/thoseJanes/tinyblog/pkg/gencode"
)

var(
	source = gencode.SourceTemplate
	mode = gencode.ModeGenIfNotExists//默认情况下保守策略，只有文件不存在才会生成代码。
	outputPath string
	inputPath string
)

var(
	// interfaceTypes []string
	// implementTypes []string
	interfaceReflection []string
)

var rootCmd = &cobra.Command{
	Use: "gencode",
	Short: "Generate code from a yaml file.",
	SilenceUsage: true,
	RunE: func(c *cobra.Command, args []string) error {
		// outputPath, _ := c.Flags().GetString("output")
		inputPath = args[0]
		// fmt.Printf("%v\n", args)
		// fmt.Printf("%v\n", outputPath)
		
		return run()
	},
	Args: cobra.ExactArgs(1),
}


func main(){
	rootCmd.PersistentFlags().StringVarP(&outputPath, "output", "o", "", "The output path of generated code file.")
	rootCmd.PersistentFlags().VarP(&source, "type", "t", "Input type. Default type is template(t). Support interface(i) whose name should be specified.")
	rootCmd.PersistentFlags().StringSliceVarP(&interfaceReflection, "reflect", "r", nil, 
		`Interface to be implemented and the name of its implement splited by ":". Support multiple reflection connected by ",".`)
	rootCmd.PersistentFlags().VarP(&mode, "mode", "m", "Mode to generate code. Overwrite(o), append(a) or Generate if not exists(n).")
	// rootCmd.Flags().Var()
	rootCmd.Execute()
}


func run() error {
	switch source {
	case gencode.SourceTemplate:
		return gencode.GenerateFromTemplate(inputPath, outputPath, mode)
	case gencode.SourceInterface:
		var reflection = map[string]string{}
		for _, item := range interfaceReflection {
			pair := strings.Split(item, ":")
			if len(pair) != 2 {
				return errors.New("invalid interface reflection")
			}
			reflection[pair[0]] = pair[1]
		}
		return gencode.GenerateFromInterface(inputPath, outputPath, mode, reflection)
	}

	return errors.New("invalid source type")
}




