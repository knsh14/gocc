package gocc

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyer(t *testing.T) {
	Analyzer.Flags.Set("threshold", "1")
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "complexity")
}

func TestFilter(t *testing.T) {
	Analyzer.Flags.Set("threshold", "0")
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "filter")
}

func TestFilterVariables(t *testing.T) {
	Analyzer.Flags.Set("threshold", "1")
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "filter/variables")
}
