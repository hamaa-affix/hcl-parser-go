package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	// "strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
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
		blockPath := createBlockPath(block)
		fmt.Println(blockPath)

		// blockのbodyに新しいblockを追加
		testBlock := createNewBlock(block.Body(), "test", nil)
		testBlock.Body().SetAttributeValue("hoge", cty.StringVal("fuga"))

		fmt.Printf("--------------------\n")
		block.Body().SetAttributeTraversal("traversal", hcl.Traversal{
			hcl.TraverseRoot{Name: "root"},
			hcl.TraverseAttr{Name: "attr"},
			hcl.TraverseAttr{Name: "hgoe"},
		})

		// blockのbodyのattributeを取得
		rewriteToken(block.Body().Attributes(), "10.0.0.0/16", "127.0.0.0/32")

		// 特定のattributeを取得
		attr := block.Body().GetAttribute("instance_tenancy")
		if attr != nil {
			bytes := attr.Expr().BuildTokens(nil).Bytes()
			fmt.Printf("--------------------\n")
			fmt.Println(string(bytes))
		}
	}
	fmt.Printf("--------------------\n")
	updated := hclFile.BuildTokens(nil).Bytes()
	// formatを整えてくれる
	output := hclwrite.Format(updated)
	fmt.Fprint(os.Stdout, string(output))
}


func createBlockPath(block *hclwrite.Block) string {
	lables := block.Labels()
	resourceType := block.Type()

	var blockPathStrings []string
	blockPathStrings = append(blockPathStrings, resourceType)
	blockPathStrings = append(blockPathStrings, lables...)
	return strings.Join(blockPathStrings, ".")
}

func createNewBlock(body *hclwrite.Body, resourceType string, labels []string) *hclwrite.Block {
	if labels == nil {
		labels = []string{}
	}
	block := body.AppendNewBlock(resourceType, labels)
	return block
}

func rewriteToken(attributes map[string]*hclwrite.Attribute, conditionString, replaceString string) {
	// blockのbodyのattributeを取得
	for _, attribute := range attributes {
		// attributeのexpressionを取得
		tokens := attribute.Expr().BuildTokens(nil)
		for tokenIndex, token := range tokens {
			hclToken := string(token.Bytes)
			if hclToken == conditionString {
				tokens[tokenIndex].Bytes = []byte(strings.ReplaceAll(hclToken, hclToken, replaceString))
			}
		}
	}
}
