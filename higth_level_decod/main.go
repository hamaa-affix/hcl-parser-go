package main

import (
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)


func main() {
	spec := hcldec.ObjectSpec{
		"io_mode": &hcldec.AttrSpec{
			Name: "io_mode",
			Type: cty.String,
		},
		"services": &hcldec.BlockMapSpec{
			TypeName:   "service",
			LabelNames: []string{"type", "name"},
			Nested:     hcldec.ObjectSpec{
				"listen_addr": &hcldec.AttrSpec{
					Name:     "listen_addr",
					Type:     cty.String,
					Required: true,
				},
				"processes": &hcldec.BlockMapSpec{
					TypeName:   "service",
					LabelNames: []string{"name"},
					Nested:     hcldec.ObjectSpec{
						"command": &hcldec.AttrSpec{
							Name:     "command",
							Type:     cty.List(cty.String),
							Required: true,
						},
					},
				},
			},
		},
	}
	p := hclparse.NewParser()
	p.ParseHCLFile()
	val, moreDiags := hcldec.Decode(f.Body, spec, nil)
	diags = append(diags, moreDiags...)

}
