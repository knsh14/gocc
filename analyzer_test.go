package gocc

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "complexity")
}
