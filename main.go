package formatcss 

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	flag "github.com/ogier/pflag"
)

var replaceFlag, removeNestedCalcFlag, minFlag bool

func main() {
	flag.BoolVarP(&replaceFlag, "replace", "r", false, "Find and Replace CSS Variables with their value")
	flag.BoolVarP(&removeNestedCalcFlag, "nested-calc", "n", false, "Find and remove nested CSS 'calc()' expressions so 'calc' is only called once")
	flag.BoolVarP(&minFlag, "minimize", "m", false, "Minimize CSS file")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Fprintf(os.Stderr, "Error: You must provide a file path to process\n")
		fmt.Fprintf(os.Stdout, "%s foo.css\n", os.Args[0])
		flag.PrintDefaults()
		return
	}
	inputPath := flag.Args()[0]

	cssByte, err := ioutil.ReadFile(inputPath)
	if err != nil {
		log.Panic(err)
	}
	css := string(cssByte)

	if replaceFlag { // Remove and replace CSS variables
		css = RemoveCSSVariables(css)
		if removeNestedCalcFlag {
			css = RemoveNestedCalc(css)
		}
		if minFlag { // Minimize CSS file
			css = Minimize(css)
		}
		// Setup output file path
		outputDir := filepath.Dir(inputPath)
		inputFile := filepath.Base(inputPath)
		outputFile := inputFile[0:len(inputFile)-3] + "min.css"
		output := filepath.Join(outputDir, outputFile)
		// Write CSS to file
		ioutil.WriteFile(output, []byte(css), 0644)

	} else if minFlag { // Minimize CSS file
		css = Minimize(css)
		// Setup output file path
		outputDir := filepath.Dir(inputPath)
		inputFile := filepath.Base(inputPath)
		outputFile := inputFile[0:len(inputFile)-3] + "min.css"
		output := filepath.Join(outputDir, outputFile)
		// Write CSS to file
		ioutil.WriteFile(output, []byte(css), 0644)

	} else { // Format CSS file
		css = FormatCSS(css)
		// Write CSS to file
		ioutil.WriteFile(inputPath, []byte(css), 0644)
	}
}
