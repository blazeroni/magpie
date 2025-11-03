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

//go:embed nrgba.go.tmpl
var funcsTemplateFS embed.FS

type TemplateData struct {
	Name                      string
	KernelR, KernelG, KernelB string
	KernelDocs                string
}

func main() {
	processedData := processTemplates(shared.BlendTemplates)

	t, err := template.ParseFS(funcsTemplateFS, "nrgba.go.tmpl")
	if err != nil {
		log.Fatal("[nrgba] parsing template:", err)
	}

	var buf bytes.Buffer
	if err = t.Execute(&buf, processedData); err != nil {
		log.Fatal("[nrgba] executing template:", err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal("[nrgba] formatting generated code:", err)
	}

	if err = os.WriteFile("blend_gen.go", formatted, 0644); err != nil {
		log.Fatal("[nrgba] writing output file:", err)
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
			k = strings.ReplaceAll(k, "$R", "o"+ch)

			switch ch {
			case "R":
				td.KernelR = k
			case "G":
				td.KernelG = k
			case "B":
				td.KernelB = k
			}
		}
		data[i] = td
	}
	return data
}
