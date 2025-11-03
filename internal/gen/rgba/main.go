// Copyright 2025 Magpie Contributors
// SPDX-License-Identifier: MIT

package main

import (
	"bytes"
	"embed"
	"go/format"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/blazeroni/magpie/internal/gen/shared"
)

//go:embed rgba.go.tmpl
var funcsTemplateFS embed.FS

// BlendMode holds the template strings for a single blend mode.
// It can contain a Kernel (for non-premultiplied logic) and an optional Equation
// (for a fast path using premultiplied logic).
type BlendMode struct {
	Name             string
	KernelTemplate   string
	EquationTemplate string // Optional
}

// TemplateData is the final struct passed to the template after placeholders are replaced.
type TemplateData struct {
	Name                            string
	KernelR, KernelG, KernelB       string
	EquationR, EquationG, EquationB string
	KernelDocs                      string
}

// testing

func main() {
	// Process the templates to generate per-channel logic
	processedData := processTemplates(shared.BlendTemplates)

	t, err := template.ParseFS(funcsTemplateFS, "rgba.go.tmpl")
	if err != nil {
		log.Fatal("[rgba] parsing template:", err)
	}

	var buf bytes.Buffer
	if err = t.Execute(&buf, processedData); err != nil {
		log.Fatal("[rgba] executing template:", err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal("[rgba] formatting generated code:", err)
	}

	if err = os.WriteFile("blend_gen.go", formatted, 0644); err != nil {
		log.Fatal("[rgba] writing output file:", err)
	}
}

// processTemplates expands the placeholder strings in the BlendMode definitions
// into the full per-channel logic required by the final template.
func processTemplates(modes []shared.BlendTemplate) []TemplateData {
	data := make([]TemplateData, len(modes))
	channels := []string{"R", "G", "B"}

	for i, mode := range modes {
		td := TemplateData{Name: mode.Name, KernelDocs: mode.KernelDocs}

		for _, ch := range channels {
			// Process Kernel Template (non-premultiplied)
			k := mode.Kernel
			k = strings.ReplaceAll(k, "$S", "s"+ch)
			k = strings.ReplaceAll(k, "$D", "d"+ch)
			k = strings.ReplaceAll(k, "$R", "k"+ch)

			// Process Equation Template (premultiplied)
			eq := mode.EquationRGBA
			eq = strings.ReplaceAll(eq, "$Sp", "s"+ch)
			eq = strings.ReplaceAll(eq, "$Dp", "d"+ch)
			eq = strings.ReplaceAll(eq, "$sA", "sA")
			eq = strings.ReplaceAll(eq, "$dA", "dA")

			switch ch {
			case "R":
				td.KernelR = k
				td.EquationR = eq
			case "G":
				td.KernelG = k
				td.EquationG = eq
			case "B":
				td.KernelB = k
				td.EquationB = eq
			}
		}
		data[i] = td
	}
	return data
}
