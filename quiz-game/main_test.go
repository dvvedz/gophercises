package main

import (
	"testing"
)

func TestReadCsvFileCorrectCsv(t *testing.T) {

	csv, _ := readCsvFile("_testCorrectCsv.csv")

	if len(csv) != 12 {
		t.Errorf("readCsvFile('_testCorrectCsv.csv') = %d; want 12", len(csv))
	}
}

func TestReadCsvFileIncorrectCsv(t *testing.T) {
	csv, err := readCsvFile("_testIncorrectCsv.csv")

	if len(csv) != 0 && err == nil {
		t.Errorf("readCsvFile('_testIncorrectCsv.csv') = %d, %v; want [], invalid csv data", len(csv), err)
	}
}

func TestReadCsvFileEmpty(t *testing.T) {
	csv, err := readCsvFile("")

	if len(csv) != 0 && err == nil {
		t.Errorf("readCsvFile('') = %d, %v; want [], could not read file", len(csv), err)
	}
}
