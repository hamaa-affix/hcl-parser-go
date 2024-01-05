package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func main() {
	filepath := "../terraform/main.tf"
	fileByteData, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	hclFile, diag := hclwrite.ParseConfig(fileByteData, filepath, hcl.InitialPos)
	if diag != nil {
		log.Fatal(diag)
	}

	// hcl fileの解析、body, blocksの取得
	for _, block := range hclFile.Body().Blocks() {
		// blockのbodyのattributeを取得
		for _, attribute := range block.Body().Attributes() {
			// attributeのexpressionを取得
			fmt.Println(string(attribute.Expr().BuildTokens(nil).Bytes()))
		}
	}
}
