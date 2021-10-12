/*
Copyright 2021 The pdfcpu Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

// CreateFormFromJSON reads a form definition from rd, creates a multiple page PDF form and writes it to outFile.
func CreateFormFromJSON(rd io.Reader, outFile string, conf *pdfcpu.Configuration) error {

	bb, err := ioutil.ReadAll(rd)
	if err != nil {
		return err
	}

	xRefTable, err := pdfcpu.CreateFormXRefFromJSON(bb)
	if err != nil {
		return err
	}

	return CreatePDFFile(xRefTable, outFile, conf)
}

// CreateFormFromJSONFile reads a form definition from inFile, creates a multi page PDF form and writes it to outFile.
func CreateFormFromJSONFile(inFile, outFile string, conf *pdfcpu.Configuration) error {

	f, err := os.Open(inFile)
	if err != nil {
		return err
	}

	return CreateFormFromJSON(f, outFile, conf)
}

// ExtractForm extracts form data from inFile into jsonFile.
func ExtractForm(inFile, jsonFile string, conf *pdfcpu.Configuration) error {
	return nil
}

// FillForm reads form data from jsonFile and fills the existing form in outfile.
func FillForm(jsonFile, outFile string, conf *pdfcpu.Configuration) error {
	return nil
}
