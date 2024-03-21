package main

import (
	"encoding/json"
	"fmt"
	htmlTemplate "html/template"
	"io"
	"os"
	textTemplate "text/template"

	flag "github.com/spf13/pflag"
)

var (
	templateFile string
	inputFile    string
	outputFile   string
	mode         string
	showVersion  bool
	version      string
	commit       string
	date         string
	builtBy      string
)

type templateRunner func(io.Writer, interface{}) error

func getTemplateRunner(mode, templateFile string) (templateRunner, error) {

	templateBytes, readErr := os.ReadFile(templateFile)
	if readErr != nil {
		return nil, readErr
	}
	templateText := string(templateBytes)

	if mode == "text" {
		textRunner, textErr := textTemplate.New("greta").Parse(templateText)
		return textRunner.Execute, textErr
	} else if mode == "html" {
		htmlRunner, htmlErr := htmlTemplate.New("greta").Parse(templateText)
		return htmlRunner.Execute, htmlErr
	}
	return nil, fmt.Errorf("mode %s not supported", mode)
}

func getData(inputFile string) (interface{}, error) {
	var raw []byte
	var readErr error
	if inputFile == "-" {
		raw, readErr = io.ReadAll(os.Stdin)
	} else {
		raw, readErr = os.ReadFile(inputFile)
	}
	if readErr != nil {
		return nil, readErr
	}

	var data any
	jsonErr := json.Unmarshal(raw, &data)
	return data, jsonErr
}

func Main() int {

	if os.Getenv("GITHUB_ACTIONS") == "true" {
		templateFile = os.Getenv("INPUT_TEMPLATE")
		inputFile = os.Getenv("INPUT_INPUT")
		if inputFile == "" {
			inputFile = "-"
		}
		outputFile = os.Getenv("INPUT_OUTPUT")
		mode = os.Getenv("INPUT_MODE")
		if mode == "" {
			mode = "html"
		}
	} else {
		flag.StringVarP(&templateFile, "template", "t", "", "template file (required)")
		flag.StringVarP(&inputFile, "input", "i", "-", "input file or '-' for stdin")
		flag.StringVarP(&outputFile, "output", "o", "", "output file (or leave blank for stdout)")
		flag.StringVarP(&mode, "mode", "m", "html", "mode: text or html")
		flag.BoolVarP(&showVersion, "version", "v", false, "show version")

		flag.Parse()

		if showVersion {
			fmt.Fprintf(os.Stdout, "greta v%s (from %s on %s by %s)\n", version, commit, date, builtBy)
			return 0
		}
	}

	if templateFile == "" {
		fmt.Fprintf(os.Stderr, "ERROR: template is required")
		return 1
	}

	templateFn, templateErr := getTemplateRunner(mode, templateFile)
	if templateErr != nil {
		fmt.Fprintf(os.Stderr, "ERROR: unable to parse template '%s': %s", templateFile, templateErr)
		return 1
	}

	data, dataErr := getData(inputFile)
	if dataErr != nil {
		fmt.Fprintf(os.Stderr, "ERROR: unable to read input file '%s': %s", inputFile, dataErr)
		return 2
	}

	var output io.Writer
	if outputFile == "" {
		output = os.Stdout
	} else {
		f, createErr := os.Create(outputFile)
		if createErr != nil {
			fmt.Fprintf(os.Stderr, "ERROR: unable to create output file '%s': %s", outputFile, createErr)
			return 3
		}
		defer f.Close()
		output = f
	}
	runErr := templateFn(output, data)
	if runErr != nil {
		fmt.Fprintf(os.Stderr, "ERROR: unable to run template '%s': %s", templateFile, runErr)
		return 4
	}

	return 0
}

func main() {
	os.Exit(Main())
}
