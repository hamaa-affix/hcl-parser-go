package main

import (
	"fmt"
	"log"
	"os"
	"strings"

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
			tokens := attribute.Expr().BuildTokens(nil)
			for tokenIndex, token := range tokens {
				hclToken := string(token.Bytes)
				if hclToken == "10.0.0.0/16" {
					tokens[tokenIndex].Bytes = []byte(strings.ReplaceAll(hclToken, hclToken, "127.0.0.0/32"))
				}
			}

			fmt.Println(string(tokens.Bytes()))
		}
	}
}
