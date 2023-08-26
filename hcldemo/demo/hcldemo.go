package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

type foo struct {
    Num  int      `hcl:"num"`
    Chr  string   `hcl:"chr"`
    List []string `hcl:"list"`
}

func main() {
    Run()
}

func Run() {
    parser := hclparse.NewParser()
    file, diag := parser.ParseHCLFile("./demo.hcl")
    if diag != nil {
        log.Fatalf("parse error: %s", diag.Error())
    }

    var foo foo
    diag = gohcl.DecodeBody(file.Body, nil, &foo)
    if diag != nil {
        log.Fatalf("decode error: %s", diag.Error())
    }

    newHclFile := hclwrite.NewFile()
    gohcl.EncodeIntoBody(&foo, newHclFile.Body())
    fmt.Println(newHclFile.Bytes())

    f, err := os.Create("basic.tf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

    // write to file
    _, err = f.Write(newHclFile.Bytes())
    if err != nil {
        log.Fatal(err)
    }
}
